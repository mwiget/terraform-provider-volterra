// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/api_credential/types.proto

package api_credential

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// API Credential type
//
// x-displayName: "Credential Type"
// Types of API credential given when requesting credentials from volterra
type APICredentialType int32

const (
	// x-displayName: "User Certificate"
	// Volterra user certificate to access Volterra public API using mTLS
	// using self credential (my credential)
	API_CERTIFICATE APICredentialType = 0
	// x-displayName: "Kubernetes Config File"
	// Kubernetes config file to access Virtual Kubernetes API in Volterra
	// using self credential (my credential)
	KUBE_CONFIG APICredentialType = 1
	// x-displayName: "API Token"
	// API token to access Volterra public API
	// using self credential (my credential)
	API_TOKEN APICredentialType = 2
	// x-displayName: "Service API Token"
	// API token for service credentials
	// using service user credential (service credential)
	SERVICE_API_TOKEN APICredentialType = 3
	// x-displayName: "Service API Certificate"
	// API certificate for service credentials
	// using service user credential (service credential)
	SERVICE_API_CERTIFICATE APICredentialType = 4
	// x-displayName: " Kubernetes Config File for Service Credential"
	// Service Credential kubeconfig
	// using service user credential (service credential)
	SERVICE_KUBE_CONFIG APICredentialType = 5
	// x-displayName: "Site Global Kubeconfig"
	// Kubeconfig for accessing Site via Global Controller
	SITE_GLOBAL_KUBE_CONFIG APICredentialType = 6
)

var APICredentialType_name = map[int32]string{
	0: "API_CERTIFICATE",
	1: "KUBE_CONFIG",
	2: "API_TOKEN",
	3: "SERVICE_API_TOKEN",
	4: "SERVICE_API_CERTIFICATE",
	5: "SERVICE_KUBE_CONFIG",
	6: "SITE_GLOBAL_KUBE_CONFIG",
}

var APICredentialType_value = map[string]int32{
	"API_CERTIFICATE":         0,
	"KUBE_CONFIG":             1,
	"API_TOKEN":               2,
	"SERVICE_API_TOKEN":       3,
	"SERVICE_API_CERTIFICATE": 4,
	"SERVICE_KUBE_CONFIG":     5,
	"SITE_GLOBAL_KUBE_CONFIG": 6,
}

func (APICredentialType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7c0ad745f7a8c3b, []int{0}
}

