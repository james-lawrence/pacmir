package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/james-lawrence/pacmir"
	"github.com/james-lawrence/pacmir/internal/httputilx"
)

func main() {
	var (
		middleware = alice.New(
			httputilx.RouteInvokedHandler,
		)
		router = mux.NewRouter()
	)

	cconfig := pacmir.NewCachedConfig("/etc/pacman.conf")
	prouter := router.PathPrefix("/{repo}/os/{arch}").Subrouter()
	fallback := proxied{
		pacman: cconfig,
	}
	fallback.Bind(middleware, prouter)

	torrent{
		packager: fspackager{
			cached: cconfig,
		},
		fallback: http.HandlerFunc(fallback.proxy),
	}.Bind(middleware.Append(
		httputilx.DumpRequestHandler,
	), prouter)

	httputilx.NotFound(middleware).Bind(router)

	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatalln(err)
	}
}

type fspackager struct {
	cached *pacmir.CachedConfig
}

func (t fspackager) Package(name string) (io.ReadCloser, error) {
	config := t.cached.Current()
	if config == nil {
		return nil, errors.New("missing pacman configuration")
	}

	for _, d := range config.CacheDir {
		path := filepath.Join(d, name)
		log.Println("checking", path)
		if _, err := os.Stat(path); err == nil {
			return os.Open(path)
		}
	}

	return nil, errors.New("package not found")
}
