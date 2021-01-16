package main

import (
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/james-lawrence/pacmir/internal/httputilx"
	"github.com/justinas/alice"
)

// Mirror command
type Mirror struct {
	RootDirectory string `default:"/var/cache/pacmir" help:"mirror root directory" env:"CACHE_DIRECTORY"`
	HTTPBind      string `default:":4002" help:"HTTP address to bind the mirror"`
}

// Run the command
func (t *Mirror) Run(ctx *CmdContext) (err error) {
	var (
		// tsocket    *utp.Socket
		// tclient    *torrent.Client
		middleware = alice.New(
			httputilx.RouteInvokedHandler,
		)
		router = mux.NewRouter()
		// mirror torrents.Mirror
	)

	// if tsocket, err = utp.NewSocket("udp", t.HTTPBind); err != nil {
	// 	return err
	// }

	// tclient, err = torrent.NewClient(
	// 	torrent.NewDefaultClientConfig(
	// 		torrent.ClientConfigSeed(true),
	// 		func(c *torrent.ClientConfig) {
	// 			c.DhtStartingNodes = func() ([]dht.Addr, error) {
	// 				return []dht.Addr{
	// 					dht.NewAddr(&net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 4002}),
	// 				}, nil
	// 			}
	// 			c.DataDir = filepath.Join(t.RootDirectory, "pool", "packages")
	// 		},
	// 		func(c *torrent.ClientConfig) {
	// 			c.DHTOnQuery = func(query *krpc.Msg, source net.Addr) bool {
	// 				log.Println("query", source.String(), query.Q, spew.Sdump(query.A))
	// 				return true
	// 			}
	// 			c.DHTAnnouncePeer = func(infoHash metainfo.Hash, ip net.IP, port int, portOK bool) {
	// 				log.Printf("announce peer %s - %s:%d %t\n", infoHash.String(), ip.String(), port, portOK)
	// 			}
	// 		},
	// 	),
	// )
	// if tclient, err = torrent.NewSocketsBind(sockets.New(tsocket, &net.Dialer{LocalAddr: tsocket.Addr()})).Bind(tclient, err); err != nil {
	// 	return err
	// }

	// go timex.NowAndEvery(20*time.Second, func() {
	// 	tclient.WriteStatus(os.Stderr)
	// 	log.Println(len(tclient.Torrents()), "torrents running")
	// })

	// if mirror, err = torrents.NewMirror(tclient, filepath.Join(t.RootDirectory, "pacmir-torrents")); err != nil {
	// 	return err
	// }

	// if err = mirror.Bind(t.RootDirectory); err != nil {
	// 	return err
	// }

	// go mirror.AutoTorrents(context.Background())

	log.Println("initiating mirror daemon", spew.Sdump(t))
	router.Handle("/mirror", middleware.Then(http.FileServer(http.Dir(t.RootDirectory))))
	httputilx.NotFound(middleware).Bind(router)

	return http.ListenAndServe(t.HTTPBind, router)
}
