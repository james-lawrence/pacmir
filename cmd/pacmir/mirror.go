package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

type mirror struct {
	HTTPBind string `default:"localhost:4000" help:"HTTP address to bind the mirror"`
}

func (t mirror) Run(ctx *context) error {
	log.Println("mirror", spew.Sdump(t), spew.Sdump(ctx))
	return nil
}

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	alpm "github.com/Jguer/go-alpm/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/james-lawrence/pacmir"
	"github.com/james-lawrence/pacmir/internal/httputilx"
	"github.com/james-lawrence/pacmir/localmir"
	"github.com/james-lawrence/pacmir/mirrors"
	"github.com/james-lawrence/torrent"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

// Mirror command
type Mirror struct {
	HTTPBind string   `default:"localhost:4000" help:"HTTP address to bind the mirror"`
}

// Run the command
func (t *Mirror) Run(ctx *context) (err error) {
	var (
		tclient    *torrent.Client
		middleware = alice.New(
			httputilx.RouteInvokedHandler,
		)
		router = mux.NewRouter()
	)

	tclient, err = torrent.NewClient(
		torrent.NewDefaultClientConfig(
			torrent.ClientConfigSeed(true),
		),
	)

	log.Println("initiating mirror daemon", t.HTTPBind)

	cconfig := pacmir.NewCachedConfig(ctx.Config)
	cconfig.Current()
	prouter := router.PathPrefix("/{repo}/os/{arch}").Subrouter()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

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

