// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/dc_cluster_group/types.proto

package dc_cluster_group

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
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

// Global Specification
//
// x-displayName: "Global Specification"
// DC Cluster Group specification
type GlobalSpecType struct {
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a620df2814ae1a05, []int{0}
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

// Create DC Cluster group
//
// x-displayName: "Create DC Cluster Group"
// Create DC Cluster group in given namespace
type CreateSpecType struct {
}

func (m *CreateSpecType) Reset()      { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage() {}
func (*CreateSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a620df2814ae1a05, []int{1}
}
func (m *CreateSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CreateSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSpecType.Merge(m, src)
}
func (m *CreateSpecType) XXX_Size() int {
	return m.Size()
}
func (m *CreateSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSpecType proto.InternalMessageInfo

// Replace DC Cluster Group
//
// x-displayName: "Replace DC Cluster Group"
// Replace given DC Cluster Group in given namespace
type ReplaceSpecType struct {
}

func (m *ReplaceSpecType) Reset()      { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage() {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a620df2814ae1a05, []int{2}
}
func (m *ReplaceSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplaceSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplaceSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceSpecType.Merge(m, src)
}
func (m *ReplaceSpecType) XXX_Size() int {
	return m.Size()
}
func (m *ReplaceSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceSpecType proto.InternalMessageInfo

// Get DC Cluster Group
//
// x-displayName: "Get DC Cluster Group"
// Gets DC Cluster Group in given namespace
type GetSpecType struct {
}

func (m *GetSpecType) Reset()      { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage() {}
func (*GetSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a620df2814ae1a05, []int{3}
}
func (m *GetSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GetSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecType.Merge(m, src)
}
func (m *GetSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GetSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecType proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.dc_cluster_group.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.dc_cluster_group.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.dc_cluster_group.CreateSpecType")
	golang_proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.dc_cluster_group.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.dc_cluster_group.ReplaceSpecType")
	golang_proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.dc_cluster_group.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.dc_cluster_group.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.dc_cluster_group.GetSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/dc_cluster_group/types.proto", fileDescriptor_a620df2814ae1a05)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/dc_cluster_group/types.proto", fileDescriptor_a620df2814ae1a05)
}

var fileDescriptor_a620df2814ae1a05 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xed, 0xe5, 0x37, 0xe4, 0x27, 0x95, 0x3f, 0x0b, 0xa2, 0x48, 0x27, 0x54, 0x98, 0x90,
	0x6a, 0x0f, 0x6c, 0x0c, 0x0c, 0x30, 0x74, 0x87, 0x4e, 0x2c, 0x55, 0xe2, 0x5e, 0xdd, 0x88, 0xb4,
	0x67, 0x39, 0x4e, 0x45, 0x37, 0x1e, 0x01, 0xf1, 0x14, 0x3c, 0x06, 0x23, 0x62, 0xea, 0xd8, 0x91,
	0x3a, 0x0b, 0x63, 0x1f, 0x01, 0xc9, 0x49, 0x25, 0x52, 0x89, 0x6e, 0x77, 0xfe, 0x7c, 0x6c, 0x9f,
	0xbe, 0x17, 0x5d, 0xcc, 0x30, 0x17, 0x29, 0xc9, 0x5c, 0x8d, 0x71, 0x12, 0xcb, 0xa1, 0x1a, 0xa8,
	0xac, 0xc8, 0x1d, 0xda, 0x81, 0xb6, 0x54, 0x18, 0xe9, 0xe6, 0x06, 0x73, 0x61, 0x2c, 0x39, 0x3a,
	0x84, 0xca, 0x15, 0x95, 0x2b, 0xb6, 0xdd, 0x76, 0x57, 0xa7, 0x6e, 0x5c, 0x24, 0x42, 0xd1, 0x44,
	0x6a, 0xd2, 0x24, 0xc3, 0xb5, 0xa4, 0x18, 0x85, 0x2e, 0x34, 0xa1, 0xaa, 0x9e, 0x6b, 0x1f, 0x35,
	0xbf, 0x9e, 0xa2, 0xab, 0xc1, 0x49, 0x13, 0x90, 0x71, 0x29, 0x4d, 0xeb, 0x21, 0xda, 0xc7, 0x4d,
	0xf8, 0x6b, 0xbe, 0xce, 0x7e, 0xd4, 0xea, 0x65, 0x94, 0xc4, 0xd9, 0xbd, 0x41, 0xd5, 0x9f, 0x1b,
	0xec, 0x9c, 0x45, 0xad, 0x5b, 0x8b, 0xb1, 0xc3, 0xcd, 0xc9, 0xd5, 0xc1, 0xe7, 0xf5, 0xb6, 0x74,
	0x1e, 0xed, 0xdd, 0xa1, 0xc9, 0x62, 0xb5, 0xd3, 0x3a, 0x8d, 0xfe, 0xf7, 0xd0, 0xed, 0x30, 0x6e,
	0x5e, 0xf9, 0x62, 0x05, 0x6c, 0xb9, 0x02, 0xb6, 0x5e, 0x01, 0x7f, 0xf6, 0xc0, 0xdf, 0x3c, 0xf0,
	0x0f, 0x0f, 0x7c, 0xe1, 0x81, 0x2f, 0x3d, 0xf0, 0x2f, 0x0f, 0xfc, 0xdb, 0x03, 0x5b, 0x7b, 0xe0,
	0x2f, 0x25, 0xb0, 0xf7, 0x12, 0xf8, 0xa2, 0x04, 0xb6, 0x2c, 0x81, 0x3d, 0xf4, 0x35, 0x99, 0x47,
	0x2d, 0x66, 0x94, 0x39, 0xb4, 0x36, 0x16, 0x45, 0x2e, 0x43, 0x31, 0x22, 0x3b, 0xe9, 0x1a, 0x4b,
	0xb3, 0x74, 0x88, 0xb6, 0xbb, 0xc1, 0xd2, 0x24, 0x9a, 0x24, 0x3e, 0xb9, 0x3a, 0x87, 0x3f, 0xf6,
	0x97, 0xfc, 0x0b, 0xd1, 0x5c, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x32, 0xf3, 0x85, 0xc4, 0xe8,
	0x01, 0x00, 0x00,
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
	return true
}
func (this *CreateSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSpecType)
	if !ok {
		that2, ok := that.(CreateSpecType)
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
	return true
}
func (this *ReplaceSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceSpecType)
	if !ok {
		that2, ok := that.(ReplaceSpecType)
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
	return true
}
func (this *GetSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType)
	if !ok {
		that2, ok := that.(GetSpecType)
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
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.GlobalSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.CreateSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.ReplaceSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.GetSpecType{")
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
	return len(dAtA) - i, nil
}

func (m *CreateSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ReplaceSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplaceSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
	return n
}

func (m *CreateSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
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
	s := strings.Join([]string{`&GlobalSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType{`,
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
func (m *CreateSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *ReplaceSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReplaceSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *GetSpecType) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
