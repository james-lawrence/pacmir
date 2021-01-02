package httputilx

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

// Binder interface for binding handlers to routes.
type Binder interface {
	Bind(r *mux.Router)
}

// BinderFunc pure function based binder
type BinderFunc func(r *mux.Router)

// Bind binds some handler to the provided router
func (t BinderFunc) Bind(r *mux.Router) {
	t(r)
}

// NotFound handles paths that do not exist.
func NotFound(autowrap alice.Chain) Binder {
	return BinderFunc(func(r *mux.Router) {
		notFound := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			raw, _ := httputil.DumpRequest(req, false)
			log.Println("requested endpoint not found", string(raw))
			resp.WriteHeader(http.StatusNotFound)
		})

		r.NotFoundHandler = autowrap.Then(notFound)
	})
}

// RouteInvokedHandler wraps a http.Handler and emits route invocations.
func RouteInvokedHandler(original http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		p := req.Host + req.URL.Path
		if route := mux.CurrentRoute(req); route != nil && len(route.GetName()) > 0 {
			p = route.GetName()
		}
		started := time.Now()
		log.Println(p, "invoked")
		original.ServeHTTP(resp, req)
		log.Println(p, "completed", time.Since(started))
	})
}

// DumpRequestHandler dumps the request to STDERR.
func DumpRequestHandler(original http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		raw, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Println(errors.Wrap(err, "failed to dump request"))
		} else {
			log.Println(string(raw))
		}
		original.ServeHTTP(resp, req)
	})
}
