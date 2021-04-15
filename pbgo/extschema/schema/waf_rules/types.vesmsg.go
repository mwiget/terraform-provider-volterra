//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package waf_rules

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_waf_rule_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rule_list"
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

func (v *ValidateCreateSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.WafModeType)
		return int32(i)
	}
	// ves_io_schema.WafModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.WafModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) AnomalyScoreThresholdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for anomaly_score_threshold")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) ParanoiaLevelValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for paranoia_level")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) RuleIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_waf_rule_list.WafRuleID)
		return int32(i)
	}
	// ves_io_schema_waf_rule_list.WafRuleID_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema_waf_rule_list.WafRuleID_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema_waf_rule_list.WafRuleID, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for rule_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema_waf_rule_list.WafRuleID)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema_waf_rule_list.WafRuleID, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated rule_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items rule_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) RuleListTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(RuleModeType)
		return int32(i)
	}
	// RuleModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, RuleModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_list_type")
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

	if fv, exists := v.FldValidators["anomaly_score_threshold"]; exists {

		vOpts := append(opts, db.WithValidateField("anomaly_score_threshold"))
		if err := fv(ctx, m.GetAnomalyScoreThreshold(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["paranoia_level"]; exists {

		vOpts := append(opts, db.WithValidateField("paranoia_level"))
		if err := fv(ctx, m.GetParanoiaLevel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("rule_ids"))
		if err := fv(ctx, m.GetRuleIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_list_type"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_list_type"))
		if err := fv(ctx, m.GetRuleListType(), vOpts...); err != nil {
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

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAnomalyScoreThreshold := v.AnomalyScoreThresholdValidationRuleHandler
	rulesAnomalyScoreThreshold := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
	}
	vFn, err = vrhAnomalyScoreThreshold(rulesAnomalyScoreThreshold)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.anomaly_score_threshold: %s", err)
		panic(errMsg)
	}
	v.FldValidators["anomaly_score_threshold"] = vFn

	vrhParanoiaLevel := v.ParanoiaLevelValidationRuleHandler
	rulesParanoiaLevel := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
		"ves.io.schema.rules.uint32.lte":       "4",
	}
	vFn, err = vrhParanoiaLevel(rulesParanoiaLevel)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.paranoia_level: %s", err)
		panic(errMsg)
	}
	v.FldValidators["paranoia_level"] = vFn

	vrhRuleIds := v.RuleIdsValidationRuleHandler
	rulesRuleIds := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "256",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhRuleIds(rulesRuleIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.rule_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_ids"] = vFn

	vrhRuleListType := v.RuleListTypeValidationRuleHandler
	rulesRuleListType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhRuleListType(rulesRuleListType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.rule_list_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_list_type"] = vFn

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

func (v *ValidateGetSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.WafModeType)
		return int32(i)
	}
	// ves_io_schema.WafModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.WafModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) AnomalyScoreThresholdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for anomaly_score_threshold")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) ParanoiaLevelValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for paranoia_level")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) RuleIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_waf_rule_list.WafRuleID)
		return int32(i)
	}
	// ves_io_schema_waf_rule_list.WafRuleID_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema_waf_rule_list.WafRuleID_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema_waf_rule_list.WafRuleID, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for rule_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema_waf_rule_list.WafRuleID)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema_waf_rule_list.WafRuleID, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated rule_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items rule_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) RuleListTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(RuleModeType)
		return int32(i)
	}
	// RuleModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, RuleModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_list_type")
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

	if fv, exists := v.FldValidators["anomaly_score_threshold"]; exists {

		vOpts := append(opts, db.WithValidateField("anomaly_score_threshold"))
		if err := fv(ctx, m.GetAnomalyScoreThreshold(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["paranoia_level"]; exists {

		vOpts := append(opts, db.WithValidateField("paranoia_level"))
		if err := fv(ctx, m.GetParanoiaLevel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("rule_ids"))
		if err := fv(ctx, m.GetRuleIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_list_type"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_list_type"))
		if err := fv(ctx, m.GetRuleListType(), vOpts...); err != nil {
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

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAnomalyScoreThreshold := v.AnomalyScoreThresholdValidationRuleHandler
	rulesAnomalyScoreThreshold := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
	}
	vFn, err = vrhAnomalyScoreThreshold(rulesAnomalyScoreThreshold)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.anomaly_score_threshold: %s", err)
		panic(errMsg)
	}
	v.FldValidators["anomaly_score_threshold"] = vFn

	vrhParanoiaLevel := v.ParanoiaLevelValidationRuleHandler
	rulesParanoiaLevel := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
		"ves.io.schema.rules.uint32.lte":       "4",
	}
	vFn, err = vrhParanoiaLevel(rulesParanoiaLevel)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.paranoia_level: %s", err)
		panic(errMsg)
	}
	v.FldValidators["paranoia_level"] = vFn

	vrhRuleIds := v.RuleIdsValidationRuleHandler
	rulesRuleIds := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "256",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhRuleIds(rulesRuleIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.rule_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_ids"] = vFn

	vrhRuleListType := v.RuleListTypeValidationRuleHandler
	rulesRuleListType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhRuleListType(rulesRuleListType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.rule_list_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_list_type"] = vFn

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
	if fdrInfos, err := m.GetWafDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *GlobalSpecType) GetWafDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, ref := range m.GetWaf() {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.waf[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "waf.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "waf",
			Ref:        ref,
		})
	}

	return drInfos, nil
}

// GetWafDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetWafDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "waf.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: waf")
	}
	for _, ref := range m.GetWaf() {
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

func (v *ValidateGlobalSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.WafModeType)
		return int32(i)
	}
	// ves_io_schema.WafModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.WafModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) AnomalyScoreThresholdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for anomaly_score_threshold")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ParanoiaLevelValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for paranoia_level")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) RuleIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_waf_rule_list.WafRuleID)
		return int32(i)
	}
	// ves_io_schema_waf_rule_list.WafRuleID_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema_waf_rule_list.WafRuleID_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema_waf_rule_list.WafRuleID, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for rule_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema_waf_rule_list.WafRuleID)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema_waf_rule_list.WafRuleID, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated rule_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items rule_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) RuleListTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(RuleModeType)
		return int32(i)
	}
	// RuleModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, RuleModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_list_type")
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

	if fv, exists := v.FldValidators["anomaly_score_threshold"]; exists {

		vOpts := append(opts, db.WithValidateField("anomaly_score_threshold"))
		if err := fv(ctx, m.GetAnomalyScoreThreshold(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["paranoia_level"]; exists {

		vOpts := append(opts, db.WithValidateField("paranoia_level"))
		if err := fv(ctx, m.GetParanoiaLevel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("rule_ids"))
		if err := fv(ctx, m.GetRuleIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_list_type"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_list_type"))
		if err := fv(ctx, m.GetRuleListType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["waf"]; exists {

		vOpts := append(opts, db.WithValidateField("waf"))
		for idx, item := range m.GetWaf() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
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

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAnomalyScoreThreshold := v.AnomalyScoreThresholdValidationRuleHandler
	rulesAnomalyScoreThreshold := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
	}
	vFn, err = vrhAnomalyScoreThreshold(rulesAnomalyScoreThreshold)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.anomaly_score_threshold: %s", err)
		panic(errMsg)
	}
	v.FldValidators["anomaly_score_threshold"] = vFn

	vrhParanoiaLevel := v.ParanoiaLevelValidationRuleHandler
	rulesParanoiaLevel := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
		"ves.io.schema.rules.uint32.lte":       "4",
	}
	vFn, err = vrhParanoiaLevel(rulesParanoiaLevel)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.paranoia_level: %s", err)
		panic(errMsg)
	}
	v.FldValidators["paranoia_level"] = vFn

	vrhRuleIds := v.RuleIdsValidationRuleHandler
	rulesRuleIds := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "256",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhRuleIds(rulesRuleIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.rule_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_ids"] = vFn

	vrhRuleListType := v.RuleListTypeValidationRuleHandler
	rulesRuleListType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhRuleListType(rulesRuleListType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.rule_list_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_list_type"] = vFn

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

func (v *ValidateReplaceSpecType) ModeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.WafModeType)
		return int32(i)
	}
	// ves_io_schema.WafModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.WafModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mode")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) AnomalyScoreThresholdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for anomaly_score_threshold")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) ParanoiaLevelValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for paranoia_level")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) RuleIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema_waf_rule_list.WafRuleID)
		return int32(i)
	}
	// ves_io_schema_waf_rule_list.WafRuleID_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema_waf_rule_list.WafRuleID_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema_waf_rule_list.WafRuleID, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for rule_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema_waf_rule_list.WafRuleID)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema_waf_rule_list.WafRuleID, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated rule_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items rule_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) RuleListTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(RuleModeType)
		return int32(i)
	}
	// RuleModeType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, RuleModeType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for rule_list_type")
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

	if fv, exists := v.FldValidators["anomaly_score_threshold"]; exists {

		vOpts := append(opts, db.WithValidateField("anomaly_score_threshold"))
		if err := fv(ctx, m.GetAnomalyScoreThreshold(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["paranoia_level"]; exists {

		vOpts := append(opts, db.WithValidateField("paranoia_level"))
		if err := fv(ctx, m.GetParanoiaLevel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("rule_ids"))
		if err := fv(ctx, m.GetRuleIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_list_type"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_list_type"))
		if err := fv(ctx, m.GetRuleListType(), vOpts...); err != nil {
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

	vrhMode := v.ModeValidationRuleHandler
	rulesMode := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMode(rulesMode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.mode: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mode"] = vFn

	vrhAnomalyScoreThreshold := v.AnomalyScoreThresholdValidationRuleHandler
	rulesAnomalyScoreThreshold := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
	}
	vFn, err = vrhAnomalyScoreThreshold(rulesAnomalyScoreThreshold)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.anomaly_score_threshold: %s", err)
		panic(errMsg)
	}
	v.FldValidators["anomaly_score_threshold"] = vFn

	vrhParanoiaLevel := v.ParanoiaLevelValidationRuleHandler
	rulesParanoiaLevel := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.uint32.gt":        "0",
		"ves.io.schema.rules.uint32.lte":       "4",
	}
	vFn, err = vrhParanoiaLevel(rulesParanoiaLevel)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.paranoia_level: %s", err)
		panic(errMsg)
	}
	v.FldValidators["paranoia_level"] = vFn

	vrhRuleIds := v.RuleIdsValidationRuleHandler
	rulesRuleIds := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "256",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhRuleIds(rulesRuleIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.rule_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_ids"] = vFn

	vrhRuleListType := v.RuleListTypeValidationRuleHandler
	rulesRuleListType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhRuleListType(rulesRuleListType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.rule_list_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["rule_list_type"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Rules) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Rules) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Rules) DeepCopy() *Rules {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Rules{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Rules) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Rules) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RulesValidator().Validate(ctx, m, opts...)
}

