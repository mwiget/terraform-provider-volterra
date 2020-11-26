// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/vesenv/quota_resource_keys.proto

package vesenv

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type QuotaResourceKeyChoice struct {
	// Types that are valid to be assigned to Choice:
	//	*QuotaResourceKeyChoice_VirtualHostPublic
	//	*QuotaResourceKeyChoice_HealthcheckMinimumInterval
	Choice isQuotaResourceKeyChoice_Choice `protobuf_oneof:"choice"`
}

func (m *QuotaResourceKeyChoice) Reset()      { *m = QuotaResourceKeyChoice{} }
func (*QuotaResourceKeyChoice) ProtoMessage() {}
func (*QuotaResourceKeyChoice) Descriptor() ([]byte, []int) {
	return fileDescriptorQuotaResourceKeys, []int{0}
}

type isQuotaResourceKeyChoice_Choice interface {
	isQuotaResourceKeyChoice_Choice()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type QuotaResourceKeyChoice_VirtualHostPublic struct {
	VirtualHostPublic bool `protobuf:"varint,1,opt,name=virtual_host_public,json=virtualHostPublic,proto3,oneof"`
}
type QuotaResourceKeyChoice_HealthcheckMinimumInterval struct {
	HealthcheckMinimumInterval bool `protobuf:"varint,2,opt,name=healthcheck_minimum_interval,json=healthcheckMinimumInterval,proto3,oneof"`
}

func (*QuotaResourceKeyChoice_VirtualHostPublic) isQuotaResourceKeyChoice_Choice()          {}
func (*QuotaResourceKeyChoice_HealthcheckMinimumInterval) isQuotaResourceKeyChoice_Choice() {}

func (m *QuotaResourceKeyChoice) GetChoice() isQuotaResourceKeyChoice_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

func (m *QuotaResourceKeyChoice) GetVirtualHostPublic() bool {
	if x, ok := m.GetChoice().(*QuotaResourceKeyChoice_VirtualHostPublic); ok {
		return x.VirtualHostPublic
	}
	return false
}

func (m *QuotaResourceKeyChoice) GetHealthcheckMinimumInterval() bool {
	if x, ok := m.GetChoice().(*QuotaResourceKeyChoice_HealthcheckMinimumInterval); ok {
		return x.HealthcheckMinimumInterval
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*QuotaResourceKeyChoice) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _QuotaResourceKeyChoice_OneofMarshaler, _QuotaResourceKeyChoice_OneofUnmarshaler, _QuotaResourceKeyChoice_OneofSizer, []interface{}{
		(*QuotaResourceKeyChoice_VirtualHostPublic)(nil),
		(*QuotaResourceKeyChoice_HealthcheckMinimumInterval)(nil),
	}
}

func _QuotaResourceKeyChoice_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*QuotaResourceKeyChoice)
	// choice
	switch x := m.Choice.(type) {
	case *QuotaResourceKeyChoice_VirtualHostPublic:
		t := uint64(0)
		if x.VirtualHostPublic {
			t = 1
		}
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *QuotaResourceKeyChoice_HealthcheckMinimumInterval:
		t := uint64(0)
		if x.HealthcheckMinimumInterval {
			t = 1
		}
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("QuotaResourceKeyChoice.Choice has unexpected type %T", x)
	}
	return nil
}

