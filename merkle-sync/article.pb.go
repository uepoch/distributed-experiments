// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	article.proto

It has these top-level messages:
	Article
	Category
	MerkleTree
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ArticleKind int32

const (
	ArticleKind_TOMATO ArticleKind = 0
	ArticleKind_POTATO ArticleKind = 1
	ArticleKind_DRUG   ArticleKind = 2
)

var ArticleKind_name = map[int32]string{
	0: "TOMATO",
	1: "POTATO",
	2: "DRUG",
}
var ArticleKind_value = map[string]int32{
	"TOMATO": 0,
	"POTATO": 1,
	"DRUG":   2,
}

func (x ArticleKind) String() string {
	return proto.EnumName(ArticleKind_name, int32(x))
}
func (ArticleKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Article struct {
	Name string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind ArticleKind `protobuf:"varint,2,opt,name=kind,enum=main.ArticleKind" json:"kind,omitempty"`
}

func (m *Article) Reset()                    { *m = Article{} }
func (m *Article) String() string            { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()               {}
func (*Article) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Article) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Article) GetKind() ArticleKind {
	if m != nil {
		return m.Kind
	}
	return ArticleKind_TOMATO
}

type Category struct {
	Articles []*Article  `protobuf:"bytes,1,rep,name=articles" json:"articles,omitempty"`
	Kind     ArticleKind `protobuf:"varint,2,opt,name=kind,enum=main.ArticleKind" json:"kind,omitempty"`
}

func (m *Category) Reset()                    { *m = Category{} }
func (m *Category) String() string            { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()               {}
func (*Category) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Category) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

func (m *Category) GetKind() ArticleKind {
	if m != nil {
		return m.Kind
	}
	return ArticleKind_TOMATO
}

type MerkleTree struct {
	Path   string                 `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
	Sum    string                 `protobuf:"bytes,2,opt,name=Sum" json:"Sum,omitempty"`
	Hasher int32                  `protobuf:"varint,3,opt,name=Hasher" json:"Hasher,omitempty"`
	Parent *MerkleTree            `protobuf:"bytes,4,opt,name=Parent" json:"Parent,omitempty"`
	Tree   map[string]*MerkleTree `protobuf:"bytes,5,rep,name=Tree" json:"Tree,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Leaf   bool                   `protobuf:"varint,6,opt,name=Leaf" json:"Leaf,omitempty"`
}

func (m *MerkleTree) Reset()                    { *m = MerkleTree{} }
func (m *MerkleTree) String() string            { return proto.CompactTextString(m) }
func (*MerkleTree) ProtoMessage()               {}
func (*MerkleTree) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MerkleTree) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MerkleTree) GetSum() string {
	if m != nil {
		return m.Sum
	}
	return ""
}

func (m *MerkleTree) GetHasher() int32 {
	if m != nil {
		return m.Hasher
	}
	return 0
}

func (m *MerkleTree) GetParent() *MerkleTree {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *MerkleTree) GetTree() map[string]*MerkleTree {
	if m != nil {
		return m.Tree
	}
	return nil
}

func (m *MerkleTree) GetLeaf() bool {
	if m != nil {
		return m.Leaf
	}
	return false
}

