package swarm

import (
	"context"
	"io"
)

// Node in the swarm
type Node interface {
	Upload(ctx context.Context, src io.Reader) (string, error)
	Download(ctx context.Context, cid string, dst io.Writer) error
	Connect(ctx context.Context, peers ...string) error
}
