// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/views/terraform_parameters/object.proto

package terraform_parameters

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema4 "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	_ "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Object
//
// x-displayName: "Object"
// view terraform parameters object
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *ves_io_schema4.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *ves_io_schema4.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the desired behavior of the tenant
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{0} }

func (m *Object) GetMetadata() *ves_io_schema4.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *ves_io_schema4.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

type SpecType struct {
	GcSpec *GlobalSpecType `protobuf:"bytes,1,opt,name=gc_spec,json=gcSpec" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()                    { *m = SpecType{} }
func (*SpecType) ProtoMessage()               {}
func (*SpecType) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{1} }

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

// Status Object
//
// x-displayName: "Status"
// view terraform parameters status object
type StatusObject struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard status's metadata
	Metadata *ves_io_schema4.StatusMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// object_refs
	//
	// x-displayName: "Config Object"
	// View terraform parameters object reference for which this status object is generated
	ObjectRefs []*ves_io_schema4.ObjectRefType `protobuf:"bytes,2,rep,name=object_refs,json=objectRefs" json:"object_refs,omitempty"`
	// Plan Status
	//
	// x-displayName: "Plan Status"
	// Status of Plan action
	PlanStatus *PlanStatus `protobuf:"bytes,10,opt,name=plan_status,json=planStatus" json:"plan_status,omitempty"`
	// Apply Status
	//
	// x-displayName: "Apply Status"
	// Status of Apply or Destroy action
	ApplyStatus *ApplyStatus `protobuf:"bytes,11,opt,name=apply_status,json=applyStatus" json:"apply_status,omitempty"`
}

func (m *StatusObject) Reset()                    { *m = StatusObject{} }
func (*StatusObject) ProtoMessage()               {}
func (*StatusObject) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{2} }

func (m *StatusObject) GetMetadata() *ves_io_schema4.StatusMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatusObject) GetObjectRefs() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.ObjectRefs
	}
	return nil
}

func (m *StatusObject) GetPlanStatus() *PlanStatus {
	if m != nil {
		return m.PlanStatus
	}
	return nil
}

func (m *StatusObject) GetApplyStatus() *ApplyStatus {
	if m != nil {
		return m.ApplyStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.views.terraform_parameters.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.views.terraform_parameters.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.views.terraform_parameters.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.views.terraform_parameters.SpecType")
	proto.RegisterType((*StatusObject)(nil), "ves.io.schema.views.terraform_parameters.StatusObject")
	golang_proto.RegisterType((*StatusObject)(nil), "ves.io.schema.views.terraform_parameters.StatusObject")
}
func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
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
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *StatusObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StatusObject)
	if !ok {
		that2, ok := that.(StatusObject)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.ObjectRefs) != len(that1.ObjectRefs) {
		return false
	}
	for i := range this.ObjectRefs {
		if !this.ObjectRefs[i].Equal(that1.ObjectRefs[i]) {
			return false
		}
	}
	if !this.PlanStatus.Equal(that1.PlanStatus) {
		return false
	}
	if !this.ApplyStatus.Equal(that1.ApplyStatus) {
		return false
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&terraform_parameters.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&terraform_parameters.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StatusObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&terraform_parameters.StatusObject{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.ObjectRefs != nil {
		s = append(s, "ObjectRefs: "+fmt.Sprintf("%#v", this.ObjectRefs)+",\n")
	}
	if this.PlanStatus != nil {
		s = append(s, "PlanStatus: "+fmt.Sprintf("%#v", this.PlanStatus)+",\n")
	}
	if this.ApplyStatus != nil {
		s = append(s, "ApplyStatus: "+fmt.Sprintf("%#v", this.ApplyStatus)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SystemMetadata != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.SystemMetadata.Size()))
		n2, err := m.SystemMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Spec != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Spec.Size()))
		n3, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.GcSpec.Size()))
		n4, err := m.GcSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *StatusObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n5, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.ObjectRefs) > 0 {
		for _, msg := range m.ObjectRefs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintObject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.PlanStatus != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.PlanStatus.Size()))
		n6, err := m.PlanStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.ApplyStatus != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.ApplyStatus.Size()))
		n7, err := m.ApplyStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedObject(r randyObject, easy bool) *Object {
	this := &Object{}
	if r.Intn(10) != 0 {
		this.Metadata = ves_io_schema4.NewPopulatedObjectMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		this.SystemMetadata = ves_io_schema4.NewPopulatedSystemObjectMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Spec = NewPopulatedSpecType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSpecType(r randyObject, easy bool) *SpecType {
	this := &SpecType{}
	if r.Intn(10) != 0 {
		this.GcSpec = NewPopulatedGlobalSpecType(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStatusObject(r randyObject, easy bool) *StatusObject {
	this := &StatusObject{}
	if r.Intn(10) != 0 {
		this.Metadata = ves_io_schema4.NewPopulatedStatusMetaType(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.ObjectRefs = make([]*ves_io_schema4.ObjectRefType, v1)
		for i := 0; i < v1; i++ {
			this.ObjectRefs[i] = ves_io_schema4.NewPopulatedObjectRefType(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		this.PlanStatus = NewPopulatedPlanStatus(r, easy)
	}
	if r.Intn(10) != 0 {
		this.ApplyStatus = NewPopulatedApplyStatus(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyObject interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneObject(r randyObject) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringObject(r randyObject) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneObject(r)
	}
	return string(tmps)
}
func randUnrecognizedObject(r randyObject, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldObject(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldObject(dAtA []byte, r randyObject, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateObject(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateObject(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateObject(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateObject(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Object) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *StatusObject) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if len(m.ObjectRefs) > 0 {
		for _, e := range m.ObjectRefs {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if m.PlanStatus != nil {
		l = m.PlanStatus.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.ApplyStatus != nil {
		l = m.ApplyStatus.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func sovObject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "ves_io_schema4.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "ves_io_schema4.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(fmt.Sprintf("%v", this.Spec), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatusObject{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "StatusMetaType", "ves_io_schema4.StatusMetaType", 1) + `,`,
		`ObjectRefs:` + strings.Replace(fmt.Sprintf("%v", this.ObjectRefs), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`PlanStatus:` + strings.Replace(fmt.Sprintf("%v", this.PlanStatus), "PlanStatus", "PlanStatus", 1) + `,`,
		`ApplyStatus:` + strings.Replace(fmt.Sprintf("%v", this.ApplyStatus), "ApplyStatus", "ApplyStatus", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &ves_io_schema4.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func (m *StatusObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: StatusObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.StatusMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectRefs = append(m.ObjectRefs, &ves_io_schema4.ObjectRefType{})
			if err := m.ObjectRefs[len(m.ObjectRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlanStatus == nil {
				m.PlanStatus = &PlanStatus{}
			}
			if err := m.PlanStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplyStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApplyStatus == nil {
				m.ApplyStatus = &ApplyStatus{}
			}
			if err := m.ApplyStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowObject
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
				next, err := skipObject(dAtA[start:])
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
	ErrInvalidLengthObject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/views/terraform_parameters/object.proto", fileDescriptorObject)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/views/terraform_parameters/object.proto", fileDescriptorObject)
}

var fileDescriptorObject = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x8d, 0x9b, 0xb8, 0xe7, 0xa8, 0x58, 0x06, 0x89, 0x50, 0xe0, 0x14, 0x85, 0x25,
	0x4b, 0xce, 0x52, 0x68, 0x25, 0x60, 0xa8, 0x44, 0x07, 0x90, 0x2c, 0x2a, 0x20, 0x01, 0x84, 0x90,
	0xc0, 0x3a, 0x3b, 0x17, 0xd7, 0x60, 0xe7, 0x4e, 0xbe, 0x4b, 0x20, 0x03, 0x12, 0x1f, 0x01, 0x66,
	0xbe, 0x00, 0xe2, 0x53, 0x74, 0x64, 0xac, 0x98, 0xca, 0x46, 0xdc, 0x85, 0xb1, 0x23, 0x23, 0xca,
	0x39, 0x0e, 0x4d, 0x89, 0x84, 0xb7, 0xf7, 0xfc, 0xde, 0xff, 0xe7, 0xf7, 0xfe, 0xf6, 0x83, 0x3b,
	0x63, 0x2a, 0x70, 0xc8, 0x6c, 0xe1, 0x1f, 0xd0, 0x98, 0xd8, 0xe3, 0x90, 0xbe, 0x15, 0xb6, 0xa4,
	0x49, 0x42, 0x06, 0x2c, 0x89, 0x5d, 0x4e, 0x12, 0x12, 0x53, 0x49, 0x13, 0x61, 0x33, 0xef, 0x35,
	0xf5, 0x25, 0xe6, 0x09, 0x93, 0xcc, 0x6a, 0x65, 0x32, 0x9c, 0xc9, 0xb0, 0x92, 0xe1, 0x55, 0xb2,
	0xad, 0x76, 0x10, 0xca, 0x83, 0x91, 0x87, 0x7d, 0x16, 0xdb, 0x01, 0x0b, 0x98, 0xad, 0x00, 0xde,
	0x68, 0xa0, 0x32, 0x95, 0xa8, 0x28, 0x03, 0x6f, 0x5d, 0x5d, 0x9e, 0x87, 0x71, 0x19, 0xb2, 0xa1,
	0x98, 0x17, 0xaf, 0x2c, 0x17, 0xe5, 0x84, 0xd3, 0xbc, 0xd4, 0xf8, 0x77, 0x0f, 0x77, 0x59, 0xbc,
	0x5d, 0x78, 0xd3, 0x33, 0xdc, 0xe6, 0x0f, 0x00, 0x2b, 0x0f, 0xd5, 0xe6, 0xd6, 0x6d, 0xa8, 0xc7,
	0x54, 0x92, 0x3e, 0x91, 0xa4, 0x0e, 0x1a, 0xa0, 0x65, 0x74, 0xae, 0xe3, 0x65, 0x1b, 0xb2, 0xc6,
	0x7d, 0x2a, 0xc9, 0x93, 0x09, 0xa7, 0xdd, 0x45, 0xbb, 0xf5, 0x00, 0x5e, 0x10, 0x13, 0x21, 0x69,
	0xec, 0x2e, 0x08, 0x6b, 0x8a, 0x70, 0xe3, 0x1c, 0xa1, 0xa7, 0xba, 0xce, 0x71, 0x36, 0x33, 0xed,
	0x7e, 0x4e, 0xbb, 0x07, 0x35, 0xc1, 0xa9, 0x5f, 0x2f, 0x2b, 0x44, 0x07, 0x17, 0xfd, 0x16, 0xb8,
	0xc7, 0xa9, 0xaf, 0x88, 0x4a, 0xdf, 0x7c, 0x09, 0xf5, 0xfc, 0x89, 0xf5, 0x18, 0x56, 0x03, 0xdf,
	0x55, 0xd8, 0x6c, 0xb7, 0x5b, 0xc5, 0xb1, 0xf7, 0x23, 0xe6, 0x91, 0x68, 0x01, 0xaf, 0x04, 0xfe,
	0x2c, 0x6e, 0x7e, 0x2a, 0xc3, 0x5a, 0x4f, 0x12, 0x39, 0x12, 0x85, 0x0d, 0xcc, 0xda, 0x57, 0x18,
	0xf8, 0x0c, 0x1a, 0xd9, 0xff, 0xe7, 0x26, 0x74, 0x20, 0xea, 0x6b, 0x8d, 0x72, 0xcb, 0xe8, 0x5c,
	0x5b, 0x69, 0x7f, 0x97, 0x0e, 0x66, 0xe2, 0xbd, 0xcb, 0x5f, 0xdf, 0x5f, 0x5a, 0x35, 0x6e, 0x17,
	0xb2, 0xbc, 0x4f, 0x58, 0x4f, 0xa1, 0xc1, 0x23, 0x32, 0x74, 0x85, 0x7a, 0x71, 0x1d, 0xaa, 0xa9,
	0xb6, 0x8b, 0xaf, 0xfe, 0x28, 0x22, 0xc3, 0x6c, 0xe8, 0x2e, 0xe4, 0x8b, 0xd8, 0x7a, 0x0e, 0x6b,
	0x84, 0xf3, 0x68, 0x92, 0x73, 0x0d, 0xc5, 0xdd, 0x29, 0xce, 0xbd, 0x3b, 0x53, 0xcf, 0xc1, 0x06,
	0xf9, 0x9b, 0xdc, 0xb9, 0xf8, 0x7d, 0xd7, 0x84, 0x9b, 0xb0, 0x96, 0x3b, 0x83, 0x47, 0x61, 0xdf,
	0xd1, 0xf4, 0xb2, 0xa9, 0x39, 0x9a, 0xae, 0x99, 0xeb, 0x8e, 0xa6, 0xaf, 0x9b, 0x15, 0x47, 0xd3,
	0x2b, 0x66, 0xd5, 0xd1, 0xf4, 0xaa, 0xa9, 0x3b, 0x9a, 0xae, 0x9b, 0x1b, 0x8e, 0xa6, 0x6f, 0x98,
	0x70, 0xef, 0x33, 0x38, 0xdc, 0x2d, 0x1d, 0x4d, 0x51, 0xe9, 0x78, 0x8a, 0x4a, 0xa7, 0x53, 0x04,
	0x7e, 0x4f, 0x11, 0xf8, 0x90, 0x22, 0xf0, 0x25, 0x45, 0xe0, 0x5b, 0x8a, 0xc0, 0x51, 0x8a, 0xc0,
	0x71, 0x8a, 0xc0, 0xcf, 0x14, 0x81, 0x5f, 0x29, 0x2a, 0x9d, 0xa6, 0x08, 0x7c, 0x3c, 0x41, 0xa5,
	0xc3, 0x13, 0x04, 0x5e, 0xbc, 0x0a, 0x18, 0x7f, 0x13, 0xe0, 0x31, 0x8b, 0xd4, 0xd0, 0x78, 0x74,
	0xe6, 0x80, 0xda, 0x3c, 0x61, 0xe3, 0xb0, 0x4f, 0x93, 0x76, 0x5e, 0xb6, 0xb9, 0x17, 0x30, 0x9b,
	0xbe, 0x93, 0xf3, 0x8b, 0xfb, 0xef, 0xe1, 0x79, 0x15, 0x75, 0x73, 0x37, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x32, 0x43, 0xa8, 0x24, 0x95, 0x04, 0x00, 0x00,
}
