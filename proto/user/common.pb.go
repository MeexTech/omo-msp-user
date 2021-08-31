// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/common.proto

package user

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

type ResultCode int32

const (
	ResultCode_Success     ResultCode = 0
	ResultCode_MaxLimit    ResultCode = 1
	ResultCode_Repeated    ResultCode = 2
	ResultCode_NotExisted  ResultCode = 3
	ResultCode_DBException ResultCode = 4
	ResultCode_Empty       ResultCode = 5
)

var ResultCode_name = map[int32]string{
	0: "Success",
	1: "MaxLimit",
	2: "Repeated",
	3: "NotExisted",
	4: "DBException",
	5: "Empty",
}

var ResultCode_value = map[string]int32{
	"Success":     0,
	"MaxLimit":    1,
	"Repeated":    2,
	"NotExisted":  3,
	"DBException": 4,
	"Empty":       5,
}

func (x ResultCode) String() string {
	return proto.EnumName(ResultCode_name, int32(x))
}

func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{0}
}

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Entity               string   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{0}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestIDInfo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestIDInfo) Reset()         { *m = RequestIDInfo{} }
func (m *RequestIDInfo) String() string { return proto.CompactTextString(m) }
func (*RequestIDInfo) ProtoMessage()    {}
func (*RequestIDInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{1}
}

func (m *RequestIDInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestIDInfo.Unmarshal(m, b)
}
func (m *RequestIDInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestIDInfo.Marshal(b, m, deterministic)
}
func (m *RequestIDInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestIDInfo.Merge(m, src)
}
func (m *RequestIDInfo) XXX_Size() int {
	return xxx_messageInfo_RequestIDInfo.Size(m)
}
func (m *RequestIDInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestIDInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestIDInfo proto.InternalMessageInfo

func (m *RequestIDInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestIDInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestPage struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPage) Reset()         { *m = RequestPage{} }
func (m *RequestPage) String() string { return proto.CompactTextString(m) }
func (*RequestPage) ProtoMessage()    {}
func (*RequestPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{2}
}

func (m *RequestPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPage.Unmarshal(m, b)
}
func (m *RequestPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPage.Marshal(b, m, deterministic)
}
func (m *RequestPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPage.Merge(m, src)
}
func (m *RequestPage) XXX_Size() int {
	return xxx_messageInfo_RequestPage.Size(m)
}
func (m *RequestPage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPage proto.InternalMessageInfo

func (m *RequestPage) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestPage) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type RequestList struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	List                 []string `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{3}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyStatus struct {
	Code                 ResultCode `protobuf:"varint,1,opt,name=code,proto3,enum=omo.msp.user.ResultCode" json:"code,omitempty"`
	Msg                  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{4}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() ResultCode {
	if m != nil {
		return m.Code
	}
	return ResultCode_Success
}

func (m *ReplyStatus) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ReplyInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Uid                  string       `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_27de64fbdd9e62a7, []int{5}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterEnum("omo.msp.user.ResultCode", ResultCode_name, ResultCode_value)
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.user.RequestInfo")
	proto.RegisterType((*RequestIDInfo)(nil), "omo.msp.user.RequestIDInfo")
	proto.RegisterType((*RequestPage)(nil), "omo.msp.user.RequestPage")
	proto.RegisterType((*RequestList)(nil), "omo.msp.user.RequestList")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.user.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.user.ReplyInfo")
}

func init() {
	proto.RegisterFile("proto/user/common.proto", fileDescriptor_27de64fbdd9e62a7)
}

