// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: discovery

package discovery

import (
	"fmt"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// An AdId is a globally unique identifier of an advertisement.
type AdId [16]byte

func (AdId) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.AdId"`
}) {
}

func (m *AdId) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromBytes([]byte((*m)[:]), tt); err != nil {
		return err
	}
	return nil
}

func (m *AdId) MakeVDLTarget() vdl.Target {
	return &AdIdTarget{Value: m}
}

type AdIdTarget struct {
	Value *AdId
	vdl.TargetBase
}

func (t *AdIdTarget) FromBytes(src []byte, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*AdId)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	copy((*t.Value)[:], src)

	return nil
}

// Attributes represents service attributes as a key/value pair.
type Attributes map[string]string

func (Attributes) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attributes"`
}) {
}

func (m *Attributes) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	mapTarget1, err := t.StartMap(tt, len((*m)))
	if err != nil {
		return err
	}
	for key3, value5 := range *m {
		keyTarget2, err := mapTarget1.StartKey()
		if err != nil {
			return err
		}
		if err := keyTarget2.FromString(string(key3), tt.NonOptional().Key()); err != nil {
			return err
		}
		valueTarget4, err := mapTarget1.FinishKeyStartField(keyTarget2)
		if err != nil {
			return err
		}
		if err := valueTarget4.FromString(string(value5), tt.NonOptional().Elem()); err != nil {
			return err
		}
		if err := mapTarget1.FinishField(keyTarget2, valueTarget4); err != nil {
			return err
		}
	}
	if err := t.FinishMap(mapTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Attributes) MakeVDLTarget() vdl.Target {
	return &AttributesTarget{Value: m}
}

type AttributesTarget struct {
	Value      *Attributes
	currKey    string
	currElem   string
	keyTarget  vdl.StringTarget
	elemTarget vdl.StringTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *AttributesTarget) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*Attributes)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(Attributes)
	return t, nil
}
func (t *AttributesTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *AttributesTarget) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = ""
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *AttributesTarget) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *AttributesTarget) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// Attachments represents service attachments as a key/value pair.
type Attachments map[string][]byte

func (Attachments) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attachments"`
}) {
}

func (m *Attachments) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	mapTarget1, err := t.StartMap(tt, len((*m)))
	if err != nil {
		return err
	}
	for key3, value5 := range *m {
		keyTarget2, err := mapTarget1.StartKey()
		if err != nil {
			return err
		}
		if err := keyTarget2.FromString(string(key3), tt.NonOptional().Key()); err != nil {
			return err
		}
		valueTarget4, err := mapTarget1.FinishKeyStartField(keyTarget2)
		if err != nil {
			return err
		}

		if err := valueTarget4.FromBytes([]byte(value5), tt.NonOptional().Elem()); err != nil {
			return err
		}
		if err := mapTarget1.FinishField(keyTarget2, valueTarget4); err != nil {
			return err
		}
	}
	if err := t.FinishMap(mapTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Attachments) MakeVDLTarget() vdl.Target {
	return &AttachmentsTarget{Value: m}
}

type AttachmentsTarget struct {
	Value      *Attachments
	currKey    string
	currElem   []byte
	keyTarget  vdl.StringTarget
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *AttachmentsTarget) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*Attachments)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(Attachments)
	return t, nil
}
func (t *AttachmentsTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *AttachmentsTarget) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = []byte(nil)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *AttachmentsTarget) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *AttachmentsTarget) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// Advertisement represents a feed into advertiser to broadcast its contents
// to scanners.
//
// A large advertisement may require additional RPC calls causing delay in
// discovery. We limit the maximum size of an advertisement to 512 bytes
// excluding id and attachments.
type Advertisement struct {
	// Universal unique identifier of the advertisement.
	// If this is not specified, a random unique identifier will be assigned.
	Id AdId
	// Interface name that the advertised service implements.
	// E.g., 'v.io/v23/services/vtrace.Store'.
	InterfaceName string
	// Addresses (vanadium object names) that the advertised service is served on.
	// E.g., '/host:port/a/b/c', '/ns.dev.v.io:8101/blah/blah'.
	Addresses []string
	// Attributes as a key/value pair.
	// E.g., {'resolution': '1024x768'}.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	Attributes Attributes
	// Attachments as a key/value pair.
	// E.g., {'thumbnail': binary_data }.
	//
	// Unlike attributes, attachments are for binary data and they are not queryable.
	// We limit the maximum size of a single attachment to 4K bytes.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	Attachments Attachments
}

func (Advertisement) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Advertisement"`
}) {
}

func (m *Advertisement) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Id.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("InterfaceName")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.InterfaceName), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Addresses")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget8, err := fieldTarget7.StartList(tt.NonOptional().Field(2).Type, len(m.Addresses))
		if err != nil {
			return err
		}
		for i, elem10 := range m.Addresses {
			elemTarget9, err := listTarget8.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget9.FromString(string(elem10), tt.NonOptional().Field(2).Type.Elem()); err != nil {
				return err
			}
			if err := listTarget8.FinishElem(elemTarget9); err != nil {
				return err
			}
		}
		if err := fieldTarget7.FinishList(listTarget8); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("Attributes")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Attributes.FillVDLTarget(fieldTarget12, tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
			return err
		}
	}
	keyTarget13, fieldTarget14, err := fieldsTarget1.StartField("Attachments")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Attachments.FillVDLTarget(fieldTarget14, tt.NonOptional().Field(4).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget13, fieldTarget14); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Advertisement) MakeVDLTarget() vdl.Target {
	return &AdvertisementTarget{Value: m}
}

type AdvertisementTarget struct {
	Value               *Advertisement
	idTarget            AdIdTarget
	interfaceNameTarget vdl.StringTarget
	addressesTarget     vdl.StringSliceTarget
	attributesTarget    AttributesTarget
	attachmentsTarget   AttachmentsTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *AdvertisementTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Advertisement)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *AdvertisementTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "InterfaceName":
		t.interfaceNameTarget.Value = &t.Value.InterfaceName
		target, err := &t.interfaceNameTarget, error(nil)
		return nil, target, err
	case "Addresses":
		t.addressesTarget.Value = &t.Value.Addresses
		target, err := &t.addressesTarget, error(nil)
		return nil, target, err
	case "Attributes":
		t.attributesTarget.Value = &t.Value.Attributes
		target, err := &t.attributesTarget, error(nil)
		return nil, target, err
	case "Attachments":
		t.attachmentsTarget.Value = &t.Value.Attachments
		target, err := &t.attachmentsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/discovery.Advertisement", name)
	}
}
func (t *AdvertisementTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *AdvertisementTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}

	// Register types.
	vdl.Register((*AdId)(nil))
	vdl.Register((*Attributes)(nil))
	vdl.Register((*Attachments)(nil))
	vdl.Register((*Advertisement)(nil))

	return struct{}{}
}