// Code generated by protoc-gen-go.
// source: game.proto
// DO NOT EDIT!

package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// client->game
type EnterGameREQ struct {
}

func (m *EnterGameREQ) Reset()                    { *m = EnterGameREQ{} }
func (m *EnterGameREQ) String() string            { return proto.CompactTextString(m) }
func (*EnterGameREQ) ProtoMessage()               {}
func (*EnterGameREQ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// game->client
type EnterGameACK struct {
}

func (m *EnterGameACK) Reset()                    { *m = EnterGameACK{} }
func (m *EnterGameACK) String() string            { return proto.CompactTextString(m) }
func (*EnterGameACK) ProtoMessage()               {}
func (*EnterGameACK) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*EnterGameREQ)(nil), "gamedef.EnterGameREQ")
	proto.RegisterType((*EnterGameACK)(nil), "gamedef.EnterGameACK")
}

var fileDescriptor2 = []byte{
	// 67 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0xb1, 0x53, 0x52, 0xd3, 0x94, 0xf8, 0xb8,
	0x78, 0x5c, 0xf3, 0x4a, 0x52, 0x8b, 0xdc, 0x81, 0xfc, 0x20, 0xd7, 0x40, 0x14, 0xbe, 0xa3, 0xb3,
	0x77, 0x12, 0x1b, 0x58, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x17, 0xad, 0xb0, 0x3d,
	0x00, 0x00, 0x00,
}