func _QuotaResourceKeyChoice_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*QuotaResourceKeyChoice)
	switch tag {
	case 1: // choice.virtual_host_public
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Choice = &QuotaResourceKeyChoice_VirtualHostPublic{x != 0}
		return true, err
	case 2: // choice.healthcheck_minimum_interval
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Choice = &QuotaResourceKeyChoice_HealthcheckMinimumInterval{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _QuotaResourceKeyChoice_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*QuotaResourceKeyChoice)
	// choice
	switch x := m.Choice.(type) {
	case *QuotaResourceKeyChoice_VirtualHostPublic:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *QuotaResourceKeyChoice_HealthcheckMinimumInterval:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*QuotaResourceKeyChoice)(nil), "ves.io.schema.vesenv.QuotaResourceKeyChoice")
}
func (this *QuotaResourceKeyChoice) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuotaResourceKeyChoice)
	if !ok {
		that2, ok := that.(QuotaResourceKeyChoice)
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
	if that1.Choice == nil {
		if this.Choice != nil {
			return false
		}
	} else if this.Choice == nil {
		return false
	} else if !this.Choice.Equal(that1.Choice) {
		return false
	}
	return true
}
func (this *QuotaResourceKeyChoice_VirtualHostPublic) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuotaResourceKeyChoice_VirtualHostPublic)
	if !ok {
		that2, ok := that.(QuotaResourceKeyChoice_VirtualHostPublic)
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
	if this.VirtualHostPublic != that1.VirtualHostPublic {
		return false
	}
	return true
}
func (this *QuotaResourceKeyChoice_HealthcheckMinimumInterval) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuotaResourceKeyChoice_HealthcheckMinimumInterval)
	if !ok {
		that2, ok := that.(QuotaResourceKeyChoice_HealthcheckMinimumInterval)
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
	if this.HealthcheckMinimumInterval != that1.HealthcheckMinimumInterval {
		return false
	}
	return true
}
func (this *QuotaResourceKeyChoice) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&vesenv.QuotaResourceKeyChoice{")
	if this.Choice != nil {
		s = append(s, "Choice: "+fmt.Sprintf("%#v", this.Choice)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QuotaResourceKeyChoice_VirtualHostPublic) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&vesenv.QuotaResourceKeyChoice_VirtualHostPublic{` +
		`VirtualHostPublic:` + fmt.Sprintf("%#v", this.VirtualHostPublic) + `}`}, ", ")
	return s
}
func (this *QuotaResourceKeyChoice_HealthcheckMinimumInterval) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&vesenv.QuotaResourceKeyChoice_HealthcheckMinimumInterval{` +
		`HealthcheckMinimumInterval:` + fmt.Sprintf("%#v", this.HealthcheckMinimumInterval) + `}`}, ", ")
	return s
}
func valueToGoStringQuotaResourceKeys(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QuotaResourceKeyChoice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotaResourceKeyChoice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Choice != nil {
		nn1, err := m.Choice.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *QuotaResourceKeyChoice_VirtualHostPublic) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x8
	i++
	if m.VirtualHostPublic {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *QuotaResourceKeyChoice_HealthcheckMinimumInterval) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x10
	i++
	if m.HealthcheckMinimumInterval {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func encodeVarintQuotaResourceKeys(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QuotaResourceKeyChoice) Size() (n int) {
	var l int
	_ = l
	if m.Choice != nil {
		n += m.Choice.Size()
	}
	return n
}

func (m *QuotaResourceKeyChoice_VirtualHostPublic) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *QuotaResourceKeyChoice_HealthcheckMinimumInterval) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}

