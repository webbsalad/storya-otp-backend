// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: api/otp/otp_service.proto

package otp

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendOtpRequest) Reset() {
	*x = SendOtpRequest{}
	mi := &file_api_otp_otp_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOtpRequest) ProtoMessage() {}

func (x *SendOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_otp_otp_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOtpRequest.ProtoReflect.Descriptor instead.
func (*SendOtpRequest) Descriptor() ([]byte, []int) {
	return file_api_otp_otp_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendOtpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ConfirmOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	OtpCode string `protobuf:"bytes,2,opt,name=otp_code,json=otpCode,proto3" json:"otp_code,omitempty"`
}

func (x *ConfirmOtpRequest) Reset() {
	*x = ConfirmOtpRequest{}
	mi := &file_api_otp_otp_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmOtpRequest) ProtoMessage() {}

func (x *ConfirmOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_otp_otp_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmOtpRequest.ProtoReflect.Descriptor instead.
func (*ConfirmOtpRequest) Descriptor() ([]byte, []int) {
	return file_api_otp_otp_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmOtpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ConfirmOtpRequest) GetOtpCode() string {
	if x != nil {
		return x.OtpCode
	}
	return ""
}

type ConfirmOtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailId string `protobuf:"bytes,1,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
}

func (x *ConfirmOtpResponse) Reset() {
	*x = ConfirmOtpResponse{}
	mi := &file_api_otp_otp_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmOtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmOtpResponse) ProtoMessage() {}

func (x *ConfirmOtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_otp_otp_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmOtpResponse.ProtoReflect.Descriptor instead.
func (*ConfirmOtpResponse) Descriptor() ([]byte, []int) {
	return file_api_otp_otp_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmOtpResponse) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

var File_api_otp_otp_service_proto protoreflect.FileDescriptor

var file_api_otp_otp_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x74, 0x70,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4d, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x74, 0x70,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x74, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x39, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x32,
	0xb2, 0x01, 0x0a, 0x0a, 0x4f, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x12, 0x13, 0x2e, 0x6f, 0x74, 0x70, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2d, 0x6f, 0x74, 0x70, 0x12, 0x56, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x74, 0x70, 0x12, 0x16, 0x2e, 0x6f, 0x74, 0x70,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x2d, 0x6f, 0x74, 0x70, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x62, 0x73, 0x61, 0x6c, 0x61, 0x64, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x61, 0x2d, 0x6f, 0x74, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x6f, 0x74, 0x70, 0x3b, 0x6f, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_otp_otp_service_proto_rawDescOnce sync.Once
	file_api_otp_otp_service_proto_rawDescData = file_api_otp_otp_service_proto_rawDesc
)

func file_api_otp_otp_service_proto_rawDescGZIP() []byte {
	file_api_otp_otp_service_proto_rawDescOnce.Do(func() {
		file_api_otp_otp_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_otp_otp_service_proto_rawDescData)
	})
	return file_api_otp_otp_service_proto_rawDescData
}

var file_api_otp_otp_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_otp_otp_service_proto_goTypes = []any{
	(*SendOtpRequest)(nil),     // 0: otp.SendOtpRequest
	(*ConfirmOtpRequest)(nil),  // 1: otp.ConfirmOtpRequest
	(*ConfirmOtpResponse)(nil), // 2: otp.ConfirmOtpResponse
	(*emptypb.Empty)(nil),      // 3: google.protobuf.Empty
}
var file_api_otp_otp_service_proto_depIdxs = []int32{
	0, // 0: otp.OtpService.SendOtp:input_type -> otp.SendOtpRequest
	1, // 1: otp.OtpService.ConfirmOtp:input_type -> otp.ConfirmOtpRequest
	3, // 2: otp.OtpService.SendOtp:output_type -> google.protobuf.Empty
	2, // 3: otp.OtpService.ConfirmOtp:output_type -> otp.ConfirmOtpResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_otp_otp_service_proto_init() }
func file_api_otp_otp_service_proto_init() {
	if File_api_otp_otp_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_otp_otp_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_otp_otp_service_proto_goTypes,
		DependencyIndexes: file_api_otp_otp_service_proto_depIdxs,
		MessageInfos:      file_api_otp_otp_service_proto_msgTypes,
	}.Build()
	File_api_otp_otp_service_proto = out.File
	file_api_otp_otp_service_proto_rawDesc = nil
	file_api_otp_otp_service_proto_goTypes = nil
	file_api_otp_otp_service_proto_depIdxs = nil
}
