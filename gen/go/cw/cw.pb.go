// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: cw/cw.proto

package cwv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_cw_cw_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_cw_cw_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_cw_cw_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_cw_cw_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IsAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IsAdminRequest) Reset() {
	*x = IsAdminRequest{}
	mi := &file_cw_cw_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminRequest) ProtoMessage() {}

func (x *IsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminRequest.ProtoReflect.Descriptor instead.
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{4}
}

func (x *IsAdminRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *IsAdminResponse) Reset() {
	*x = IsAdminResponse{}
	mi := &file_cw_cw_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminResponse) ProtoMessage() {}

func (x *IsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminResponse.ProtoReflect.Descriptor instead.
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{5}
}

func (x *IsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type ListPhotosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope string `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *ListPhotosRequest) Reset() {
	*x = ListPhotosRequest{}
	mi := &file_cw_cw_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPhotosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhotosRequest) ProtoMessage() {}

func (x *ListPhotosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhotosRequest.ProtoReflect.Descriptor instead.
func (*ListPhotosRequest) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{6}
}

func (x *ListPhotosRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type ListPhotosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhotoName string `protobuf:"bytes,1,opt,name=photo_name,json=photoName,proto3" json:"photo_name,omitempty"`
	Chunk     []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *ListPhotosResponse) Reset() {
	*x = ListPhotosResponse{}
	mi := &file_cw_cw_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPhotosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhotosResponse) ProtoMessage() {}

func (x *ListPhotosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cw_cw_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhotosResponse.ProtoReflect.Descriptor instead.
func (*ListPhotosResponse) Descriptor() ([]byte, []int) {
	return file_cw_cw_proto_rawDescGZIP(), []int{7}
}

func (x *ListPhotosResponse) GetPhotoName() string {
	if x != nil {
		return x.PhotoName
	}
	return ""
}

func (x *ListPhotosResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_cw_cw_proto protoreflect.FileDescriptor

var file_cw_cw_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x77, 0x2f, 0x63, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x22, 0x2b, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29,
	0x0a, 0x0e, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x69, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x22, 0x49, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xf1, 0x01,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x17, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x13, 0x5a, 0x11, 0x79, 0x6b, 0x77, 0x61, 0x69, 0x73, 0x2e, 0x63, 0x77, 0x2e, 0x76,
	0x31, 0x3b, 0x63, 0x77, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cw_cw_proto_rawDescOnce sync.Once
	file_cw_cw_proto_rawDescData = file_cw_cw_proto_rawDesc
)

func file_cw_cw_proto_rawDescGZIP() []byte {
	file_cw_cw_proto_rawDescOnce.Do(func() {
		file_cw_cw_proto_rawDescData = protoimpl.X.CompressGZIP(file_cw_cw_proto_rawDescData)
	})
	return file_cw_cw_proto_rawDescData
}

var file_cw_cw_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cw_cw_proto_goTypes = []any{
	(*RegisterResponse)(nil),   // 0: auth.RegisterResponse
	(*RegisterRequest)(nil),    // 1: auth.RegisterRequest
	(*LoginRequest)(nil),       // 2: auth.LoginRequest
	(*LoginResponse)(nil),      // 3: auth.LoginResponse
	(*IsAdminRequest)(nil),     // 4: auth.isAdminRequest
	(*IsAdminResponse)(nil),    // 5: auth.isAdminResponse
	(*ListPhotosRequest)(nil),  // 6: auth.ListPhotosRequest
	(*ListPhotosResponse)(nil), // 7: auth.ListPhotosResponse
}
var file_cw_cw_proto_depIdxs = []int32{
	1, // 0: auth.Service.Register:input_type -> auth.RegisterRequest
	2, // 1: auth.Service.Login:input_type -> auth.LoginRequest
	4, // 2: auth.Service.isAdmin:input_type -> auth.isAdminRequest
	6, // 3: auth.Service.ListPhotos:input_type -> auth.ListPhotosRequest
	0, // 4: auth.Service.Register:output_type -> auth.RegisterResponse
	3, // 5: auth.Service.Login:output_type -> auth.LoginResponse
	5, // 6: auth.Service.isAdmin:output_type -> auth.isAdminResponse
	7, // 7: auth.Service.ListPhotos:output_type -> auth.ListPhotosResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cw_cw_proto_init() }
func file_cw_cw_proto_init() {
	if File_cw_cw_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cw_cw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cw_cw_proto_goTypes,
		DependencyIndexes: file_cw_cw_proto_depIdxs,
		MessageInfos:      file_cw_cw_proto_msgTypes,
	}.Build()
	File_cw_cw_proto = out.File
	file_cw_cw_proto_rawDesc = nil
	file_cw_cw_proto_goTypes = nil
	file_cw_cw_proto_depIdxs = nil
}