var fileDescriptor_27de64fbdd9e62a7 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x95, 0xb6, 0x20, 0x9d, 0x02, 0x36, 0x7b, 0xd0, 0xea, 0x89, 0xf4, 0x44, 0x8c, 0x29, 0x11,
	0x4e, 0xc6, 0x1b, 0xc2, 0xc1, 0x04, 0x0c, 0x59, 0x6e, 0xde, 0x4a, 0x3b, 0x92, 0x4d, 0xd8, 0xee,
	0xda, 0xdd, 0x26, 0xf0, 0xf7, 0x66, 0x37, 0xa5, 0x28, 0xf1, 0xf6, 0xde, 0xcc, 0xbc, 0xf7, 0x66,
	0x36, 0x0b, 0x77, 0xb2, 0x14, 0x5a, 0x8c, 0x2b, 0x85, 0xe5, 0x38, 0x13, 0x9c, 0x8b, 0x22, 0xb1,
	0x15, 0xd2, 0x13, 0x5c, 0x24, 0x5c, 0xc9, 0xc4, 0xb4, 0xe2, 0x0d, 0x04, 0x14, 0xbf, 0x2b, 0x54,
	0xfa, 0xbd, 0xf8, 0x12, 0x24, 0x04, 0xb7, 0x62, 0x79, 0xd4, 0x1a, 0xb6, 0x46, 0x3e, 0x35, 0x90,
	0xdc, 0x42, 0x07, 0x0b, 0xcd, 0xf4, 0x31, 0x72, 0x6c, 0xb1, 0x66, 0xe4, 0x01, 0xba, 0x42, 0x62,
	0x99, 0x6a, 0x51, 0x46, 0xae, 0xed, 0x34, 0x3c, 0x7e, 0x85, 0xfe, 0xc9, 0x74, 0x6e, 0x6d, 0x07,
	0xe0, 0xd4, 0xae, 0x1e, 0x75, 0x58, 0xfe, 0x47, 0xec, 0x5c, 0x88, 0x5f, 0x9a, 0x8d, 0xd6, 0xe9,
	0x0e, 0x4d, 0x7e, 0x51, 0xf1, 0x2d, 0x96, 0x56, 0xde, 0xa7, 0x35, 0x23, 0x04, 0x3c, 0x99, 0xee,
	0xd0, 0xca, 0xfb, 0xd4, 0xe2, 0x78, 0xda, 0x48, 0x97, 0x4c, 0xe9, 0x7f, 0x8e, 0x21, 0xe0, 0xed,
	0x99, 0xd2, 0x91, 0x33, 0x74, 0x47, 0x3e, 0xb5, 0x38, 0x5e, 0x19, 0x91, 0xdc, 0x1f, 0x37, 0x3a,
	0xd5, 0x95, 0x22, 0x4f, 0xe0, 0x65, 0x22, 0x47, 0xab, 0x1a, 0x4c, 0xa2, 0xe4, 0xf7, 0x6b, 0x25,
	0x14, 0x55, 0xb5, 0xd7, 0x6f, 0x22, 0x47, 0x6a, 0xa7, 0x4c, 0x04, 0x57, 0xbb, 0xfa, 0x06, 0x03,
	0xe3, 0x35, 0xf8, 0xd6, 0xce, 0xde, 0xfd, 0x0c, 0x1d, 0x65, 0x6d, 0xad, 0x5d, 0x30, 0xb9, 0xbf,
	0xb4, 0x6b, 0x72, 0x69, 0x3d, 0x78, 0x5a, 0xda, 0x69, 0x96, 0x7e, 0x4c, 0x01, 0xce, 0xb9, 0x24,
	0x80, 0xeb, 0x4d, 0x95, 0x65, 0xa8, 0x54, 0x78, 0x45, 0x7a, 0xd0, 0x5d, 0xa5, 0x87, 0x25, 0xe3,
	0x4c, 0x87, 0x2d, 0xc3, 0x28, 0x4a, 0x4c, 0x35, 0xe6, 0xa1, 0x43, 0x06, 0x00, 0x1f, 0x42, 0x2f,
	0x0e, 0x4c, 0x19, 0xee, 0x92, 0x1b, 0x08, 0xe6, 0xb3, 0xc5, 0x21, 0x43, 0xa9, 0x99, 0x28, 0x42,
	0x8f, 0xf8, 0xd0, 0x5e, 0x70, 0xa9, 0x8f, 0x61, 0x7b, 0xd6, 0xfb, 0x84, 0xf3, 0x77, 0xd9, 0x76,
	0x2c, 0x9e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xac, 0x61, 0xc9, 0x03, 0x43, 0x02, 0x00, 0x00,
}
