// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vdl

package vdl

import (
	"fmt"
	"reflect"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// WireRetryCode is the suggested retry behavior for the receiver of an error.
// If the receiver doesn't know how to handle the specific error, it should
// attempt the suggested retry behavior.
type WireRetryCode int

const (
	WireRetryCodeNoRetry WireRetryCode = iota
	WireRetryCodeRetryConnection
	WireRetryCodeRetryRefetch
	WireRetryCodeRetryBackoff
)

// WireRetryCodeAll holds all labels for WireRetryCode.
var WireRetryCodeAll = [...]WireRetryCode{WireRetryCodeNoRetry, WireRetryCodeRetryConnection, WireRetryCodeRetryRefetch, WireRetryCodeRetryBackoff}

// WireRetryCodeFromString creates a WireRetryCode from a string label.
func WireRetryCodeFromString(label string) (x WireRetryCode, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *WireRetryCode) Set(label string) error {
	switch label {
	case "NoRetry", "noretry":
		*x = WireRetryCodeNoRetry
		return nil
	case "RetryConnection", "retryconnection":
		*x = WireRetryCodeRetryConnection
		return nil
	case "RetryRefetch", "retryrefetch":
		*x = WireRetryCodeRetryRefetch
		return nil
	case "RetryBackoff", "retrybackoff":
		*x = WireRetryCodeRetryBackoff
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdl.WireRetryCode", label)
}

// String returns the string label of x.
func (x WireRetryCode) String() string {
	switch x {
	case WireRetryCodeNoRetry:
		return "NoRetry"
	case WireRetryCodeRetryConnection:
		return "RetryConnection"
	case WireRetryCodeRetryRefetch:
		return "RetryRefetch"
	case WireRetryCodeRetryBackoff:
		return "RetryBackoff"
	}
	return ""
}

func (WireRetryCode) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vdl.WireRetryCode"`
	Enum struct{ NoRetry, RetryConnection, RetryRefetch, RetryBackoff string }
}) {
}

func (m *WireRetryCode) FillVDLTarget(t Target, tt *Type) error {
	if err := t.FromEnumLabel((*m).String(), tt); err != nil {
		return err
	}
	return nil
}

func (m *WireRetryCode) MakeVDLTarget() Target {
	return &WireRetryCodeTarget{Value: m}
}

type WireRetryCodeTarget struct {
	Value *WireRetryCode
	TargetBase
}

func (t *WireRetryCodeTarget) FromEnumLabel(src string, tt *Type) error {

	if ttWant := TypeOf((*WireRetryCode)(nil)); !Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	switch src {
	case "NoRetry":
		*t.Value = 0
	case "RetryConnection":
		*t.Value = 1
	case "RetryRefetch":
		*t.Value = 2
	case "RetryBackoff":
		*t.Value = 3
	default:
		return fmt.Errorf("label %s not in enum WireRetryCode", src)
	}

	return nil
}

func (x WireRetryCode) VDLIsZero() bool {
	return x == WireRetryCodeNoRetry
}

func (x WireRetryCode) VDLWrite(enc Encoder) error {
	if err := enc.StartValue(TypeOf((*WireRetryCode)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeString(x.String()); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireRetryCode) VDLRead(dec Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	enum, err := dec.DecodeString()
	if err != nil {
		return err
	}
	if err := x.Set(enum); err != nil {
		return err
	}
	return dec.FinishValue()
}

// WireError is the wire representation for the built-in error type.  Errors and
// exceptions in each programming environment are converted to this type to
// ensure wire compatibility.  Generated code for each environment provides
// automatic conversions into idiomatic native representations.
type WireError struct {
	Id        string        // Error Id, used to uniquely identify each error.
	RetryCode WireRetryCode // Retry behavior suggested for the receiver.
	Msg       string        // Error message, may be empty.
	ParamList []*Value      // Variadic parameters contained in the error.
}

func (WireError) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vdl.WireError"`
}) {
}

