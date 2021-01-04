package torrent

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/james-lawrence/torrent/metainfo"
	"github.com/pkg/errors"
)

func newDigests(retrieve func(int) *Piece, complete func(int, error)) digests {
	return digests{
		retrieve: retrieve,
		complete: complete,
		pending:  newBitQueue(),
		c:        sync.NewCond(&sync.Mutex{}),
	}
}

// digests is responsible correctness of received data.
type digests struct {
	retrieve func(int) *Piece
	complete func(int, error)
	// marks whether digest is actively processing.
	reaping int64
	// cache of the pieces that need to be verified.
	pending bitQueue
	c       *sync.Cond
}

// Enqueue a piece to check its completed digest.
func (t *digests) Enqueue(idx int) {
	t.pending.Push(idx)
	t.verify()
}

func (t *digests) verify() {
	if atomic.AddInt64(&t.reaping, 1) > int64(runtime.NumCPU()) {
		atomic.AddInt64(&t.reaping, -1)
		return
	}

	go func() {
		for idx, ok := t.pending.Pop(); ok; idx, ok = t.pending.Pop() {
			t.check(idx)
		}

		atomic.AddInt64(&t.reaping, -1)
	}()
}

func (t *digests) check(idx int) {
	var (
		err    error
		digest metainfo.Hash
		p      *Piece
	)

	if p = t.retrieve(idx); p == nil {
		t.complete(idx, fmt.Errorf("piece %d not found during digest", idx))
		return
	}

	if digest, err = t.compute(p); err != nil {
		t.complete(idx, err)
		return
	}

	if digest != *p.hash {
		t.complete(idx, fmt.Errorf("piece %d digest mismatch %s != %s", idx, hex.EncodeToString(digest[:]), hex.EncodeToString(p.hash[:])))
		return
	}

	t.complete(idx, nil)
}

func (t *digests) compute(p *Piece) (ret metainfo.Hash, err error) {
	c := sha1.New()
	p.waitNoPendingWrites()

	pl := int64(p.length())

	n, err := io.Copy(c, io.NewSectionReader(p.Storage(), 0, pl))
	if err != nil {
		return ret, errors.Wrapf(err, "piece %d digest failed:", p.index)
	}

	if n != int64(pl) {
		return ret, fmt.Errorf("piece digest failed short copy %d: %d != %d", p.index, n, pl)
	}

	copy(ret[:], c.Sum(nil))
	return ret, nil
}
