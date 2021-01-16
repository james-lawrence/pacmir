package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/james-lawrence/pacmir/swarm"

	"github.com/pkg/errors"
	// This package is needed so that all the preloaded plugins are loaded automatically
)

// Spike command
type Spike struct {
	HTTPBind string   `default:"localhost:4000" help:"HTTP address to bind the mirror"`
	Mirrors  []string `default:"/etc/pacman.d/mirrorlist" help:"mirror list files to rewrite"`
}

// Run the command
func (t *Spike) Run(cctx *CmdContext) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// cconfig := pacmir.NewCachedConfig(ctx.Config)

	// torrentpackager{
	// 	cached: cconfig,
	// }.Package("community", "0ad")
	n, err := swarm.NewNode(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to initialize node")
	}

	log.Println("node is running")

	err = n.Connect(ctx,
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
		// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
	)

	if err != nil {
		return err
	}

	// cid := "/ipfs/QmV4eunfXteYNYLNpbvAWUVW1t1nNmWHY3HJxzvmZiqM2S"
	sfile, err := os.Open("linux-5.10.5.arch1-1-x86_64.pkg.tar.zst")
	if err != nil {
		return err
	}
	defer sfile.Close()

	cid, err := n.Upload(ctx, sfile)
	if err != nil {
		return err
	}

	// log.Println("CID", cid)
	dfile, err := os.Create("linux-5.10.5.arch1-1-x86_64.pkg.tar.zst.download")
	if err != nil {
		return err
	}
	defer dfile.Close()

	if err = n.Download(ctx, cid, dfile); err != nil {
		return err
	}

	return nil
}
