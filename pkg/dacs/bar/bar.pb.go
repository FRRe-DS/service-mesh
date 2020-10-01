// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bar.proto

package bar

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Bar message
type Bar struct {
	// Id of the bar
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the Bar
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Bar
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{0}
}

func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bar) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// CreateBarRequest is a request for creating a Bar
type CreateBarRequest struct {
	// Bar to create
	Bar                  *Bar     `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBarRequest) Reset()         { *m = CreateBarRequest{} }
func (m *CreateBarRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBarRequest) ProtoMessage()    {}
func (*CreateBarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{1}
}

func (m *CreateBarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBarRequest.Unmarshal(m, b)
}
func (m *CreateBarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBarRequest.Marshal(b, m, deterministic)
}
func (m *CreateBarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBarRequest.Merge(m, src)
}
func (m *CreateBarRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBarRequest.Size(m)
}
func (m *CreateBarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBarRequest proto.InternalMessageInfo

func (m *CreateBarRequest) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

// CreateBarResponse is the response after creating a Bar
type CreateBarResponse struct {
	// Created bar
	Bar                  *Bar     `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBarResponse) Reset()         { *m = CreateBarResponse{} }
func (m *CreateBarResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBarResponse) ProtoMessage()    {}
func (*CreateBarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{2}
}

func (m *CreateBarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBarResponse.Unmarshal(m, b)
}
func (m *CreateBarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBarResponse.Marshal(b, m, deterministic)
}
func (m *CreateBarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBarResponse.Merge(m, src)
}
func (m *CreateBarResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBarResponse.Size(m)
}
func (m *CreateBarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBarResponse proto.InternalMessageInfo

func (m *CreateBarResponse) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

// ListBarsRequest is a request for listing bars
type ListBarsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBarsRequest) Reset()         { *m = ListBarsRequest{} }
func (m *ListBarsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBarsRequest) ProtoMessage()    {}
func (*ListBarsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{3}
}

