// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_center.proto

package user_center

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

type CommonResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{0}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{1}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{2}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	EMail                string   `protobuf:"bytes,3,opt,name=e_mail,json=eMail,proto3" json:"e_mail,omitempty"`
	College              string   `protobuf:"bytes,4,opt,name=college,proto3" json:"college,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{3}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CreateUserRequest) GetEMail() string {
	if m != nil {
		return m.EMail
	}
	return ""
}

func (m *CreateUserRequest) GetCollege() string {
	if m != nil {
		return m.College
	}
	return ""
}

type CreateUserResponse struct {
	Basic                *CommonResponse `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`
	User                 *UserInfo       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{4}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetBasic() *CommonResponse {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *CreateUserResponse) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserInfoByIDRequest struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoByIDRequest) Reset()         { *m = GetUserInfoByIDRequest{} }
func (m *GetUserInfoByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoByIDRequest) ProtoMessage()    {}
func (*GetUserInfoByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{5}
}

func (m *GetUserInfoByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoByIDRequest.Unmarshal(m, b)
}
func (m *GetUserInfoByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoByIDRequest.Merge(m, src)
}
func (m *GetUserInfoByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoByIDRequest.Size(m)
}
func (m *GetUserInfoByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoByIDRequest proto.InternalMessageInfo

func (m *GetUserInfoByIDRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetUserInfoByIDResponse struct {
	Basic                *CommonResponse `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`
	User                 *UserInfo       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetUserInfoByIDResponse) Reset()         { *m = GetUserInfoByIDResponse{} }
func (m *GetUserInfoByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoByIDResponse) ProtoMessage()    {}
func (*GetUserInfoByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{6}
}

func (m *GetUserInfoByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoByIDResponse.Unmarshal(m, b)
}
func (m *GetUserInfoByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetUserInfoByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoByIDResponse.Merge(m, src)
}
func (m *GetUserInfoByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoByIDResponse.Size(m)
}
func (m *GetUserInfoByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoByIDResponse proto.InternalMessageInfo

func (m *GetUserInfoByIDResponse) GetBasic() *CommonResponse {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *GetUserInfoByIDResponse) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type UserInfo struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	College              string   `protobuf:"bytes,4,opt,name=college,proto3" json:"college,omitempty"`
	EMail                string   `protobuf:"bytes,5,opt,name=e_mail,json=eMail,proto3" json:"e_mail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{7}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetCollege() string {
	if m != nil {
		return m.College
	}
	return ""
}

func (m *UserInfo) GetEMail() string {
	if m != nil {
		return m.EMail
	}
	return ""
}

type VerifyUserIdentityRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyUserIdentityRequest) Reset()         { *m = VerifyUserIdentityRequest{} }
func (m *VerifyUserIdentityRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyUserIdentityRequest) ProtoMessage()    {}
func (*VerifyUserIdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{8}
}

func (m *VerifyUserIdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserIdentityRequest.Unmarshal(m, b)
}
func (m *VerifyUserIdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserIdentityRequest.Marshal(b, m, deterministic)
}
func (m *VerifyUserIdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserIdentityRequest.Merge(m, src)
}
func (m *VerifyUserIdentityRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyUserIdentityRequest.Size(m)
}
func (m *VerifyUserIdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserIdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserIdentityRequest proto.InternalMessageInfo

func (m *VerifyUserIdentityRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *VerifyUserIdentityRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type VerifyUserIdentityResponse struct {
	Basic                *CommonResponse `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`
	User                 *UserInfo       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VerifyUserIdentityResponse) Reset()         { *m = VerifyUserIdentityResponse{} }
func (m *VerifyUserIdentityResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyUserIdentityResponse) ProtoMessage()    {}
func (*VerifyUserIdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99b116ca098c6071, []int{9}
}

func (m *VerifyUserIdentityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserIdentityResponse.Unmarshal(m, b)
}
func (m *VerifyUserIdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserIdentityResponse.Marshal(b, m, deterministic)
}
func (m *VerifyUserIdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserIdentityResponse.Merge(m, src)
}
func (m *VerifyUserIdentityResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyUserIdentityResponse.Size(m)
}
func (m *VerifyUserIdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserIdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserIdentityResponse proto.InternalMessageInfo

func (m *VerifyUserIdentityResponse) GetBasic() *CommonResponse {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *VerifyUserIdentityResponse) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*CommonResponse)(nil), "CommonResponse")
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "CreateUserResponse")
	proto.RegisterType((*GetUserInfoByIDRequest)(nil), "GetUserInfoByIDRequest")
	proto.RegisterType((*GetUserInfoByIDResponse)(nil), "GetUserInfoByIDResponse")
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterType((*VerifyUserIdentityRequest)(nil), "VerifyUserIdentityRequest")
	proto.RegisterType((*VerifyUserIdentityResponse)(nil), "VerifyUserIdentityResponse")
}

