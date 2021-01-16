package dagpb

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
)

// Code generated go-ipld-prime DO NOT EDIT.

var _ ipld.Node = String{}
var _ schema.TypedNode = String{}

func (String) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_String
}
func (String) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (String) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (String) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}
}
func (String) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_String}
}
func (String) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "String", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}}
}
func (String) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "String", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}}
}
func (String) Length() int {
	return -1
}
func (String) IsUndefined() bool {
	return false
}
func (String) IsNull() bool {
	return false
}
func (String) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_String}
}
func (String) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_String}
}
func (String) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_String}
}
func (x String) AsString() (string, error) {
	return x.x, nil
}
func (String) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_String}
}
func (String) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_String}
}

type String struct{ x string }

func (x String) String() string {
	return x.x
}

type String__Content struct {
	Value string
}

func (b String__Content) Build() (String, error) {
	x := String{
		b.Value,
	}
	// FUTURE : want to support customizable validation.
	//   but 'if v, ok := x.(schema.Validatable); ok {' doesn't fly: need a way to work on concrete types.
	return x, nil
}
func (b String__Content) MustBuild() String {
	if x, err := b.Build(); err != nil {
		panic(err)
	} else {
		return x
	}
}

type MaybeString struct {
	Maybe schema.Maybe
	Value String
}

func (m MaybeString) Must() String {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (String) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (String) Representation() ipld.Node {
	panic("TODO representation")
}

var _ ipld.Node = Int{}
var _ schema.TypedNode = Int{}

func (Int) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Int
}
func (Int) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}
}
func (Int) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}
}
func (Int) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Int}
}
func (Int) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Int}
}
func (Int) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "Int", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}}
}
func (Int) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "Int", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Int}}
}
func (Int) Length() int {
	return -1
}
func (Int) IsUndefined() bool {
	return false
}
func (Int) IsNull() bool {
	return false
}
func (Int) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Int}
}
func (x Int) AsInt() (int, error) {
	return x.x, nil
}
func (Int) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Int}
}
func (Int) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Int}
}
func (Int) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Int}
}
func (Int) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Int}
}

type Int struct{ x int }

func (x Int) Int() int {
	return x.x
}

// TODO generateKindInt.EmitNativeBuilder
type MaybeInt struct {
	Maybe schema.Maybe
	Value Int
}

func (m MaybeInt) Must() Int {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (Int) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Int) Representation() ipld.Node {
	panic("TODO representation")
}

var _ ipld.Node = Bytes{}
var _ schema.TypedNode = Bytes{}

func (Bytes) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Bytes
}
func (Bytes) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}}
}
func (Bytes) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}}
}
func (Bytes) Length() int {
	return -1
}
func (Bytes) IsUndefined() bool {
	return false
}
func (Bytes) IsNull() bool {
	return false
}
func (Bytes) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Bytes}
}
func (x Bytes) AsBytes() ([]byte, error) {
	return x.x, nil
}
func (Bytes) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Bytes}
}

type Bytes struct{ x []byte }

// TODO generateKindBytes.EmitNativeAccessors
// TODO generateKindBytes.EmitNativeBuilder
type MaybeBytes struct {
	Maybe schema.Maybe
	Value Bytes
}

func (m MaybeBytes) Must() Bytes {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (Bytes) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Bytes) Representation() ipld.Node {
	panic("TODO representation")
}

var _ ipld.Node = Link{}
var _ schema.TypedNode = Link{}

func (Link) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Link
}
func (Link) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}
}
func (Link) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}
}
func (Link) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Link}
}
func (Link) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Link}
}
func (Link) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "Link", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}}
}
func (Link) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "Link", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Link}}
}
func (Link) Length() int {
	return -1
}
func (Link) IsUndefined() bool {
	return false
}
func (Link) IsNull() bool {
	return false
}
func (Link) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Link}
}
func (x Link) AsLink() (ipld.Link, error) {
	return x.x, nil
}

type Link struct{ x ipld.Link }

// TODO generateKindLink.EmitNativeAccessors
// TODO generateKindLink.EmitNativeBuilder
type MaybeLink struct {
	Maybe schema.Maybe
	Value Link
}

func (m MaybeLink) Must() Link {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (Link) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Link) Representation() ipld.Node {
	panic("TODO representation")
}

var _ ipld.Node = PBLink{}
var _ schema.TypedNode = PBLink{}

func (PBLink) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (x PBLink) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Hash":
		if x.d.Hash.Maybe == schema.Maybe_Absent {
			return ipld.Undef, nil
		}
		return x.d.Hash.Value, nil
	case "Name":
		if x.d.Name.Maybe == schema.Maybe_Absent {
			return ipld.Undef, nil
		}
		return x.d.Name.Value, nil
	case "Tsize":
		if x.d.Tsize.Maybe == schema.Maybe_Absent {
			return ipld.Undef, nil
		}
		return x.d.Tsize.Value, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (x PBLink) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return x.LookupString(ks)
}
func (PBLink) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n PBLink) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (x PBLink) MapIterator() ipld.MapIterator {
	return &_PBLink__Itr{&x, 0}
}

