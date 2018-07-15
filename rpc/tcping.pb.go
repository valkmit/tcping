// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tcping.proto

package rpc

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_tcping_678b7ef5c6e9fe84, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Probe struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe) Reset()         { *m = Probe{} }
func (m *Probe) String() string { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()    {}
func (*Probe) Descriptor() ([]byte, []int) {
	return fileDescriptor_tcping_678b7ef5c6e9fe84, []int{1}
}
func (m *Probe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe.Unmarshal(m, b)
}
func (m *Probe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe.Marshal(b, m, deterministic)
}
func (dst *Probe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe.Merge(dst, src)
}
func (m *Probe) XXX_Size() int {
	return xxx_messageInfo_Probe.Size(m)
}
func (m *Probe) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe.DiscardUnknown(m)
}

var xxx_messageInfo_Probe proto.InternalMessageInfo

func (m *Probe) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Probe) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Probe) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Probe) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ProbeSchedule struct {
	Probe                *Probe   `protobuf:"bytes,1,opt,name=probe,proto3" json:"probe,omitempty"`
	Schedule             string   `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeSchedule) Reset()         { *m = ProbeSchedule{} }
func (m *ProbeSchedule) String() string { return proto.CompactTextString(m) }
func (*ProbeSchedule) ProtoMessage()    {}
func (*ProbeSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_tcping_678b7ef5c6e9fe84, []int{2}
}
func (m *ProbeSchedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeSchedule.Unmarshal(m, b)
}
func (m *ProbeSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeSchedule.Marshal(b, m, deterministic)
}
func (dst *ProbeSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeSchedule.Merge(dst, src)
}
func (m *ProbeSchedule) XXX_Size() int {
	return xxx_messageInfo_ProbeSchedule.Size(m)
}
func (m *ProbeSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeSchedule proto.InternalMessageInfo

func (m *ProbeSchedule) GetProbe() *Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProbeSchedule) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

type ProbeQuery struct {
	Probe                *Probe   `protobuf:"bytes,1,opt,name=probe,proto3" json:"probe,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64    `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeQuery) Reset()         { *m = ProbeQuery{} }
func (m *ProbeQuery) String() string { return proto.CompactTextString(m) }
func (*ProbeQuery) ProtoMessage()    {}
func (*ProbeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_tcping_678b7ef5c6e9fe84, []int{3}
}
func (m *ProbeQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeQuery.Unmarshal(m, b)
}
func (m *ProbeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeQuery.Marshal(b, m, deterministic)
}
func (dst *ProbeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeQuery.Merge(dst, src)
}
func (m *ProbeQuery) XXX_Size() int {
	return xxx_messageInfo_ProbeQuery.Size(m)
}
func (m *ProbeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeQuery proto.InternalMessageInfo

func (m *ProbeQuery) GetProbe() *Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProbeQuery) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ProbeQuery) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type ProbeQueryResult struct {
	Probe                *Probe   `protobuf:"bytes,1,opt,name=probe,proto3" json:"probe,omitempty"`
	Latency              int64    `protobuf:"varint,2,opt,name=latency,proto3" json:"latency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeQueryResult) Reset()         { *m = ProbeQueryResult{} }
func (m *ProbeQueryResult) String() string { return proto.CompactTextString(m) }
func (*ProbeQueryResult) ProtoMessage()    {}
func (*ProbeQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_tcping_678b7ef5c6e9fe84, []int{4}
}
func (m *ProbeQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeQueryResult.Unmarshal(m, b)
}
func (m *ProbeQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeQueryResult.Marshal(b, m, deterministic)
}
func (dst *ProbeQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeQueryResult.Merge(dst, src)
}
func (m *ProbeQueryResult) XXX_Size() int {
	return xxx_messageInfo_ProbeQueryResult.Size(m)
}
func (m *ProbeQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeQueryResult proto.InternalMessageInfo

func (m *ProbeQueryResult) GetProbe() *Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProbeQueryResult) GetLatency() int64 {
	if m != nil {
		return m.Latency
	}
	return 0
}

type ProbeQueryResults struct {
	Results              []*ProbeQueryResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProbeQueryResults) Reset()         { *m = ProbeQueryResults{} }
func (m *ProbeQueryResults) String() string { return proto.CompactTextString(m) }
func (*ProbeQueryResults) ProtoMessage()    {}
func (*ProbeQueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_tcping_678b7ef5c6e9fe84, []int{5}
}
func (m *ProbeQueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeQueryResults.Unmarshal(m, b)
}
func (m *ProbeQueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeQueryResults.Marshal(b, m, deterministic)
}
func (dst *ProbeQueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeQueryResults.Merge(dst, src)
}
func (m *ProbeQueryResults) XXX_Size() int {
	return xxx_messageInfo_ProbeQueryResults.Size(m)
}
func (m *ProbeQueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeQueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeQueryResults proto.InternalMessageInfo

func (m *ProbeQueryResults) GetResults() []*ProbeQueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "rpc.Empty")
	proto.RegisterType((*Probe)(nil), "rpc.Probe")
	proto.RegisterType((*ProbeSchedule)(nil), "rpc.ProbeSchedule")
	proto.RegisterType((*ProbeQuery)(nil), "rpc.ProbeQuery")
	proto.RegisterType((*ProbeQueryResult)(nil), "rpc.ProbeQueryResult")
	proto.RegisterType((*ProbeQueryResults)(nil), "rpc.ProbeQueryResults")
}

func init() { proto.RegisterFile("tcping.proto", fileDescriptor_tcping_678b7ef5c6e9fe84) }

var fileDescriptor_tcping_678b7ef5c6e9fe84 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0xea, 0x50,
	0x10, 0xed, 0x07, 0x7d, 0x7d, 0x0c, 0x8f, 0x87, 0x4e, 0xd4, 0x34, 0x8d, 0x8b, 0xe6, 0x6e, 0x64,
	0x85, 0x06, 0x56, 0x46, 0x77, 0x6a, 0x5c, 0x69, 0xb4, 0xc0, 0x0f, 0x28, 0xed, 0x28, 0x4d, 0x4a,
	0xdb, 0xdc, 0x7b, 0x31, 0xe1, 0xf7, 0xf8, 0x47, 0x0d, 0xf7, 0xb6, 0xb6, 0x12, 0x12, 0xd9, 0xcd,
	0x9c, 0x73, 0x7a, 0x4e, 0xe7, 0xce, 0xc0, 0x3f, 0x19, 0x97, 0x69, 0xfe, 0x3e, 0x2a, 0x79, 0x21,
	0x0b, 0xb4, 0x79, 0x19, 0x33, 0x17, 0x9c, 0x87, 0x55, 0x29, 0x37, 0x6c, 0x0e, 0xce, 0x0b, 0x2f,
	0x16, 0x84, 0xff, 0xc1, 0x4a, 0x13, 0xcf, 0x0c, 0xcc, 0x61, 0x37, 0xb4, 0xd2, 0x04, 0x4f, 0xc0,
	0xc9, 0xa2, 0x05, 0x65, 0x9e, 0xa5, 0x20, 0xdd, 0x20, 0x42, 0x67, 0x59, 0x08, 0xe9, 0xd9, 0x0a,
	0x54, 0xf5, 0x16, 0x2b, 0x0b, 0x2e, 0xbd, 0x4e, 0x60, 0x0e, 0x9d, 0x50, 0xd5, 0xec, 0x09, 0xfa,
	0xca, 0x76, 0x1a, 0x2f, 0x29, 0x59, 0x67, 0x84, 0x01, 0x38, 0xe5, 0x16, 0x50, 0x09, 0xbd, 0x31,
	0x8c, 0x78, 0x19, 0x8f, 0x94, 0x24, 0xd4, 0x04, 0xfa, 0xf0, 0x57, 0x54, 0xea, 0x2a, 0xf3, 0xbb,
	0x67, 0x6f, 0x00, 0x4a, 0xfb, 0xba, 0x26, 0xbe, 0x39, 0xc0, 0xeb, 0x1c, 0xba, 0x42, 0x46, 0x5c,
	0xce, 0xd2, 0x95, 0x36, 0xb3, 0xc3, 0x06, 0x40, 0x0f, 0x5c, 0xca, 0x13, 0xc5, 0xd9, 0x8a, 0xab,
	0x5b, 0xf6, 0x0c, 0x47, 0x4d, 0x4e, 0x48, 0x62, 0x9d, 0xc9, 0x03, 0xd2, 0x3c, 0x70, 0xb3, 0x48,
	0x52, 0x1e, 0x6f, 0xaa, 0xac, 0xba, 0x65, 0xf7, 0x70, 0xbc, 0xeb, 0x27, 0xf0, 0x12, 0x5c, 0xae,
	0x4b, 0xcf, 0x0c, 0xec, 0x61, 0x6f, 0x7c, 0xda, 0x58, 0xb6, 0x84, 0x61, 0xad, 0x1a, 0x7f, 0x5a,
	0xd0, 0x9f, 0xa9, 0x15, 0x4e, 0x89, 0x7f, 0xa4, 0x31, 0xe1, 0x05, 0xf4, 0xee, 0x38, 0x45, 0x92,
	0xf4, 0xee, 0x5a, 0xff, 0xe4, 0xb7, 0x6a, 0x66, 0xe0, 0x35, 0xf4, 0xeb, 0x15, 0x68, 0x29, 0x36,
	0x74, 0x4d, 0xf8, 0x7b, 0x30, 0x66, 0xe0, 0x04, 0x06, 0xf3, 0x5c, 0xfc, 0xfa, 0xb1, 0xce, 0xd3,
	0xc7, 0x64, 0xe0, 0x2d, 0x0c, 0x1e, 0x49, 0xea, 0xd7, 0xa9, 0xc6, 0x1d, 0xec, 0x4c, 0xe7, 0x9f,
	0xed, 0x1d, 0x57, 0x30, 0x03, 0x6f, 0x00, 0xa7, 0x92, 0x53, 0xb4, 0xfa, 0x61, 0xd0, 0x9e, 0x6e,
	0xff, 0x53, 0x31, 0xe3, 0xca, 0x5c, 0xfc, 0x51, 0xe7, 0x3d, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x48, 0x37, 0x07, 0xf6, 0xee, 0x02, 0x00, 0x00,
}
