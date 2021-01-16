// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package graphsync_message_pb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	// the actual data included in this message
	CompleteRequestList bool               `protobuf:"varint,1,opt,name=completeRequestList,proto3" json:"completeRequestList,omitempty"`
	Requests            []Message_Request  `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests"`
	Responses           []Message_Response `protobuf:"bytes,3,rep,name=responses,proto3" json:"responses"`
	Data                []Message_Block    `protobuf:"bytes,4,rep,name=data,proto3" json:"data"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetCompleteRequestList() bool {
	if m != nil {
		return m.CompleteRequestList
	}
	return false
}

func (m *Message) GetRequests() []Message_Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *Message) GetResponses() []Message_Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

func (m *Message) GetData() []Message_Block {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message_Request struct {
	Id         int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Root       []byte            `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Selector   []byte            `protobuf:"bytes,3,opt,name=selector,proto3" json:"selector,omitempty"`
	Extensions map[string][]byte `protobuf:"bytes,4,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Priority   int32             `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Cancel     bool              `protobuf:"varint,6,opt,name=cancel,proto3" json:"cancel,omitempty"`
	Update     bool              `protobuf:"varint,7,opt,name=update,proto3" json:"update,omitempty"`
}

func (m *Message_Request) Reset()         { *m = Message_Request{} }
func (m *Message_Request) String() string { return proto.CompactTextString(m) }
func (*Message_Request) ProtoMessage()    {}
func (*Message_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}
func (m *Message_Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Request.Merge(m, src)
}
func (m *Message_Request) XXX_Size() int {
	return m.Size()
}
func (m *Message_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Request proto.InternalMessageInfo

func (m *Message_Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message_Request) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Message_Request) GetSelector() []byte {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *Message_Request) GetExtensions() map[string][]byte {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *Message_Request) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Message_Request) GetCancel() bool {
	if m != nil {
		return m.Cancel
	}
	return false
}

func (m *Message_Request) GetUpdate() bool {
	if m != nil {
		return m.Update
	}
	return false
}

type Message_Response struct {
	Id         int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status     int32             `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Extensions map[string][]byte `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Message_Response) Reset()         { *m = Message_Response{} }
