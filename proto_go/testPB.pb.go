// Code generated by protoc-gen-go.
// source: testPB.proto
// DO NOT EDIT!

/*
Package testPB is a generated protocol buffer package.

It is generated from these files:
	testPB.proto

It has these top-level messages:
	RequestMatch
	MatchResult
	Player
*/
package testPB

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgId int32

const (
	MsgId_REQUEST_MATCH MsgId = 1
	MsgId_MATCH_RESULT  MsgId = 2
)

var MsgId_name = map[int32]string{
	1: "REQUEST_MATCH",
	2: "MATCH_RESULT",
}
var MsgId_value = map[string]int32{
	"REQUEST_MATCH": 1,
	"MATCH_RESULT":  2,
}

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}
func (x MsgId) String() string {
	return proto.EnumName(MsgId_name, int32(x))
}
func (x *MsgId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgId_value, data, "MsgId")
	if err != nil {
		return err
	}
	*x = MsgId(value)
	return nil
}
func (MsgId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 请求匹配
type RequestMatch struct {
	UId              *int32 `protobuf:"varint,1,req,name=UId" json:"UId,omitempty"`
	Score            *int32 `protobuf:"varint,2,req,name=Score" json:"Score,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestMatch) Reset()                    { *m = RequestMatch{} }
func (m *RequestMatch) String() string            { return proto.CompactTextString(m) }
func (*RequestMatch) ProtoMessage()               {}
func (*RequestMatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestMatch) GetUId() int32 {
	if m != nil && m.UId != nil {
		return *m.UId
	}
	return 0
}

func (m *RequestMatch) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// 比赛结果下发
type MatchResult struct {
	IsTeam1Win       *bool     `protobuf:"varint,1,req,name=IsTeam1Win" json:"IsTeam1Win,omitempty"`
	Team1            []*Player `protobuf:"bytes,2,rep,name=Team1" json:"Team1,omitempty"`
	Team2            []*Player `protobuf:"bytes,3,rep,name=Team2" json:"Team2,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MatchResult) Reset()                    { *m = MatchResult{} }
func (m *MatchResult) String() string            { return proto.CompactTextString(m) }
func (*MatchResult) ProtoMessage()               {}
func (*MatchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MatchResult) GetIsTeam1Win() bool {
	if m != nil && m.IsTeam1Win != nil {
		return *m.IsTeam1Win
	}
	return false
}

func (m *MatchResult) GetTeam1() []*Player {
	if m != nil {
		return m.Team1
	}
	return nil
}

func (m *MatchResult) GetTeam2() []*Player {
	if m != nil {
		return m.Team2
	}
	return nil
}

// 玩家
type Player struct {
	UId              *int32 `protobuf:"varint,1,req,name=UId" json:"UId,omitempty"`
	Score            *int32 `protobuf:"varint,2,req,name=Score" json:"Score,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Player) GetUId() int32 {
	if m != nil && m.UId != nil {
		return *m.UId
	}
	return 0
}

func (m *Player) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestMatch)(nil), "proto_go.RequestMatch")
	proto.RegisterType((*MatchResult)(nil), "proto_go.MatchResult")
	proto.RegisterType((*Player)(nil), "proto_go.Player")
	proto.RegisterEnum("proto_go.MsgId", MsgId_name, MsgId_value)
}

func init() { proto.RegisterFile("testPB.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x49, 0x2d, 0x2e,
	0x09, 0x70, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xf1, 0xe9, 0xf9, 0x4a,
	0x5a, 0x5c, 0x3c, 0x41, 0xa9, 0x85, 0xa5, 0x40, 0x49, 0xdf, 0xc4, 0x92, 0xe4, 0x0c, 0x21, 0x6e,
	0x2e, 0xe6, 0x50, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x56, 0x21, 0x5e, 0x2e, 0xd6, 0xe0,
	0xe4, 0xfc, 0xa2, 0x54, 0x09, 0x26, 0x10, 0x57, 0x29, 0x99, 0x8b, 0x1b, 0xac, 0x28, 0x28, 0xb5,
	0xb8, 0x34, 0xa7, 0x44, 0x48, 0x88, 0x8b, 0xcb, 0xb3, 0x38, 0x24, 0x35, 0x31, 0xd7, 0x30, 0x3c,
	0x33, 0x0f, 0xac, 0x83, 0x43, 0x48, 0x9e, 0x8b, 0x15, 0x2c, 0x02, 0xd4, 0xc1, 0xac, 0xc1, 0x6d,
	0x24, 0xa0, 0x07, 0xb3, 0x48, 0x2f, 0x20, 0x27, 0xb1, 0x32, 0xb5, 0x08, 0xa6, 0xc0, 0x48, 0x82,
	0x19, 0xbb, 0x02, 0x25, 0x15, 0x2e, 0x36, 0xa8, 0x52, 0x3c, 0x4e, 0xd1, 0xd2, 0xe1, 0x62, 0xf5,
	0x2d, 0x4e, 0xf7, 0x4c, 0x11, 0x12, 0xe4, 0xe2, 0x0d, 0x72, 0x0d, 0x0c, 0x75, 0x0d, 0x0e, 0x89,
	0xf7, 0x75, 0x0c, 0x71, 0xf6, 0x10, 0x60, 0x14, 0x12, 0xe0, 0xe2, 0x01, 0x33, 0xe3, 0x83, 0x5c,
	0x83, 0x43, 0x7d, 0x42, 0x04, 0x98, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x7e, 0x14, 0xb2,
	0xfd, 0x00, 0x00, 0x00,
}
