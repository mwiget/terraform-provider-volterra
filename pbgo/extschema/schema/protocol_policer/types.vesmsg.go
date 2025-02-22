//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package protocol_policer

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *CreateSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProtocolPolicerDRefInfo()

}

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetProtocolPolicerDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProtocolPolicer() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetProtocolPolicer() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetProtocolPolicer() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("protocol_policer[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) ProtocolPolicerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ProtocolPolicerType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ProtocolPolicerTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for protocol_policer")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ProtocolPolicerType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ProtocolPolicerType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated protocol_policer")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items protocol_policer")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["protocol_policer"]; exists {
		vOpts := append(opts, db.WithValidateField("protocol_policer"))
		if err := fv(ctx, m.GetProtocolPolicer(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhProtocolPolicer := v.ProtocolPolicerValidationRuleHandler
	rulesProtocolPolicer := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "64",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhProtocolPolicer(rulesProtocolPolicer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.protocol_policer: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol_policer"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DnsType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DnsType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DnsType) DeepCopy() *DnsType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DnsType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DnsType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DnsType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DnsTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDnsType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDnsType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DnsType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DnsType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDnsTypeValidator = func() *ValidateDnsType {
	v := &ValidateDnsType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DnsTypeValidator() db.Validator {
	return DefaultDnsTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *GetSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProtocolPolicerDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetSpecType) GetProtocolPolicerDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProtocolPolicer() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetProtocolPolicer() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetProtocolPolicer() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("protocol_policer[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) ProtocolPolicerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ProtocolPolicerType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ProtocolPolicerTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for protocol_policer")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ProtocolPolicerType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ProtocolPolicerType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated protocol_policer")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items protocol_policer")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["protocol_policer"]; exists {
		vOpts := append(opts, db.WithValidateField("protocol_policer"))
		if err := fv(ctx, m.GetProtocolPolicer(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhProtocolPolicer := v.ProtocolPolicerValidationRuleHandler
	rulesProtocolPolicer := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "64",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhProtocolPolicer(rulesProtocolPolicer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.protocol_policer: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol_policer"] = vFn

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProtocolPolicerDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetProtocolPolicerDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProtocolPolicer() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetProtocolPolicer() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetProtocolPolicer() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("protocol_policer[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) ProtocolPolicerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ProtocolPolicerType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ProtocolPolicerTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for protocol_policer")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ProtocolPolicerType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ProtocolPolicerType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated protocol_policer")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items protocol_policer")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["protocol_policer"]; exists {
		vOpts := append(opts, db.WithValidateField("protocol_policer"))
		if err := fv(ctx, m.GetProtocolPolicer(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhProtocolPolicer := v.ProtocolPolicerValidationRuleHandler
	rulesProtocolPolicer := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "64",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhProtocolPolicer(rulesProtocolPolicer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.protocol_policer: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol_policer"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *IcmpType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IcmpType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IcmpType) DeepCopy() *IcmpType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IcmpType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IcmpType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IcmpType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IcmpTypeValidator().Validate(ctx, m, opts...)
}

type ValidateIcmpType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIcmpType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IcmpType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IcmpType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		for idx, item := range m.GetType() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIcmpTypeValidator = func() *ValidateIcmpType {
	v := &ValidateIcmpType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func IcmpTypeValidator() db.Validator {
	return DefaultIcmpTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ProtocolPolicerType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ProtocolPolicerType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ProtocolPolicerType) DeepCopy() *ProtocolPolicerType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ProtocolPolicerType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ProtocolPolicerType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ProtocolPolicerType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ProtocolPolicerTypeValidator().Validate(ctx, m, opts...)
}

func (m *ProtocolPolicerType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetPolicerDRefInfo()

}

func (m *ProtocolPolicerType) GetPolicerDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetPolicer()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("ProtocolPolicerType.policer[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "policer.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "policer",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetPolicerDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *ProtocolPolicerType) GetPolicerDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "policer.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: policer")
	}
	for _, ref := range m.GetPolicer() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateProtocolPolicerType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateProtocolPolicerType) ProtocolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for protocol")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateProtocolPolicerType) PolicerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for policer")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated policer")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items policer")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateProtocolPolicerType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ProtocolPolicerType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ProtocolPolicerType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["policer"]; exists {
		vOpts := append(opts, db.WithValidateField("policer"))
		if err := fv(ctx, m.GetPolicer(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["protocol"]; exists {

		vOpts := append(opts, db.WithValidateField("protocol"))
		if err := fv(ctx, m.GetProtocol(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultProtocolPolicerTypeValidator = func() *ValidateProtocolPolicerType {
	v := &ValidateProtocolPolicerType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhProtocol := v.ProtocolValidationRuleHandler
	rulesProtocol := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhProtocol(rulesProtocol)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ProtocolPolicerType.protocol: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol"] = vFn

	vrhPolicer := v.PolicerValidationRuleHandler
	rulesPolicer := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhPolicer(rulesPolicer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ProtocolPolicerType.policer: %s", err)
		panic(errMsg)
	}
	v.FldValidators["policer"] = vFn

	return v
}()

func ProtocolPolicerTypeValidator() db.Validator {
	return DefaultProtocolPolicerTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ProtocolType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ProtocolType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ProtocolType) DeepCopy() *ProtocolType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ProtocolType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ProtocolType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ProtocolType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ProtocolTypeValidator().Validate(ctx, m, opts...)
}

type ValidateProtocolType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateProtocolType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ProtocolType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ProtocolType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetType().(type) {
	case *ProtocolType_Tcp:
		if fv, exists := v.FldValidators["type.tcp"]; exists {
			val := m.GetType().(*ProtocolType_Tcp).Tcp
			vOpts := append(opts,
				db.WithValidateField("type"),
				db.WithValidateField("tcp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ProtocolType_Icmp:
		if fv, exists := v.FldValidators["type.icmp"]; exists {
			val := m.GetType().(*ProtocolType_Icmp).Icmp
			vOpts := append(opts,
				db.WithValidateField("type"),
				db.WithValidateField("icmp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ProtocolType_Udp:
		if fv, exists := v.FldValidators["type.udp"]; exists {
			val := m.GetType().(*ProtocolType_Udp).Udp
			vOpts := append(opts,
				db.WithValidateField("type"),
				db.WithValidateField("udp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ProtocolType_Dns:
		if fv, exists := v.FldValidators["type.dns"]; exists {
			val := m.GetType().(*ProtocolType_Dns).Dns
			vOpts := append(opts,
				db.WithValidateField("type"),
				db.WithValidateField("dns"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultProtocolTypeValidator = func() *ValidateProtocolType {
	v := &ValidateProtocolType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ProtocolTypeValidator() db.Validator {
	return DefaultProtocolTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *ReplaceSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetProtocolPolicerDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ReplaceSpecType) GetProtocolPolicerDRefInfo() ([]db.DRefInfo, error) {
	if m.GetProtocolPolicer() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetProtocolPolicer() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetProtocolPolicer() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("protocol_policer[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) ProtocolPolicerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ProtocolPolicerType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ProtocolPolicerTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for protocol_policer")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ProtocolPolicerType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ProtocolPolicerType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated protocol_policer")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items protocol_policer")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["protocol_policer"]; exists {
		vOpts := append(opts, db.WithValidateField("protocol_policer"))
		if err := fv(ctx, m.GetProtocolPolicer(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhProtocolPolicer := v.ProtocolPolicerValidationRuleHandler
	rulesProtocolPolicer := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "64",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhProtocolPolicer(rulesProtocolPolicer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.protocol_policer: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol_policer"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *TcpType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *TcpType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *TcpType) DeepCopy() *TcpType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &TcpType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *TcpType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *TcpType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return TcpTypeValidator().Validate(ctx, m, opts...)
}

type ValidateTcpType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateTcpType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*TcpType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *TcpType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["flags"]; exists {

		vOpts := append(opts, db.WithValidateField("flags"))
		for idx, item := range m.GetFlags() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultTcpTypeValidator = func() *ValidateTcpType {
	v := &ValidateTcpType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func TcpTypeValidator() db.Validator {
	return DefaultTcpTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *UdpType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UdpType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UdpType) DeepCopy() *UdpType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UdpType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UdpType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UdpType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UdpTypeValidator().Validate(ctx, m, opts...)
}

type ValidateUdpType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUdpType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UdpType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UdpType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUdpTypeValidator = func() *ValidateUdpType {
	v := &ValidateUdpType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UdpTypeValidator() db.Validator {
	return DefaultUdpTypeValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.ProtocolPolicer = f.GetProtocolPolicer()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.ProtocolPolicer = m1.ProtocolPolicer
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.ProtocolPolicer = f.GetProtocolPolicer()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.ProtocolPolicer = m1.ProtocolPolicer
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.ProtocolPolicer = f.GetProtocolPolicer()
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.ProtocolPolicer = m1.ProtocolPolicer
}
