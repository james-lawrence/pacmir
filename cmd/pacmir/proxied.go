package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/james-lawrence/pacmir"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

type proxied struct {
	pacman *pacmir.CachedConfig
}

func (t proxied) Bind(c alice.Chain, r *mux.Router) {
	r.Handle("/{package}.db", c.ThenFunc(t.proxy))
	r.Handle("/{package}.sig", c.ThenFunc(t.proxy))
}

func (t proxied) proxy(resp http.ResponseWriter, req *http.Request) {
	var (
		err     error
		proxied *http.Response
		params  = mux.Vars(req)
		rname   = params["repo"]
		arch    = params["arch"]
	)

	mirrors := t.pacman.Mirrors(rname)

	if len(mirrors) == 0 {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	for _, s := range mirrors {
		if proxied != nil {
			// reset previous proxied request because it was rejected.
			proxied.Body.Close()
			proxied = nil
		}

		proxieduri := strings.ReplaceAll(s, fmt.Sprintf("/%s/os/%s", rname, arch), req.RequestURI)
		if strings.Contains(s, "localhost") {
			continue
		}

		if proxied, err = http.Get(proxieduri); err != nil {
			log.Println("skipping", proxieduri, err)
			continue
		}

		if proxied.StatusCode != http.StatusOK {
			log.Println("skipping", proxieduri, proxied.StatusCode, proxied.Status)
			continue
		}

		break
	}
	defer proxied.Body.Close()

	for k, v := range proxied.Header {
		resp.Header()[k] = v
	}
	resp.WriteHeader(proxied.StatusCode)

	if n, err := io.CopyN(resp, proxied.Body, proxied.ContentLength); err != nil {
		// silence broken pipe errors. pacman is evil and just nukes the connection if it doesn't need
		// to request all the data.
		if cause := new(syscall.Errno); errors.As(err, cause) && cause.Error() == "broken pipe" {
			return
		}

		log.Println("proxy failed", n, proxied.ContentLength, err)
		log.Printf("%T\n", err)
	}
}
