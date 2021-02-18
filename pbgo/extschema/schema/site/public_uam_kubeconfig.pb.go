// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/public_uam_kubeconfig.proto

package site

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/googleapis/google/api"
import google_api2 "google.golang.org/genproto/googleapis/api/httpbody"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UamKubeConfigAPI service

type UamKubeConfigAPIClient interface {
	// Create K8s Cluster Global Kube Config
	//
	// x-displayName: "Create K8s Cluster Global Kube Config"
	// Down load kube config for global k8s cluster access
	CreateGlobalKubeConfig(ctx context.Context, in *CreateKubeConfigReq, opts ...grpc.CallOption) (*google_api2.HttpBody, error)
	// List Local Kube Configs
	//
	// x-displayName: "List Global Kube Configs"
	// Returns list of all global active kubeconfig minted for this site
	ListGlobalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error)
}

type uamKubeConfigAPIClient struct {
	cc *grpc.ClientConn
}

func NewUamKubeConfigAPIClient(cc *grpc.ClientConn) UamKubeConfigAPIClient {
	return &uamKubeConfigAPIClient{cc}
}

func (c *uamKubeConfigAPIClient) CreateGlobalKubeConfig(ctx context.Context, in *CreateKubeConfigReq, opts ...grpc.CallOption) (*google_api2.HttpBody, error) {
	out := new(google_api2.HttpBody)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.UamKubeConfigAPI/CreateGlobalKubeConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uamKubeConfigAPIClient) ListGlobalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error) {
	out := new(ListKubeConfigRsp)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.UamKubeConfigAPI/ListGlobalKubeConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UamKubeConfigAPI service

type UamKubeConfigAPIServer interface {
	// Create K8s Cluster Global Kube Config
	//
	// x-displayName: "Create K8s Cluster Global Kube Config"
	// Down load kube config for global k8s cluster access
	CreateGlobalKubeConfig(context.Context, *CreateKubeConfigReq) (*google_api2.HttpBody, error)
	// List Local Kube Configs
	//
	// x-displayName: "List Global Kube Configs"
	// Returns list of all global active kubeconfig minted for this site
	ListGlobalKubeConfig(context.Context, *ListKubeConfigReq) (*ListKubeConfigRsp, error)
}

func RegisterUamKubeConfigAPIServer(s *grpc.Server, srv UamKubeConfigAPIServer) {
	s.RegisterService(&_UamKubeConfigAPI_serviceDesc, srv)
}

func _UamKubeConfigAPI_CreateGlobalKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKubeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UamKubeConfigAPIServer).CreateGlobalKubeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.UamKubeConfigAPI/CreateGlobalKubeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UamKubeConfigAPIServer).CreateGlobalKubeConfig(ctx, req.(*CreateKubeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UamKubeConfigAPI_ListGlobalKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UamKubeConfigAPIServer).ListGlobalKubeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.UamKubeConfigAPI/ListGlobalKubeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UamKubeConfigAPIServer).ListGlobalKubeConfig(ctx, req.(*ListKubeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UamKubeConfigAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.UamKubeConfigAPI",
	HandlerType: (*UamKubeConfigAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGlobalKubeConfig",
			Handler:    _UamKubeConfigAPI_CreateGlobalKubeConfig_Handler,
		},
		{
			MethodName: "ListGlobalKubeConfig",
			Handler:    _UamKubeConfigAPI_ListGlobalKubeConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/public_uam_kubeconfig.proto",
}

func init() {
	proto.RegisterFile("ves.io/schema/site/public_uam_kubeconfig.proto", fileDescriptorPublicUamKubeconfig)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/public_uam_kubeconfig.proto", fileDescriptorPublicUamKubeconfig)
}

var fileDescriptorPublicUamKubeconfig = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x29, 0x78, 0xc8, 0xa9, 0x2c, 0x45, 0xda, 0x28, 0x73, 0x08, 0x48, 0x41, 0xd8,
	0x19, 0xd0, 0x9b, 0xa0, 0x68, 0x7a, 0xb0, 0xa2, 0xa0, 0x08, 0x5e, 0x44, 0xac, 0x33, 0x9b, 0x97,
	0xc9, 0xd8, 0xdd, 0x7d, 0xe3, 0xce, 0xec, 0x62, 0x94, 0x82, 0xf4, 0x13, 0x08, 0x7e, 0x09, 0xbf,
	0x80, 0x20, 0xf6, 0xd2, 0x83, 0x60, 0x4f, 0x52, 0xf4, 0xe2, 0xd1, 0x4c, 0x3d, 0x78, 0xec, 0x47,
	0x90, 0xce, 0x26, 0xa6, 0x21, 0x11, 0x84, 0xde, 0xde, 0xdb, 0xff, 0x6f, 0x1f, 0xff, 0xff, 0xbc,
	0xd7, 0x62, 0x15, 0x58, 0xa6, 0x91, 0xdb, 0x64, 0x00, 0x99, 0xe0, 0x56, 0x3b, 0xe0, 0xa6, 0x94,
	0xa9, 0x4e, 0xb6, 0x4a, 0x91, 0x6d, 0x6d, 0x97, 0x12, 0x12, 0xcc, 0xfb, 0x5a, 0x31, 0x53, 0xa0,
	0xc3, 0x28, 0xaa, 0x79, 0x56, 0xf3, 0xec, 0x84, 0x6f, 0xc7, 0x4a, 0xbb, 0x41, 0x29, 0x59, 0x82,
	0x19, 0x57, 0xa8, 0x90, 0x07, 0x54, 0x96, 0xfd, 0xd0, 0x85, 0x26, 0x54, 0xf5, 0x88, 0xf6, 0x45,
	0x85, 0xa8, 0x52, 0xe0, 0xc2, 0x68, 0x2e, 0xf2, 0x1c, 0x9d, 0x70, 0x1a, 0x73, 0x3b, 0x56, 0xd7,
	0x4e, 0xa9, 0x03, 0xe7, 0x8c, 0xc4, 0xde, 0x70, 0x2c, 0x5d, 0x98, 0xf5, 0x8a, 0xe6, 0xf4, 0x7f,
	0x74, 0x41, 0x10, 0x37, 0x34, 0x30, 0xd1, 0x3b, 0xb3, 0x7a, 0x05, 0x16, 0xf2, 0x6a, 0x76, 0xc6,
	0x95, 0x4f, 0x4b, 0xad, 0xe5, 0x47, 0x22, 0xbb, 0x5b, 0x4a, 0xd8, 0x08, 0xa1, 0x6f, 0x3d, 0xb8,
	0x13, 0x7d, 0x20, 0xad, 0xf3, 0x1b, 0x05, 0x08, 0x07, 0xb7, 0x53, 0x94, 0x22, 0x9d, 0xaa, 0xd1,
	0x3a, 0x9b, 0x7f, 0x0d, 0x56, 0xb3, 0x53, 0xea, 0x21, 0xbc, 0x68, 0xaf, 0xb0, 0x3a, 0x15, 0x13,
	0x46, 0xb3, 0x4d, 0xe7, 0x4c, 0x17, 0x7b, 0xc3, 0xce, 0x13, 0xff, 0x65, 0x75, 0xb9, 0x02, 0x1b,
	0x6b, 0x8c, 0xb5, 0x10, 0x36, 0x4e, 0x84, 0xb0, 0xbb, 0xdf, 0x7f, 0xbd, 0x6b, 0x76, 0x3b, 0xd7,
	0xc7, 0x5b, 0xe0, 0xb9, 0xc8, 0xc0, 0x1a, 0x91, 0x80, 0xe5, 0xaf, 0xff, 0xd6, 0x3b, 0x21, 0xdf,
	0xf8, 0xcb, 0x0e, 0x57, 0xc1, 0x58, 0x3c, 0x5d, 0xd6, 0x35, 0x72, 0x39, 0xfa, 0x4c, 0x5a, 0x2b,
	0xf7, 0xb4, 0x75, 0x73, 0xae, 0x2f, 0x2d, 0x72, 0x7d, 0x42, 0xce, 0x7a, 0xfe, 0x1f, 0xcc, 0x9a,
	0xce, 0xd3, 0x83, 0x8f, 0x4d, 0xf2, 0xcf, 0x20, 0x37, 0xa3, 0x1b, 0x67, 0x0a, 0x62, 0xdb, 0xeb,
	0xfb, 0x7b, 0x64, 0xe9, 0xdb, 0x1e, 0x59, 0x5b, 0xe0, 0xe6, 0xbe, 0x7c, 0x0e, 0x89, 0xdb, 0xfd,
	0xba, 0xda, 0x7c, 0x46, 0xba, 0xaf, 0x0e, 0x47, 0xb4, 0xf1, 0x63, 0x44, 0x1b, 0xc7, 0x23, 0x4a,
	0xde, 0x78, 0x4a, 0xde, 0x7b, 0x4a, 0x0e, 0x3c, 0x25, 0x87, 0x9e, 0x92, 0x9f, 0x9e, 0x92, 0xdf,
	0x9e, 0x36, 0x8e, 0x3d, 0x25, 0x6f, 0x8f, 0x68, 0x63, 0xff, 0x88, 0x92, 0xc7, 0x9b, 0x0a, 0xcd,
	0xb6, 0x62, 0x15, 0xa6, 0x0e, 0x8a, 0x42, 0xb0, 0xd2, 0xf2, 0x50, 0xf4, 0xb1, 0xc8, 0x62, 0x53,
	0x60, 0xa5, 0x7b, 0x50, 0xc4, 0x13, 0x99, 0x1b, 0xa9, 0x90, 0xc3, 0x4b, 0x37, 0xb9, 0xb0, 0xe9,
	0xa1, 0xc9, 0x73, 0xe1, 0x7e, 0xae, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x19, 0xcc, 0x1c, 0x91,
	0x4e, 0x03, 0x00, 0x00,
}
