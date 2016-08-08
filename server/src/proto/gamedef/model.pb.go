// Code generated by protoc-gen-go.
// source: model.proto
// DO NOT EDIT!

package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginModel struct {
	Account string `protobuf:"bytes,1,opt,name=Account,json=account" json:"Account,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=Address,json=address" json:"Address,omitempty"`
}

func (m *LoginModel) Reset()                    { *m = LoginModel{} }
func (m *LoginModel) String() string            { return proto.CompactTextString(m) }
func (*LoginModel) ProtoMessage()               {}
func (*LoginModel) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type ModelACK struct {
	Login      *LoginModel       `protobuf:"bytes,1,opt,name=Login,json=login" json:"Login,omitempty"`
	ServerList []*ServerInfo     `protobuf:"bytes,2,rep,name=ServerList,json=serverList" json:"ServerList,omitempty"`
	CharList   []*SimpleCharInfo `protobuf:"bytes,3,rep,name=CharList,json=charList" json:"CharList,omitempty"`
}

func (m *ModelACK) Reset()                    { *m = ModelACK{} }
func (m *ModelACK) String() string            { return proto.CompactTextString(m) }
func (*ModelACK) ProtoMessage()               {}
func (*ModelACK) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ModelACK) GetLogin() *LoginModel {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *ModelACK) GetServerList() []*ServerInfo {
	if m != nil {
		return m.ServerList
	}
	return nil
}

func (m *ModelACK) GetCharList() []*SimpleCharInfo {
	if m != nil {
		return m.CharList
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginModel)(nil), "gamedef.LoginModel")
	proto.RegisterType((*ModelACK)(nil), "gamedef.ModelACK")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4f, 0xcc, 0x4d, 0x4d, 0x49, 0x4d,
	0x93, 0xe2, 0xce, 0xc9, 0x4f, 0xcf, 0xcc, 0x83, 0x88, 0x4a, 0x71, 0x81, 0x44, 0x21, 0x6c, 0x25,
	0x07, 0x2e, 0x2e, 0x1f, 0x90, 0x94, 0x2f, 0x48, 0x97, 0x90, 0x04, 0x17, 0xbb, 0x63, 0x72, 0x72,
	0x7e, 0x69, 0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x7b, 0x22, 0x84, 0x0b, 0x96,
	0x49, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x82, 0xca, 0x40, 0xb8, 0x4a, 0xf3, 0x19, 0xb9,
	0x38, 0xc0, 0xba, 0x1d, 0x9d, 0xbd, 0x85, 0x34, 0xb9, 0x58, 0xc1, 0xc6, 0x81, 0xb5, 0x73, 0x1b,
	0x09, 0xeb, 0x41, 0x1d, 0xa0, 0x87, 0xb0, 0x24, 0x88, 0x15, 0xec, 0x16, 0x21, 0x63, 0x2e, 0xae,
	0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0x9f, 0xcc, 0xe2, 0x12, 0xa0, 0xa1, 0xcc, 0x28, 0xea, 0x21,
	0x52, 0x9e, 0x79, 0x69, 0xf9, 0x41, 0x5c, 0xc5, 0x70, 0x65, 0x40, 0x4d, 0x1c, 0xce, 0x19, 0x89,
	0x10, 0x2d, 0xcc, 0x60, 0x2d, 0xe2, 0x08, 0x2d, 0x99, 0xb9, 0x05, 0x39, 0xa9, 0x20, 0x69, 0xb0,
	0x36, 0x8e, 0x64, 0xa8, 0xc2, 0x24, 0x36, 0xb0, 0x57, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x81, 0x0c, 0xa7, 0x1b, 0x01, 0x00, 0x00,
}