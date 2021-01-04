package main

import (
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/james-lawrence/pacmir/internal/httputilx"
	"github.com/james-lawrence/torrent"
	"github.com/justinas/alice"
)

// Mirror command
type Mirror struct {
	RootDirectory string `default:"/var/cache/pacmir" help:"mirror root directory"`
	HTTPBind      string `default:"localhost:4001" help:"HTTP address to bind the mirror"`
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

	if err != nil {
		return err
	}

	_ = tclient

	log.Println("initiating mirror daemon", spew.Sdump(t))
	router.Handle("/mirror", middleware.Then(http.FileServer(http.Dir(t.RootDirectory))))
	httputilx.NotFound(middleware).Bind(router)

	return http.ListenAndServe(t.HTTPBind, router)
}
