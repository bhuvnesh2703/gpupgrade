// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub_to_agent.proto

package idl

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UpgradePrimariesRequest struct {
	SourceBinDir         string         `protobuf:"bytes,1,opt,name=SourceBinDir,proto3" json:"SourceBinDir,omitempty"`
	TargetBinDir         string         `protobuf:"bytes,2,opt,name=TargetBinDir,proto3" json:"TargetBinDir,omitempty"`
	TargetVersion        string         `protobuf:"bytes,3,opt,name=TargetVersion,proto3" json:"TargetVersion,omitempty"`
	DataDirPairs         []*DataDirPair `protobuf:"bytes,4,rep,name=DataDirPairs,proto3" json:"DataDirPairs,omitempty"`
	CheckOnly            bool           `protobuf:"varint,5,opt,name=CheckOnly,proto3" json:"CheckOnly,omitempty"`
	UseLinkMode          bool           `protobuf:"varint,6,opt,name=UseLinkMode,proto3" json:"UseLinkMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpgradePrimariesRequest) Reset()         { *m = UpgradePrimariesRequest{} }
func (m *UpgradePrimariesRequest) String() string { return proto.CompactTextString(m) }
func (*UpgradePrimariesRequest) ProtoMessage()    {}
func (*UpgradePrimariesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{0}
}

func (m *UpgradePrimariesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradePrimariesRequest.Unmarshal(m, b)
}
func (m *UpgradePrimariesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradePrimariesRequest.Marshal(b, m, deterministic)
}
func (m *UpgradePrimariesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradePrimariesRequest.Merge(m, src)
}
func (m *UpgradePrimariesRequest) XXX_Size() int {
	return xxx_messageInfo_UpgradePrimariesRequest.Size(m)
}
func (m *UpgradePrimariesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradePrimariesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradePrimariesRequest proto.InternalMessageInfo

func (m *UpgradePrimariesRequest) GetSourceBinDir() string {
	if m != nil {
		return m.SourceBinDir
	}
	return ""
}

func (m *UpgradePrimariesRequest) GetTargetBinDir() string {
	if m != nil {
		return m.TargetBinDir
	}
	return ""
}

func (m *UpgradePrimariesRequest) GetTargetVersion() string {
	if m != nil {
		return m.TargetVersion
	}
	return ""
}

func (m *UpgradePrimariesRequest) GetDataDirPairs() []*DataDirPair {
	if m != nil {
		return m.DataDirPairs
	}
	return nil
}

func (m *UpgradePrimariesRequest) GetCheckOnly() bool {
	if m != nil {
		return m.CheckOnly
	}
	return false
}

func (m *UpgradePrimariesRequest) GetUseLinkMode() bool {
	if m != nil {
		return m.UseLinkMode
	}
	return false
}

type DataDirPair struct {
	SourceDataDir        string   `protobuf:"bytes,1,opt,name=SourceDataDir,proto3" json:"SourceDataDir,omitempty"`
	TargetDataDir        string   `protobuf:"bytes,2,opt,name=TargetDataDir,proto3" json:"TargetDataDir,omitempty"`
	SourcePort           int32    `protobuf:"varint,3,opt,name=SourcePort,proto3" json:"SourcePort,omitempty"`
	TargetPort           int32    `protobuf:"varint,4,opt,name=TargetPort,proto3" json:"TargetPort,omitempty"`
	Content              int32    `protobuf:"varint,5,opt,name=Content,proto3" json:"Content,omitempty"`
	DBID                 int32    `protobuf:"varint,6,opt,name=DBID,proto3" json:"DBID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataDirPair) Reset()         { *m = DataDirPair{} }
func (m *DataDirPair) String() string { return proto.CompactTextString(m) }
func (*DataDirPair) ProtoMessage()    {}
func (*DataDirPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{1}
}

func (m *DataDirPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataDirPair.Unmarshal(m, b)
}
func (m *DataDirPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataDirPair.Marshal(b, m, deterministic)
}
func (m *DataDirPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataDirPair.Merge(m, src)
}
func (m *DataDirPair) XXX_Size() int {
	return xxx_messageInfo_DataDirPair.Size(m)
}
func (m *DataDirPair) XXX_DiscardUnknown() {
	xxx_messageInfo_DataDirPair.DiscardUnknown(m)
}

var xxx_messageInfo_DataDirPair proto.InternalMessageInfo

func (m *DataDirPair) GetSourceDataDir() string {
	if m != nil {
		return m.SourceDataDir
	}
	return ""
}

func (m *DataDirPair) GetTargetDataDir() string {
	if m != nil {
		return m.TargetDataDir
	}
	return ""
}

func (m *DataDirPair) GetSourcePort() int32 {
	if m != nil {
		return m.SourcePort
	}
	return 0
}

func (m *DataDirPair) GetTargetPort() int32 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

func (m *DataDirPair) GetContent() int32 {
	if m != nil {
		return m.Content
	}
	return 0
}

func (m *DataDirPair) GetDBID() int32 {
	if m != nil {
		return m.DBID
	}
	return 0
}

type UpgradePrimariesReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradePrimariesReply) Reset()         { *m = UpgradePrimariesReply{} }
func (m *UpgradePrimariesReply) String() string { return proto.CompactTextString(m) }
func (*UpgradePrimariesReply) ProtoMessage()    {}
func (*UpgradePrimariesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{2}
}