type _PBLink__Itr struct {
	node *PBLink
	idx  int
}

func (itr *_PBLink__Itr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = String{"Hash"}
		if itr.node.d.Hash.Maybe == schema.Maybe_Absent {
			v = ipld.Undef
			break
		}
		v = itr.node.d.Hash.Value
	case 1:
		k = String{"Name"}
		if itr.node.d.Name.Maybe == schema.Maybe_Absent {
			v = ipld.Undef
			break
		}
		v = itr.node.d.Name.Value
	case 2:
		k = String{"Tsize"}
		if itr.node.d.Tsize.Maybe == schema.Maybe_Absent {
			v = ipld.Undef
			break
		}
		v = itr.node.d.Tsize.Value
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_PBLink__Itr) Done() bool {
	return itr.idx >= 3
}

func (PBLink) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (PBLink) Length() int {
	return 3
}
func (PBLink) IsUndefined() bool {
	return false
}
func (PBLink) IsNull() bool {
	return false
}
func (PBLink) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}

type PBLink struct {
	d PBLink__Content
}

func (x PBLink) FieldHash() MaybeLink {
	return x.d.Hash
}
func (x PBLink) FieldName() MaybeString {
	return x.d.Name
}
func (x PBLink) FieldTsize() MaybeInt {
	return x.d.Tsize
}

type PBLink__Content struct {
	Hash  MaybeLink
	Name  MaybeString
	Tsize MaybeInt
}

func (b PBLink__Content) Build() (PBLink, error) {
	x := PBLink{b}
	// FUTURE : want to support customizable validation.
	//   but 'if v, ok := x.(schema.Validatable); ok {' doesn't fly: need a way to work on concrete types.
	return x, nil
}
func (b PBLink__Content) MustBuild() PBLink {
	if x, err := b.Build(); err != nil {
		panic(err)
	} else {
		return x
	}
}

type MaybePBLink struct {
	Maybe schema.Maybe
	Value PBLink
}

func (m MaybePBLink) Must() PBLink {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (PBLink) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n PBLink) Representation() ipld.Node {
	return _PBLink__Repr{&n}
}

var _ ipld.Node = _PBLink__Repr{}

type _PBLink__Repr struct {
	n *PBLink
}

func (_PBLink__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (rn _PBLink__Repr) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Hash":
		if rn.n.d.Hash.Maybe == schema.Maybe_Absent {
			return ipld.Undef, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return rn.n.d.Hash.Value, nil
	case "Name":
		if rn.n.d.Name.Maybe == schema.Maybe_Absent {
			return ipld.Undef, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return rn.n.d.Name.Value, nil
	case "Tsize":
		if rn.n.d.Tsize.Maybe == schema.Maybe_Absent {
			return ipld.Undef, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return rn.n.d.Tsize.Value, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (rn _PBLink__Repr) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return rn.LookupString(ks)
}
func (_PBLink__Repr) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n _PBLink__Repr) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (rn _PBLink__Repr) MapIterator() ipld.MapIterator {
	return &_PBLink__ReprItr{rn.n, 0}
}

type _PBLink__ReprItr struct {
	node *PBLink
	idx  int
}

func (itr *_PBLink__ReprItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	for {
		switch itr.idx {
		case 0:
			k = String{"Hash"}
			if itr.node.d.Hash.Maybe == schema.Maybe_Absent {
				itr.idx++
				continue
			}
			v = itr.node.d.Hash.Value
		case 1:
			k = String{"Name"}
			if itr.node.d.Name.Maybe == schema.Maybe_Absent {
				itr.idx++
				continue
			}
			v = itr.node.d.Name.Value
		case 2:
			k = String{"Tsize"}
			if itr.node.d.Tsize.Maybe == schema.Maybe_Absent {
				itr.idx++
				continue
			}
			v = itr.node.d.Tsize.Value
		default:
			panic("unreachable")
		}
	}
	itr.idx++
	return
}
func (itr *_PBLink__ReprItr) Done() bool {
	return itr.idx >= 3
}

func (_PBLink__Repr) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (rn _PBLink__Repr) Length() int {
	l := 3
	if rn.n.d.Hash.Maybe == schema.Maybe_Absent {
		l--
	}
	if rn.n.d.Name.Maybe == schema.Maybe_Absent {
		l--
	}
	if rn.n.d.Tsize.Maybe == schema.Maybe_Absent {
		l--
	}
	return l
}
func (_PBLink__Repr) IsUndefined() bool {
	return false
}
func (_PBLink__Repr) IsNull() bool {
	return false
}
func (_PBLink__Repr) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}

var _ ipld.Node = PBLinks{}
var _ schema.TypedNode = PBLinks{}

func (PBLinks) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (PBLinks) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_List}
}
func (x PBLinks) Lookup(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return x.LookupIndex(ki)
}
func (x PBLinks) LookupIndex(index int) (ipld.Node, error) {
	if index >= len(x.x) {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(index)}
	}
	return x.x[index], nil
}
func (n PBLinks) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	idx, err := seg.Index()
	if err != nil {
		return nil, err
	}
	return n.LookupIndex(idx)
}
func (PBLinks) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_List}}
}
func (x PBLinks) ListIterator() ipld.ListIterator {
	return &_PBLinks__Itr{&x, 0}
}

