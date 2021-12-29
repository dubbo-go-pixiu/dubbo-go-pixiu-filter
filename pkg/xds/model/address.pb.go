// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Address the address
type Address struct {
	SocketAddress        *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3" json:"socket_address,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetSocketAddress() *SocketAddress {
	if m != nil {
		return m.SocketAddress
	}
	return nil
}

func (m *Address) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// SocketAddress specify either a logical or physical address and port, which are
// used to tell server where to bind/listen, connect to upstream and find
// management servers
type SocketAddress struct {
	ProtocolStr          string   `protobuf:"bytes,1,opt,name=protocol_str,json=protocolStr,proto3" json:"protocol_str,omitempty"`
	Protocol             int32    `protobuf:"varint,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int64    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	ResolverName         string   `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	Domains              []string `protobuf:"bytes,6,rep,name=domains,proto3" json:"domains,omitempty"`
	CertsDir             string   `protobuf:"bytes,7,opt,name=certsDir,proto3" json:"certsDir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (m *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(m, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

func (m *SocketAddress) GetProtocolStr() string {
	if m != nil {
		return m.ProtocolStr
	}
	return ""
}

func (m *SocketAddress) GetProtocol() int32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SocketAddress) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *SocketAddress) GetCertsDir() string {
	if m != nil {
		return m.CertsDir
	}
	return ""
}

func init() {
	proto.RegisterType((*Address)(nil), "pixiu.api.v1.Address")
	proto.RegisterType((*SocketAddress)(nil), "pixiu.api.v1.SocketAddress")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0xd2, 0x36, 0xf4, 0x9a, 0x30, 0x78, 0xb2, 0x60, 0x09, 0x65, 0xc9, 0x64, 0x09,
	0x78, 0x02, 0x10, 0x33, 0x83, 0xfb, 0x00, 0x91, 0x49, 0x3c, 0x58, 0x34, 0x75, 0x74, 0x67, 0x2a,
	0xde, 0x94, 0xd7, 0x41, 0x3e, 0xe2, 0x2a, 0x6c, 0xf7, 0x9d, 0xfd, 0xfd, 0xf7, 0x43, 0x6d, 0x87,
	0x01, 0x1d, 0x91, 0x9e, 0x30, 0xc4, 0x20, 0xab, 0xc9, 0x7f, 0xfb, 0x2f, 0x6d, 0x27, 0xaf, 0xcf,
	0x8f, 0x7b, 0x0b, 0xe5, 0xcb, 0xdf, 0xb3, 0x7c, 0x85, 0x1b, 0x0a, 0xfd, 0xa7, 0x8b, 0xdd, 0x2c,
	0x28, 0xd1, 0x88, 0x76, 0xf7, 0x74, 0xa7, 0x97, 0x86, 0x3e, 0xf0, 0x9f, 0x59, 0x32, 0x35, 0x2d,
	0x51, 0x4a, 0x58, 0x9d, 0xec, 0xe8, 0xd4, 0x55, 0x23, 0xda, 0xad, 0xe1, 0x79, 0xff, 0x23, 0xa0,
	0xfe, 0x27, 0xc9, 0x7b, 0xa8, 0xb8, 0x4b, 0x1f, 0x8e, 0x1d, 0x45, 0xe4, 0x3b, 0x5b, 0xb3, 0xcb,
	0xbb, 0x43, 0x44, 0x79, 0x0b, 0xd7, 0x19, 0x39, 0x6c, 0x6d, 0x2e, 0x2c, 0x15, 0x94, 0xb9, 0x61,
	0xc1, 0x66, 0xc6, 0x74, 0x7e, 0x0a, 0x18, 0xd5, 0xaa, 0x11, 0x6d, 0x61, 0x78, 0x96, 0x0f, 0x50,
	0xa3, 0xa3, 0x70, 0x3c, 0x3b, 0xec, 0xb8, 0xdb, 0x9a, 0x9d, 0x2a, 0x2f, 0xdf, 0xed, 0xe8, 0x52,
	0xe4, 0x10, 0x46, 0xeb, 0x4f, 0xa4, 0x36, 0x4d, 0x91, 0x22, 0x67, 0x4c, 0x45, 0x7a, 0x87, 0x91,
	0xde, 0x3c, 0xaa, 0x92, 0xcd, 0x0b, 0x7f, 0x6c, 0xb8, 0xd2, 0xf3, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x09, 0x7d, 0xb6, 0x62, 0x01, 0x00, 0x00,
}
