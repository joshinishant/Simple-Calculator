// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: BasicCalculator.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CalculationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number1 int64 `protobuf:"varint,1,opt,name=Number1,proto3" json:"Number1,omitempty"`
	Number2 int64 `protobuf:"varint,2,opt,name=Number2,proto3" json:"Number2,omitempty"`
}

func (x *CalculationRequest) Reset() {
	*x = CalculationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BasicCalculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationRequest) ProtoMessage() {}

func (x *CalculationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BasicCalculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationRequest.ProtoReflect.Descriptor instead.
func (*CalculationRequest) Descriptor() ([]byte, []int) {
	return file_BasicCalculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculationRequest) GetNumber1() int64 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *CalculationRequest) GetNumber2() int64 {
	if x != nil {
		return x.Number2
	}
	return 0
}

type CalculatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *CalculatedResponse) Reset() {
	*x = CalculatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BasicCalculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatedResponse) ProtoMessage() {}

func (x *CalculatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BasicCalculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatedResponse.ProtoReflect.Descriptor instead.
func (*CalculatedResponse) Descriptor() ([]byte, []int) {
	return file_BasicCalculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatedResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_BasicCalculator_proto protoreflect.FileDescriptor

var file_BasicCalculator_proto_rawDesc = []byte{
	0x0a, 0x15, 0x42, 0x61, 0x73, 0x69, 0x63, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x32, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xf3, 0x02, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x54, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x54, 0x0a, 0x08, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x47, 0x52, 0x50, 0x43, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BasicCalculator_proto_rawDescOnce sync.Once
	file_BasicCalculator_proto_rawDescData = file_BasicCalculator_proto_rawDesc
)

func file_BasicCalculator_proto_rawDescGZIP() []byte {
	file_BasicCalculator_proto_rawDescOnce.Do(func() {
		file_BasicCalculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_BasicCalculator_proto_rawDescData)
	})
	return file_BasicCalculator_proto_rawDescData
}

var file_BasicCalculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_BasicCalculator_proto_goTypes = []interface{}{
	(*CalculationRequest)(nil), // 0: basiccalculator.CalculationRequest
	(*CalculatedResponse)(nil), // 1: basiccalculator.CalculatedResponse
}
var file_BasicCalculator_proto_depIdxs = []int32{
	0, // 0: basiccalculator.BasicCalculator.Addition:input_type -> basiccalculator.CalculationRequest
	0, // 1: basiccalculator.BasicCalculator.Substraction:input_type -> basiccalculator.CalculationRequest
	0, // 2: basiccalculator.BasicCalculator.Multiplication:input_type -> basiccalculator.CalculationRequest
	0, // 3: basiccalculator.BasicCalculator.Division:input_type -> basiccalculator.CalculationRequest
	1, // 4: basiccalculator.BasicCalculator.Addition:output_type -> basiccalculator.CalculatedResponse
	1, // 5: basiccalculator.BasicCalculator.Substraction:output_type -> basiccalculator.CalculatedResponse
	1, // 6: basiccalculator.BasicCalculator.Multiplication:output_type -> basiccalculator.CalculatedResponse
	1, // 7: basiccalculator.BasicCalculator.Division:output_type -> basiccalculator.CalculatedResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BasicCalculator_proto_init() }
func file_BasicCalculator_proto_init() {
	if File_BasicCalculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BasicCalculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BasicCalculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatedResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BasicCalculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BasicCalculator_proto_goTypes,
		DependencyIndexes: file_BasicCalculator_proto_depIdxs,
		MessageInfos:      file_BasicCalculator_proto_msgTypes,
	}.Build()
	File_BasicCalculator_proto = out.File
	file_BasicCalculator_proto_rawDesc = nil
	file_BasicCalculator_proto_goTypes = nil
	file_BasicCalculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BasicCalculatorClient is the client API for BasicCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasicCalculatorClient interface {
	Addition(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error)
	Substraction(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error)
	Multiplication(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error)
	Division(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error)
}

type basicCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicCalculatorClient(cc grpc.ClientConnInterface) BasicCalculatorClient {
	return &basicCalculatorClient{cc}
}

func (c *basicCalculatorClient) Addition(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error) {
	out := new(CalculatedResponse)
	err := c.cc.Invoke(ctx, "/basiccalculator.BasicCalculator/Addition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicCalculatorClient) Substraction(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error) {
	out := new(CalculatedResponse)
	err := c.cc.Invoke(ctx, "/basiccalculator.BasicCalculator/Substraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicCalculatorClient) Multiplication(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error) {
	out := new(CalculatedResponse)
	err := c.cc.Invoke(ctx, "/basiccalculator.BasicCalculator/Multiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicCalculatorClient) Division(ctx context.Context, in *CalculationRequest, opts ...grpc.CallOption) (*CalculatedResponse, error) {
	out := new(CalculatedResponse)
	err := c.cc.Invoke(ctx, "/basiccalculator.BasicCalculator/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicCalculatorServer is the server API for BasicCalculator service.
type BasicCalculatorServer interface {
	Addition(context.Context, *CalculationRequest) (*CalculatedResponse, error)
	Substraction(context.Context, *CalculationRequest) (*CalculatedResponse, error)
	Multiplication(context.Context, *CalculationRequest) (*CalculatedResponse, error)
	Division(context.Context, *CalculationRequest) (*CalculatedResponse, error)
}

// UnimplementedBasicCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedBasicCalculatorServer struct {
}

func (*UnimplementedBasicCalculatorServer) Addition(context.Context, *CalculationRequest) (*CalculatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addition not implemented")
}
func (*UnimplementedBasicCalculatorServer) Substraction(context.Context, *CalculationRequest) (*CalculatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Substraction not implemented")
}
func (*UnimplementedBasicCalculatorServer) Multiplication(context.Context, *CalculationRequest) (*CalculatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiplication not implemented")
}
func (*UnimplementedBasicCalculatorServer) Division(context.Context, *CalculationRequest) (*CalculatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Division not implemented")
}

func RegisterBasicCalculatorServer(s *grpc.Server, srv BasicCalculatorServer) {
	s.RegisterService(&_BasicCalculator_serviceDesc, srv)
}

func _BasicCalculator_Addition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicCalculatorServer).Addition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basiccalculator.BasicCalculator/Addition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicCalculatorServer).Addition(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicCalculator_Substraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicCalculatorServer).Substraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basiccalculator.BasicCalculator/Substraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicCalculatorServer).Substraction(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicCalculator_Multiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicCalculatorServer).Multiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basiccalculator.BasicCalculator/Multiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicCalculatorServer).Multiplication(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicCalculator_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicCalculatorServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basiccalculator.BasicCalculator/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicCalculatorServer).Division(ctx, req.(*CalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasicCalculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "basiccalculator.BasicCalculator",
	HandlerType: (*BasicCalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addition",
			Handler:    _BasicCalculator_Addition_Handler,
		},
		{
			MethodName: "Substraction",
			Handler:    _BasicCalculator_Substraction_Handler,
		},
		{
			MethodName: "Multiplication",
			Handler:    _BasicCalculator_Multiplication_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _BasicCalculator_Division_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BasicCalculator.proto",
}