type ValidateRules struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRules) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Rules)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Rules got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mode"]; exists {

		vOpts := append(opts, db.WithValidateField("mode"))
		if err := fv(ctx, m.GetMode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["severity"]; exists {

		vOpts := append(opts, db.WithValidateField("severity"))
		if err := fv(ctx, m.GetSeverity(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tags"]; exists {

		vOpts := append(opts, db.WithValidateField("tags"))
		for idx, item := range m.GetTags() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRulesValidator = func() *ValidateRules {
	v := &ValidateRules{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RulesValidator() db.Validator {
	return DefaultRulesValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.AnomalyScoreThreshold = f.GetAnomalyScoreThreshold()
	m.Mode = f.GetMode()
	m.ParanoiaLevel = f.GetParanoiaLevel()
	m.RuleIds = f.GetRuleIds()
	m.RuleListType = f.GetRuleListType()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.AnomalyScoreThreshold = m1.AnomalyScoreThreshold
	f.Mode = m1.Mode
	f.ParanoiaLevel = m1.ParanoiaLevel
	f.RuleIds = m1.RuleIds
	f.RuleListType = m1.RuleListType
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.AnomalyScoreThreshold = f.GetAnomalyScoreThreshold()
	m.Mode = f.GetMode()
	m.ParanoiaLevel = f.GetParanoiaLevel()
	m.RuleIds = f.GetRuleIds()
	m.RuleListType = f.GetRuleListType()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.AnomalyScoreThreshold = m1.AnomalyScoreThreshold
	f.Mode = m1.Mode
	f.ParanoiaLevel = m1.ParanoiaLevel
	f.RuleIds = m1.RuleIds
	f.RuleListType = m1.RuleListType
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.AnomalyScoreThreshold = f.GetAnomalyScoreThreshold()
	m.Mode = f.GetMode()
	m.ParanoiaLevel = f.GetParanoiaLevel()
	m.RuleIds = f.GetRuleIds()
	m.RuleListType = f.GetRuleListType()
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.AnomalyScoreThreshold = m1.AnomalyScoreThreshold
	f.Mode = m1.Mode
	f.ParanoiaLevel = m1.ParanoiaLevel
	f.RuleIds = m1.RuleIds
	f.RuleListType = m1.RuleListType
}
