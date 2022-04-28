// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package messaging

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

type InitReq struct {
	FileSize             uint64   `protobuf:"varint,1,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileHash             string   `protobuf:"bytes,2,opt,name=FileHash,proto3" json:"FileHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitReq) Reset()         { *m = InitReq{} }
func (m *InitReq) String() string { return proto.CompactTextString(m) }
func (*InitReq) ProtoMessage()    {}
func (*InitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *InitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitReq.Unmarshal(m, b)
}
func (m *InitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitReq.Marshal(b, m, deterministic)
}
func (m *InitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitReq.Merge(m, src)
}
func (m *InitReq) XXX_Size() int {
	return xxx_messageInfo_InitReq.Size(m)
}
func (m *InitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitReq proto.InternalMessageInfo

func (m *InitReq) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *InitReq) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

type InitAck struct {
	Code                 int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *FileInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InitAck) Reset()         { *m = InitAck{} }
func (m *InitAck) String() string { return proto.CompactTextString(m) }
func (*InitAck) ProtoMessage()    {}
func (*InitAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *InitAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitAck.Unmarshal(m, b)
}
func (m *InitAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitAck.Marshal(b, m, deterministic)
}
func (m *InitAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitAck.Merge(m, src)
}
func (m *InitAck) XXX_Size() int {
	return xxx_messageInfo_InitAck.Size(m)
}
func (m *InitAck) XXX_DiscardUnknown() {
	xxx_messageInfo_InitAck.DiscardUnknown(m)
}

var xxx_messageInfo_InitAck proto.InternalMessageInfo

func (m *InitAck) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InitAck) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *InitAck) GetData() *FileInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type CompleteReq struct {
	FileSize             uint64   `protobuf:"varint,1,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileHash             string   `protobuf:"bytes,2,opt,name=FileHash,proto3" json:"FileHash,omitempty"`
	FileName             string   `protobuf:"bytes,3,opt,name=FileName,proto3" json:"FileName,omitempty"`
	UploadID             string   `protobuf:"bytes,4,opt,name=UploadID,proto3" json:"UploadID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteReq) Reset()         { *m = CompleteReq{} }
func (m *CompleteReq) String() string { return proto.CompactTextString(m) }
func (*CompleteReq) ProtoMessage()    {}
func (*CompleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *CompleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteReq.Unmarshal(m, b)
}
func (m *CompleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteReq.Marshal(b, m, deterministic)
}
func (m *CompleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteReq.Merge(m, src)
}
func (m *CompleteReq) XXX_Size() int {
	return xxx_messageInfo_CompleteReq.Size(m)
}
func (m *CompleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteReq proto.InternalMessageInfo

func (m *CompleteReq) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *CompleteReq) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

func (m *CompleteReq) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *CompleteReq) GetUploadID() string {
	if m != nil {
		return m.UploadID
	}
	return ""
}

type CompleteAck struct {
	Code                 int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *FileInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CompleteAck) Reset()         { *m = CompleteAck{} }
func (m *CompleteAck) String() string { return proto.CompactTextString(m) }
func (*CompleteAck) ProtoMessage()    {}
func (*CompleteAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *CompleteAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteAck.Unmarshal(m, b)
}
func (m *CompleteAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteAck.Marshal(b, m, deterministic)
}
func (m *CompleteAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteAck.Merge(m, src)
}
func (m *CompleteAck) XXX_Size() int {
	return xxx_messageInfo_CompleteAck.Size(m)
}
func (m *CompleteAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteAck.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteAck proto.InternalMessageInfo

func (m *CompleteAck) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CompleteAck) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CompleteAck) GetData() *FileInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileInfo struct {
	ChunkSize            uint64   `protobuf:"varint,1,opt,name=ChunkSize,proto3" json:"ChunkSize,omitempty"`
	ChunkCount           uint64   `protobuf:"varint,2,opt,name=ChunkCount,proto3" json:"ChunkCount,omitempty"`
	ChunkExists          []uint64 `protobuf:"varint,3,rep,packed,name=ChunkExists,proto3" json:"ChunkExists,omitempty"`
	UploadID             string   `protobuf:"bytes,4,opt,name=UploadID,proto3" json:"UploadID,omitempty"`
	FileSize             uint64   `protobuf:"varint,5,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileHash             string   `protobuf:"bytes,6,opt,name=FileHash,proto3" json:"FileHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetChunkSize() uint64 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *FileInfo) GetChunkCount() uint64 {
	if m != nil {
		return m.ChunkCount
	}
	return 0
}

