// +build utils

/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

package main

import proto "google.golang.org/protobuf/proto"
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

type PingMessage struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	DelayNanos           int64    `protobuf:"varint,4,opt,name=delayNanos,proto3" json:"delayNanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ping_ba7bbf995167e761, []int{0}
}
func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (dst *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(dst, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PingMessage) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *PingMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PingMessage) GetDelayNanos() int64 {
	if m != nil {
		return m.DelayNanos
	}
	return 0
}

func init() {
	proto.RegisterType((*PingMessage)(nil), "fgrpc.PingMessage")
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_ping_ba7bbf995167e761) }

var fileDescriptor_ping_ba7bbf995167e761 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4b, 0x2f, 0x2a, 0x48, 0x56, 0xca, 0xe4,
	0xe2, 0x0e, 0x00, 0x0a, 0xfa, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x09, 0x70, 0x31, 0x17,
	0xa7, 0x16, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81, 0x98, 0x42, 0x7c, 0x5c, 0x4c, 0x25,
	0xc5, 0x12, 0x4c, 0x60, 0x01, 0x20, 0x4b, 0x48, 0x82, 0x8b, 0xbd, 0x20, 0xb1, 0x32, 0x27, 0x3f,
	0x31, 0x45, 0x82, 0x19, 0x28, 0xc8, 0x19, 0x04, 0xe3, 0x0a, 0xc9, 0x71, 0x71, 0xa5, 0xa4, 0xe6,
	0x24, 0x56, 0xfa, 0x25, 0xe6, 0xe5, 0x17, 0x4b, 0xb0, 0x80, 0x75, 0x20, 0x89, 0x18, 0xd9, 0x71,
	0x71, 0x81, 0xac, 0x0a, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x12, 0x32, 0xe0, 0x62, 0x01, 0xf1, 0x84,
	0x84, 0xf4, 0xc0, 0x0e, 0xd1, 0x43, 0x72, 0x85, 0x14, 0x16, 0x31, 0x25, 0x86, 0x24, 0x36, 0xb0,
	0xc3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x54, 0x22, 0xa0, 0xfe, 0xc6, 0x00, 0x00, 0x00,
}