// API credentials
//
// x-displayName: "API Credentials"
// Keeps track of user requested API credentials
type GlobalSpecType struct {
	// API Credential type
	//
	// x-displayName: "Credential Type"
	// Type of API credential
	Type APICredentialType `protobuf:"varint,1,opt,name=type,proto3,enum=ves.io.schema.api_credential.APICredentialType" json:"type,omitempty"`
	// user requesting credential
	//
	// x-displayName: "User"
	// Reference to user for whom API credential is created
	Users []*schema.ObjectRefType `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	// Virtual K8s namespace
	//
	// x-displayName: "Namespace"
	// Namespace of virtual_k8s
	VirtualK8SNamespace string `protobuf:"bytes,3,opt,name=virtual_k8s_namespace,json=virtualK8sNamespace,proto3" json:"virtual_k8s_namespace,omitempty"`
	// Virtual K8s
	//
	// x-displayName: "Virtual K8s"
	// Name of virtual K8s cluster
	VirtualK8SName string `protobuf:"bytes,4,opt,name=virtual_k8s_name,json=virtualK8sName,proto3" json:"virtual_k8s_name,omitempty"`
	// Digest sha1
	//
	// x-displayName: "Digest sha1"
	// Digest sha1 of credential
	Digest string `protobuf:"bytes,5,opt,name=digest,proto3" json:"digest,omitempty"`
	// Created timestamp
	//
	// x-displayName: "Created timestamp"
	// Timestamp of credential creation
	CreatedTimestamp *types.Timestamp `protobuf:"bytes,6,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp,omitempty"`
	// Expiry timestamp
	//
	// x-displayName: "Expiry timestamp"
	// Timestamp of credential expiration
	ExpirationTimestamp *types.Timestamp `protobuf:"bytes,7,opt,name=expiration_timestamp,json=expirationTimestamp,proto3" json:"expiration_timestamp,omitempty"`
	// Active
	//
	// x-displayName: "Active"
	// Possibility to deactivate/activate credential with no deletion
	Active bool `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
	// Certificate Serial Number
	//
	// x-displayName: "Certificate Serial Number"
	// Serial number of the client certificate part of credential type API Certificate or Kubeconfig
	CertificateSerialNum string `protobuf:"bytes,9,opt,name=certificate_serial_num,json=certificateSerialNum,proto3" json:"certificate_serial_num,omitempty"`
	// Site Name
	//
	// x-displayName: "Site Name"
	// Site name when global kubeconfig is issued for physical k8s site
	SiteName string `protobuf:"bytes,10,opt,name=site_name,json=siteName,proto3" json:"site_name,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7c0ad745f7a8c3b, []int{0}
}
func (m *GlobalSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GlobalSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSpecType.Merge(m, src)
}
func (m *GlobalSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSpecType proto.InternalMessageInfo

func (m *GlobalSpecType) GetType() APICredentialType {
	if m != nil {
		return m.Type
	}
	return API_CERTIFICATE
}

func (m *GlobalSpecType) GetUsers() []*schema.ObjectRefType {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GlobalSpecType) GetVirtualK8SNamespace() string {
	if m != nil {
		return m.VirtualK8SNamespace
	}
	return ""
}

func (m *GlobalSpecType) GetVirtualK8SName() string {
	if m != nil {
		return m.VirtualK8SName
	}
	return ""
}

func (m *GlobalSpecType) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *GlobalSpecType) GetCreatedTimestamp() *types.Timestamp {
	if m != nil {
		return m.CreatedTimestamp
	}
	return nil
}

func (m *GlobalSpecType) GetExpirationTimestamp() *types.Timestamp {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return nil
}

func (m *GlobalSpecType) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *GlobalSpecType) GetCertificateSerialNum() string {
	if m != nil {
		return m.CertificateSerialNum
	}
	return ""
}

func (m *GlobalSpecType) GetSiteName() string {
	if m != nil {
		return m.SiteName
	}
	return ""
}

func init() {
	proto.RegisterEnum("ves.io.schema.api_credential.APICredentialType", APICredentialType_name, APICredentialType_value)
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.api_credential.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/api_credential/types.proto", fileDescriptor_a7c0ad745f7a8c3b)
}

var fileDescriptor_a7c0ad745f7a8c3b = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0x73, 0xf9, 0xf7, 0x6b, 0xae, 0xfa, 0xb5, 0xc9, 0xa5, 0x7f, 0x4c, 0x5a, 0x1d, 0x11,
	0x53, 0x84, 0x54, 0x5b, 0x0a, 0x4b, 0x37, 0x94, 0x44, 0x6e, 0x14, 0xb5, 0x24, 0x95, 0x1b, 0x18,
	0x58, 0xac, 0x8b, 0x73, 0x71, 0x8f, 0xda, 0x39, 0xeb, 0x7c, 0x89, 0xda, 0x01, 0x89, 0x17, 0xc0,
	0x80, 0x78, 0x05, 0x8c, 0x88, 0x91, 0x91, 0x89, 0x91, 0xb1, 0x63, 0x47, 0xea, 0x2e, 0xb0, 0xf5,
	0x25, 0x20, 0xff, 0x49, 0x93, 0x14, 0x04, 0xdb, 0x3d, 0xcf, 0xf7, 0xf3, 0xfd, 0xea, 0xee, 0x79,
	0x64, 0xc3, 0xda, 0x94, 0xfa, 0x2a, 0xe3, 0x9a, 0x6f, 0x9d, 0x52, 0x97, 0x68, 0xc4, 0x63, 0xa6,
	0x25, 0xe8, 0x90, 0x8e, 0x25, 0x23, 0x8e, 0x26, 0x2f, 0x3c, 0xea, 0xab, 0x9e, 0xe0, 0x92, 0xa3,
	0xdd, 0x98, 0x54, 0x63, 0x52, 0x5d, 0x26, 0x2b, 0x7b, 0x36, 0x93, 0xa7, 0x93, 0x81, 0x6a, 0x71,
	0x57, 0xb3, 0xb9, 0xcd, 0xb5, 0xc8, 0x34, 0x98, 0x8c, 0xa2, 0x2a, 0x2a, 0xa2, 0x53, 0x1c, 0x56,
	0x79, 0x68, 0x73, 0x6e, 0x3b, 0x74, 0x4e, 0x49, 0xe6, 0x52, 0x5f, 0x12, 0xd7, 0x4b, 0x80, 0x9d,
	0xe5, 0x7b, 0x71, 0x4f, 0x32, 0x3e, 0x4e, 0xae, 0x52, 0x79, 0xb0, 0x2c, 0x2e, 0xdc, 0xb2, 0xb2,
	0xbb, 0x2c, 0x4d, 0x89, 0xc3, 0x86, 0x44, 0xd2, 0x58, 0x7d, 0xf4, 0x21, 0x0b, 0xd7, 0xda, 0x0e,
	0x1f, 0x10, 0xe7, 0xc4, 0xa3, 0x56, 0xff, 0xc2, 0xa3, 0xa8, 0x05, 0xb3, 0xa1, 0x5f, 0x01, 0x55,
	0x50, 0x5b, 0xab, 0x6b, 0xea, 0xdf, 0x5e, 0xa9, 0x36, 0x8e, 0x3b, 0xad, 0xbb, 0x2a, 0xb4, 0x1b,
	0x91, 0x19, 0xe9, 0x30, 0x37, 0xf1, 0xa9, 0xf0, 0x95, 0x74, 0x35, 0x53, 0x5b, 0xad, 0xef, 0xde,
	0x4b, 0xe9, 0x0d, 0x5e, 0x51, 0x4b, 0x1a, 0x74, 0x14, 0x5a, 0x9a, 0xa5, 0x4f, 0xaf, 0xb3, 0x21,
	0xfd, 0xe5, 0xe7, 0xd7, 0x4c, 0xee, 0x3d, 0x48, 0x17, 0x81, 0x11, 0xbb, 0x51, 0x1d, 0x6e, 0x4e,
	0x99, 0x90, 0x13, 0xe2, 0x98, 0x67, 0xfb, 0xbe, 0x39, 0x26, 0x2e, 0xf5, 0x3d, 0x62, 0x51, 0x25,
	0x53, 0x05, 0xb5, 0x82, 0x51, 0x4e, 0xc4, 0xc3, 0x7d, 0xbf, 0x3b, 0x93, 0x50, 0x0d, 0x16, 0xef,
	0x7b, 0x94, 0x6c, 0x84, 0xaf, 0x2d, 0xe3, 0x68, 0x0b, 0xe6, 0x87, 0xcc, 0xa6, 0xbe, 0x54, 0x72,
	0x91, 0x9e, 0x54, 0xa8, 0x0d, 0x4b, 0x96, 0xa0, 0x44, 0xd2, 0xa1, 0x79, 0xb7, 0x05, 0x25, 0x5f,
	0x05, 0xb5, 0xd5, 0x7a, 0x45, 0x8d, 0xf7, 0xa4, 0xce, 0xf6, 0xa4, 0xf6, 0x67, 0x84, 0x51, 0x4c,
	0x4c, 0x77, 0x1d, 0xf4, 0x0c, 0x6e, 0xd0, 0x73, 0x8f, 0x09, 0x12, 0xee, 0x6a, 0x21, 0xeb, 0xbf,
	0x7f, 0x66, 0x95, 0xe7, 0xbe, 0x79, 0xdc, 0x16, 0xcc, 0x13, 0x4b, 0xb2, 0x29, 0x55, 0x56, 0xaa,
	0xa0, 0xb6, 0x62, 0x24, 0x15, 0x7a, 0x0a, 0xb7, 0x2c, 0x2a, 0x24, 0x1b, 0x31, 0x8b, 0x48, 0x6a,
	0xfa, 0x54, 0x30, 0xe2, 0x98, 0xe3, 0x89, 0xab, 0x14, 0xc2, 0x77, 0x35, 0x0b, 0xe1, 0x60, 0xb3,
	0x22, 0xad, 0x54, 0x8d, 0x8d, 0x05, 0xf0, 0x24, 0xe2, 0xba, 0x13, 0x17, 0xed, 0xc0, 0x82, 0xcf,
	0x24, 0x8d, 0x67, 0x05, 0xa3, 0x59, 0xac, 0x84, 0x8d, 0x70, 0x4a, 0x8f, 0x3f, 0x03, 0x58, 0xfa,
	0x6d, 0xcd, 0xa8, 0x0c, 0xd7, 0x1b, 0xc7, 0x1d, 0xb3, 0xa5, 0x1b, 0xfd, 0xce, 0x41, 0xa7, 0xd5,
	0xe8, 0xeb, 0xc5, 0x14, 0x5a, 0x87, 0xab, 0x87, 0xcf, 0x9b, 0xba, 0xd9, 0xea, 0x75, 0x0f, 0x3a,
	0xed, 0x22, 0x40, 0xff, 0xc3, 0x42, 0x48, 0xf5, 0x7b, 0x87, 0x7a, 0xb7, 0x98, 0x46, 0x9b, 0xb0,
	0x74, 0xa2, 0x1b, 0x2f, 0x3a, 0x2d, 0xdd, 0x9c, 0xb7, 0x33, 0x68, 0x07, 0x6e, 0x2f, 0xb6, 0x17,
	0x33, 0xb3, 0x68, 0x1b, 0x96, 0x67, 0xe2, 0x62, 0x76, 0x2e, 0x72, 0x75, 0xfa, 0xba, 0xd9, 0x3e,
	0xea, 0x35, 0x1b, 0x47, 0x4b, 0x62, 0xbe, 0xf9, 0x16, 0x5c, 0x5e, 0xe3, 0xd4, 0xd5, 0x35, 0x4e,
	0xdd, 0x5e, 0x63, 0xf0, 0x26, 0xc0, 0xe0, 0x63, 0x80, 0xc1, 0xb7, 0x00, 0x83, 0xcb, 0x00, 0x83,
	0xab, 0x00, 0x83, 0xef, 0x01, 0x06, 0x3f, 0x02, 0x9c, 0xba, 0x0d, 0x30, 0x78, 0x77, 0x83, 0x53,
	0x97, 0x37, 0x38, 0x75, 0x75, 0x83, 0x53, 0x2f, 0x0d, 0x9b, 0x7b, 0x67, 0xb6, 0x3a, 0xe5, 0x8e,
	0xa4, 0x42, 0x10, 0x75, 0xe2, 0x6b, 0xd1, 0x61, 0xc4, 0x85, 0xbb, 0xe7, 0x09, 0x3e, 0x65, 0x43,
	0x2a, 0xf6, 0x66, 0xb2, 0xe6, 0x0d, 0x6c, 0xae, 0xd1, 0x73, 0x99, 0x7c, 0x5f, 0x7f, 0xfc, 0x6d,
	0x0c, 0xf2, 0xd1, 0x8a, 0x9f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x49, 0xf8, 0x5b, 0x5d,
	0x04, 0x00, 0x00,
}

func (x APICredentialType) String() string {
	s, ok := APICredentialType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
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
	if this.Type != that1.Type {
		return false
	}
	if len(this.Users) != len(that1.Users) {
		return false
	}
	for i := range this.Users {
		if !this.Users[i].Equal(that1.Users[i]) {
			return false
		}
	}
	if this.VirtualK8SNamespace != that1.VirtualK8SNamespace {
		return false
	}
	if this.VirtualK8SName != that1.VirtualK8SName {
		return false
	}
	if this.Digest != that1.Digest {
		return false
	}
	if !this.CreatedTimestamp.Equal(that1.CreatedTimestamp) {
		return false
	}
	if !this.ExpirationTimestamp.Equal(that1.ExpirationTimestamp) {
		return false
	}
	if this.Active != that1.Active {
		return false
	}
	if this.CertificateSerialNum != that1.CertificateSerialNum {
		return false
	}
	if this.SiteName != that1.SiteName {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&api_credential.GlobalSpecType{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	if this.Users != nil {
		s = append(s, "Users: "+fmt.Sprintf("%#v", this.Users)+",\n")
	}
	s = append(s, "VirtualK8SNamespace: "+fmt.Sprintf("%#v", this.VirtualK8SNamespace)+",\n")
	s = append(s, "VirtualK8SName: "+fmt.Sprintf("%#v", this.VirtualK8SName)+",\n")
	s = append(s, "Digest: "+fmt.Sprintf("%#v", this.Digest)+",\n")
	if this.CreatedTimestamp != nil {
		s = append(s, "CreatedTimestamp: "+fmt.Sprintf("%#v", this.CreatedTimestamp)+",\n")
	}
	if this.ExpirationTimestamp != nil {
		s = append(s, "ExpirationTimestamp: "+fmt.Sprintf("%#v", this.ExpirationTimestamp)+",\n")
	}
	s = append(s, "Active: "+fmt.Sprintf("%#v", this.Active)+",\n")
	s = append(s, "CertificateSerialNum: "+fmt.Sprintf("%#v", this.CertificateSerialNum)+",\n")
	s = append(s, "SiteName: "+fmt.Sprintf("%#v", this.SiteName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SiteName) > 0 {
		i -= len(m.SiteName)
		copy(dAtA[i:], m.SiteName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SiteName)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.CertificateSerialNum) > 0 {
		i -= len(m.CertificateSerialNum)
		copy(dAtA[i:], m.CertificateSerialNum)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CertificateSerialNum)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.ExpirationTimestamp != nil {
		{
			size, err := m.ExpirationTimestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.CreatedTimestamp != nil {
		{
			size, err := m.CreatedTimestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.VirtualK8SName) > 0 {
		i -= len(m.VirtualK8SName)
		copy(dAtA[i:], m.VirtualK8SName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.VirtualK8SName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VirtualK8SNamespace) > 0 {
		i -= len(m.VirtualK8SNamespace)
		copy(dAtA[i:], m.VirtualK8SNamespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.VirtualK8SNamespace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Users[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.VirtualK8SNamespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.VirtualK8SName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.CreatedTimestamp != nil {
		l = m.CreatedTimestamp.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ExpirationTimestamp != nil {
		l = m.ExpirationTimestamp.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Active {
		n += 2
	}
	l = len(m.CertificateSerialNum)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SiteName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForUsers := "[]*ObjectRefType{"
	for _, f := range this.Users {
		repeatedStringForUsers += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForUsers += "}"
	s := strings.Join([]string{`&GlobalSpecType{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Users:` + repeatedStringForUsers + `,`,
		`VirtualK8SNamespace:` + fmt.Sprintf("%v", this.VirtualK8SNamespace) + `,`,
		`VirtualK8SName:` + fmt.Sprintf("%v", this.VirtualK8SName) + `,`,
		`Digest:` + fmt.Sprintf("%v", this.Digest) + `,`,
		`CreatedTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.CreatedTimestamp), "Timestamp", "types.Timestamp", 1) + `,`,
		`ExpirationTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.ExpirationTimestamp), "Timestamp", "types.Timestamp", 1) + `,`,
		`Active:` + fmt.Sprintf("%v", this.Active) + `,`,
		`CertificateSerialNum:` + fmt.Sprintf("%v", this.CertificateSerialNum) + `,`,
		`SiteName:` + fmt.Sprintf("%v", this.SiteName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= APICredentialType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &schema.ObjectRefType{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualK8SNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualK8SNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualK8SName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualK8SName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedTimestamp == nil {
				m.CreatedTimestamp = &types.Timestamp{}
			}
			if err := m.CreatedTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTimestamp == nil {
				m.ExpirationTimestamp = &types.Timestamp{}
			}
			if err := m.ExpirationTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateSerialNum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificateSerialNum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SiteName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SiteName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
