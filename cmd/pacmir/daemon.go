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

	// if tsocket, err = utp.NewSocket("udp", ":0"); err != nil {
	// 	return err
	// }

	// tclient, err = torrent.NewClient(
	// 	torrent.NewDefaultClientConfig(
	// 		torrent.ClientConfigSeed(true),
	// 		func(c *torrent.ClientConfig) {
	// 			c.DHTOnQuery = func(query *krpc.Msg, source net.Addr) bool {
	// 				log.Println("query", source.String(), spew.Sdump(query))
	// 				return true
	// 			}
	// 			c.DHTAnnouncePeer = func(infoHash metainfo.Hash, ip net.IP, port int, portOK bool) {
	// 				log.Printf("announce peer %s - %s:%d %t\n", infoHash.String(), ip.String(), port, portOK)
	// 			}
	// 			c.DhtStartingNodes = func() (peers []dht.Addr, err error) {
	// 				peers = []dht.Addr{
	// 					dht.NewAddr(&net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 4002}),
	// 				}
	// 				return peers, err
	// 			}
	// 		},
	// 	),
	// )
	// if tclient, err = torrent.NewSocketsBind(sockets.New(tsocket, &net.Dialer{LocalAddr: tsocket.Addr()})).Bind(tclient, err); err != nil {
	// 	return errors.Wrap(err, "unable to create torrent service")
	// }

	log.Println("initiating local mirror daemon", t.HTTPBind)
	// for _, mirror := range t.Mirrors {
	// 	if _, err = os.Stat(mirror); err != nil {
	// 		log.Println("ignored mirror file (missing)", mirror)
	// 		continue
	// 	}

	// 	log.Println("rewriting mirror file", mirror)

	// 	if err = mirrors.Rewrite(t.HTTPBind, mirror); err != nil {
	// 		return err
	// 	}
	// }
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

	// torrentpackager{
	// 	client: tclient,
	// 	cached: cconfig,
	// }.Package("core", "openssh")
	return http.ListenAndServe(t.HTTPBind, router)
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