func init() {
	proto.RegisterType((*Article)(nil), "main.Article")
	proto.RegisterType((*Category)(nil), "main.Category")
	proto.RegisterType((*MerkleTree)(nil), "main.MerkleTree")
	proto.RegisterEnum("main.ArticleKind", ArticleKind_name, ArticleKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Synchronizer service

type SynchronizerClient interface {
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

// Server API for Synchronizer service

type SynchronizerServer interface {
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "article.proto",
}

func init() { proto.RegisterFile("article.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4f, 0xc2, 0x30,
	0x18, 0x87, 0xed, 0x28, 0x73, 0xbc, 0x08, 0x99, 0x3d, 0x98, 0x86, 0xd3, 0x42, 0xa2, 0xa9, 0x1e,
	0x66, 0x32, 0x2f, 0xc6, 0x1b, 0x11, 0xa3, 0x46, 0x09, 0xa4, 0xcc, 0x9b, 0x97, 0x0a, 0x55, 0x16,
	0xa0, 0x33, 0xa5, 0x98, 0xcc, 0xab, 0xff, 0xb8, 0x69, 0x37, 0xf0, 0x23, 0x1e, 0xbc, 0x2c, 0x4f,
	0xf7, 0x7b, 0x3f, 0x9e, 0xad, 0xd0, 0x12, 0xda, 0x64, 0x93, 0x85, 0x8c, 0x5f, 0x75, 0x6e, 0x72,
	0x82, 0x97, 0x22, 0x53, 0xdd, 0x3e, 0xec, 0xf6, 0xca, 0xd7, 0x84, 0x00, 0x56, 0x62, 0x29, 0x29,
	0x8a, 0x10, 0x6b, 0x70, 0xc7, 0xe4, 0x10, 0xf0, 0x3c, 0x53, 0x53, 0xea, 0x45, 0x88, 0xb5, 0x93,
	0xfd, 0xd8, 0xf6, 0xc4, 0x55, 0xc3, 0x5d, 0xa6, 0xa6, 0xdc, 0xc5, 0xdd, 0x47, 0x08, 0x2e, 0x85,
	0x91, 0x2f, 0xb9, 0x2e, 0xc8, 0x31, 0x04, 0xd5, 0xa2, 0x15, 0x45, 0x51, 0x8d, 0x35, 0x93, 0xd6,
	0x8f, 0x36, 0xbe, 0x8d, 0xff, 0x3b, 0xfd, 0xc3, 0x03, 0x18, 0x48, 0x3d, 0x5f, 0xc8, 0x54, 0x4b,
	0xe7, 0x39, 0x12, 0x66, 0xb6, 0xf1, 0xb4, 0x4c, 0x42, 0xa8, 0x8d, 0xd7, 0x4b, 0x37, 0xa8, 0xc1,
	0x2d, 0x92, 0x03, 0xf0, 0x6f, 0xc4, 0x6a, 0x26, 0x35, 0xad, 0x45, 0x88, 0xd5, 0x79, 0x75, 0x22,
	0x0c, 0xfc, 0x91, 0xd0, 0x52, 0x19, 0x8a, 0x23, 0xc4, 0x9a, 0x49, 0x58, 0x6e, 0xfd, 0x9a, 0xcf,
	0xab, 0x9c, 0xc4, 0x80, 0xed, 0x99, 0xd6, 0xdd, 0x47, 0x74, 0x7e, 0xd7, 0xc5, 0xf6, 0x71, 0xa5,
	0x8c, 0x2e, 0x38, 0xde, 0x78, 0xdd, 0x4b, 0xf1, 0x4c, 0xfd, 0x08, 0xb1, 0x80, 0x3b, 0xee, 0xdc,
	0x42, 0x63, 0x5b, 0x66, 0x25, 0xe7, 0xb2, 0xa8, 0xbc, 0x2d, 0x92, 0x23, 0xa8, 0xbf, 0x89, 0xc5,
	0x5a, 0x3a, 0xf1, 0xbf, 0x5c, 0xca, 0xf8, 0xc2, 0x3b, 0x47, 0x27, 0xa7, 0xd0, 0xfc, 0xf6, 0x6b,
	0x08, 0x80, 0x9f, 0x0e, 0x07, 0xbd, 0x74, 0x18, 0xee, 0x58, 0x1e, 0x0d, 0x53, 0xcb, 0x88, 0x04,
	0x80, 0xfb, 0xfc, 0xe1, 0x3a, 0xf4, 0x92, 0x36, 0xec, 0x8d, 0x0b, 0x35, 0x99, 0xe9, 0x5c, 0x65,
	0xef, 0x52, 0x3f, 0xf9, 0xee, 0xde, 0xcf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x53, 0x64,
	0x82, 0x08, 0x02, 0x00, 0x00,
}
