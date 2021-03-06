// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

//包名

package models

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

// 用户
type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Phones               []string `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetPhones() []string {
	if m != nil {
		return m.Phones
	}
	return nil
}

// 组
type UserGroup struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGroup) Reset()         { *m = UserGroup{} }
func (m *UserGroup) String() string { return proto.CompactTextString(m) }
func (*UserGroup) ProtoMessage()    {}
func (*UserGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *UserGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroup.Unmarshal(m, b)
}
func (m *UserGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroup.Marshal(b, m, deterministic)
}
func (m *UserGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroup.Merge(m, src)
}
func (m *UserGroup) XXX_Size() int {
	return xxx_messageInfo_UserGroup.Size(m)
}
func (m *UserGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroup proto.InternalMessageInfo

func (m *UserGroup) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserGroup) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "models.User")
	proto.RegisterType((*UserGroup)(nil), "models.UserGroup")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x56, 0x0a,
	0xe1, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x0b, 0xca, 0x53, 0x24, 0x98,
	0xc1, 0x42, 0x20, 0xa6, 0x90, 0x18, 0x17, 0x5b, 0x41, 0x46, 0x7e, 0x5e, 0x6a, 0xb1, 0x04, 0x8b,
	0x02, 0xb3, 0x06, 0x67, 0x10, 0x94, 0xa7, 0x14, 0xcc, 0xc5, 0x09, 0x32, 0xd5, 0xbd, 0x28, 0xbf,
	0xb4, 0x80, 0x28, 0xa3, 0x95, 0xb8, 0x58, 0x4b, 0x8b, 0x53, 0x8b, 0x8a, 0x25, 0x98, 0x15, 0x98,
	0x35, 0xb8, 0x8d, 0x78, 0xf4, 0x20, 0xce, 0xd3, 0x03, 0x99, 0x12, 0x04, 0x91, 0x4a, 0x62, 0x03,
	0xbb, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x05, 0x4d, 0x24, 0xc7, 0x00, 0x00, 0x00,
}