type _PBLinks__Itr struct {
	node *PBLinks
	idx  int
}

func (itr *_PBLinks__Itr) Next() (idx int, value ipld.Node, _ error) {
	if itr.idx >= len(itr.node.x) {
		return 0, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	value = itr.node.x[idx]
	itr.idx++
	return
}

func (itr *_PBLinks__Itr) Done() bool {
	return itr.idx >= len(itr.node.x)
}

func (x PBLinks) Length() int {
	return len(x.x)
}
func (PBLinks) IsUndefined() bool {
	return false
}
func (PBLinks) IsNull() bool {
	return false
}
func (PBLinks) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_List}
}

type PBLinks struct {
	x []PBLink
}

// TODO generateKindList.EmitNativeAccessors
// TODO generateKindList.EmitNativeBuilder
type MaybePBLinks struct {
	Maybe schema.Maybe
	Value PBLinks
}

func (m MaybePBLinks) Must() PBLinks {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (PBLinks) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n PBLinks) Representation() ipld.Node {
	panic("TODO representation")
}

var _ ipld.Node = PBNode{}
var _ schema.TypedNode = PBNode{}

func (PBNode) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (x PBNode) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Links":
		return x.d.Links, nil
	case "Data":
		return x.d.Data, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (x PBNode) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return x.LookupString(ks)
}
func (PBNode) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n PBNode) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (x PBNode) MapIterator() ipld.MapIterator {
	return &_PBNode__Itr{&x, 0}
}

type _PBNode__Itr struct {
	node *PBNode
	idx  int
}

func (itr *_PBNode__Itr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = String{"Links"}
		v = itr.node.d.Links
	case 1:
		k = String{"Data"}
		v = itr.node.d.Data
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_PBNode__Itr) Done() bool {
	return itr.idx >= 2
}

func (PBNode) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (PBNode) Length() int {
	return 2
}
func (PBNode) IsUndefined() bool {
	return false
}
func (PBNode) IsNull() bool {
	return false
}
func (PBNode) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}

type PBNode struct {
	d PBNode__Content
}

func (x PBNode) FieldLinks() PBLinks {
	return x.d.Links
}
func (x PBNode) FieldData() Bytes {
	return x.d.Data
}

type PBNode__Content struct {
	Links PBLinks
	Data  Bytes
}

func (b PBNode__Content) Build() (PBNode, error) {
	x := PBNode{b}
	// FUTURE : want to support customizable validation.
	//   but 'if v, ok := x.(schema.Validatable); ok {' doesn't fly: need a way to work on concrete types.
	return x, nil
}
func (b PBNode__Content) MustBuild() PBNode {
	if x, err := b.Build(); err != nil {
		panic(err)
	} else {
		return x
	}
}

type MaybePBNode struct {
	Maybe schema.Maybe
	Value PBNode
}

func (m MaybePBNode) Must() PBNode {
	if m.Maybe != schema.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

func (PBNode) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n PBNode) Representation() ipld.Node {
	return _PBNode__Repr{&n}
}

var _ ipld.Node = _PBNode__Repr{}

type _PBNode__Repr struct {
	n *PBNode
}

func (_PBNode__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (rn _PBNode__Repr) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Links":
		return rn.n.d.Links, nil
	case "Data":
		return rn.n.d.Data, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (rn _PBNode__Repr) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return rn.LookupString(ks)
}
func (_PBNode__Repr) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n _PBNode__Repr) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (rn _PBNode__Repr) MapIterator() ipld.MapIterator {
	return &_PBNode__ReprItr{rn.n, 0}
}

type _PBNode__ReprItr struct {
	node *PBNode
	idx  int
}

func (itr *_PBNode__ReprItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	for {
		switch itr.idx {
		case 0:
			k = String{"Links"}
			v = itr.node.d.Links
		case 1:
			k = String{"Data"}
			v = itr.node.d.Data
		default:
			panic("unreachable")
		}
	}
	itr.idx++
	return
}
func (itr *_PBNode__ReprItr) Done() bool {
	return itr.idx >= 2
}

func (_PBNode__Repr) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (rn _PBNode__Repr) Length() int {
	l := 2
	return l
}
func (_PBNode__Repr) IsUndefined() bool {
	return false
}
func (_PBNode__Repr) IsNull() bool {
	return false
}
func (_PBNode__Repr) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
