// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: cv_service/api/api.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cv_service_api_api_proto protoreflect.FileDescriptor

var file_cv_service_api_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x76, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x2f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x76, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa7, 0x04, 0x0a, 0x09, 0x43, 0x56, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x49, 0x43, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x49, 0x43, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x49, 0x43, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x87, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63,
	0x65, 0x12, 0x28, 0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x76,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x10, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x43, 0x49, 0x43, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27,
	0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x49, 0x43, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43,
	0x49, 0x43, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x3a, 0x01, 0x2a,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_cv_service_api_api_proto_goTypes = []interface{}{
	(*GetCICByUserIdRequest)(nil),     // 0: cv_service.api.GetCICByUserIdRequest
	(*RegisterUserFaceRequest)(nil),   // 1: cv_service.api.RegisterUserFaceRequest
	(*AuthorizeUserFaceRequest)(nil),  // 2: cv_service.api.AuthorizeUserFaceRequest
	(*UpsertCICForUserRequest)(nil),   // 3: cv_service.api.UpsertCICForUserRequest
	(*GetCICByUserIdResponse)(nil),    // 4: cv_service.api.GetCICByUserIdResponse
	(*RegisterUserFaceResponse)(nil),  // 5: cv_service.api.RegisterUserFaceResponse
	(*AuthorizeUserFaceResponse)(nil), // 6: cv_service.api.AuthorizeUserFaceResponse
	(*UpsertCICForUserResponse)(nil),  // 7: cv_service.api.UpsertCICForUserResponse
}
var file_cv_service_api_api_proto_depIdxs = []int32{
	0, // 0: cv_service.api.CVService.GetCICByUserId:input_type -> cv_service.api.GetCICByUserIdRequest
	1, // 1: cv_service.api.CVService.RegisterUserFace:input_type -> cv_service.api.RegisterUserFaceRequest
	2, // 2: cv_service.api.CVService.AuthorizeUserFace:input_type -> cv_service.api.AuthorizeUserFaceRequest
	3, // 3: cv_service.api.CVService.UpsertCICForUser:input_type -> cv_service.api.UpsertCICForUserRequest
	4, // 4: cv_service.api.CVService.GetCICByUserId:output_type -> cv_service.api.GetCICByUserIdResponse
	5, // 5: cv_service.api.CVService.RegisterUserFace:output_type -> cv_service.api.RegisterUserFaceResponse
	6, // 6: cv_service.api.CVService.AuthorizeUserFace:output_type -> cv_service.api.AuthorizeUserFaceResponse
	7, // 7: cv_service.api.CVService.UpsertCICForUser:output_type -> cv_service.api.UpsertCICForUserResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cv_service_api_api_proto_init() }
func file_cv_service_api_api_proto_init() {
	if File_cv_service_api_api_proto != nil {
		return
	}
	file_cv_service_api_data_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cv_service_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cv_service_api_api_proto_goTypes,
		DependencyIndexes: file_cv_service_api_api_proto_depIdxs,
	}.Build()
	File_cv_service_api_api_proto = out.File
	file_cv_service_api_api_proto_rawDesc = nil
	file_cv_service_api_api_proto_goTypes = nil
	file_cv_service_api_api_proto_depIdxs = nil
}
