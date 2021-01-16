package dagpb

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ipfs/go-cid"
	merkledag_pb "github.com/ipfs/go-merkledag/pb"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/schema"
)

// DecodeDagProto is a fast path decoding to protobuf
// from PBNode__NodeBuilders
func (nb _PBNode__NodeBuilder) DecodeDagProto(r io.Reader) error {
	var pbn merkledag_pb.PBNode
	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("io error during unmarshal. %v", err)
	}
	if err := pbn.Unmarshal(encoded); err != nil {
		return fmt.Errorf("unmarshal failed. %v", err)
	}
	pbLinks := make([]PBLink, 0, len(pbn.Links))
	for _, link := range pbn.Links {
		hash, err := cid.Cast(link.GetHash())

		if err != nil {
			return fmt.Errorf("unmarshal failed. %v", err)
		}
		pbLinks = append(pbLinks, PBLink{
			d: PBLink__Content{
				Hash: MaybeLink{
					Maybe: schema.Maybe_Value,
					Value: Link{cidlink.Link{Cid: hash}},
				},
				Name: MaybeString{
					Maybe: schema.Maybe_Value,
					Value: String{link.GetName()},
				},
				Tsize: MaybeInt{
					Maybe: schema.Maybe_Value,
					Value: Int{int(link.GetTsize())},
				},
			},
		})
	}
	nb.nd.d.Links.x = pbLinks
	nb.nd.d.Data.x = pbn.GetData()
	return nil
}

// EncodeDagProto is a fast path encoding to protobuf
// for PBNode types
func (nd PBNode) EncodeDagProto(w io.Writer) error {
	pbn := new(merkledag_pb.PBNode)
	pbn.Links = make([]*merkledag_pb.PBLink, 0, len(nd.d.Links.x))
	for _, link := range nd.d.Links.x {
		var hash []byte
		if link.d.Hash.Maybe == schema.Maybe_Value {
			cid := link.d.Hash.Value.x.(cidlink.Link).Cid
			if cid.Defined() {
				hash = cid.Bytes()
			}
		}
		var name *string
		if link.d.Name.Maybe == schema.Maybe_Value {
			tmp := link.d.Name.Value.x
			name = &tmp
		}
		var tsize *uint64
		if link.d.Tsize.Maybe == schema.Maybe_Value {
			tmp := uint64(link.d.Tsize.Value.x)
			tsize = &tmp
		}
		pbn.Links = append(pbn.Links, &merkledag_pb.PBLink{
			Hash:  hash,
			Name:  name,
			Tsize: tsize})
	}
	pbn.Data = nd.d.Data.x
	data, err := pbn.Marshal()
	if err != nil {
		return fmt.Errorf("marshal failed. %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf(" error during marshal. %v", err)
	}
	return nil
}

// DecodeDagRaw is a fast path decoding to protobuf
// from RawNode__NodeBuilders
func (nb _RawNode__NodeBuilder) DecodeDagRaw(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("io error during unmarshal. %v", err)
	}
	nb.nd.x = data
	return nil
}

// EncodeDagRaw is a fast path encoding to protobuf
// for RawNode types
func (nd RawNode) EncodeDagRaw(w io.Writer) error {
	_, err := w.Write(nd.x)
	if err != nil {
		return fmt.Errorf(" error during marshal. %v", err)
	}
	return nil
}