func (m *FileInfo) GetChunkExists() []uint64 {
	if m != nil {
		return m.ChunkExists
	}
	return nil
}

func (m *FileInfo) GetUploadID() string {
	if m != nil {
		return m.UploadID
	}
	return ""
}

func (m *FileInfo) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *FileInfo) GetFileHash() string {
	if m != nil {
		return m.FileHash
	}
	return ""
}

type Chunk struct {
	// Types that are valid to be assigned to Data:
	//	*Chunk_Content
	//	*Chunk_Info
	Data                 isChunk_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

type isChunk_Data interface {
	isChunk_Data()
}

type Chunk_Content struct {
	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3,oneof"`
}

type Chunk_Info struct {
	Info *ChunkInfo `protobuf:"bytes,2,opt,name=Info,proto3,oneof"`
}

func (*Chunk_Content) isChunk_Data() {}

func (*Chunk_Info) isChunk_Data() {}

func (m *Chunk) GetData() isChunk_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Chunk) GetContent() []byte {
	if x, ok := m.GetData().(*Chunk_Content); ok {
		return x.Content
	}
	return nil
}

func (m *Chunk) GetInfo() *ChunkInfo {
	if x, ok := m.GetData().(*Chunk_Info); ok {
		return x.Info
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Chunk) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Chunk_Content)(nil),
		(*Chunk_Info)(nil),
	}
}