func (m *WireError) FillVDLTarget(t Target, tt *Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Id == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Id"); err != nil && err != ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
		if err != ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Id), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.RetryCode == WireRetryCodeNoRetry)
	if var7 {
		if err := fieldsTarget1.ZeroField("RetryCode"); err != nil && err != ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("RetryCode")
		if err != ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.RetryCode.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var10 := (m.Msg == "")
	if var10 {
		if err := fieldsTarget1.ZeroField("Msg"); err != nil && err != ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Msg")
		if err != ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget9.FromString(string(m.Msg), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var var13 bool
	if len(m.ParamList) == 0 {
		var13 = true
	}
	if var13 {
		if err := fieldsTarget1.ZeroField("ParamList"); err != nil && err != ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("ParamList")
		if err != ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget14, err := fieldTarget12.StartList(tt.NonOptional().Field(3).Type, len(m.ParamList))
			if err != nil {
				return err
			}
			for i, elem16 := range m.ParamList {
				elemTarget15, err := listTarget14.StartElem(i)
				if err != nil {
					return err
				}

				if err := FromValue(elemTarget15, elem16); err != nil {
					return err
				}
				if err := listTarget14.FinishElem(elemTarget15); err != nil {
					return err
				}
			}
			if err := fieldTarget12.FinishList(listTarget14); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireError) MakeVDLTarget() Target {
	return &WireErrorTarget{Value: m}
}

type WireErrorTarget struct {
	Value           *WireError
	idTarget        StringTarget
	retryCodeTarget WireRetryCodeTarget
	msgTarget       StringTarget
	paramListTarget __VDLTarget1_list
	TargetBase
	FieldsTargetBase
}

func (t *WireErrorTarget) StartFields(tt *Type) (FieldsTarget, error) {

	if ttWant := TypeOf((*WireError)(nil)).Elem(); !Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *WireErrorTarget) StartField(name string) (key, field Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "RetryCode":
		t.retryCodeTarget.Value = &t.Value.RetryCode
		target, err := &t.retryCodeTarget, error(nil)
		return nil, target, err
	case "Msg":
		t.msgTarget.Value = &t.Value.Msg
		target, err := &t.msgTarget, error(nil)
		return nil, target, err
	case "ParamList":
		t.paramListTarget.Value = &t.Value.ParamList
		target, err := &t.paramListTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/vdl.WireError", name)
	}
}
func (t *WireErrorTarget) FinishField(_, _ Target) error {
	return nil
}
func (t *WireErrorTarget) ZeroField(name string) error {
	switch name {
	case "Id":
		t.Value.Id = ""
		return nil
	case "RetryCode":
		t.Value.RetryCode = WireRetryCodeNoRetry
		return nil
	case "Msg":
		t.Value.Msg = ""
		return nil
	case "ParamList":
		t.Value.ParamList = []*Value(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/vdl.WireError", name)
	}
}
func (t *WireErrorTarget) FinishFields(_ FieldsTarget) error {

	return nil
}

// []*Value
type __VDLTarget1_list struct {
	Value *[]*Value

	TargetBase
	ListTargetBase
}

func (t *__VDLTarget1_list) StartList(tt *Type, len int) (ListTarget, error) {

	if ttWant := TypeOf((*[]*Value)(nil)); !Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]*Value, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget1_list) StartElem(index int) (elem Target, _ error) {
	target, err := ReflectTarget(reflect.ValueOf(&(*t.Value)[index]))
	return target, err
}
func (t *__VDLTarget1_list) FinishElem(elem Target) error {
	return nil
}
func (t *__VDLTarget1_list) FinishList(elem ListTarget) error {

	return nil
}

func (x WireError) VDLIsZero() bool {
	if x.Id != "" {
		return false
	}
	if x.RetryCode != WireRetryCodeNoRetry {
		return false
	}
	if x.Msg != "" {
		return false
	}
	if len(x.ParamList) != 0 {
		return false
	}
	return true
}

func (x WireError) VDLWrite(enc Encoder) error {
	if err := enc.StartValue(TypeOf((*WireError)(nil)).Elem()); err != nil {
		return err
	}
	if x.Id != "" {
		if err := enc.NextField("Id"); err != nil {
			return err
		}
		if err := enc.StartValue(StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Id); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.RetryCode != WireRetryCodeNoRetry {
		if err := enc.NextField("RetryCode"); err != nil {
			return err
		}
		if err := x.RetryCode.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Msg != "" {
		if err := enc.NextField("Msg"); err != nil {
			return err
		}
		if err := enc.StartValue(StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Msg); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.ParamList) != 0 {
		if err := enc.NextField("ParamList"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.ParamList); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc Encoder, x []*Value) error {
	if err := enc.StartValue(TypeOf((*[]*Value)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for i := 0; i < len(x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if x[i] == nil {
			if err := enc.NilValue(AnyType); err != nil {
				return err
			}
		} else {
			if err := x[i].VDLWrite(enc); err != nil {
				return err
			}
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireError) VDLRead(dec Decoder) error {
	*x = WireError{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !Compatible(TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Id":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Id, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "RetryCode":
			if err := x.RetryCode.VDLRead(dec); err != nil {
				return err
			}
		case "Msg":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Msg, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "ParamList":
			if err := __VDLReadAnon_list_1(dec, &x.ParamList); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec Decoder, x *[]*Value) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !Compatible(TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([]*Value, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem *Value
		elem = new(Value)
		if err := elem.VDLRead(dec); err != nil {
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
	Register((*WireRetryCode)(nil))
	Register((*WireError)(nil))

	return struct{}{}
}
