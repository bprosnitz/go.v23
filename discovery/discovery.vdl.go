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

func (x *AdId) VDLRead(dec vdl.Decoder) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	bytes := x[:]
	if err = dec.DecodeBytes(16, &bytes); err != nil {
		return err
	}
	return dec.FinishValue()
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

func (x *Attributes) VDLRead(dec vdl.Decoder) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if k := dec.Type().Kind(); k != vdl.Map {
		return fmt.Errorf("incompatible map %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len == 0:
		*x = nil
		return dec.FinishValue()
	case len > 0:
		*x = make(Attributes, len)
	default:
		*x = make(Attributes)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var key string
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		var elem string
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if elem, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		(*x)[key] = elem
	}
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

func (x *Attachments) VDLRead(dec vdl.Decoder) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if k := dec.Type().Kind(); k != vdl.Map {
		return fmt.Errorf("incompatible map %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len == 0:
		*x = nil
		return dec.FinishValue()
	case len > 0:
		*x = make(Attachments, len)
	default:
		*x = make(Attachments)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var key string
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		var elem []byte
		{
			if err = dec.StartValue(); err != nil {
				return err
			}
			if err = dec.DecodeBytes(-1, &elem); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		}
		(*x)[key] = elem
	}
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
	//
	// We limit the maximum number of attachments to 32.
	Attributes Attributes
	// Attachments as a key/value pair.
	// E.g., {'thumbnail': binary_data }.
	//
	// Unlike attributes, attachments are for binary data and they are not queryable.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	//
	// We limit the maximum number of attachments to 32 and the maximum size of each
	// attachment is 4K bytes.
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
	var4 := (m.Id == AdId{})
	if var4 {
		if err := fieldsTarget1.ZeroField("Id"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Id.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.InterfaceName == "")
	if var7 {
		if err := fieldsTarget1.ZeroField("InterfaceName"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("InterfaceName")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromString(string(m.InterfaceName), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var var10 bool
	if len(m.Addresses) == 0 {
		var10 = true
	}
	if var10 {
		if err := fieldsTarget1.ZeroField("Addresses"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Addresses")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget11, err := fieldTarget9.StartList(tt.NonOptional().Field(2).Type, len(m.Addresses))
			if err != nil {
				return err
			}
			for i, elem13 := range m.Addresses {
				elemTarget12, err := listTarget11.StartElem(i)
				if err != nil {
					return err
				}
				if err := elemTarget12.FromString(string(elem13), tt.NonOptional().Field(2).Type.Elem()); err != nil {
					return err
				}
				if err := listTarget11.FinishElem(elemTarget12); err != nil {
					return err
				}
			}
			if err := fieldTarget9.FinishList(listTarget11); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var var16 bool
	if len(m.Attributes) == 0 {
		var16 = true
	}
	if var16 {
		if err := fieldsTarget1.ZeroField("Attributes"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget14, fieldTarget15, err := fieldsTarget1.StartField("Attributes")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Attributes.FillVDLTarget(fieldTarget15, tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget14, fieldTarget15); err != nil {
				return err
			}
		}
	}
	var var19 bool
	if len(m.Attachments) == 0 {
		var19 = true
	}
	if var19 {
		if err := fieldsTarget1.ZeroField("Attachments"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget17, fieldTarget18, err := fieldsTarget1.StartField("Attachments")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Attachments.FillVDLTarget(fieldTarget18, tt.NonOptional().Field(4).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget17, fieldTarget18); err != nil {
				return err
			}
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
func (t *AdvertisementTarget) ZeroField(name string) error {
	switch name {
	case "Id":
		t.Value.Id = AdId{}
		return nil
	case "InterfaceName":
		t.Value.InterfaceName = ""
		return nil
	case "Addresses":
		t.Value.Addresses = []string(nil)
		return nil
	case "Attributes":
		t.Value.Attributes = Attributes(nil)
		return nil
	case "Attachments":
		t.Value.Attachments = Attachments(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/discovery.Advertisement", name)
	}
}
func (t *AdvertisementTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x *Advertisement) VDLRead(dec vdl.Decoder) error {
	*x = Advertisement{}
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if dec.Type().Kind() != vdl.Struct {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	match := 0
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			if match == 0 && dec.Type().NumField() > 0 {
				return fmt.Errorf("no matching fields in struct %T, from %v", *x, dec.Type())
			}
			return dec.FinishValue()
		case "Id":
			match++
			if err = x.Id.VDLRead(dec); err != nil {
				return err
			}
		case "InterfaceName":
			match++
			if err = dec.StartValue(); err != nil {
				return err
			}
			if x.InterfaceName, err = dec.DecodeString(); err != nil {
				return err
			}
			if err = dec.FinishValue(); err != nil {
				return err
			}
		case "Addresses":
			match++
			if err = __VDLRead1_list(dec, &x.Addresses); err != nil {
				return err
			}
		case "Attributes":
			match++
			if err = x.Attributes.VDLRead(dec); err != nil {
				return err
			}
		case "Attachments":
			match++
			if err = x.Attachments.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err = dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLRead1_list(dec vdl.Decoder, x *[]string) error {
	var err error
	if err = dec.StartValue(); err != nil {
		return err
	}
	if k := dec.Type().Kind(); k != vdl.Array && k != vdl.List {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len == 0:
		*x = nil
	case len > 0:
		*x = make([]string, 0, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem string
		if err = dec.StartValue(); err != nil {
			return err
		}
		if elem, err = dec.DecodeString(); err != nil {
			return err
		}
		if err = dec.FinishValue(); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
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
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*AdId)(nil))
	vdl.Register((*Attributes)(nil))
	vdl.Register((*Attachments)(nil))
	vdl.Register((*Advertisement)(nil))

	return struct{}{}
}
