// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as messages for now, in the future we may want to allow this to be specified by size in bytes
	Messages int32 `protobuf:"varint,1,opt,name=messages" json:"messages,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x85, 0x29, 0xfc, 0xbf, 0xda, 0x80, 0x4b, 0x5c, 0xa4, 0x2e, 0x52, 0x07, 0x1d, 0xa4, 0x19,
	0x7c, 0x83, 0xfa, 0x06, 0xd5, 0xc9, 0x2d, 0x4d, 0x6f, 0xd3, 0x80, 0xc9, 0x0d, 0x37, 0xe9, 0x50,
	0x9f, 0x5e, 0x48, 0x8b, 0xd3, 0xbd, 0x1f, 0xdf, 0x81, 0x73, 0xd8, 0x01, 0xa9, 0x03, 0x02, 0x12,
	0x0a, 0x5d, 0x6f, 0xf4, 0x48, 0x32, 0x1a, 0x74, 0x95, 0x27, 0x8c, 0xc8, 0xd7, 0x8b, 0x2c, 0x76,
	0x0a, 0xad, 0x45, 0x27, 0xe6, 0x33, 0xdb, 0xf2, 0xc4, 0xb6, 0x77, 0x74, 0x01, 0x5c, 0x18, 0xc3,
	0x73, 0xf2, 0xc0, 0x39, 0xfb, 0x8b, 0x93, 0x87, 0x7d, 0x76, 0xcc, 0x2e, 0x79, 0x93, 0xfe, 0xf2,
	0xcc, 0xf2, 0x5a, 0x46, 0x35, 0x3c, 0xcc, 0x07, 0x78, 0xc1, 0x36, 0x16, 0x42, 0x90, 0x1a, 0x42,
	0x0a, 0xfd, 0x37, 0x3f, 0xae, 0xab, 0xd7, 0x55, 0x9b, 0x38, 0x8c, 0x6d, 0xa5, 0xd0, 0x8a, 0x61,
	0xf2, 0x40, 0x6f, 0xe8, 0x34, 0x90, 0xe8, 0x65, 0x4b, 0x46, 0x89, 0x54, 0x1a, 0xc4, 0x32, 0xa9,
	0x5d, 0x25, 0xbe, 0x7d, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x10, 0x57, 0x86, 0xc1, 0x00, 0x00,
	0x00,
}