func (m *UpgradePrimariesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradePrimariesReply.Unmarshal(m, b)
}
func (m *UpgradePrimariesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradePrimariesReply.Marshal(b, m, deterministic)
}
func (m *UpgradePrimariesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradePrimariesReply.Merge(m, src)
}
func (m *UpgradePrimariesReply) XXX_Size() int {
	return xxx_messageInfo_UpgradePrimariesReply.Size(m)
}
func (m *UpgradePrimariesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradePrimariesReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradePrimariesReply proto.InternalMessageInfo

type CreateSegmentDataDirRequest struct {
	Datadirs             []string `protobuf:"bytes,1,rep,name=datadirs,proto3" json:"datadirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSegmentDataDirRequest) Reset()         { *m = CreateSegmentDataDirRequest{} }
func (m *CreateSegmentDataDirRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSegmentDataDirRequest) ProtoMessage()    {}
func (*CreateSegmentDataDirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{3}
}

func (m *CreateSegmentDataDirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSegmentDataDirRequest.Unmarshal(m, b)
}
func (m *CreateSegmentDataDirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSegmentDataDirRequest.Marshal(b, m, deterministic)
}
func (m *CreateSegmentDataDirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSegmentDataDirRequest.Merge(m, src)
}
func (m *CreateSegmentDataDirRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSegmentDataDirRequest.Size(m)
}
func (m *CreateSegmentDataDirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSegmentDataDirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSegmentDataDirRequest proto.InternalMessageInfo

func (m *CreateSegmentDataDirRequest) GetDatadirs() []string {
	if m != nil {
		return m.Datadirs
	}
	return nil
}

type CreateSegmentDataDirReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSegmentDataDirReply) Reset()         { *m = CreateSegmentDataDirReply{} }
func (m *CreateSegmentDataDirReply) String() string { return proto.CompactTextString(m) }
func (*CreateSegmentDataDirReply) ProtoMessage()    {}
func (*CreateSegmentDataDirReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{4}
}

func (m *CreateSegmentDataDirReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSegmentDataDirReply.Unmarshal(m, b)
}
func (m *CreateSegmentDataDirReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSegmentDataDirReply.Marshal(b, m, deterministic)
}
func (m *CreateSegmentDataDirReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSegmentDataDirReply.Merge(m, src)
}
func (m *CreateSegmentDataDirReply) XXX_Size() int {
	return xxx_messageInfo_CreateSegmentDataDirReply.Size(m)
}
func (m *CreateSegmentDataDirReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSegmentDataDirReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSegmentDataDirReply proto.InternalMessageInfo

type CopyMasterRequest struct {
	MasterDir            string   `protobuf:"bytes,1,opt,name=masterDir,proto3" json:"masterDir,omitempty"`
	Datadirs             []string `protobuf:"bytes,2,rep,name=datadirs,proto3" json:"datadirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyMasterRequest) Reset()         { *m = CopyMasterRequest{} }
func (m *CopyMasterRequest) String() string { return proto.CompactTextString(m) }
func (*CopyMasterRequest) ProtoMessage()    {}
func (*CopyMasterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{5}
}

func (m *CopyMasterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyMasterRequest.Unmarshal(m, b)
}
func (m *CopyMasterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyMasterRequest.Marshal(b, m, deterministic)
}
func (m *CopyMasterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyMasterRequest.Merge(m, src)
}
func (m *CopyMasterRequest) XXX_Size() int {
	return xxx_messageInfo_CopyMasterRequest.Size(m)
}
func (m *CopyMasterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyMasterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CopyMasterRequest proto.InternalMessageInfo

func (m *CopyMasterRequest) GetMasterDir() string {
	if m != nil {
		return m.MasterDir
	}
	return ""
}

func (m *CopyMasterRequest) GetDatadirs() []string {
	if m != nil {
		return m.Datadirs
	}
	return nil
}

type CopyMasterReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyMasterReply) Reset()         { *m = CopyMasterReply{} }
func (m *CopyMasterReply) String() string { return proto.CompactTextString(m) }
func (*CopyMasterReply) ProtoMessage()    {}
func (*CopyMasterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{6}
}

func (m *CopyMasterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyMasterReply.Unmarshal(m, b)
}
func (m *CopyMasterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyMasterReply.Marshal(b, m, deterministic)
}
func (m *CopyMasterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyMasterReply.Merge(m, src)
}
func (m *CopyMasterReply) XXX_Size() int {
	return xxx_messageInfo_CopyMasterReply.Size(m)
}
func (m *CopyMasterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyMasterReply.DiscardUnknown(m)
}

var xxx_messageInfo_CopyMasterReply proto.InternalMessageInfo

type StopAgentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentRequest) Reset()         { *m = StopAgentRequest{} }
func (m *StopAgentRequest) String() string { return proto.CompactTextString(m) }
func (*StopAgentRequest) ProtoMessage()    {}
func (*StopAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{7}
}

func (m *StopAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAgentRequest.Unmarshal(m, b)
}
func (m *StopAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAgentRequest.Marshal(b, m, deterministic)
}
func (m *StopAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAgentRequest.Merge(m, src)
}
func (m *StopAgentRequest) XXX_Size() int {
	return xxx_messageInfo_StopAgentRequest.Size(m)
}
func (m *StopAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopAgentRequest proto.InternalMessageInfo

type StopAgentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentReply) Reset()         { *m = StopAgentReply{} }
func (m *StopAgentReply) String() string { return proto.CompactTextString(m) }
func (*StopAgentReply) ProtoMessage()    {}
func (*StopAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{8}
}

func (m *StopAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAgentReply.Unmarshal(m, b)
}
func (m *StopAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAgentReply.Marshal(b, m, deterministic)
}
func (m *StopAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAgentReply.Merge(m, src)
}
func (m *StopAgentReply) XXX_Size() int {
	return xxx_messageInfo_StopAgentReply.Size(m)
}
func (m *StopAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_StopAgentReply proto.InternalMessageInfo

type CheckSegmentDiskSpaceRequest struct {
	Request              *CheckDiskSpaceRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Datadirs             []string               `protobuf:"bytes,2,rep,name=datadirs,proto3" json:"datadirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CheckSegmentDiskSpaceRequest) Reset()         { *m = CheckSegmentDiskSpaceRequest{} }
func (m *CheckSegmentDiskSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*CheckSegmentDiskSpaceRequest) ProtoMessage()    {}
func (*CheckSegmentDiskSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e73bb06acc917d8, []int{9}
}

func (m *CheckSegmentDiskSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSegmentDiskSpaceRequest.Unmarshal(m, b)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSegmentDiskSpaceRequest.Marshal(b, m, deterministic)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSegmentDiskSpaceRequest.Merge(m, src)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_CheckSegmentDiskSpaceRequest.Size(m)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSegmentDiskSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSegmentDiskSpaceRequest proto.InternalMessageInfo

func (m *CheckSegmentDiskSpaceRequest) GetRequest() *CheckDiskSpaceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *CheckSegmentDiskSpaceRequest) GetDatadirs() []string {
	if m != nil {
		return m.Datadirs
	}
	return nil
}

func init() {
	proto.RegisterType((*UpgradePrimariesRequest)(nil), "idl.UpgradePrimariesRequest")
	proto.RegisterType((*DataDirPair)(nil), "idl.DataDirPair")
	proto.RegisterType((*UpgradePrimariesReply)(nil), "idl.UpgradePrimariesReply")
	proto.RegisterType((*CreateSegmentDataDirRequest)(nil), "idl.CreateSegmentDataDirRequest")
	proto.RegisterType((*CreateSegmentDataDirReply)(nil), "idl.CreateSegmentDataDirReply")
	proto.RegisterType((*CopyMasterRequest)(nil), "idl.CopyMasterRequest")
	proto.RegisterType((*CopyMasterReply)(nil), "idl.CopyMasterReply")
	proto.RegisterType((*StopAgentRequest)(nil), "idl.StopAgentRequest")
	proto.RegisterType((*StopAgentReply)(nil), "idl.StopAgentReply")
	proto.RegisterType((*CheckSegmentDiskSpaceRequest)(nil), "idl.CheckSegmentDiskSpaceRequest")
}

func init() { proto.RegisterFile("hub_to_agent.proto", fileDescriptor_9e73bb06acc917d8) }

var fileDescriptor_9e73bb06acc917d8 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xad, 0x01, 0x27, 0x61, 0x48, 0x53, 0x67, 0xdb, 0x34, 0xae, 0x83, 0x22, 0x6a, 0xf5, 0x90,
	0x13, 0x07, 0x9a, 0x4b, 0xa4, 0x5e, 0x1a, 0xb8, 0x54, 0x2a, 0x0d, 0x32, 0x4d, 0xaf, 0xd1, 0x62,
	0xaf, 0xc8, 0x0a, 0xf0, 0xba, 0xeb, 0xe5, 0xc0, 0x0f, 0xe8, 0x6f, 0xeb, 0xaf, 0xaa, 0xd4, 0xfd,
	0xb0, 0xf1, 0x47, 0x80, 0x9b, 0xf7, 0xcd, 0x9b, 0x37, 0x6f, 0xf6, 0x2d, 0x00, 0x7a, 0x5e, 0xcf,
	0x9e, 0x04, 0x7b, 0xc2, 0x73, 0x12, 0x8b, 0x7e, 0xc2, 0x99, 0x60, 0xa8, 0x49, 0xa3, 0xa5, 0xe7,
	0x84, 0x4b, 0xaa, 0x0a, 0xb2, 0x6e, 0x60, 0xff, 0x9f, 0x05, 0x97, 0x8f, 0xc9, 0x9c, 0xe3, 0x88,
	0x4c, 0x38, 0x5d, 0x61, 0x4e, 0x49, 0x1a, 0x90, 0xdf, 0x6b, 0x92, 0x0a, 0xe4, 0xc3, 0xe9, 0x94,
	0xad, 0x79, 0x48, 0xee, 0x69, 0x3c, 0xa2, 0xdc, 0xb5, 0x7a, 0xd6, 0x4d, 0x3b, 0xa8, 0x60, 0x8a,
	0xf3, 0x13, 0xf3, 0x39, 0x11, 0x19, 0xa7, 0x61, 0x38, 0x65, 0x0c, 0x7d, 0x82, 0xd7, 0xe6, 0xfc,
	0x8b, 0xf0, 0x94, 0xb2, 0xd8, 0x6d, 0x6a, 0x52, 0x15, 0x44, 0xb7, 0x70, 0x3a, 0xc2, 0x02, 0xcb,
	0x86, 0x09, 0xa6, 0x3c, 0x75, 0x5b, 0xbd, 0xe6, 0x4d, 0x67, 0xe0, 0xf4, 0xa5, 0xef, 0x7e, 0xa9,
	0x10, 0x54, 0x58, 0xa8, 0x0b, 0xed, 0xe1, 0x33, 0x09, 0x17, 0x0f, 0xf1, 0x72, 0xe3, 0xda, 0x52,
	0xf7, 0x24, 0x28, 0x00, 0xd4, 0x83, 0xce, 0x63, 0x4a, 0xbe, 0xd3, 0x78, 0x31, 0x66, 0x11, 0x71,
	0x8f, 0x74, 0xbd, 0x0c, 0xf9, 0x7f, 0x2d, 0xe8, 0x94, 0x04, 0x95, 0x57, 0xb3, 0x5f, 0x06, 0x66,
	0x4b, 0x57, 0xc1, 0x62, 0xa3, 0x9c, 0xd5, 0x28, 0x6f, 0x94, 0xb3, 0xae, 0x01, 0x4c, 0xdb, 0x84,
	0x71, 0xa1, 0x97, 0xb6, 0x83, 0x12, 0xa2, 0xea, 0xa6, 0x41, 0xd7, 0x5b, 0xa6, 0x5e, 0x20, 0xc8,
	0x85, 0xe3, 0x21, 0x8b, 0x85, 0xcc, 0x50, 0x6f, 0x66, 0x07, 0xf9, 0x11, 0x21, 0x68, 0x8d, 0xee,
	0xbf, 0x8d, 0xf4, 0x42, 0x76, 0xa0, 0xbf, 0xfd, 0x4b, 0xb8, 0x78, 0x19, 0x64, 0xb2, 0xdc, 0xf8,
	0x77, 0x70, 0x35, 0xe4, 0x04, 0x0b, 0x32, 0x25, 0xf3, 0x95, 0xec, 0xce, 0xec, 0xe5, 0x29, 0x7b,
	0x70, 0x12, 0x49, 0x24, 0x52, 0x77, 0x6e, 0xc9, 0x3b, 0x6f, 0x07, 0xdb, 0xb3, 0x7f, 0x05, 0x1f,
	0x76, 0xb7, 0x2a, 0xdd, 0x31, 0x9c, 0x0f, 0x59, 0xb2, 0x19, 0xe3, 0x54, 0x90, 0xad, 0x9a, 0xcc,
	0x63, 0xa5, 0x81, 0xe2, 0xee, 0x0a, 0xa0, 0x32, 0xab, 0x51, 0x9b, 0x75, 0x0e, 0x6f, 0xca, 0x72,
	0x6a, 0x02, 0x02, 0x67, 0x2a, 0x58, 0xf2, 0x55, 0x3d, 0xe3, 0x6c, 0x80, 0xef, 0xc0, 0x59, 0x09,
	0x53, 0xac, 0x04, 0xba, 0x3a, 0xf1, 0xdc, 0x23, 0x4d, 0x17, 0xd3, 0x04, 0x87, 0x24, 0xb7, 0x74,
	0x0b, 0xc7, 0xdc, 0x7c, 0x6a, 0x43, 0x9d, 0x81, 0xa7, 0xdf, 0x94, 0xee, 0xa9, 0x93, 0x83, 0x9c,
	0x7a, 0xc8, 0xea, 0xe0, 0x4f, 0x13, 0x6c, 0x6d, 0x00, 0x3d, 0xc0, 0x59, 0x55, 0x07, 0x7d, 0x2c,
	0xc4, 0xf7, 0x18, 0xf2, 0xdc, 0x9d, 0xf3, 0xd5, 0x2a, 0xaf, 0xd0, 0x0f, 0x70, 0xea, 0x29, 0xa2,
	0xae, 0xe6, 0xef, 0xf9, 0x95, 0x7a, 0xde, 0x9e, 0xaa, 0xd1, 0x9b, 0xc9, 0xcb, 0xd9, 0x91, 0x20,
	0x09, 0x05, 0xd3, 0xda, 0x3d, 0xe3, 0x65, 0xff, 0xfb, 0xf0, 0xae, 0x0f, 0x30, 0xcc, 0x8c, 0x2f,
	0x00, 0x45, 0x72, 0xe8, 0xbd, 0xe1, 0xd7, 0x5f, 0x86, 0xf7, 0xee, 0x05, 0x6e, 0xba, 0xef, 0xa0,
	0xbd, 0x0d, 0x14, 0x5d, 0x68, 0x52, 0x3d, 0x74, 0xef, 0x6d, 0x1d, 0xd6, 0xad, 0xb3, 0x23, 0xfd,
	0x1f, 0xf6, 0xf9, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x55, 0x1d, 0xe7, 0xf0, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	CheckDiskSpace(ctx context.Context, in *CheckSegmentDiskSpaceRequest, opts ...grpc.CallOption) (*CheckDiskSpaceReply, error)
	UpgradePrimaries(ctx context.Context, in *UpgradePrimariesRequest, opts ...grpc.CallOption) (*UpgradePrimariesReply, error)
	CreateSegmentDataDirectories(ctx context.Context, in *CreateSegmentDataDirRequest, opts ...grpc.CallOption) (*CreateSegmentDataDirReply, error)
	CopyMaster(ctx context.Context, in *CopyMasterRequest, opts ...grpc.CallOption) (*CopyMasterReply, error)
	StopAgent(ctx context.Context, in *StopAgentRequest, opts ...grpc.CallOption) (*StopAgentReply, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CheckDiskSpace(ctx context.Context, in *CheckSegmentDiskSpaceRequest, opts ...grpc.CallOption) (*CheckDiskSpaceReply, error) {
	out := new(CheckDiskSpaceReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/CheckDiskSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpgradePrimaries(ctx context.Context, in *UpgradePrimariesRequest, opts ...grpc.CallOption) (*UpgradePrimariesReply, error) {
	out := new(UpgradePrimariesReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/UpgradePrimaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateSegmentDataDirectories(ctx context.Context, in *CreateSegmentDataDirRequest, opts ...grpc.CallOption) (*CreateSegmentDataDirReply, error) {
	out := new(CreateSegmentDataDirReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/CreateSegmentDataDirectories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CopyMaster(ctx context.Context, in *CopyMasterRequest, opts ...grpc.CallOption) (*CopyMasterReply, error) {
	out := new(CopyMasterReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/CopyMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StopAgent(ctx context.Context, in *StopAgentRequest, opts ...grpc.CallOption) (*StopAgentReply, error) {
	out := new(StopAgentReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/StopAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	CheckDiskSpace(context.Context, *CheckSegmentDiskSpaceRequest) (*CheckDiskSpaceReply, error)
	UpgradePrimaries(context.Context, *UpgradePrimariesRequest) (*UpgradePrimariesReply, error)
	CreateSegmentDataDirectories(context.Context, *CreateSegmentDataDirRequest) (*CreateSegmentDataDirReply, error)
	CopyMaster(context.Context, *CopyMasterRequest) (*CopyMasterReply, error)
	StopAgent(context.Context, *StopAgentRequest) (*StopAgentReply, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) CheckDiskSpace(ctx context.Context, req *CheckSegmentDiskSpaceRequest) (*CheckDiskSpaceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDiskSpace not implemented")
}
func (*UnimplementedAgentServer) UpgradePrimaries(ctx context.Context, req *UpgradePrimariesRequest) (*UpgradePrimariesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradePrimaries not implemented")
}
func (*UnimplementedAgentServer) CreateSegmentDataDirectories(ctx context.Context, req *CreateSegmentDataDirRequest) (*CreateSegmentDataDirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSegmentDataDirectories not implemented")
}
func (*UnimplementedAgentServer) CopyMaster(ctx context.Context, req *CopyMasterRequest) (*CopyMasterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyMaster not implemented")
}
func (*UnimplementedAgentServer) StopAgent(ctx context.Context, req *StopAgentRequest) (*StopAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAgent not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_CheckDiskSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSegmentDiskSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CheckDiskSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CheckDiskSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CheckDiskSpace(ctx, req.(*CheckSegmentDiskSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpgradePrimaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradePrimariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpgradePrimaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/UpgradePrimaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpgradePrimaries(ctx, req.(*UpgradePrimariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateSegmentDataDirectories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSegmentDataDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateSegmentDataDirectories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CreateSegmentDataDirectories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateSegmentDataDirectories(ctx, req.(*CreateSegmentDataDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CopyMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CopyMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CopyMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CopyMaster(ctx, req.(*CopyMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StopAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StopAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/StopAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StopAgent(ctx, req.(*StopAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckDiskSpace",
			Handler:    _Agent_CheckDiskSpace_Handler,
		},
		{
			MethodName: "UpgradePrimaries",
			Handler:    _Agent_UpgradePrimaries_Handler,
		},
		{
			MethodName: "CreateSegmentDataDirectories",
			Handler:    _Agent_CreateSegmentDataDirectories_Handler,
		},
		{
			MethodName: "CopyMaster",
			Handler:    _Agent_CopyMaster_Handler,
		},
		{
			MethodName: "StopAgent",
			Handler:    _Agent_StopAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub_to_agent.proto",
}