func (m *Message_Response) String() string { return proto.CompactTextString(m) }
func (*Message_Response) ProtoMessage()    {}
func (*Message_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}
func (m *Message_Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Response.Merge(m, src)
}
func (m *Message_Response) XXX_Size() int {
	return m.Size()
}
func (m *Message_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Response proto.InternalMessageInfo

func (m *Message_Response) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message_Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Message_Response) GetExtensions() map[string][]byte {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type Message_Block struct {
	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Message_Block) Reset()         { *m = Message_Block{} }
func (m *Message_Block) String() string { return proto.CompactTextString(m) }
func (*Message_Block) ProtoMessage()    {}
func (*Message_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 2}
}
func (m *Message_Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Block.Merge(m, src)
}
func (m *Message_Block) XXX_Size() int {
	return m.Size()
}
func (m *Message_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Block proto.InternalMessageInfo

func (m *Message_Block) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Message_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "graphsync.message.pb.Message")
	proto.RegisterType((*Message_Request)(nil), "graphsync.message.pb.Message.Request")
	proto.RegisterMapType((map[string][]byte)(nil), "graphsync.message.pb.Message.Request.ExtensionsEntry")
	proto.RegisterType((*Message_Response)(nil), "graphsync.message.pb.Message.Response")
	proto.RegisterMapType((map[string][]byte)(nil), "graphsync.message.pb.Message.Response.ExtensionsEntry")
	proto.RegisterType((*Message_Block)(nil), "graphsync.message.pb.Message.Block")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x9d, 0xd8, 0x49, 0x2f, 0xe5, 0xa1, 0xa1, 0xaa, 0x46, 0x5e, 0x98, 0x08, 0x04, 0xca,
	0x06, 0x17, 0x51, 0x81, 0x10, 0x52, 0x37, 0x91, 0x2a, 0x24, 0x04, 0x9b, 0x91, 0x60, 0xef, 0x38,
	0xb7, 0xae, 0x55, 0xc7, 0x63, 0x66, 0xc6, 0xa8, 0xfe, 0x0b, 0xfe, 0x83, 0x7f, 0x60, 0x5d, 0x76,
	0x5d, 0xb2, 0x42, 0x28, 0xf9, 0x11, 0xe4, 0x3b, 0x83, 0x79, 0x55, 0xa5, 0x12, 0xbb, 0x7b, 0xee,
	0xcc, 0x39, 0xf7, 0x71, 0x66, 0xe0, 0xfa, 0x0a, 0xb5, 0x4e, 0x73, 0x4c, 0x6a, 0x25, 0x8d, 0x64,
	0x3b, 0xb9, 0x4a, 0xeb, 0x63, 0xdd, 0x56, 0x59, 0xd2, 0x1f, 0x2c, 0xa2, 0x87, 0x79, 0x61, 0x8e,
	0x9b, 0x45, 0x92, 0xc9, 0xd5, 0x5e, 0x2e, 0x73, 0xb9, 0x47, 0x97, 0x17, 0xcd, 0x11, 0x21, 0x02,
	0x14, 0x59, 0x91, 0xbb, 0x9f, 0x42, 0x18, 0xbf, 0xb6, 0x6c, 0xf6, 0x08, 0x6e, 0x67, 0x72, 0x55,
	0x97, 0x68, 0x50, 0xe0, 0xbb, 0x06, 0xb5, 0x79, 0x55, 0x68, 0xc3, 0xbd, 0xa9, 0x37, 0x9b, 0x88,
	0x8b, 0x8e, 0xd8, 0x0b, 0x98, 0x28, 0x0b, 0x35, 0xf7, 0xa7, 0xc3, 0xd9, 0xb5, 0xc7, 0xf7, 0x93,
	0x8b, 0xba, 0x4a, 0x5c, 0x89, 0xc4, 0x91, 0xe7, 0xa3, 0xb3, 0xaf, 0x77, 0x06, 0xa2, 0x27, 0xb3,
	0x97, 0xb0, 0xa5, 0x50, 0xd7, 0xb2, 0xd2, 0xa8, 0xf9, 0x90, 0x94, 0x1e, 0xfc, 0x4b, 0xc9, 0x5e,
	0x77, 0x52, 0x3f, 0xe9, 0xec, 0x00, 0x46, 0xcb, 0xd4, 0xa4, 0x7c, 0x44, 0x32, 0xf7, 0x2e, 0x97,
	0x99, 0x97, 0x32, 0x3b, 0x71, 0x1a, 0x44, 0x8b, 0x3e, 0xfa, 0x30, 0x76, 0x6d, 0xb2, 0x1b, 0xe0,
	0x17, 0x4b, 0x5a, 0x40, 0x20, 0xfc, 0x62, 0xc9, 0x18, 0x8c, 0x94, 0x94, 0x86, 0xfb, 0x53, 0x6f,
	0xb6, 0x2d, 0x28, 0x66, 0x11, 0x4c, 0x34, 0x96, 0x98, 0x19, 0xa9, 0xf8, 0x90, 0xf2, 0x3d, 0x66,
	0x6f, 0x00, 0xf0, 0xd4, 0x60, 0xa5, 0x0b, 0x59, 0x69, 0xd7, 0xd0, 0x93, 0x2b, 0x6d, 0x28, 0x39,
	0xec, 0x79, 0x87, 0x95, 0x51, 0xad, 0xf8, 0x45, 0xa8, 0x2b, 0x59, 0xab, 0x42, 0xaa, 0xc2, 0xb4,
	0x3c, 0xa0, 0xe6, 0x7a, 0xcc, 0x76, 0x21, 0xcc, 0xd2, 0x2a, 0xc3, 0x92, 0x87, 0xe4, 0x9b, 0x43,
	0x5d, 0xbe, 0xa9, 0x97, 0xa9, 0x41, 0x3e, 0xb6, 0x79, 0x8b, 0xa2, 0x03, 0xb8, 0xf9, 0x47, 0x29,
	0x76, 0x0b, 0x86, 0x27, 0xd8, 0xd2, 0xd8, 0x5b, 0xa2, 0x0b, 0xd9, 0x0e, 0x04, 0xef, 0xd3, 0xb2,
	0x41, 0x37, 0xb8, 0x05, 0xcf, 0xfd, 0x67, 0x5e, 0xf4, 0xd9, 0x83, 0xc9, 0x0f, 0x2b, 0xfe, 0x5a,
	0xd7, 0x2e, 0x84, 0xda, 0xa4, 0xa6, 0xd1, 0xc4, 0x0b, 0x84, 0x43, 0xec, 0xed, 0x6f, 0x6b, 0xb1,
	0x76, 0x3f, 0xbd, 0x9a, 0xdd, 0x97, 0xed, 0xe5, 0x7f, 0x67, 0xd9, 0x87, 0x80, 0x9e, 0x43, 0xd7,
	0x77, 0xad, 0xf0, 0xa8, 0x38, 0x25, 0xde, 0xb6, 0x70, 0xa8, 0xb3, 0x9f, 0x5e, 0x96, 0xb3, 0xbf,
	0x8b, 0xe7, 0xfc, 0x6c, 0x1d, 0x7b, 0xe7, 0xeb, 0xd8, 0xfb, 0xb6, 0x8e, 0xbd, 0x0f, 0x9b, 0x78,
	0x70, 0xbe, 0x89, 0x07, 0x5f, 0x36, 0xf1, 0x60, 0x11, 0xd2, 0x0f, 0xdb, 0xff, 0x1e, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x5e, 0x01, 0x7e, 0xb7, 0x03, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Responses) > 0 {
		for iNdEx := len(m.Responses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Responses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CompleteRequestList {
		i--
		if m.CompleteRequestList {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Message_Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Update {
		i--
		if m.Update {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Cancel {
		i--
		if m.Cancel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Priority != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Extensions) > 0 {
		for k := range m.Extensions {
			v := m.Extensions[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintMessage(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMessage(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMessage(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Selector) > 0 {
		i -= len(m.Selector)
		copy(dAtA[i:], m.Selector)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Selector)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Message_Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extensions) > 0 {
		for k := range m.Extensions {
			v := m.Extensions[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintMessage(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMessage(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMessage(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Status != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Message_Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message_Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CompleteRequestList {
		n += 2
	}
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	return n
}

func (m *Message_Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMessage(uint64(m.Id))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Selector)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Extensions) > 0 {
		for k, v := range m.Extensions {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovMessage(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovMessage(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovMessage(uint64(mapEntrySize))
		}
	}
	if m.Priority != 0 {
		n += 1 + sovMessage(uint64(m.Priority))
	}
	if m.Cancel {
		n += 2
	}
	if m.Update {
		n += 2
	}
	return n
}

func (m *Message_Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMessage(uint64(m.Id))
	}
	if m.Status != 0 {
		n += 1 + sovMessage(uint64(m.Status))
	}
	if len(m.Extensions) > 0 {
		for k, v := range m.Extensions {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovMessage(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovMessage(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovMessage(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Message_Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompleteRequestList", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CompleteRequestList = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, Message_Request{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, Message_Response{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, Message_Block{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Selector = append(m.Selector[:0], dAtA[iNdEx:postIndex]...)
			if m.Selector == nil {
				m.Selector = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Extensions == nil {
				m.Extensions = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMessage
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMessage
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthMessage
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthMessage
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMessage(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMessage
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Extensions[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Cancel = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Update", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Update = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Extensions == nil {
				m.Extensions = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMessage
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMessage
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthMessage
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthMessage
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMessage(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMessage
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Extensions[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message_Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append(m.Prefix[:0], dAtA[iNdEx:postIndex]...)
			if m.Prefix == nil {
				m.Prefix = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)