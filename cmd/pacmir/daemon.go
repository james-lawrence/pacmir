package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/james-lawrence/pacmir"
	"github.com/james-lawrence/pacmir/internal/httputilx"
	"github.com/james-lawrence/pacmir/localmir"
	"github.com/james-lawrence/pacmir/mirrors"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

// Daemon command
type Daemon struct {
	HTTPBind string   `default:"localhost:4000" help:"HTTP address to bind the mirror"`
	Mirrors  []string `default:"/etc/pacman.d/mirrorlist" help:"mirror list files to rewrite"`
}

// Run the command
func (t *Daemon) Run(ctx *context) (err error) {
	var (
		middleware = alice.New(
			httputilx.RouteInvokedHandler,
		)
		router = mux.NewRouter()
	)

	log.Println("initiating local mirror daemon", t.HTTPBind)
	for _, mirror := range t.Mirrors {
		if _, err = os.Stat(mirror); err != nil {
			log.Println("ignored mirror file (missing)", mirror)
			continue
		}

		log.Println("rewriting mirror file", mirror)

		if err = mirrors.Rewrite(t.HTTPBind, mirror); err != nil {
			return err
		}
	}
	cconfig := pacmir.NewCachedConfig(ctx.Config)

	prouter := router.PathPrefix("/{repo}/os/{arch}").Subrouter()
	fallback := localmir.Proxied{
		HTTPAddress: t.HTTPBind,
		Pacman:      cconfig,
	}
	fallback.Bind(middleware, prouter)

	localmir.Download{
		Downloader: fspackager{
			cached: cconfig,
		},
		Fallback: http.HandlerFunc(fallback.Proxy),
	}.Bind(middleware.Append(
		httputilx.DumpRequestHandler,
	), prouter)

	httputilx.NotFound(middleware).Bind(router)

	return http.ListenAndServe(t.HTTPBind, router)
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
