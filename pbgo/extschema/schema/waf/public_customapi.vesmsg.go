//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package waf

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_waf_rules "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rules"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *VirtualHostWafStatusReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *VirtualHostWafStatusReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *VirtualHostWafStatusReq) DeepCopy() *VirtualHostWafStatusReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &VirtualHostWafStatusReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *VirtualHostWafStatusReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *VirtualHostWafStatusReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return VirtualHostWafStatusReqValidator().Validate(ctx, m, opts...)
}

type ValidateVirtualHostWafStatusReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVirtualHostWafStatusReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*VirtualHostWafStatusReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *VirtualHostWafStatusReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultVirtualHostWafStatusReqValidator = func() *ValidateVirtualHostWafStatusReq {
	v := &ValidateVirtualHostWafStatusReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func VirtualHostWafStatusReqValidator() db.Validator {
	return DefaultVirtualHostWafStatusReqValidator
}

// augmented methods on protoc/std generated struct

func (m *VirtualHostWafStatusRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *VirtualHostWafStatusRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *VirtualHostWafStatusRsp) DeepCopy() *VirtualHostWafStatusRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &VirtualHostWafStatusRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *VirtualHostWafStatusRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *VirtualHostWafStatusRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return VirtualHostWafStatusRspValidator().Validate(ctx, m, opts...)
}

func (m *VirtualHostWafStatusRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetWafStatusDRefInfo()

}

// GetDRefInfo for the field's type
func (m *VirtualHostWafStatusRsp) GetWafStatusDRefInfo() ([]db.DRefInfo, error) {
	if m.GetWafStatus() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetWafStatus() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetWafStatus() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("waf_status[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateVirtualHostWafStatusRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVirtualHostWafStatusRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*VirtualHostWafStatusRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *VirtualHostWafStatusRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["waf_status"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_status"))
		for idx, item := range m.GetWafStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultVirtualHostWafStatusRspValidator = func() *ValidateVirtualHostWafStatusRsp {
	v := &ValidateVirtualHostWafStatusRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["waf_status"] = WafStatusValidator().Validate

	return v
}()

func VirtualHostWafStatusRspValidator() db.Validator {
	return DefaultVirtualHostWafStatusRspValidator
}

// augmented methods on protoc/std generated struct

func (m *WafStatus) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *WafStatus) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *WafStatus) DeepCopy() *WafStatus {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &WafStatus{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *WafStatus) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *WafStatus) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return WafStatusValidator().Validate(ctx, m, opts...)
}

func (m *WafStatus) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetWafRefDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetWafRefDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetWafRulesRefDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetWafRulesRefDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *WafStatus) GetWafRefDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetWafRef()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("WafStatus.waf_ref[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "waf.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "waf_ref",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetWafRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *WafStatus) GetWafRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "waf.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: waf")
	}
	for _, ref := range m.GetWafRef() {
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

func (m *WafStatus) GetWafRulesRefDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetWafRulesRef()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("WafStatus.waf_rules_ref[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "waf_rules.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "waf_rules_ref",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetWafRulesRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *WafStatus) GetWafRulesRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "waf_rules.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: waf_rules")
	}
	for _, ref := range m.GetWafRulesRef() {
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

type ValidateWafStatus struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateWafStatus) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*WafStatus)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *WafStatus got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["waf_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_ref"))
		for idx, item := range m.GetWafRef() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["waf_rules_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_rules_ref"))
		for idx, item := range m.GetWafRulesRef() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["waf_rules_status"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_rules_status"))
		for idx, item := range m.GetWafRulesStatus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultWafStatusValidator = func() *ValidateWafStatus {
	v := &ValidateWafStatus{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["waf_rules_status"] = ves_io_schema_waf_rules.WafRulesStatusValidator().Validate

	return v
}()

func WafStatusValidator() db.Validator {
	return DefaultWafStatusValidator
}
