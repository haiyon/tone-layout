// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/interface/v1/greeter.proto

package iV1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v1 "sample/api/schema/v1"
	v11 "sample/api/shared/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_interface_v1_greeter_proto protoreflect.FileDescriptor

var file_api_interface_v1_greeter_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xea, 0x02, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x43, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x43, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x42, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0xa6, 0x01, 0x5a, 0x1b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x69, 0x56, 0x31, 0x92, 0x41, 0x85, 0x01, 0x12, 0x82, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x73, 0x22,
	0x24, 0x0a, 0x0e, 0x54, 0x6f, 0x6e, 0x65, 0x20, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6d,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x3e, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2d,
	0x32, 0x2e, 0x30, 0x12, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2d, 0x32, 0x2e, 0x30,
	0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x32, 0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_api_interface_v1_greeter_proto_goTypes = []interface{}{
	(*v1.GreeterRequest)(nil),      // 0: schema.v1.GreeterRequest
	(*v1.GetGreeterRequest)(nil),   // 1: schema.v1.GetGreeterRequest
	(*v1.ListGreetersRequest)(nil), // 2: schema.v1.ListGreetersRequest
	(*v1.GreeterReply)(nil),        // 3: schema.v1.GreeterReply
	(*v11.Response)(nil),           // 4: shared.v1.Response
	(*v1.ListGreetersReply)(nil),   // 5: schema.v1.ListGreetersReply
}
var file_api_interface_v1_greeter_proto_depIdxs = []int32{
	0, // 0: sample.v1.Greeter.CreateGreeter:input_type -> schema.v1.GreeterRequest
	1, // 1: sample.v1.Greeter.GetGreeter:input_type -> schema.v1.GetGreeterRequest
	0, // 2: sample.v1.Greeter.UpdateGreeter:input_type -> schema.v1.GreeterRequest
	1, // 3: sample.v1.Greeter.DeleteGreeter:input_type -> schema.v1.GetGreeterRequest
	2, // 4: sample.v1.Greeter.ListGreeters:input_type -> schema.v1.ListGreetersRequest
	3, // 5: sample.v1.Greeter.CreateGreeter:output_type -> schema.v1.GreeterReply
	3, // 6: sample.v1.Greeter.GetGreeter:output_type -> schema.v1.GreeterReply
	3, // 7: sample.v1.Greeter.UpdateGreeter:output_type -> schema.v1.GreeterReply
	4, // 8: sample.v1.Greeter.DeleteGreeter:output_type -> shared.v1.Response
	5, // 9: sample.v1.Greeter.ListGreeters:output_type -> schema.v1.ListGreetersReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_interface_v1_greeter_proto_init() }
func file_api_interface_v1_greeter_proto_init() {
	if File_api_interface_v1_greeter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_interface_v1_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_interface_v1_greeter_proto_goTypes,
		DependencyIndexes: file_api_interface_v1_greeter_proto_depIdxs,
	}.Build()
	File_api_interface_v1_greeter_proto = out.File
	file_api_interface_v1_greeter_proto_rawDesc = nil
	file_api_interface_v1_greeter_proto_goTypes = nil
	file_api_interface_v1_greeter_proto_depIdxs = nil
}