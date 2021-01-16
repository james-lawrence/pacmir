package swarm

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	path "github.com/ipfs/interface-go-ipfs-core/path"
	peer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

// NewNode build a new node
func NewNode(ctx context.Context) (Node, error) {
	if err := plugins(""); err != nil {
		return nil, err
	}

	// Create a Temporary Repo
	repo, err := createDir(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create repository")
	}

	ipfs, err := createNode(ctx, repo)

	return node{ipfs: ipfs}, err
}

type node struct {
	ipfs icore.CoreAPI
}

func (t node) Upload(ctx context.Context, src io.Reader) (s string, err error) {
	var (
		cid path.Resolved
	)

	if cid, err = t.ipfs.Unixfs().Add(ctx, files.NewReaderFile(src)); err != nil {
		return "", errors.Wrap(err, "failed to add to storage")
	}

	return cid.String(), nil
}

func (t node) Download(ctx context.Context, id string, dst io.Writer) (err error) {
	var (
		ok      bool
		content files.File
	)

	n, err := t.ipfs.Unixfs().Get(ctx, path.New(id))
	if err != nil {
		return errors.Wrap(err, "unable to locate content")
	}

	if content, ok = n.(files.File); !ok {
		return errors.Errorf("%s is not a file", id)
	}

	if _, err = io.Copy(dst, content); err != nil {
		return errors.Wrap(err, "unable to download content")
	}

	return nil
}

func (t node) Connect(ctx context.Context, peers ...string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peerstore.PeerInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peerstore.InfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peerstore.PeerInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peerstore.PeerInfo) {
			defer wg.Done()
			log.Println("connecting", peerInfo)
			err := t.ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			}
		}(peerInfo)
	}
	wg.Wait()

	return nil
}

// IPFS internals are a mess, and their plugin system is required because
// otherwise none of their shit works. classic overengineering problems.
func plugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

func createDir(ctx context.Context) (string, error) {
	const path = ".ipfs-repo"
	err := os.MkdirAll(path, 0700)
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	// Create the repo with the config
	err = fsrepo.Init(path, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return path, nil
}

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		Repo:    repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}
