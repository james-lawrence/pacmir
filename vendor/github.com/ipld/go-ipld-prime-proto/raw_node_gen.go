package dagpb

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
)

// Code generated go-ipld-prime DO NOT EDIT.

var _ ipld.Node = RawNode{}
var _ schema.TypedNode = RawNode{}

func (RawNode) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Bytes
}
func (RawNode) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}}
}
func (RawNode) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}}
}
func (RawNode) Length() int {
	return -1
}
func (RawNode) IsUndefined() bool {
	return false
}
func (RawNode) IsNull() bool {
	return false
}
func (RawNode) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Bytes}
}
func (RawNode) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Bytes}
}
func (x RawNode) AsBytes() ([]byte, error) {
	return x.x, nil
}
func (RawNode) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "RawNode", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Bytes}
}

type RawNode struct{ x []byte }

// TODO generateKindBytes.EmitNativeAccessors
// TODO generateKindBytes.EmitNativeBuilder
type MaybeRawNode struct {
	Maybe schema.Maybe
	Value RawNode
}

func (m MaybeRawNode) Must() RawNode {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (RawNode) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (RawNode) Representation() ipld.Node {
	panic("TODO representation")
}
