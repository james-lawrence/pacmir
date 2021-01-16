package dagpb

import (
	ipld "github.com/ipld/go-ipld-prime"
)

var Style style

type style struct {
	Protobuf _PBNode__NodeStyle
	Raw      _RawNode__NodeStyle
}

type _PBNode__NodeStyle struct {
}

func (ns _PBNode__NodeStyle) NewBuilder() ipld.NodeBuilder {
	var nd PBNode
	return &_PBNode__NodeBuilder{_PBNode__NodeAssembler{nd: &nd}}
}

type _PBNode__NodeBuilder struct {
	_PBNode__NodeAssembler
}

func (nb *_PBNode__NodeBuilder) Build() ipld.Node {
	return nb.nd
}

func (nb *_PBNode__NodeBuilder) Reset() {
	var nd PBNode
	*nb = _PBNode__NodeBuilder{_PBNode__NodeAssembler{nd: &nd}}
}

type _PBNode__NodeAssembler struct {
	nd *PBNode
}

func (na *_PBNode__NodeAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignNull() error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignBool(_ bool) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignInt(_ int) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignFloat(_ float64) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignString(_ string) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignBytes(_ []byte) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignLink(_ ipld.Link) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) AssignNode(_ ipld.Node) error {
	panic("not implemented")
}

func (na *_PBNode__NodeAssembler) Style() ipld.NodeStyle {
	return _PBNode__NodeStyle{}
}

func (nd PBNode) Style() ipld.NodeStyle {
	return _PBNode__NodeStyle{}
}

func (nd _PBNode__Repr) Style() ipld.NodeStyle {
	return nil
}

func (nd PBLinks) Style() ipld.NodeStyle {
	return nil
}

func (nd PBLink) Style() ipld.NodeStyle {
	return nil
}

func (nd _PBLink__Repr) Style() ipld.NodeStyle {
	return nil
}

func (nb Link) Style() ipld.NodeStyle {
	return nil
}

func (nb Bytes) Style() ipld.NodeStyle {
	return nil
}

func (nb Int) Style() ipld.NodeStyle {
	return nil
}

func (nb String) Style() ipld.NodeStyle {
	return nil
}

type _RawNode__NodeStyle struct {
}

func (ns _RawNode__NodeStyle) NewBuilder() ipld.NodeBuilder {
	var nd RawNode
	return &_RawNode__NodeBuilder{_RawNode__NodeAssembler{nd: &nd}}
}

type _RawNode__NodeBuilder struct {
	_RawNode__NodeAssembler
}

func (nb *_RawNode__NodeBuilder) Build() ipld.Node {
	return nb.nd
}

func (nb *_RawNode__NodeBuilder) Reset() {
	var nd RawNode
	*nb = _RawNode__NodeBuilder{_RawNode__NodeAssembler{nd: &nd}}
}

type _RawNode__NodeAssembler struct {
	nd *RawNode
}

func (na *_RawNode__NodeAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignNull() error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignBool(_ bool) error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignInt(_ int) error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignFloat(_ float64) error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignString(_ string) error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignBytes(x []byte) error {
	na.nd.x = x
	return nil
}

func (na *_RawNode__NodeAssembler) AssignLink(_ ipld.Link) error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) AssignNode(_ ipld.Node) error {
	panic("not implemented")
}

func (na *_RawNode__NodeAssembler) Style() ipld.NodeStyle {
	return _RawNode__NodeStyle{}
}

func (nd RawNode) Style() ipld.NodeStyle {
	return _RawNode__NodeStyle{}
}
