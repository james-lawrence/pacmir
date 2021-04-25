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
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

// Daemon command
type Daemon struct {
	HTTPBind string   `default:"localhost:4000" help:"HTTP address to bind the mirror"`
	Mirrors  []string `default:"/etc/pacman.d/mirrorlist" help:"mirror list files to rewrite"`
}

// Run the command
func (t *Daemon) Run(ctx *CmdContext) (err error) {
	var (
		// tsocket    *utp.Socket
		// tclient    *torrent.Client
		middleware = alice.New(
			httputilx.RouteInvokedHandler,
		)
		router = mux.NewRouter()
	)

	// var (
	// 	l net.Listener
	// 	m = muxer.New()
	// 	p2ppriv    []byte
	// )

	// if p2ppriv, err = rsax.CachedAuto("p2p.key"); err != nil {
	// 	return err
	// }

	// tmpl, err := tlsx.X509Template(
	// 	10*360*24*time.Hour,
	// 	tlsx.X509OptionCA(),
	// 	tlsx.X509OptionSubject(pkix.Name{
	// 		CommonName: "pacmir.lan",
	// 	}),
	// 	tlsx.X509OptionHosts("pacmir.lan"),
	// )
	// if err != nil {
	// 	return err
	// }
	// _, certblock, err := tlsx.SelfSigned(rsax.MustDecode(p2ppriv), tmpl)
	// if err != nil {
	// 	return err
	// }

	// cert, err := tls.X509KeyPair(certblock, p2ppriv)
	// if err != nil {
	// 	log.Println("certblock", len(certblock))
	// 	return errors.Wrap(err, "failed to parse x509 keypair")
	// }

	// tlsconfig := &tls.Config{
	// 	Certificates: []tls.Certificate{
	// 		cert,
	// 	},
	// 	NextProtos: []string{"bw.mux"},
	// }
	// if l, err = net.Listen("tcp", t.HTTPBind); err != nil {
	// 	return err
	// }
	// l = tls.NewListener(l, tlsconfig)

	// go muxer.Background(context.Background(), m, l)

	log.Println("initiating local mirror daemon", t.HTTPBind)

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
	// return func(l net.Listener, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	http.Serve(l, router)
	// 	return nil
	// }(m.Default("http", l.Addr()))
}

type fspackager struct {
	cached *pacmir.CachedConfig
}

func (t fspackager) Package(repo string, name string) (io.ReadCloser, error) {
	config := t.cached.Current()
	if config == nil {
		return nil, errors.New("missing pacman configuration")
	}

	for _, d := range config.CacheDir {
		path := filepath.Join(d, name)
		if _, err := os.Stat(path); err == nil {
			return os.Open(path)
		}
	}

	return nil, errors.New("package not found")
}

// type torrentpackager struct {
// 	client   *torrent.Client
// 	cached   *pacmir.CachedConfig
// 	fallback fspackager
// }

// func (t torrentpackager) Package(repo string, name string) (_ io.ReadCloser, err error) {
// 	var (
// 		alpmh *alpm.Handle
// 		db    alpm.IDB
// 		pkg   alpm.IPackage
// 	)

// 	config := t.cached.Current()
// 	if config == nil {
// 		return nil, errors.New("missing pacman configuration")
// 	}

// 	log.Println("Downloading", repo, name)
// 	if alpmh, err = alpm.Initialize(config.RootDir, config.DBPath); err != nil {
// 		return nil, err
// 	}
// 	defer alpmh.Release()

// 	if db, err = alpmh.RegisterSyncDB(repo, 0); err != nil {
// 		return nil, err
// 	}

// 	if pkg, err = db.PkgCache().FindSatisfier(name); err != nil {
// 		return nil, err
// 	}

// 	log.Println("WAAAT", pkg.Name(), pkg.SHA256Sum(), pkg.Size())
// 	for _, s := range t.client.DhtServers() {
// 		s.Announce([20]byte{}, 4000, true)
// 	}
// 	return nil, errors.New("not implemented")
// }
