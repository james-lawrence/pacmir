package localmir

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

// Download allow downloading packages.
type Download struct {
	Downloader packager
	Fallback   http.Handler
}

// Bind to a router
func (t Download) Bind(c alice.Chain, r *mux.Router) {
	r.Handle("/{package}", c.ThenFunc(t.download))
}

func (t Download) download(resp http.ResponseWriter, req *http.Request) {
	var (
		err    error
		pdata  io.ReadCloser
		params = mux.Vars(req)
		pname  = params["package"]
	)

	log.Println("torrent download", params)
	if pdata, err = t.Downloader.Package(pname); err != nil {
		t.Fallback.ServeHTTP(resp, req)
		return
	}

	resp.WriteHeader(http.StatusOK)

	if n, err := io.Copy(resp, pdata); err != nil {
		log.Println("proxy failed", n, err)
		log.Printf("%T\n", err)
	}
}
