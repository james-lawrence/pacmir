package torrent

import (
	"fmt"
	"sync"

	"github.com/anacrolix/missinggo/bitmap"

	"github.com/james-lawrence/torrent/metainfo"
	pp "github.com/james-lawrence/torrent/peer_protocol"
	"github.com/james-lawrence/torrent/storage"
)

// Describes the importance of obtaining a particular piece.
type piecePriority byte

func (pp *piecePriority) Raise(maybe piecePriority) bool {
	if maybe > *pp {
		*pp = maybe
		return true
	}
	return false
}

// Priority for use in PriorityBitmap
func (me piecePriority) BitmapPriority() int {
	return -int(me)
}

const (
	PiecePriorityNone      piecePriority = iota // Not wanted. Must be the zero value.
	PiecePriorityNormal                         // Wanted.
	PiecePriorityHigh                           // Wanted a lot.
	PiecePriorityReadahead                      // May be required soon.
	// Succeeds a piece where a read occurred. Currently the same as Now,
	// apparently due to issues with caching.
	PiecePriorityNext
	PiecePriorityNow // A Reader is reading in this piece. Highest urgency.
)

type Piece struct {
	// The completed piece SHA1 hash, from the metainfo "pieces" field.
	hash  *metainfo.Hash
	t     *torrent
	index pieceIndex

	// Chunks we've written to since the last check. The chunk offset and
	// length can be determined by the request chunkSize in use.
	dirtyChunks bitmap.Bitmap

	hashing             bool
	numVerifies         int64
	storageCompletionOk bool

	publicPieceState PieceState
	priority         piecePriority

	pendingWritesMutex sync.Mutex
	pendingWrites      int
	noPendingWrites    sync.Cond
}

func (p *Piece) String() string {
	return fmt.Sprintf("%s/%d", p.t.infoHash.HexString(), p.index)
}

func (p *Piece) Info() metainfo.Piece {
	return p.t.info.Piece(int(p.index))
}

func (p *Piece) Storage() storage.Piece {
	return p.t.storage.Piece(p.Info())
}

func (p *Piece) pendingChunkIndex(chunkIndex int) bool {
	if p.dirtyChunks.IsEmpty() {
		return true
	}
	return !p.dirtyChunks.Contains(chunkIndex)
}

func (p *Piece) pendingChunk(cs chunkSpec, chunkSize pp.Integer) bool {
	return p.pendingChunkIndex(chunkIndex(cs, chunkSize))
}

func (p *Piece) hasDirtyChunks() bool {
	return p.dirtyChunks.Len() != 0
}

func (p *Piece) numDirtyChunks() pp.Integer {
	return pp.Integer(p.dirtyChunks.Len())
}

func (p *Piece) unpendChunkIndex(i int) {
	p.dirtyChunks.Add(i)
	p.t.tickleReaders()
}

func (p *Piece) pendChunkIndex(i int) {
	p.dirtyChunks.Remove(i)
}

func (p *Piece) numChunks() pp.Integer {
	return p.t.pieceNumChunks(p.index)
}

func (p *Piece) undirtiedChunkIndices() (ret bitmap.Bitmap) {
	ret = p.dirtyChunks.Copy()
	ret.FlipRange(0, bitmap.BitIndex(p.numChunks()))
	return
}

func (p *Piece) incrementPendingWrites() {
	p.pendingWritesMutex.Lock()
	p.pendingWrites++
	p.pendingWritesMutex.Unlock()
}

func (p *Piece) decrementPendingWrites() {
	p.pendingWritesMutex.Lock()
	if p.pendingWrites == 0 {
		panic("assertion")
	}
	p.pendingWrites--
	if p.pendingWrites == 0 {
		p.noPendingWrites.Broadcast()
	}
	p.pendingWritesMutex.Unlock()
}

func (p *Piece) waitNoPendingWrites() {
	p.pendingWritesMutex.Lock()
	for p.pendingWrites != 0 {
		p.noPendingWrites.Wait()
	}
	p.pendingWritesMutex.Unlock()
}

func (p *Piece) chunkIndexDirty(chunk pp.Integer) bool {
	if p.dirtyChunks.IsEmpty() {
		return false
	}

	if p.dirtyChunks.Len() == 0 {
		return false
	}

	return p.dirtyChunks.Contains(bitmap.BitIndex(chunk))
}

func (p *Piece) chunkIndexSpec(chunk pp.Integer) chunkSpec {
	return chunkIndexSpec(chunk, p.length(), p.chunkSize())
}

func (p *Piece) numDirtyBytes() (ret pp.Integer) {
	// defer func() {
	// 	if ret > p.length() {
	// 		panic("too many dirty bytes")
	// 	}
	// }()
	numRegularDirtyChunks := p.numDirtyChunks()
	if p.chunkIndexDirty(p.numChunks() - 1) {
		numRegularDirtyChunks--
		ret += p.chunkIndexSpec(p.lastChunkIndex()).Length
	}
	ret += pp.Integer(numRegularDirtyChunks) * p.chunkSize()
	return
}

func (p *Piece) length() pp.Integer {
	return p.t.pieceLength(p.index)
}

func (p *Piece) chunkSize() pp.Integer {
	return p.t.chunkSize
}

func (p *Piece) lastChunkIndex() pp.Integer {
	return p.numChunks() - 1
}

func (p *Piece) bytesLeft() (ret pp.Integer) {
	if p.t.pieceComplete(p.index) {
		return 0
	}
	return p.length() - p.numDirtyBytes()
}

// VerifyData forces the piece data to be rehashed.
func (p *Piece) VerifyData() {
	p.t.lock()
	defer p.t.unlock()
	target := p.numVerifies + 1
	if p.hashing {
		target++
	}

	p.t.digests.Enqueue(p.index)

	for {
		if p.numVerifies >= target {
			break
		}

		p.t.event.Wait()
	}
}

func (p *Piece) queuedForHash() bool {
	return p.t.chunks.ChunksComplete(p.index)
}

func (p *Piece) torrentBeginOffset() int64 {
	return int64(p.index) * p.t.info.PieceLength
}

func (p *Piece) torrentEndOffset() int64 {
	return p.torrentBeginOffset() + int64(p.length())
}

func (p *Piece) SetPriority(prio piecePriority) {
	p.t.lock()
	defer p.t.unlock()
	p.priority = prio
	p.t.updatePiecePriority(p.index)
}

func (p *Piece) uncachedPriority() (ret piecePriority) {
	if p.t.pieceComplete(p.index) || p.t.chunks.ChunksComplete(p.index) {
		return PiecePriorityNone
	}

	if p.t.readerNowPieces.Contains(int(p.index)) {
		ret.Raise(PiecePriorityNow)
	}

	if p.t.readerReadaheadPieces.Contains(bitmap.BitIndex(p.index)) {
		ret.Raise(PiecePriorityReadahead)
	}
	ret.Raise(p.priority)
	return
}

func (p *Piece) completion() (ret storage.Completion) {
	ret.Complete = p.t.pieceComplete(p.index)
	ret.Ok = p.storageCompletionOk
	return
}