func sovQuotaResourceKeys(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuotaResourceKeys(x uint64) (n int) {
	return sovQuotaResourceKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QuotaResourceKeyChoice) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaResourceKeyChoice{`,
		`Choice:` + fmt.Sprintf("%v", this.Choice) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaResourceKeyChoice_VirtualHostPublic) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaResourceKeyChoice_VirtualHostPublic{`,
		`VirtualHostPublic:` + fmt.Sprintf("%v", this.VirtualHostPublic) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaResourceKeyChoice_HealthcheckMinimumInterval) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaResourceKeyChoice_HealthcheckMinimumInterval{`,
		`HealthcheckMinimumInterval:` + fmt.Sprintf("%v", this.HealthcheckMinimumInterval) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQuotaResourceKeys(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QuotaResourceKeyChoice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuotaResourceKeys
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuotaResourceKeyChoice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaResourceKeyChoice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualHostPublic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaResourceKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Choice = &QuotaResourceKeyChoice_VirtualHostPublic{b}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HealthcheckMinimumInterval", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotaResourceKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Choice = &QuotaResourceKeyChoice_HealthcheckMinimumInterval{b}
		default:
			iNdEx = preIndex
			skippy, err := skipQuotaResourceKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuotaResourceKeys
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
func skipQuotaResourceKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuotaResourceKeys
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
					return 0, ErrIntOverflowQuotaResourceKeys
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuotaResourceKeys
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthQuotaResourceKeys
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuotaResourceKeys
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipQuotaResourceKeys(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthQuotaResourceKeys = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuotaResourceKeys   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/vesenv/quota_resource_keys.proto", fileDescriptorQuotaResourceKeys)
}

var fileDescriptorQuotaResourceKeys = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0x3f, 0x4e, 0xc3, 0x30,
	0x18, 0x05, 0xf0, 0x18, 0xa1, 0xaa, 0xca, 0x46, 0x41, 0x28, 0x2a, 0xc8, 0xa0, 0x8a, 0xa1, 0x4b,
	0x9d, 0x81, 0x1b, 0x94, 0xa5, 0x80, 0x90, 0x68, 0x47, 0x96, 0xc8, 0x0d, 0x1f, 0x8d, 0x95, 0x3f,
	0x5f, 0xb0, 0x1d, 0xab, 0xdd, 0x38, 0x02, 0xc7, 0x40, 0xbd, 0x42, 0x2f, 0xc0, 0xd8, 0x91, 0x91,
	0xba, 0x0b, 0x63, 0x8f, 0x80, 0x48, 0x02, 0xa2, 0x88, 0xcd, 0xd2, 0xfb, 0xf9, 0x59, 0x7e, 0x2e,
	0x33, 0xa0, 0x98, 0x40, 0x5f, 0x85, 0x11, 0xa4, 0xdc, 0x37, 0xa0, 0x20, 0x33, 0xfe, 0x63, 0x81,
	0x9a, 0x07, 0x12, 0x14, 0x16, 0x32, 0x84, 0x20, 0x86, 0x99, 0x62, 0xb9, 0x44, 0x8d, 0xad, 0x83,
	0xca, 0xb3, 0xca, 0xb3, 0xca, 0xb7, 0x3b, 0xff, 0xb6, 0x60, 0xae, 0x05, 0x66, 0xf5, 0xcd, 0xce,
	0x9a, 0xb8, 0x87, 0xc3, 0xaf, 0xde, 0x51, 0x5d, 0x7b, 0x0d, 0xb3, 0x8b, 0x08, 0x45, 0x08, 0xad,
	0xa1, 0xbb, 0x6f, 0x84, 0xd4, 0x05, 0x4f, 0x82, 0x08, 0x95, 0x0e, 0xf2, 0x62, 0x9c, 0x88, 0xd0,
	0x23, 0xa7, 0xa4, 0xdb, 0xec, 0x9f, 0xcc, 0x17, 0xde, 0xd1, 0x36, 0x60, 0x15, 0x68, 0xed, 0xf2,
	0x98, 0xcb, 0x81, 0x33, 0xda, 0xab, 0xc3, 0x01, 0x2a, 0x7d, 0x5b, 0x45, 0x89, 0x7b, 0x1c, 0x01,
	0x4f, 0x74, 0x14, 0x46, 0x10, 0xc6, 0x41, 0x2a, 0x32, 0x91, 0x16, 0x69, 0x20, 0x32, 0x0d, 0xd2,
	0xf0, 0xc4, 0xdb, 0x29, 0xbb, 0xbb, 0xf3, 0x85, 0x77, 0xb6, 0x25, 0xd9, 0x5f, 0xf9, 0xf3, 0x48,
	0xfb, 0x97, 0xba, 0xa9, 0xd0, 0x65, 0x6d, 0xfa, 0x4d, 0xb7, 0x11, 0x96, 0x5f, 0xe9, 0x4f, 0x97,
	0x2b, 0xea, 0xbc, 0xad, 0xa8, 0xb3, 0x59, 0x51, 0xf2, 0x64, 0x29, 0x79, 0xb1, 0x94, 0xbc, 0x5a,
	0x4a, 0x96, 0x96, 0x92, 0x77, 0x4b, 0xc9, 0x87, 0xa5, 0xce, 0xc6, 0x52, 0xf2, 0xbc, 0xa6, 0xce,
	0xdd, 0xd5, 0x04, 0xf3, 0x78, 0xc2, 0x0c, 0x26, 0x1a, 0xa4, 0xe4, 0xac, 0x50, 0x7e, 0x79, 0x78,
	0x40, 0x99, 0xf6, 0x72, 0x89, 0x46, 0xdc, 0x83, 0xec, 0x7d, 0xc7, 0x7e, 0x3e, 0x9e, 0xa0, 0x0f,
	0x53, 0x5d, 0xcf, 0xbc, 0xb5, 0xf6, 0xb8, 0x51, 0xce, 0x7c, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0x76, 0x60, 0xbb, 0xd2, 0x01, 0x00, 0x00,
}
