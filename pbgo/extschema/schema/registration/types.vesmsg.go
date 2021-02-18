//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package registration

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

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) InfraValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for infra")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := InfraValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) TokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for token")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) PassportValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for passport")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := PassportValidator().Validate(ctx, val, opts...); err != nil {
			return err
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

	if fv, exists := v.FldValidators["infra"]; exists {

		vOpts := append(opts, db.WithValidateField("infra"))
		if err := fv(ctx, m.GetInfra(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["passport"]; exists {

		vOpts := append(opts, db.WithValidateField("passport"))
		if err := fv(ctx, m.GetPassport(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
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

	vrhInfra := v.InfraValidationRuleHandler
	rulesInfra := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInfra(rulesInfra)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.infra: %s", err)
		panic(errMsg)
	}
	v.FldValidators["infra"] = vFn

	vrhToken := v.TokenValidationRuleHandler
	rulesToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhToken(rulesToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["token"] = vFn

	vrhPassport := v.PassportValidationRuleHandler
	rulesPassport := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassport(rulesPassport)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.passport: %s", err)
		panic(errMsg)
	}
	v.FldValidators["passport"] = vFn

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
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

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) InfraValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for infra")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := InfraValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) TokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for token")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) PassportValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for passport")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := PassportValidator().Validate(ctx, val, opts...); err != nil {
			return err
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

	if fv, exists := v.FldValidators["infra"]; exists {

		vOpts := append(opts, db.WithValidateField("infra"))
		if err := fv(ctx, m.GetInfra(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["passport"]; exists {

		vOpts := append(opts, db.WithValidateField("passport"))
		if err := fv(ctx, m.GetPassport(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
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

	vrhInfra := v.InfraValidationRuleHandler
	rulesInfra := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInfra(rulesInfra)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.infra: %s", err)
		panic(errMsg)
	}
	v.FldValidators["infra"] = vFn

	vrhToken := v.TokenValidationRuleHandler
	rulesToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhToken(rulesToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["token"] = vFn

	vrhPassport := v.PassportValidationRuleHandler
	rulesPassport := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassport(rulesPassport)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.passport: %s", err)
		panic(errMsg)
	}
	v.FldValidators["passport"] = vFn

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
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetSiteDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *GlobalSpecType) GetSiteDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, ref := range m.GetSite() {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.site[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "site.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "site",
			Ref:        ref,
		})
	}

	return drInfos, nil
}

// GetSiteDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetSiteDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "site.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: site")
	}
	for _, ref := range m.GetSite() {
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

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) InfraValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for infra")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := InfraValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for token")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) PassportValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for passport")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := PassportValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TunnelTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.SiteToSiteTunnelType)
		return int32(i)
	}
	// ves_io_schema.SiteToSiteTunnelType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.SiteToSiteTunnelType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tunnel_type")
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

	if fv, exists := v.FldValidators["infra"]; exists {

		vOpts := append(opts, db.WithValidateField("infra"))
		if err := fv(ctx, m.GetInfra(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["passport"]; exists {

		vOpts := append(opts, db.WithValidateField("passport"))
		if err := fv(ctx, m.GetPassport(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["role"]; exists {

		vOpts := append(opts, db.WithValidateField("role"))
		for idx, item := range m.GetRole() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["site"]; exists {

		vOpts := append(opts, db.WithValidateField("site"))
		for idx, item := range m.GetSite() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tunnel_type"]; exists {

		vOpts := append(opts, db.WithValidateField("tunnel_type"))
		if err := fv(ctx, m.GetTunnelType(), vOpts...); err != nil {
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

	vrhInfra := v.InfraValidationRuleHandler
	rulesInfra := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInfra(rulesInfra)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.infra: %s", err)
		panic(errMsg)
	}
	v.FldValidators["infra"] = vFn

	vrhToken := v.TokenValidationRuleHandler
	rulesToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhToken(rulesToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["token"] = vFn

	vrhPassport := v.PassportValidationRuleHandler
	rulesPassport := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassport(rulesPassport)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.passport: %s", err)
		panic(errMsg)
	}
	v.FldValidators["passport"] = vFn

	vrhTunnelType := v.TunnelTypeValidationRuleHandler
	rulesTunnelType := map[string]string{
		"ves.io.schema.rules.enum.in": "[0,1,2]",
	}
	vFn, err = vrhTunnelType(rulesTunnelType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.tunnel_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tunnel_type"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Infra) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Infra) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Infra) DeepCopy() *Infra {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Infra{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Infra) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Infra) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return InfraValidator().Validate(ctx, m, opts...)
}

type ValidateInfra struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateInfra) HostnameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for hostname")
	}

	return validatorFn, nil
}

func (v *ValidateInfra) InterfacesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemKeyRules := db.GetMapStringKeyRules(rules)
	itemKeyFn, err := db.NewStringValidationRuleHandler(itemKeyRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item key ValidationRuleHandler for interfaces")
	}
	itemsValidatorFn := func(ctx context.Context, kv map[string]*Interface, opts ...db.ValidateOpt) error {
		for key, value := range kv {
			if err := itemKeyFn(ctx, key, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element with key %v", key))
			}
			if err := InterfaceValidator().Validate(ctx, value, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("value for element with key %v", key))
			}
		}
		return nil
	}
	mapValFn, err := db.NewMapValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Map ValidationRuleHandler for interfaces")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.(map[string]*Interface)
		if !ok {
			return fmt.Errorf("Map validation expected map[ string ]*Interface, got %T", val)
		}
		if err := mapValFn(ctx, len(elems), opts...); err != nil {
			return errors.Wrap(err, "map interfaces")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items interfaces")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateInfra) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Infra)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Infra got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["availability_zone"]; exists {

		vOpts := append(opts, db.WithValidateField("availability_zone"))
		if err := fv(ctx, m.GetAvailabilityZone(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["certified_hw"]; exists {

		vOpts := append(opts, db.WithValidateField("certified_hw"))
		if err := fv(ctx, m.GetCertifiedHw(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["domain"]; exists {

		vOpts := append(opts, db.WithValidateField("domain"))
		if err := fv(ctx, m.GetDomain(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["hostname"]; exists {

		vOpts := append(opts, db.WithValidateField("hostname"))
		if err := fv(ctx, m.GetHostname(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["hw_info"]; exists {

		vOpts := append(opts, db.WithValidateField("hw_info"))
		if err := fv(ctx, m.GetHwInfo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["instance_id"]; exists {

		vOpts := append(opts, db.WithValidateField("instance_id"))
		if err := fv(ctx, m.GetInstanceId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["interfaces"]; exists {
		vOpts := append(opts, db.WithValidateField("interfaces"))
		if err := fv(ctx, m.GetInterfaces(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["internet_proxy"]; exists {

		vOpts := append(opts, db.WithValidateField("internet_proxy"))
		if err := fv(ctx, m.GetInternetProxy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["machine_id"]; exists {

		vOpts := append(opts, db.WithValidateField("machine_id"))
		if err := fv(ctx, m.GetMachineId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["provider"]; exists {

		vOpts := append(opts, db.WithValidateField("provider"))
		if err := fv(ctx, m.GetProvider(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("timestamp"))
		if err := fv(ctx, m.GetTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["zone"]; exists {

		vOpts := append(opts, db.WithValidateField("zone"))
		if err := fv(ctx, m.GetZone(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultInfraValidator = func() *ValidateInfra {
	v := &ValidateInfra{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhHostname := v.HostnameValidationRuleHandler
	rulesHostname := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhHostname(rulesHostname)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Infra.hostname: %s", err)
		panic(errMsg)
	}
	v.FldValidators["hostname"] = vFn

	vrhInterfaces := v.InterfacesValidationRuleHandler
	rulesInterfaces := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInterfaces(rulesInterfaces)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Infra.interfaces: %s", err)
		panic(errMsg)
	}
	v.FldValidators["interfaces"] = vFn

	return v
}()

func InfraValidator() db.Validator {
	return DefaultInfraValidator
}

// augmented methods on protoc/std generated struct

func (m *Interface) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Interface) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Interface) DeepCopy() *Interface {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Interface{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Interface) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Interface) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return InterfaceValidator().Validate(ctx, m, opts...)
}

type ValidateInterface struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateInterface) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Interface)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Interface got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["addr"]; exists {

		vOpts := append(opts, db.WithValidateField("addr"))
		if err := fv(ctx, m.GetAddr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["plen"]; exists {

		vOpts := append(opts, db.WithValidateField("plen"))
		if err := fv(ctx, m.GetPlen(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultInterfaceValidator = func() *ValidateInterface {
	v := &ValidateInterface{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func InterfaceValidator() db.Validator {
	return DefaultInterfaceValidator
}

// augmented methods on protoc/std generated struct

func (m *InternetProxy) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *InternetProxy) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *InternetProxy) DeepCopy() *InternetProxy {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &InternetProxy{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *InternetProxy) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *InternetProxy) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return InternetProxyValidator().Validate(ctx, m, opts...)
}

type ValidateInternetProxy struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateInternetProxy) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*InternetProxy)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *InternetProxy got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["http_proxy"]; exists {

		vOpts := append(opts, db.WithValidateField("http_proxy"))
		if err := fv(ctx, m.GetHttpProxy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["https_proxy"]; exists {

		vOpts := append(opts, db.WithValidateField("https_proxy"))
		if err := fv(ctx, m.GetHttpsProxy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["no_proxy"]; exists {

		vOpts := append(opts, db.WithValidateField("no_proxy"))
		if err := fv(ctx, m.GetNoProxy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_cacert_url"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_cacert_url"))
		if err := fv(ctx, m.GetProxyCacertUrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultInternetProxyValidator = func() *ValidateInternetProxy {
	v := &ValidateInternetProxy{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func InternetProxyValidator() db.Validator {
	return DefaultInternetProxyValidator
}

// augmented methods on protoc/std generated struct

func (m *Passport) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Passport) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Passport) DeepCopy() *Passport {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Passport{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Passport) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Passport) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PassportValidator().Validate(ctx, m, opts...)
}

type ValidatePassport struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePassport) ClusterNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cluster_name")
	}

	return validatorFn, nil
}

func (v *ValidatePassport) ClusterTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cluster_type")
	}

	return validatorFn, nil
}

func (v *ValidatePassport) LatitudeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewFloatValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for latitude")
	}

	return validatorFn, nil
}

func (v *ValidatePassport) LongitudeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewFloatValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for longitude")
	}

	return validatorFn, nil
}

func (v *ValidatePassport) ClusterSizeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewInt32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for cluster_size")
	}

	return validatorFn, nil
}

func (v *ValidatePassport) PrivateNetworkNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for private_network_name")
	}

	return validatorFn, nil
}

func (v *ValidatePassport) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Passport)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Passport got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cluster_name"]; exists {

		vOpts := append(opts, db.WithValidateField("cluster_name"))
		if err := fv(ctx, m.GetClusterName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["cluster_size"]; exists {

		vOpts := append(opts, db.WithValidateField("cluster_size"))
		if err := fv(ctx, m.GetClusterSize(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["cluster_type"]; exists {

		vOpts := append(opts, db.WithValidateField("cluster_type"))
		if err := fv(ctx, m.GetClusterType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["latitude"]; exists {

		vOpts := append(opts, db.WithValidateField("latitude"))
		if err := fv(ctx, m.GetLatitude(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["longitude"]; exists {

		vOpts := append(opts, db.WithValidateField("longitude"))
		if err := fv(ctx, m.GetLongitude(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["private_network_name"]; exists {

		vOpts := append(opts, db.WithValidateField("private_network_name"))
		if err := fv(ctx, m.GetPrivateNetworkName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vpm_version"]; exists {

		vOpts := append(opts, db.WithValidateField("vpm_version"))
		if err := fv(ctx, m.GetVpmVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPassportValidator = func() *ValidatePassport {
	v := &ValidatePassport{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhClusterName := v.ClusterNameValidationRuleHandler
	rulesClusterName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhClusterName(rulesClusterName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Passport.cluster_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cluster_name"] = vFn

	vrhClusterType := v.ClusterTypeValidationRuleHandler
	rulesClusterType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhClusterType(rulesClusterType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Passport.cluster_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cluster_type"] = vFn

	vrhLatitude := v.LatitudeValidationRuleHandler
	rulesLatitude := map[string]string{
		"ves.io.schema.rules.float.gte": "-90.0",
		"ves.io.schema.rules.float.lte": "90.0",
	}
	vFn, err = vrhLatitude(rulesLatitude)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Passport.latitude: %s", err)
		panic(errMsg)
	}
	v.FldValidators["latitude"] = vFn

	vrhLongitude := v.LongitudeValidationRuleHandler
	rulesLongitude := map[string]string{
		"ves.io.schema.rules.float.gte": "-180.0",
		"ves.io.schema.rules.float.lte": "180.0",
	}
	vFn, err = vrhLongitude(rulesLongitude)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Passport.longitude: %s", err)
		panic(errMsg)
	}
	v.FldValidators["longitude"] = vFn

	vrhClusterSize := v.ClusterSizeValidationRuleHandler
	rulesClusterSize := map[string]string{
		"ves.io.schema.rules.int32.in": "[0,1,3]",
	}
	vFn, err = vrhClusterSize(rulesClusterSize)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Passport.cluster_size: %s", err)
		panic(errMsg)
	}
	v.FldValidators["cluster_size"] = vFn

	vrhPrivateNetworkName := v.PrivateNetworkNameValidationRuleHandler
	rulesPrivateNetworkName := map[string]string{
		"ves.io.schema.rules.string.max_bytes": "64",
	}
	vFn, err = vrhPrivateNetworkName(rulesPrivateNetworkName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Passport.private_network_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["private_network_name"] = vFn

	return v
}()

func PassportValidator() db.Validator {
	return DefaultPassportValidator
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

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *WorkloadContext) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *WorkloadContext) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *WorkloadContext) DeepCopy() *WorkloadContext {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &WorkloadContext{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *WorkloadContext) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *WorkloadContext) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return WorkloadContextValidator().Validate(ctx, m, opts...)
}

type ValidateWorkloadContext struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateWorkloadContext) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*WorkloadContext)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *WorkloadContext got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["params"]; exists {

		vOpts := append(opts, db.WithValidateField("params"))
		for key, value := range m.GetParams() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultWorkloadContextValidator = func() *ValidateWorkloadContext {
	v := &ValidateWorkloadContext{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func WorkloadContextValidator() db.Validator {
	return DefaultWorkloadContextValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Infra = f.GetInfra()
	m.Passport = f.GetPassport()
	m.Token = f.GetToken()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Infra = m1.Infra
	f.Passport = m1.Passport
	f.Token = m1.Token
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Infra = f.GetInfra()
	m.Passport = f.GetPassport()
	m.Token = f.GetToken()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Infra = m1.Infra
	f.Passport = m1.Passport
	f.Token = m1.Token
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
}
