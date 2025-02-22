// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/stored_object/public_customapi.proto

// Object Store APIs
//
// x-displayName: "Stored Object"
//
//

package stored_object

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("ves.io/schema/stored_object/public_customapi.proto", fileDescriptor_1c246fd1a6e8dd5b)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/stored_object/public_customapi.proto", fileDescriptor_1c246fd1a6e8dd5b)
}

var fileDescriptor_1c246fd1a6e8dd5b = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xbf, 0x4f, 0x1b, 0x4b,
	0x10, 0xc7, 0xbd, 0xf0, 0x1e, 0x12, 0xf7, 0xa8, 0xae, 0xf2, 0x33, 0xe8, 0x9e, 0xe4, 0xe6, 0xd9,
	0x8e, 0xf6, 0x96, 0x1f, 0x55, 0xd2, 0x44, 0x01, 0x24, 0x14, 0x09, 0xe5, 0x57, 0x41, 0x81, 0x85,
	0xd0, 0xfa, 0x3c, 0x1c, 0x0b, 0xbe, 0xdb, 0xcb, 0xee, 0xfa, 0x48, 0xb0, 0x2d, 0x45, 0x74, 0x51,
	0x1a, 0xa4, 0xf0, 0x07, 0x44, 0x4a, 0x13, 0xa5, 0x48, 0x1d, 0x89, 0x86, 0x0e, 0xaa, 0x08, 0x25,
	0x0d, 0x25, 0x3e, 0xa7, 0x88, 0xa2, 0x14, 0x48, 0xf9, 0x07, 0xa2, 0xec, 0x9d, 0x89, 0x8d, 0xd0,
	0xc9, 0x28, 0x14, 0xa9, 0x3c, 0xab, 0xf9, 0x7c, 0xc7, 0xf3, 0x9d, 0x9b, 0xbb, 0x35, 0xa6, 0x43,
	0x90, 0x36, 0xe3, 0x44, 0x3a, 0xeb, 0xe0, 0x51, 0x22, 0x15, 0x17, 0x50, 0x5d, 0xe5, 0x95, 0x0d,
	0x70, 0x14, 0x09, 0xea, 0x95, 0x1a, 0x73, 0x56, 0x9d, 0xba, 0x54, 0xdc, 0xa3, 0x01, 0xb3, 0x03,
	0xc1, 0x15, 0x37, 0xc7, 0x63, 0x8d, 0x1d, 0x6b, 0xec, 0x3e, 0x4d, 0x0e, 0xbb, 0x4c, 0xad, 0xd7,
	0x2b, 0xb6, 0xc3, 0x3d, 0xe2, 0x72, 0x97, 0x13, 0xad, 0xa9, 0xd4, 0xd7, 0xf4, 0x49, 0x1f, 0x74,
	0x14, 0xd7, 0xca, 0x4d, 0xb8, 0x9c, 0xbb, 0x35, 0x20, 0x34, 0x60, 0x84, 0xfa, 0x3e, 0x57, 0x54,
	0x31, 0xee, 0xcb, 0x24, 0xfb, 0x5f, 0x92, 0x3d, 0xaf, 0xa1, 0x98, 0x07, 0x52, 0x51, 0x2f, 0x48,
	0x80, 0xf1, 0xfe, 0xf6, 0x79, 0xd0, 0xab, 0xb6, 0xd3, 0xbc, 0x39, 0xdc, 0xf3, 0xb8, 0xbf, 0xaa,
	0x9e, 0x06, 0xd0, 0xe5, 0xff, 0x4f, 0xe3, 0x7b, 0xc1, 0x89, 0x7e, 0x30, 0xa4, 0x35, 0x56, 0xa5,
	0x0a, 0x92, 0x6c, 0xfe, 0x42, 0x16, 0x24, 0xf8, 0x61, 0x7f, 0x6b, 0xd3, 0xdf, 0x0d, 0x63, 0x74,
	0x4e, 0x8f, 0xf5, 0xce, 0x83, 0xbb, 0xe6, 0x29, 0x32, 0xc6, 0xe6, 0x04, 0x50, 0x05, 0xf7, 0xf5,
	0x9f, 0x99, 0x93, 0x76, 0xca, 0x88, 0xed, 0x5e, 0xf4, 0x11, 0x3c, 0xae, 0x83, 0x54, 0xb9, 0xa9,
	0x2b, 0x28, 0x64, 0xc0, 0x7d, 0x09, 0xf9, 0xcd, 0xe8, 0x30, 0x9b, 0x0b, 0x41, 0x62, 0xc6, 0x71,
	0x8c, 0xe3, 0x18, 0xc7, 0x5b, 0x82, 0x29, 0xd8, 0xf9, 0xf4, 0xf9, 0xe5, 0xd0, 0x42, 0x6e, 0x36,
	0xd9, 0x02, 0xe2, 0x53, 0x0f, 0x64, 0x40, 0x1d, 0x90, 0xa4, 0x71, 0x1e, 0xb7, 0xfa, 0xa7, 0x24,
	0x49, 0x23, 0x0e, 0xf4, 0x5c, 0x5b, 0x31, 0xd8, 0xba, 0x85, 0x4a, 0x66, 0x1b, 0x19, 0xa3, 0x0b,
	0xa0, 0x12, 0x7f, 0x38, 0xb5, 0xdb, 0x73, 0xae, 0x6b, 0xce, 0x1e, 0x14, 0x4f, 0x9c, 0xd5, 0x8f,
	0xde, 0x0f, 0xa1, 0xe8, 0x30, 0xfb, 0xef, 0xa5, 0xee, 0x04, 0xd0, 0xaa, 0x36, 0x77, 0xcf, 0x5c,
	0xfc, 0x7d, 0x73, 0xa4, 0x11, 0x82, 0x90, 0x8c, 0xfb, 0x2d, 0xf3, 0x08, 0x19, 0xff, 0x2c, 0x32,
	0x99, 0x74, 0x23, 0x4d, 0x92, 0xda, 0x76, 0x0f, 0xd9, 0xf5, 0x39, 0x39, 0xb8, 0x20, 0x71, 0xba,
	0x34, 0x98, 0x53, 0x62, 0xe2, 0x2b, 0x39, 0x35, 0x5f, 0xfd, 0x6d, 0x8c, 0xcd, 0x43, 0x0d, 0x06,
	0xdc, 0xc8, 0x5e, 0x74, 0xb0, 0x8d, 0xec, 0x57, 0x24, 0x6e, 0xde, 0xfe, 0x95, 0xbe, 0x92, 0x5f,
	0x0f, 0xb3, 0xdf, 0x86, 0x8c, 0x7c, 0x17, 0xd0, 0x75, 0x2f, 0x70, 0x55, 0x5d, 0xd7, 0xdc, 0x43,
	0xc6, 0x48, 0x12, 0x3e, 0x47, 0x97, 0x58, 0x2f, 0x94, 0x29, 0xde, 0x5e, 0x29, 0x94, 0x31, 0xc5,
	0xdb, 0x93, 0xf8, 0xe6, 0x4a, 0xa9, 0x9c, 0x04, 0xc5, 0xdb, 0xc5, 0x8b, 0x4f, 0xbd, 0x20, 0xb7,
	0xa8, 0xeb, 0x82, 0x68, 0xba, 0xe0, 0x83, 0x60, 0x4e, 0xd3, 0x01, 0xa1, 0xd8, 0x1a, 0x73, 0xa8,
	0x82, 0xe6, 0xba, 0xf2, 0x6a, 0xcd, 0x0d, 0x1a, 0x52, 0xe9, 0x08, 0x16, 0xa8, 0x62, 0x6a, 0x65,
	0xf3, 0xdd, 0xaf, 0xb6, 0x5e, 0xff, 0x39, 0x6d, 0x91, 0x42, 0xd8, 0x5c, 0x2a, 0x96, 0x7f, 0x9e,
	0x6e, 0x14, 0xb0, 0xfe, 0x6d, 0x4c, 0xb7, 0x8a, 0x8d, 0x99, 0x96, 0x5e, 0x9e, 0x3d, 0x54, 0xba,
	0xd6, 0xf7, 0x64, 0x79, 0xbe, 0x74, 0x0d, 0x1f, 0x95, 0xdc, 0xd4, 0xc1, 0x3e, 0x1a, 0xfe, 0xb8,
	0x8f, 0xf2, 0x69, 0x6b, 0x16, 0x2f, 0xd8, 0xce, 0x87, 0xec, 0xf0, 0x2e, 0x42, 0xb3, 0x2f, 0xd0,
	0x71, 0xdb, 0xca, 0x9c, 0xb4, 0xad, 0xcc, 0x59, 0xdb, 0x42, 0xcf, 0x22, 0x0b, 0xbd, 0x89, 0x2c,
	0x74, 0x14, 0x59, 0xe8, 0x38, 0xb2, 0xd0, 0x69, 0x64, 0xa1, 0x2f, 0x91, 0x95, 0x39, 0x8b, 0x2c,
	0xb4, 0xdb, 0xb1, 0x32, 0x07, 0x1d, 0x0b, 0x1d, 0x77, 0xac, 0xcc, 0x49, 0xc7, 0xca, 0x2c, 0x3f,
	0x74, 0x79, 0xb0, 0xe9, 0xda, 0x21, 0xaf, 0x29, 0x10, 0x82, 0xda, 0x75, 0x49, 0x74, 0xb0, 0xc6,
	0x85, 0x87, 0x03, 0xc1, 0x43, 0x56, 0x05, 0x81, 0xbb, 0x69, 0x12, 0x54, 0x5c, 0x4e, 0xe0, 0x89,
	0xea, 0x5e, 0x25, 0x97, 0xdc, 0x28, 0x95, 0x11, 0x7d, 0x15, 0xcc, 0xfc, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xce, 0xa3, 0x96, 0x79, 0x83, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomAPIClient is the client API for CustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomAPIClient interface {
	// CreateObject
	//
	// x-displayName: "Create Stored Object"
	// CreateObject is an API to upload an object to generic object store. Objects are immutable, a new version is created when the content is updated.
	CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...grpc.CallOption) (*CreateObjectResponse, error)
	// GetObject
	//
	// x-displayName: "Get Stored Object"
	// GetObject is an API to download an object from object store
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	// ListObjects
	//
	// x-displayName: "Get List Of Stored Objects"
	// ListObjects is an API to list objects in object store
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error)
	// DeleteObjects
	//
	// x-displayName: "Delete Stored Object(s)"
	// DeleteObjects is an API to delete object(s) in object store
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...grpc.CallOption) (*CreateObjectResponse, error) {
	out := new(CreateObjectResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.stored_object.CustomAPI/CreateObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.stored_object.CustomAPI/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error) {
	out := new(ListObjectsResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.stored_object.CustomAPI/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.stored_object.CustomAPI/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// CreateObject
	//
	// x-displayName: "Create Stored Object"
	// CreateObject is an API to upload an object to generic object store. Objects are immutable, a new version is created when the content is updated.
	CreateObject(context.Context, *CreateObjectRequest) (*CreateObjectResponse, error)
	// GetObject
	//
	// x-displayName: "Get Stored Object"
	// GetObject is an API to download an object from object store
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	// ListObjects
	//
	// x-displayName: "Get List Of Stored Objects"
	// ListObjects is an API to list objects in object store
	ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error)
	// DeleteObjects
	//
	// x-displayName: "Delete Stored Object(s)"
	// DeleteObjects is an API to delete object(s) in object store
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) CreateObject(ctx context.Context, req *CreateObjectRequest) (*CreateObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObject not implemented")
}
func (*UnimplementedCustomAPIServer) GetObject(ctx context.Context, req *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (*UnimplementedCustomAPIServer) ListObjects(ctx context.Context, req *ListObjectsRequest) (*ListObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (*UnimplementedCustomAPIServer) DeleteObject(ctx context.Context, req *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_CreateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).CreateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.stored_object.CustomAPI/CreateObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).CreateObject(ctx, req.(*CreateObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.stored_object.CustomAPI/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.stored_object.CustomAPI/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).ListObjects(ctx, req.(*ListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.stored_object.CustomAPI/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.stored_object.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObject",
			Handler:    _CustomAPI_CreateObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _CustomAPI_GetObject_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _CustomAPI_ListObjects_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _CustomAPI_DeleteObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/stored_object/public_customapi.proto",
}
