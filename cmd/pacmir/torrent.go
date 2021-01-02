package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type packager interface {
	Package(name string) (io.ReadCloser, error)
}

type torrent struct {
	packager packager
	fallback http.Handler
}

func (t torrent) Bind(c alice.Chain, r *mux.Router) {
	r.Handle("/{package}", c.ThenFunc(t.download))
}

func (t torrent) download(resp http.ResponseWriter, req *http.Request) {
	var (
		err    error
		pdata  io.ReadCloser
		params = mux.Vars(req)
		pname  = params["package"]
	)

	log.Println("torrent download", params)
	if pdata, err = t.packager.Package(pname); err != nil {
		t.fallback.ServeHTTP(resp, req)
		return
	}

	resp.WriteHeader(http.StatusOK)

	if n, err := io.Copy(resp, pdata); err != nil {
		log.Println("proxy failed", n, err)
		log.Printf("%T\n", err)
	}
}