func (m *ListBarsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBarsRequest.Unmarshal(m, b)
}
func (m *ListBarsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBarsRequest.Marshal(b, m, deterministic)
}
func (m *ListBarsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBarsRequest.Merge(m, src)
}
func (m *ListBarsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBarsRequest.Size(m)
}
func (m *ListBarsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBarsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBarsRequest proto.InternalMessageInfo

// ListBarsResponse contains the list of bars found
type ListBarsResponse struct {
	Bars                 []*Bar   `protobuf:"bytes,1,rep,name=bars,proto3" json:"bars,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBarsResponse) Reset()         { *m = ListBarsResponse{} }
func (m *ListBarsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBarsResponse) ProtoMessage()    {}
func (*ListBarsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{4}
}

func (m *ListBarsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBarsResponse.Unmarshal(m, b)
}
func (m *ListBarsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBarsResponse.Marshal(b, m, deterministic)
}
func (m *ListBarsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBarsResponse.Merge(m, src)
}
func (m *ListBarsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBarsResponse.Size(m)
}
func (m *ListBarsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBarsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBarsResponse proto.InternalMessageInfo

func (m *ListBarsResponse) GetBars() []*Bar {
	if m != nil {
		return m.Bars
	}
	return nil
}

// GetBarRequest is a request for getting an existing bar
type GetBarRequest struct {
	// Id of the bar
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBarRequest) Reset()         { *m = GetBarRequest{} }
func (m *GetBarRequest) String() string { return proto.CompactTextString(m) }
func (*GetBarRequest) ProtoMessage()    {}
func (*GetBarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{5}
}

func (m *GetBarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBarRequest.Unmarshal(m, b)
}
func (m *GetBarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBarRequest.Marshal(b, m, deterministic)
}
func (m *GetBarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBarRequest.Merge(m, src)
}
func (m *GetBarRequest) XXX_Size() int {
	return xxx_messageInfo_GetBarRequest.Size(m)
}
func (m *GetBarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBarRequest proto.InternalMessageInfo

func (m *GetBarRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// GetBarResponse contains the response with an existing bar
type GetBarResponse struct {
	// Found bar
	Bar                  *Bar     `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBarResponse) Reset()         { *m = GetBarResponse{} }
func (m *GetBarResponse) String() string { return proto.CompactTextString(m) }
func (*GetBarResponse) ProtoMessage()    {}
func (*GetBarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{6}
}

func (m *GetBarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBarResponse.Unmarshal(m, b)
}
func (m *GetBarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBarResponse.Marshal(b, m, deterministic)
}
func (m *GetBarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBarResponse.Merge(m, src)
}
func (m *GetBarResponse) XXX_Size() int {
	return xxx_messageInfo_GetBarResponse.Size(m)
}
func (m *GetBarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBarResponse proto.InternalMessageInfo

func (m *GetBarResponse) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

// UpdateBarRequest is a request for updating an existing bar
type UpdateBarRequest struct {
	// Id of the bar
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Bar to update
	Bar                  *Bar     `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBarRequest) Reset()         { *m = UpdateBarRequest{} }
func (m *UpdateBarRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBarRequest) ProtoMessage()    {}
func (*UpdateBarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{7}
}

func (m *UpdateBarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBarRequest.Unmarshal(m, b)
}
func (m *UpdateBarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBarRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBarRequest.Merge(m, src)
}
func (m *UpdateBarRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBarRequest.Size(m)
}
func (m *UpdateBarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBarRequest proto.InternalMessageInfo

func (m *UpdateBarRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateBarRequest) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

// UpdateBarResponse contains the response with the updated bar
type UpdateBarResponse struct {
	// Updated bar
	Bar                  *Bar     `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBarResponse) Reset()         { *m = UpdateBarResponse{} }
func (m *UpdateBarResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBarResponse) ProtoMessage()    {}
func (*UpdateBarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{8}
}

func (m *UpdateBarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBarResponse.Unmarshal(m, b)
}
func (m *UpdateBarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBarResponse.Marshal(b, m, deterministic)
}
func (m *UpdateBarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBarResponse.Merge(m, src)
}
func (m *UpdateBarResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBarResponse.Size(m)
}
func (m *UpdateBarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBarResponse proto.InternalMessageInfo

func (m *UpdateBarResponse) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

// DeleteBarRequest is a request for deleteing an existing bar
type DeleteBarRequest struct {
	// Id of the bar
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBarRequest) Reset()         { *m = DeleteBarRequest{} }
func (m *DeleteBarRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBarRequest) ProtoMessage()    {}
func (*DeleteBarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{9}
}

func (m *DeleteBarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBarRequest.Unmarshal(m, b)
}
func (m *DeleteBarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBarRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBarRequest.Merge(m, src)
}
func (m *DeleteBarRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBarRequest.Size(m)
}
func (m *DeleteBarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBarRequest proto.InternalMessageInfo

func (m *DeleteBarRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// DeleteBarResponse contains the response of the delete operation
type DeleteBarResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBarResponse) Reset()         { *m = DeleteBarResponse{} }
func (m *DeleteBarResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBarResponse) ProtoMessage()    {}
func (*DeleteBarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fca82f8db7cb01, []int{10}
}

func (m *DeleteBarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBarResponse.Unmarshal(m, b)
}
func (m *DeleteBarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBarResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBarResponse.Merge(m, src)
}
func (m *DeleteBarResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBarResponse.Size(m)
}
func (m *DeleteBarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBarResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Bar)(nil), "ar.edu.utn.frre.dacs.bar.Bar")
	proto.RegisterType((*CreateBarRequest)(nil), "ar.edu.utn.frre.dacs.bar.CreateBarRequest")
	proto.RegisterType((*CreateBarResponse)(nil), "ar.edu.utn.frre.dacs.bar.CreateBarResponse")
	proto.RegisterType((*ListBarsRequest)(nil), "ar.edu.utn.frre.dacs.bar.ListBarsRequest")
	proto.RegisterType((*ListBarsResponse)(nil), "ar.edu.utn.frre.dacs.bar.ListBarsResponse")
	proto.RegisterType((*GetBarRequest)(nil), "ar.edu.utn.frre.dacs.bar.GetBarRequest")
	proto.RegisterType((*GetBarResponse)(nil), "ar.edu.utn.frre.dacs.bar.GetBarResponse")
	proto.RegisterType((*UpdateBarRequest)(nil), "ar.edu.utn.frre.dacs.bar.UpdateBarRequest")
	proto.RegisterType((*UpdateBarResponse)(nil), "ar.edu.utn.frre.dacs.bar.UpdateBarResponse")
	proto.RegisterType((*DeleteBarRequest)(nil), "ar.edu.utn.frre.dacs.bar.DeleteBarRequest")
	proto.RegisterType((*DeleteBarResponse)(nil), "ar.edu.utn.frre.dacs.bar.DeleteBarResponse")
}

func init() { proto.RegisterFile("bar.proto", fileDescriptor_b8fca82f8db7cb01) }

var fileDescriptor_b8fca82f8db7cb01 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x8a, 0xd4, 0x40,
	0x14, 0x25, 0x49, 0x33, 0x74, 0xee, 0x30, 0x33, 0xe9, 0x52, 0x21, 0x04, 0xc5, 0x50, 0x1b, 0xdb,
	0x16, 0x2a, 0xcc, 0xb8, 0x73, 0x67, 0x66, 0xc4, 0x85, 0xae, 0x7a, 0x70, 0xe3, 0xae, 0xd2, 0x29,
	0x87, 0x82, 0x31, 0x89, 0xb7, 0xaa, 0x67, 0xd3, 0x08, 0xe2, 0x2f, 0xf8, 0x69, 0xfe, 0x82, 0x5f,
	0xe0, 0x17, 0x48, 0x9e, 0x26, 0x69, 0x92, 0x7e, 0xec, 0x9a, 0xea, 0x73, 0xcf, 0x39, 0xf7, 0xdc,
	0x43, 0xc0, 0x8e, 0x38, 0xb2, 0x0c, 0x53, 0x9d, 0x12, 0x97, 0x23, 0x13, 0xf1, 0x9a, 0xad, 0x75,
	0xc2, 0xbe, 0x20, 0x0a, 0x16, 0xf3, 0x95, 0x62, 0x11, 0x47, 0xef, 0xe9, 0x5d, 0x9a, 0xde, 0xdd,
	0x8b, 0x80, 0x67, 0x32, 0xe0, 0x49, 0x92, 0x6a, 0xae, 0x65, 0x9a, 0xa8, 0x72, 0x8e, 0x7e, 0x00,
	0x2b, 0xe4, 0x48, 0xce, 0xc1, 0x94, 0xb1, 0x6b, 0xf8, 0xc6, 0xdc, 0x5a, 0x9a, 0x32, 0x26, 0x04,
	0x26, 0x09, 0xff, 0x2a, 0x5c, 0xd3, 0x37, 0xe6, 0xf6, 0xb2, 0xf8, 0x4d, 0x7c, 0x38, 0x8d, 0x85,
	0x5a, 0xa1, 0xcc, 0x72, 0x02, 0xd7, 0x2a, 0xfe, 0x6a, 0x3f, 0xd1, 0x6b, 0x70, 0xae, 0x51, 0x70,
	0x2d, 0x42, 0x8e, 0x4b, 0xf1, 0x6d, 0x2d, 0x94, 0x26, 0x01, 0x58, 0x11, 0xc7, 0x82, 0xfa, 0xf4,
	0xea, 0x19, 0x1b, 0xb2, 0xc9, 0xf2, 0x91, 0x1c, 0x49, 0x6f, 0x60, 0xd6, 0x22, 0x51, 0x59, 0x9a,
	0x28, 0x71, 0x38, 0xcb, 0x0c, 0x2e, 0x3e, 0x4a, 0xa5, 0x43, 0x8e, 0xaa, 0x72, 0x42, 0xdf, 0x81,
	0xf3, 0xff, 0xa9, 0xe2, 0xbd, 0x84, 0x49, 0xc4, 0x51, 0xb9, 0x86, 0x6f, 0xed, 0x26, 0x2e, 0xa0,
	0xf4, 0x39, 0x9c, 0xbd, 0x17, 0xba, 0xb5, 0x61, 0x2f, 0x3b, 0xfa, 0x16, 0xce, 0x6b, 0xc0, 0xb1,
	0xee, 0x6f, 0xc1, 0xf9, 0x94, 0xc5, 0xdd, 0x20, 0xfb, 0x27, 0xaa, 0x48, 0xcd, 0x43, 0x82, 0x6d,
	0x91, 0x1e, 0x6b, 0x8d, 0x82, 0x73, 0x23, 0xee, 0xc5, 0x98, 0x35, 0xfa, 0x08, 0x66, 0x2d, 0x4c,
	0xa9, 0x74, 0xf5, 0x77, 0x02, 0x10, 0x72, 0xbc, 0x15, 0xf8, 0x20, 0x57, 0x82, 0x6c, 0xc0, 0x6e,
	0xce, 0x4c, 0x16, 0xc3, 0xc2, 0xfd, 0x42, 0x79, 0xaf, 0xf6, 0xc2, 0x96, 0xa2, 0xf4, 0xc9, 0xcf,
	0xdf, 0x7f, 0x7e, 0x99, 0x17, 0x74, 0x1a, 0x3c, 0x5c, 0x06, 0xf9, 0xf9, 0xde, 0xe4, 0x4b, 0x10,
	0x05, 0xd3, 0xba, 0x0a, 0xe4, 0xe5, 0x30, 0x5f, 0xaf, 0x41, 0xde, 0x62, 0x1f, 0x68, 0xa5, 0xec,
	0x14, 0xca, 0x40, 0x1a, 0x65, 0x82, 0x70, 0x52, 0xf6, 0x82, 0xbc, 0x18, 0xe6, 0xe9, 0x54, 0xcb,
	0x9b, 0xef, 0x06, 0x76, 0x17, 0x25, 0x67, 0xb5, 0x5c, 0xb0, 0x91, 0xf1, 0x77, 0xf2, 0xc3, 0x00,
	0xbb, 0x39, 0xfa, 0x58, 0xcc, 0xfd, 0xba, 0x8d, 0xc5, 0xbc, 0xd5, 0x22, 0xea, 0x15, 0xea, 0x8f,
	0xbd, 0xae, 0x7a, 0x99, 0xf5, 0x06, 0xec, 0xa6, 0x0c, 0x63, 0x0e, 0xfa, 0xad, 0x1a, 0x73, 0xb0,
	0xd5, 0xae, 0x7a, 0xff, 0x45, 0xd7, 0x41, 0x08, 0x9f, 0xa7, 0xf9, 0x50, 0xfe, 0x12, 0x9d, 0x14,
	0x5f, 0xbc, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0x81, 0x9b, 0x4c, 0x36, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BarServiceClient is the client API for BarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BarServiceClient interface {
	// Creates a new Bar in the system
	CreateBar(ctx context.Context, in *CreateBarRequest, opts ...grpc.CallOption) (*CreateBarResponse, error)
	// List the bars in the system
	ListBars(ctx context.Context, in *ListBarsRequest, opts ...grpc.CallOption) (*ListBarsResponse, error)
	// Get an existing bars
	GetBar(ctx context.Context, in *GetBarRequest, opts ...grpc.CallOption) (*GetBarResponse, error)
	// Updates an existing bars
	UpdateBar(ctx context.Context, in *UpdateBarRequest, opts ...grpc.CallOption) (*UpdateBarResponse, error)
	// Deletes an existing bar
	DeleteBar(ctx context.Context, in *DeleteBarRequest, opts ...grpc.CallOption) (*DeleteBarResponse, error)
}

type barServiceClient struct {
	cc *grpc.ClientConn
}

func NewBarServiceClient(cc *grpc.ClientConn) BarServiceClient {
	return &barServiceClient{cc}
}

func (c *barServiceClient) CreateBar(ctx context.Context, in *CreateBarRequest, opts ...grpc.CallOption) (*CreateBarResponse, error) {
	out := new(CreateBarResponse)
	err := c.cc.Invoke(ctx, "/ar.edu.utn.frre.dacs.bar.BarService/CreateBar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barServiceClient) ListBars(ctx context.Context, in *ListBarsRequest, opts ...grpc.CallOption) (*ListBarsResponse, error) {
	out := new(ListBarsResponse)
	err := c.cc.Invoke(ctx, "/ar.edu.utn.frre.dacs.bar.BarService/ListBars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barServiceClient) GetBar(ctx context.Context, in *GetBarRequest, opts ...grpc.CallOption) (*GetBarResponse, error) {
	out := new(GetBarResponse)
	err := c.cc.Invoke(ctx, "/ar.edu.utn.frre.dacs.bar.BarService/GetBar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barServiceClient) UpdateBar(ctx context.Context, in *UpdateBarRequest, opts ...grpc.CallOption) (*UpdateBarResponse, error) {
	out := new(UpdateBarResponse)
	err := c.cc.Invoke(ctx, "/ar.edu.utn.frre.dacs.bar.BarService/UpdateBar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *barServiceClient) DeleteBar(ctx context.Context, in *DeleteBarRequest, opts ...grpc.CallOption) (*DeleteBarResponse, error) {
	out := new(DeleteBarResponse)
	err := c.cc.Invoke(ctx, "/ar.edu.utn.frre.dacs.bar.BarService/DeleteBar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BarServiceServer is the server API for BarService service.
type BarServiceServer interface {
	// Creates a new Bar in the system
	CreateBar(context.Context, *CreateBarRequest) (*CreateBarResponse, error)
	// List the bars in the system
	ListBars(context.Context, *ListBarsRequest) (*ListBarsResponse, error)
	// Get an existing bars
	GetBar(context.Context, *GetBarRequest) (*GetBarResponse, error)
	// Updates an existing bars
	UpdateBar(context.Context, *UpdateBarRequest) (*UpdateBarResponse, error)
	// Deletes an existing bar
	DeleteBar(context.Context, *DeleteBarRequest) (*DeleteBarResponse, error)
}

func RegisterBarServiceServer(s *grpc.Server, srv BarServiceServer) {
	s.RegisterService(&_BarService_serviceDesc, srv)
}

func _BarService_CreateBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServiceServer).CreateBar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ar.edu.utn.frre.dacs.bar.BarService/CreateBar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServiceServer).CreateBar(ctx, req.(*CreateBarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarService_ListBars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServiceServer).ListBars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ar.edu.utn.frre.dacs.bar.BarService/ListBars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServiceServer).ListBars(ctx, req.(*ListBarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarService_GetBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServiceServer).GetBar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ar.edu.utn.frre.dacs.bar.BarService/GetBar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServiceServer).GetBar(ctx, req.(*GetBarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarService_UpdateBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServiceServer).UpdateBar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ar.edu.utn.frre.dacs.bar.BarService/UpdateBar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServiceServer).UpdateBar(ctx, req.(*UpdateBarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BarService_DeleteBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServiceServer).DeleteBar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ar.edu.utn.frre.dacs.bar.BarService/DeleteBar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServiceServer).DeleteBar(ctx, req.(*DeleteBarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ar.edu.utn.frre.dacs.bar.BarService",
	HandlerType: (*BarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBar",
			Handler:    _BarService_CreateBar_Handler,
		},
		{
			MethodName: "ListBars",
			Handler:    _BarService_ListBars_Handler,
		},
		{
			MethodName: "GetBar",
			Handler:    _BarService_GetBar_Handler,
		},
		{
			MethodName: "UpdateBar",
			Handler:    _BarService_UpdateBar_Handler,
		},
		{
			MethodName: "DeleteBar",
			Handler:    _BarService_DeleteBar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bar.proto",
}