type ChunkInfo struct {
	FileType             string   `protobuf:"bytes,1,opt,name=FileType,proto3" json:"FileType,omitempty"`
	ChunkIndex           uint64   `protobuf:"varint,2,opt,name=ChunkIndex,proto3" json:"ChunkIndex,omitempty"`
	CheckHash            string   `protobuf:"bytes,3,opt,name=CheckHash,proto3" json:"CheckHash,omitempty"`
	UploadID             string   `protobuf:"bytes,4,opt,name=UploadID,proto3" json:"UploadID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkInfo) Reset()         { *m = ChunkInfo{} }
func (m *ChunkInfo) String() string { return proto.CompactTextString(m) }
func (*ChunkInfo) ProtoMessage()    {}
func (*ChunkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *ChunkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkInfo.Unmarshal(m, b)
}
func (m *ChunkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkInfo.Marshal(b, m, deterministic)
}
func (m *ChunkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkInfo.Merge(m, src)
}
func (m *ChunkInfo) XXX_Size() int {
	return xxx_messageInfo_ChunkInfo.Size(m)
}
func (m *ChunkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkInfo proto.InternalMessageInfo

func (m *ChunkInfo) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

func (m *ChunkInfo) GetChunkIndex() uint64 {
	if m != nil {
		return m.ChunkIndex
	}
	return 0
}

func (m *ChunkInfo) GetCheckHash() string {
	if m != nil {
		return m.CheckHash
	}
	return ""
}

func (m *ChunkInfo) GetUploadID() string {
	if m != nil {
		return m.UploadID
	}
	return ""
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=messaging.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("messaging.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*InitReq)(nil), "messaging.InitReq")
	proto.RegisterType((*InitAck)(nil), "messaging.InitAck")
	proto.RegisterType((*CompleteReq)(nil), "messaging.CompleteReq")
	proto.RegisterType((*CompleteAck)(nil), "messaging.CompleteAck")
	proto.RegisterType((*FileInfo)(nil), "messaging.FileInfo")
	proto.RegisterType((*Chunk)(nil), "messaging.Chunk")
	proto.RegisterType((*ChunkInfo)(nil), "messaging.ChunkInfo")
	proto.RegisterType((*UploadStatus)(nil), "messaging.UploadStatus")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x13, 0xc7, 0x69, 0xc6, 0xa5, 0xb2, 0x06, 0x04, 0x96, 0x41, 0xc8, 0xf2, 0x05, 0xab,
	0x87, 0x80, 0xdc, 0x03, 0x17, 0x2e, 0xc1, 0x50, 0x92, 0x03, 0x20, 0x6d, 0xa9, 0x00, 0x89, 0x8b,
	0xb1, 0x97, 0xd4, 0xb2, 0xb3, 0x1b, 0xea, 0x0d, 0x14, 0x24, 0x0e, 0xfc, 0x18, 0x7f, 0xc4, 0x3f,
	0xa0, 0xdd, 0xb5, 0x13, 0x27, 0x82, 0x08, 0x09, 0x71, 0xdb, 0xf7, 0x66, 0x66, 0xe7, 0xcd, 0x1b,
	0xaf, 0xe1, 0x5a, 0x4d, 0x2f, 0x3f, 0x15, 0x19, 0x1d, 0x2f, 0x2f, 0xb9, 0xe0, 0x38, 0x5a, 0xd0,
	0xba, 0x4e, 0xe7, 0x05, 0x9b, 0x87, 0x13, 0x18, 0xce, 0x58, 0x21, 0x08, 0xfd, 0x88, 0x3e, 0x1c,
	0x9c, 0x16, 0x15, 0x3d, 0x2b, 0xbe, 0x52, 0xcf, 0x0c, 0xcc, 0xc8, 0x22, 0x6b, 0xdc, 0xc6, 0xa6,
	0x69, 0x7d, 0xe1, 0xf5, 0x02, 0x33, 0x1a, 0x91, 0x35, 0x0e, 0xdf, 0xe8, 0x2b, 0x26, 0x59, 0x89,
	0x08, 0x56, 0xc6, 0x73, 0x5d, 0x3e, 0x20, 0xea, 0x8c, 0x2e, 0xf4, 0x17, 0xf5, 0xbc, 0xa9, 0x92,
	0x47, 0xbc, 0x07, 0x56, 0x9e, 0x8a, 0xd4, 0xeb, 0x07, 0x66, 0xe4, 0xc4, 0xd7, 0xc7, 0x6b, 0x35,
	0x63, 0x79, 0xe7, 0x8c, 0x7d, 0xe0, 0x44, 0x25, 0x84, 0xdf, 0xc0, 0x49, 0xf8, 0x62, 0x59, 0x51,
	0x41, 0xff, 0x41, 0x60, 0x1b, 0x7b, 0x91, 0x2e, 0xa8, 0xea, 0xd9, 0xc4, 0x24, 0x96, 0xb1, 0xf3,
	0x65, 0xc5, 0xd3, 0x7c, 0xf6, 0xc4, 0xb3, 0x74, 0xac, 0xc5, 0xe1, 0xbb, 0x4d, 0xfb, 0xff, 0x30,
	0xdc, 0x0f, 0x53, 0xcb, 0x92, 0x14, 0xde, 0x81, 0x51, 0x72, 0xb1, 0x62, 0x65, 0x67, 0xb6, 0x0d,
	0x81, 0x77, 0x01, 0x14, 0x48, 0xf8, 0x8a, 0x09, 0xd5, 0xcc, 0x22, 0x1d, 0x06, 0x03, 0x70, 0x14,
	0x7a, 0x7a, 0x55, 0xd4, 0xa2, 0xf6, 0xfa, 0x41, 0x3f, 0xb2, 0x48, 0x97, 0xda, 0x37, 0xe6, 0x96,
	0xad, 0x83, 0x3d, 0xb6, 0xda, 0x3b, 0x7b, 0x7f, 0x0d, 0x03, 0xd5, 0x02, 0x7d, 0x18, 0x26, 0x9c,
	0x09, 0xca, 0x84, 0x92, 0x7e, 0x38, 0x35, 0x48, 0x4b, 0xe0, 0x31, 0x58, 0x72, 0x40, 0x25, 0xda,
	0x89, 0x6f, 0x74, 0xec, 0x50, 0xb5, 0x32, 0x36, 0x35, 0x88, 0xca, 0x79, 0x6c, 0x6b, 0xeb, 0xc2,
	0xef, 0x66, 0xe3, 0x86, 0xb2, 0xa6, 0x91, 0xf0, 0xea, 0xcb, 0x52, 0x3b, 0xd3, 0x48, 0x90, 0x78,
	0x6d, 0xcc, 0x8c, 0xe5, 0xf4, 0x6a, 0xcb, 0x18, 0xc5, 0x68, 0x5b, 0x69, 0x56, 0x2a, 0xfd, 0x7a,
	0xf5, 0x1b, 0x62, 0xef, 0xee, 0xdf, 0xc2, 0xa1, 0x3e, 0x9f, 0x89, 0x54, 0xac, 0x6a, 0xf4, 0x60,
	0xf8, 0x5c, 0x49, 0x6f, 0x45, 0xb4, 0x10, 0xef, 0x83, 0x95, 0xc8, 0xcf, 0x42, 0x76, 0x3f, 0x8a,
	0x6f, 0x77, 0x26, 0xec, 0x5e, 0x20, 0x53, 0x88, 0x4a, 0x3c, 0x3e, 0x01, 0x77, 0x37, 0x82, 0x0e,
	0x0c, 0xcf, 0x59, 0xc9, 0xf8, 0x67, 0xe6, 0x1a, 0x68, 0x43, 0xef, 0x65, 0xe9, 0x9a, 0x08, 0x60,
	0x9f, 0xa6, 0x45, 0x45, 0x73, 0xb7, 0x17, 0xff, 0x34, 0xe1, 0xe8, 0xd9, 0x4a, 0x97, 0xe9, 0xb7,
	0x8c, 0x0f, 0xc1, 0xd6, 0xf7, 0xa0, 0xbb, 0x6b, 0xab, 0x7f, 0xeb, 0x0f, 0x32, 0x42, 0x23, 0x32,
	0xf1, 0x81, 0xdc, 0x49, 0x21, 0x10, 0x3b, 0x49, 0xcd, 0x4f, 0xc0, 0xdf, 0xe5, 0x26, 0x59, 0x19,
	0x1a, 0xf8, 0x08, 0x0e, 0xda, 0x97, 0x80, 0x37, 0xbb, 0xcd, 0x36, 0xaf, 0xd3, 0xff, 0x1d, 0xaf,
	0xab, 0x63, 0xb0, 0x93, 0x94, 0x65, 0xb4, 0xfa, 0xfb, 0x8e, 0xef, 0x6d, 0xf5, 0xa7, 0x3a, 0xf9,
	0x15, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xde, 0xb9, 0x0a, 0xba, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GuploadServiceClient is the client API for GuploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error)
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitAck, error)
	Complete(ctx context.Context, in *CompleteReq, opts ...grpc.CallOption) (*CompleteAck, error)
	Cancel(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitAck, error)
}

type guploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewGuploadServiceClient(cc *grpc.ClientConn) GuploadServiceClient {
	return &guploadServiceClient{cc}
}

func (c *guploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (GuploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GuploadService_serviceDesc.Streams[0], "/messaging.GuploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &guploadServiceUploadClient{stream}
	return x, nil
}

type GuploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type guploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *guploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guploadServiceClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitAck, error) {
	out := new(InitAck)
	err := c.cc.Invoke(ctx, "/messaging.GuploadService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guploadServiceClient) Complete(ctx context.Context, in *CompleteReq, opts ...grpc.CallOption) (*CompleteAck, error) {
	out := new(CompleteAck)
	err := c.cc.Invoke(ctx, "/messaging.GuploadService/Complete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guploadServiceClient) Cancel(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitAck, error) {
	out := new(InitAck)
	err := c.cc.Invoke(ctx, "/messaging.GuploadService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuploadServiceServer is the server API for GuploadService service.
type GuploadServiceServer interface {
	Upload(GuploadService_UploadServer) error
	Init(context.Context, *InitReq) (*InitAck, error)
	Complete(context.Context, *CompleteReq) (*CompleteAck, error)
	Cancel(context.Context, *InitReq) (*InitAck, error)
}

// UnimplementedGuploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuploadServiceServer struct {
}

func (*UnimplementedGuploadServiceServer) Upload(srv GuploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedGuploadServiceServer) Init(ctx context.Context, req *InitReq) (*InitAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedGuploadServiceServer) Complete(ctx context.Context, req *CompleteReq) (*CompleteAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (*UnimplementedGuploadServiceServer) Cancel(ctx context.Context, req *InitReq) (*InitAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterGuploadServiceServer(s *grpc.Server, srv GuploadServiceServer) {
	s.RegisterService(&_GuploadService_serviceDesc, srv)
}

func _GuploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuploadServiceServer).Upload(&guploadServiceUploadServer{stream})
}

type GuploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type guploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *guploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GuploadService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuploadServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.GuploadService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuploadServiceServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuploadService_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuploadServiceServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.GuploadService/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuploadServiceServer).Complete(ctx, req.(*CompleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuploadService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuploadServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.GuploadService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuploadServiceServer).Cancel(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.GuploadService",
	HandlerType: (*GuploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _GuploadService_Init_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _GuploadService_Complete_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _GuploadService_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _GuploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
