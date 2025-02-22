// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/http_loadbalancer/public_customapi.proto

// HTTP loadbalancer
//
// x-displayName: "Configure HTTP Loadbalancer"
// HTTP Loadbalancer view defines a required parameters that can be used in CRUD, to create and manage HTTP loadbalancer.
// It can be used to create HTTP loadbalancer and HTTPS loadbalancer.
//
// View will create following child objects.
//
// * Virtual-host
// * routes
// * clusters
// * endpoints
// * advertise policy
//

package http_loadbalancer

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	virtual_host_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Get DNS info Request
//
// x-displayName: "Get DNS Info Request"
// Request message for get-dns-info API
type GetDnsInfoRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Namespace for the HTTP loadbalancer
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "value"
	// Name of the HTTP loadbalancer
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetDnsInfoRequest) Reset()      { *m = GetDnsInfoRequest{} }
func (*GetDnsInfoRequest) ProtoMessage() {}
func (*GetDnsInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f78b47e1ecaa96, []int{0}
}
func (m *GetDnsInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDnsInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDnsInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDnsInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDnsInfoRequest.Merge(m, src)
}
func (m *GetDnsInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetDnsInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDnsInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDnsInfoRequest proto.InternalMessageInfo

func (m *GetDnsInfoRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetDnsInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GetDnsInfoResponse
//
// x-displayName: "Get DNS Info Response"
// Response for get-dns-info API
type GetDnsInfoResponse struct {
	// DNS information
	//
	// x-displayName: "DNS information"
	// DNS information object for this HTTP loadbalancer
	DnsInfo *virtual_host_dns_info.GlobalSpecType `protobuf:"bytes,1,opt,name=dns_info,json=dnsInfo,proto3" json:"dns_info,omitempty"`
}

func (m *GetDnsInfoResponse) Reset()      { *m = GetDnsInfoResponse{} }
func (*GetDnsInfoResponse) ProtoMessage() {}
func (*GetDnsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f78b47e1ecaa96, []int{1}
}
func (m *GetDnsInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDnsInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDnsInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDnsInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDnsInfoResponse.Merge(m, src)
}
func (m *GetDnsInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetDnsInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDnsInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDnsInfoResponse proto.InternalMessageInfo

func (m *GetDnsInfoResponse) GetDnsInfo() *virtual_host_dns_info.GlobalSpecType {
	if m != nil {
		return m.DnsInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDnsInfoRequest)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoRequest")
	golang_proto.RegisterType((*GetDnsInfoRequest)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoRequest")
	proto.RegisterType((*GetDnsInfoResponse)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoResponse")
	golang_proto.RegisterType((*GetDnsInfoResponse)(nil), "ves.io.schema.views.http_loadbalancer.GetDnsInfoResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/views/http_loadbalancer/public_customapi.proto", fileDescriptor_e0f78b47e1ecaa96)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/http_loadbalancer/public_customapi.proto", fileDescriptor_e0f78b47e1ecaa96)
}

var fileDescriptor_e0f78b47e1ecaa96 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0x05, 0x04, 0xc4, 0x2c, 0xd4, 0x53, 0x14, 0xa2, 0x13, 0x8a, 0x84, 0xc4, 0x80, 0xef,
	0x50, 0xbb, 0x00, 0x62, 0xe1, 0x47, 0x15, 0x75, 0x01, 0x14, 0x98, 0x60, 0x88, 0xce, 0xf6, 0x8b,
	0x63, 0x70, 0xee, 0x1d, 0xbe, 0xb3, 0x69, 0x85, 0x2a, 0xa1, 0xfe, 0x05, 0x48, 0x88, 0xff, 0x81,
	0x9d, 0x05, 0xa9, 0x4b, 0x37, 0x3a, 0xa1, 0x08, 0x96, 0x8e, 0xc4, 0x61, 0x80, 0xad, 0x7f, 0x02,
	0xca, 0x39, 0xfd, 0x41, 0x82, 0x50, 0xd5, 0xed, 0x3d, 0x7f, 0xfa, 0xbe, 0x7b, 0xef, 0x7b, 0x9f,
	0xdd, 0x3b, 0x05, 0x68, 0x96, 0x20, 0xd7, 0xe1, 0x00, 0x86, 0x82, 0x17, 0x09, 0xbc, 0xd6, 0x7c,
	0x60, 0x8c, 0xea, 0xa5, 0x28, 0xa2, 0x40, 0xa4, 0x42, 0x86, 0x90, 0x71, 0x95, 0x07, 0x69, 0x12,
	0xf6, 0xc2, 0x5c, 0x1b, 0x1c, 0x0a, 0x95, 0x30, 0x95, 0xa1, 0x41, 0xef, 0x6a, 0xc5, 0x66, 0x15,
	0x9b, 0x59, 0x36, 0x5b, 0x60, 0x37, 0xfd, 0x38, 0x31, 0x83, 0x3c, 0x60, 0x21, 0x0e, 0x79, 0x8c,
	0x31, 0x72, 0xcb, 0x0e, 0xf2, 0xbe, 0xed, 0x6c, 0x63, 0xab, 0x4a, 0xb5, 0xd9, 0x8a, 0x11, 0xe3,
	0x14, 0xb8, 0x50, 0x09, 0x17, 0x52, 0xa2, 0x11, 0x26, 0x41, 0xa9, 0x67, 0xe8, 0xe5, 0xbf, 0x27,
	0x46, 0x75, 0x1c, 0x6c, 0xcd, 0xad, 0x23, 0xd2, 0x24, 0x12, 0x06, 0x66, 0x68, 0x7b, 0x0e, 0x05,
	0x0d, 0xb2, 0x98, 0x53, 0xb8, 0x31, 0x6f, 0x48, 0x66, 0x72, 0x91, 0xf6, 0x06, 0xa8, 0x4d, 0x2f,
	0x92, 0xba, 0x97, 0xc8, 0x3e, 0x72, 0x0c, 0x5e, 0x40, 0x68, 0x2a, 0x46, 0x7b, 0xd5, 0x5d, 0xea,
	0x80, 0x79, 0x20, 0xf5, 0x9a, 0xec, 0x63, 0x17, 0x5e, 0xe5, 0xa0, 0x8d, 0xd7, 0x72, 0xeb, 0x52,
	0x0c, 0x41, 0x2b, 0x11, 0x42, 0x83, 0x5c, 0x21, 0xd7, 0xea, 0xdd, 0xa3, 0x0f, 0x9e, 0xe7, 0x9e,
	0x9d, 0x36, 0x8d, 0x9a, 0x05, 0x6c, 0xdd, 0x8e, 0x5c, 0xef, 0xb8, 0x8c, 0x56, 0x28, 0x35, 0x78,
	0x0f, 0xdd, 0x0b, 0x07, 0xaf, 0x5a, 0x99, 0x8b, 0xcb, 0x2b, 0x6c, 0xde, 0xf4, 0x7f, 0x4c, 0xc8,
	0x3a, 0x29, 0x06, 0x22, 0x7d, 0xa2, 0x20, 0x7c, 0xba, 0xa1, 0xa0, 0x7b, 0x3e, 0xaa, 0x74, 0x97,
	0x3f, 0xd5, 0xdc, 0xfa, 0x7d, 0x7b, 0xc5, 0xbb, 0x8f, 0xd7, 0xbc, 0xdf, 0xc4, 0x75, 0x8f, 0x1e,
	0xf5, 0x6e, 0xb2, 0x13, 0xdd, 0x93, 0x2d, 0xac, 0xdb, 0xbc, 0x75, 0x0a, 0x66, 0xb5, 0x61, 0xbb,
	0xbf, 0xfb, 0xb9, 0x46, 0xca, 0x2f, 0x8d, 0xa5, 0x02, 0xb4, 0x9f, 0xa0, 0xaf, 0x32, 0x5c, 0xdf,
	0xf0, 0x33, 0x10, 0xd1, 0xd6, 0xf7, 0x9f, 0xef, 0x6b, 0x1d, 0x6f, 0x75, 0x96, 0x3e, 0x7e, 0xe8,
	0x9f, 0xe6, 0x6f, 0x0e, 0xeb, 0xcd, 0xc5, 0xb4, 0xce, 0xe0, 0x4d, 0x1e, 0x83, 0xf1, 0x23, 0xa9,
	0xfd, 0xa9, 0x23, 0xcd, 0xdb, 0x3b, 0xdb, 0xe4, 0xcc, 0xb7, 0x6d, 0x72, 0xfd, 0x64, 0x93, 0x3e,
	0xb2, 0x27, 0xde, 0xfa, 0xda, 0xa8, 0x5d, 0x22, 0xf7, 0x3e, 0x90, 0xd1, 0x98, 0x3a, 0x7b, 0x63,
	0xea, 0xec, 0x8f, 0x29, 0x79, 0x5b, 0x52, 0xf2, 0xb1, 0xa4, 0x64, 0xb7, 0xa4, 0x64, 0x54, 0x52,
	0xf2, 0xa3, 0xa4, 0xe4, 0x57, 0x49, 0x9d, 0xfd, 0x92, 0x92, 0x77, 0x13, 0xea, 0xec, 0x4c, 0x28,
	0x19, 0x4d, 0xa8, 0xb3, 0x37, 0xa1, 0xce, 0xb3, 0xe7, 0x31, 0xaa, 0x97, 0x31, 0x2b, 0x30, 0x35,
	0x90, 0x65, 0x82, 0xe5, 0x9a, 0xdb, 0xa2, 0x8f, 0xd9, 0x70, 0xba, 0x6b, 0x91, 0x44, 0x90, 0xf9,
	0x07, 0x30, 0x57, 0x41, 0x8c, 0x1c, 0xd6, 0xcd, 0x2c, 0x7e, 0xff, 0xff, 0x2d, 0x83, 0x73, 0x36,
	0x81, 0x2b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x87, 0x93, 0x40, 0x13, 0xc6, 0x03, 0x00, 0x00,
}

func (this *GetDnsInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDnsInfoRequest)
	if !ok {
		that2, ok := that.(GetDnsInfoRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *GetDnsInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDnsInfoResponse)
	if !ok {
		that2, ok := that.(GetDnsInfoResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DnsInfo.Equal(that1.DnsInfo) {
		return false
	}
	return true
}
func (this *GetDnsInfoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&http_loadbalancer.GetDnsInfoRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetDnsInfoResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&http_loadbalancer.GetDnsInfoResponse{")
	if this.DnsInfo != nil {
		s = append(s, "DnsInfo: "+fmt.Sprintf("%#v", this.DnsInfo)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
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
	// GetDnsInfo
	//
	// x-displayName: "Get DNS Info"
	// GetDnsInfo is an API to get DNS information for a given HTTP loadbalancer
	GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error) {
	out := new(GetDnsInfoResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.views.http_loadbalancer.CustomAPI/GetDnsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// GetDnsInfo
	//
	// x-displayName: "Get DNS Info"
	// GetDnsInfo is an API to get DNS information for a given HTTP loadbalancer
	GetDnsInfo(context.Context, *GetDnsInfoRequest) (*GetDnsInfoResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) GetDnsInfo(ctx context.Context, req *GetDnsInfoRequest) (*GetDnsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDnsInfo not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_GetDnsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDnsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetDnsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.views.http_loadbalancer.CustomAPI/GetDnsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetDnsInfo(ctx, req.(*GetDnsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.views.http_loadbalancer.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDnsInfo",
			Handler:    _CustomAPI_GetDnsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/views/http_loadbalancer/public_customapi.proto",
}

func (m *GetDnsInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDnsInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDnsInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetDnsInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDnsInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDnsInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DnsInfo != nil {
		{
			size, err := m.DnsInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetDnsInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *GetDnsInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DnsInfo != nil {
		l = m.DnsInfo.Size()
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetDnsInfoRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDnsInfoRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetDnsInfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDnsInfoResponse{`,
		`DnsInfo:` + strings.Replace(fmt.Sprintf("%v", this.DnsInfo), "GlobalSpecType", "virtual_host_dns_info.GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetDnsInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDnsInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDnsInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetDnsInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDnsInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDnsInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsInfo == nil {
				m.DnsInfo = &virtual_host_dns_info.GlobalSpecType{}
			}
			if err := m.DnsInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPublicCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapi = fmt.Errorf("proto: unexpected end of group")
)