func init() { proto.RegisterFile("user_center.proto", fileDescriptor_99b116ca098c6071) }

var fileDescriptor_99b116ca098c6071 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x9d, 0xd8, 0x2e, 0xcd, 0x94, 0xb6, 0x74, 0xf8, 0xa8, 0x31, 0x42, 0xaa, 0x16, 0x21,
	0x55, 0x20, 0xed, 0x21, 0x1c, 0xe0, 0x4c, 0x10, 0x90, 0x03, 0x1f, 0xda, 0x08, 0x0e, 0x5c, 0x22,
	0xc7, 0x9e, 0x80, 0x25, 0x7b, 0x37, 0x78, 0x1d, 0x45, 0xfe, 0x01, 0xfc, 0x6f, 0xe4, 0xb5, 0x9d,
	0x38, 0x38, 0xe6, 0x94, 0xdb, 0xcc, 0xbb, 0xe3, 0x99, 0xf1, 0xb3, 0xaf, 0x16, 0xae, 0xd6, 0x9a,
	0xb2, 0x79, 0x48, 0x32, 0xa7, 0x8c, 0xaf, 0x32, 0x95, 0x2b, 0xf6, 0x06, 0x2e, 0x26, 0x2a, 0x4d,
	0x95, 0x14, 0xa4, 0x57, 0x4a, 0x6a, 0x42, 0x04, 0x27, 0x54, 0x11, 0x79, 0x83, 0x9b, 0xc1, 0xad,
	0x2b, 0x4c, 0x5c, 0x6a, 0xb1, 0x5c, 0x2a, 0x6f, 0x78, 0x33, 0xb8, 0x1d, 0x09, 0x13, 0x33, 0x06,
	0x77, 0x3f, 0x52, 0x92, 0x28, 0x41, 0xbf, 0xd7, 0xa4, 0xf3, 0xb2, 0x46, 0x06, 0x69, 0xf5, 0xdd,
	0x48, 0x98, 0x98, 0x3d, 0x83, 0xf3, 0xba, 0x66, 0xd7, 0xbc, 0x53, 0x24, 0xe1, 0x6a, 0x92, 0x51,
	0x90, 0xd3, 0x37, 0x4d, 0xd9, 0x7f, 0xba, 0xe1, 0x03, 0x70, 0x57, 0xbf, 0x94, 0xa4, 0x7a, 0x8d,
	0x2a, 0xc1, 0x87, 0x70, 0x42, 0xf3, 0x34, 0x88, 0x13, 0xcf, 0xae, 0x64, 0xfa, 0x14, 0xc4, 0x09,
	0x7a, 0x70, 0x27, 0x54, 0x49, 0x42, 0x3f, 0xc9, 0x73, 0x8c, 0xde, 0xa4, 0xec, 0x07, 0x60, 0x7b,
	0x5e, 0xbd, 0xd9, 0x73, 0x70, 0x17, 0x81, 0x8e, 0x43, 0x33, 0xf1, 0x6c, 0x7c, 0xc9, 0xf7, 0xb1,
	0x88, 0xea, 0x14, 0x9f, 0x82, 0x53, 0x42, 0x34, 0x2b, 0x9c, 0x8d, 0x47, 0xbc, 0xec, 0x31, 0x95,
	0x4b, 0x25, 0x8c, 0xcc, 0x5e, 0xc0, 0xa3, 0x0f, 0x94, 0x37, 0xe2, 0xdb, 0x62, 0xfa, 0xae, 0xf9,
	0xa1, 0x7b, 0x60, 0xaf, 0xe3, 0xc8, 0x74, 0xb7, 0x45, 0x19, 0xb2, 0x39, 0x5c, 0x77, 0x6a, 0x8f,
	0xba, 0xcc, 0x06, 0x4e, 0x1b, 0xa5, 0x3b, 0x7e, 0x4b, 0x78, 0x78, 0x88, 0xb0, 0xdd, 0x26, 0xdc,
	0x8b, 0xb2, 0xc5, 0xde, 0x6d, 0xb1, 0x67, 0x33, 0x78, 0xfc, 0x9d, 0xb2, 0x78, 0x59, 0x98, 0xf1,
	0x11, 0xc9, 0x3c, 0xce, 0x8b, 0x06, 0x84, 0x5f, 0x6d, 0xf5, 0x79, 0x77, 0xbb, 0xdb, 0xbc, 0x3c,
	0xfb, 0x1a, 0x68, 0xbd, 0x51, 0x59, 0x54, 0xef, 0xb5, 0xcd, 0xd9, 0x02, 0xfc, 0x43, 0x4d, 0x8f,
	0x49, 0x6c, 0xfc, 0x67, 0x08, 0x4e, 0x29, 0xe1, 0x4b, 0x38, 0x9d, 0x05, 0x85, 0xf1, 0x2e, 0x9e,
	0xf3, 0xb6, 0xcf, 0xfd, 0x0b, 0xbe, 0x67, 0x69, 0x66, 0xe1, 0x6b, 0x80, 0x9d, 0xa1, 0x10, 0x79,
	0xc7, 0xcd, 0xfe, 0x7d, 0xde, 0x75, 0x1c, 0xb3, 0xf0, 0x3d, 0x5c, 0xfe, 0xe3, 0x00, 0xbc, 0xe6,
	0x87, 0xfd, 0xe3, 0x7b, 0xbc, 0xc7, 0x2c, 0xcc, 0xc2, 0x2f, 0x80, 0x5d, 0x34, 0xe8, 0xf3, 0xde,
	0x4b, 0xf0, 0x9f, 0xf0, 0x7e, 0x96, 0xcc, 0x5a, 0x9c, 0x98, 0xc7, 0xe1, 0xd5, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd0, 0x36, 0x99, 0x6d, 0x31, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserInfoByID(ctx context.Context, in *GetUserInfoByIDRequest, opts ...grpc.CallOption) (*GetUserInfoByIDResponse, error)
	VerifyUserIdentity(ctx context.Context, in *VerifyUserIdentityRequest, opts ...grpc.CallOption) (*VerifyUserIdentityResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/User/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInfoByID(ctx context.Context, in *GetUserInfoByIDRequest, opts ...grpc.CallOption) (*GetUserInfoByIDResponse, error) {
	out := new(GetUserInfoByIDResponse)
	err := c.cc.Invoke(ctx, "/User/GetUserInfoByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyUserIdentity(ctx context.Context, in *VerifyUserIdentityRequest, opts ...grpc.CallOption) (*VerifyUserIdentityResponse, error) {
	out := new(VerifyUserIdentityResponse)
	err := c.cc.Invoke(ctx, "/User/VerifyUserIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserInfoByID(context.Context, *GetUserInfoByIDRequest) (*GetUserInfoByIDResponse, error)
	VerifyUserIdentity(context.Context, *VerifyUserIdentityRequest) (*VerifyUserIdentityResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedUserServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServer) GetUserInfoByID(ctx context.Context, req *GetUserInfoByIDRequest) (*GetUserInfoByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByID not implemented")
}
func (*UnimplementedUserServer) VerifyUserIdentity(ctx context.Context, req *VerifyUserIdentityRequest) (*VerifyUserIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserIdentity not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInfoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserInfoByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfoByID(ctx, req.(*GetUserInfoByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyUserIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyUserIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/VerifyUserIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyUserIdentity(ctx, req.(*VerifyUserIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _User_SayHello_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "GetUserInfoByID",
			Handler:    _User_GetUserInfoByID_Handler,
		},
		{
			MethodName: "VerifyUserIdentity",
			Handler:    _User_VerifyUserIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_center.proto",
}
