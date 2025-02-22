//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package virtual_site

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
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

func (m *CreateSpecType) GetSRefInfo() ([]db.SelrFldInfo, error) {
	if m == nil {
		return nil, nil
	}
	return m.GetSiteSelectorSRefInfo()

}

// GetSiteSelectorSRefInfo returns the selector info (fld-name/val, selectee-type) of this field
func (m *CreateSpecType) GetSiteSelectorSRefInfo() ([]db.SelrFldInfo, error) {
	f := m.GetSiteSelector()
	if f == nil {
		return nil, nil
	}
	return []db.SelrFldInfo{
		{
			Name:  "CreateSpecType.site_selector",
			Kind:  "site.Object",
			Value: strings.Join(f.Expressions, ","),
			Ref:   f,
		},
	}, nil
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) SiteSelectorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for site_selector")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.LabelSelectorTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) SiteTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_site.SiteType)
		return int32(i)
	}
	// ves_io_schema_site.SiteType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema_site.SiteType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site_type")
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

	if fv, exists := v.FldValidators["site_selector"]; exists {

		vOpts := append(opts, db.WithValidateField("site_selector"))
		if err := fv(ctx, m.GetSiteSelector(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_type"]; exists {

		vOpts := append(opts, db.WithValidateField("site_type"))
		if err := fv(ctx, m.GetSiteType(), vOpts...); err != nil {
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

	vrhSiteSelector := v.SiteSelectorValidationRuleHandler
	rulesSiteSelector := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteSelector(rulesSiteSelector)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.site_selector: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_selector"] = vFn

	vrhSiteType := v.SiteTypeValidationRuleHandler
	rulesSiteType := map[string]string{
		"ves.io.schema.rules.enum.not_in":      "0",
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteType(rulesSiteType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.site_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_type"] = vFn

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

func (m *GetSpecType) GetSRefInfo() ([]db.SelrFldInfo, error) {
	if m == nil {
		return nil, nil
	}
	return m.GetSiteSelectorSRefInfo()

}

// GetSiteSelectorSRefInfo returns the selector info (fld-name/val, selectee-type) of this field
func (m *GetSpecType) GetSiteSelectorSRefInfo() ([]db.SelrFldInfo, error) {
	f := m.GetSiteSelector()
	if f == nil {
		return nil, nil
	}
	return []db.SelrFldInfo{
		{
			Name:  "GetSpecType.site_selector",
			Kind:  "site.Object",
			Value: strings.Join(f.Expressions, ","),
			Ref:   f,
		},
	}, nil
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) SiteSelectorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for site_selector")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.LabelSelectorTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) SiteTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_site.SiteType)
		return int32(i)
	}
	// ves_io_schema_site.SiteType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema_site.SiteType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site_type")
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

	if fv, exists := v.FldValidators["site_selector"]; exists {

		vOpts := append(opts, db.WithValidateField("site_selector"))
		if err := fv(ctx, m.GetSiteSelector(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_type"]; exists {

		vOpts := append(opts, db.WithValidateField("site_type"))
		if err := fv(ctx, m.GetSiteType(), vOpts...); err != nil {
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

	vrhSiteSelector := v.SiteSelectorValidationRuleHandler
	rulesSiteSelector := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteSelector(rulesSiteSelector)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.site_selector: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_selector"] = vFn

	vrhSiteType := v.SiteTypeValidationRuleHandler
	rulesSiteType := map[string]string{
		"ves.io.schema.rules.enum.not_in":      "0",
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteType(rulesSiteType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.site_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_type"] = vFn

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

func (m *GlobalSpecType) GetSRefInfo() ([]db.SelrFldInfo, error) {
	if m == nil {
		return nil, nil
	}
	return m.GetSiteSelectorSRefInfo()

}

// GetSiteSelectorSRefInfo returns the selector info (fld-name/val, selectee-type) of this field
func (m *GlobalSpecType) GetSiteSelectorSRefInfo() ([]db.SelrFldInfo, error) {
	f := m.GetSiteSelector()
	if f == nil {
		return nil, nil
	}
	return []db.SelrFldInfo{
		{
			Name:  "GlobalSpecType.site_selector",
			Kind:  "site.Object",
			Value: strings.Join(f.Expressions, ","),
			Ref:   f,
		},
	}, nil
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) SiteSelectorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for site_selector")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := ves_io_schema.LabelSelectorTypeValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) SiteTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_site.SiteType)
		return int32(i)
	}
	// ves_io_schema_site.SiteType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema_site.SiteType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site_type")
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

	if fv, exists := v.FldValidators["site_selector"]; exists {

		vOpts := append(opts, db.WithValidateField("site_selector"))
		if err := fv(ctx, m.GetSiteSelector(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_type"]; exists {

		vOpts := append(opts, db.WithValidateField("site_type"))
		if err := fv(ctx, m.GetSiteType(), vOpts...); err != nil {
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

	vrhSiteSelector := v.SiteSelectorValidationRuleHandler
	rulesSiteSelector := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteSelector(rulesSiteSelector)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.site_selector: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_selector"] = vFn

	vrhSiteType := v.SiteTypeValidationRuleHandler
	rulesSiteType := map[string]string{
		"ves.io.schema.rules.enum.not_in":      "0",
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteType(rulesSiteType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.site_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_type"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
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

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.SiteSelector = f.GetSiteSelector()
	m.SiteType = f.GetSiteType()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.SiteSelector = m1.SiteSelector
	f.SiteType = m1.SiteType
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.SiteSelector = f.GetSiteSelector()
	m.SiteType = f.GetSiteType()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.SiteSelector = m1.SiteSelector
	f.SiteType = m1.SiteType
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
