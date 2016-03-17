// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: rpc

package rpc

import (
	"fmt"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/vdlroot/time"
	"v.io/v23/verror"
	"v.io/v23/vtrace"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Request describes the request header sent by the client to the server.  A
// non-zero request header is sent at the beginning of the RPC call, followed by
// the positional args.  Thereafter a zero request header is sent before each
// streaming arg, terminated by a non-zero request header with EndStreamArgs set
// to true.
type Request struct {
	// Suffix of the name used to identify the object hosting the service.
	Suffix string
	// Method to invoke on the service.
	Method string
	// NumPosArgs is the number of positional arguments, which follow this message
	// (and any blessings) on the request stream.
	NumPosArgs uint64
	// EndStreamArgs is true iff no more streaming arguments will be sent.  No
	// more data will be sent on the request stream.
	//
	// NOTE(bprosnitz): We can support multiple stream values per request (+response) header
	// efficiently by adding a NumExtraStreamArgs (+NumExtraStreamResults to response) field
	// that is the uint64 (number of stream args to send) - 1. The request is then zero when
	// exactly one streaming arg is sent. Since the request and response headers are small,
	// this is only likely necessary for frequently streaming small values.
	// See implementation in CL: 3913
	EndStreamArgs bool
	// Deadline after which the request should be cancelled.  This is a hint to
	// the server, to avoid wasted work.
	Deadline time.Deadline
	// GrantedBlessings are blessings bound to the principal running the server,
	// provided by the client.
	GrantedBlessings security.Blessings
	// TraceRequest maintains the vtrace context between clients and servers
	// and specifies additional parameters that control how tracing behaves.
	TraceRequest vtrace.Request
	// Language indicates the language of the instegator of the RPC.
	// By convention it should be an IETF language tag:
	// http://en.wikipedia.org/wiki/IETF_language_tag
	Language string
}

func (Request) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/rpc.Request"`
}) {
}

func (m *Request) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Suffix")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Suffix), tt.NonOptional().Field(0).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Method")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Method), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("NumPosArgs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromUint(uint64(m.NumPosArgs), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("EndStreamArgs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget9.FromBool(bool(m.EndStreamArgs), tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	var wireValue10 time.WireDeadline
	if err := time.WireDeadlineFromNative(&wireValue10, m.Deadline); err != nil {
		return err
	}

	keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("Deadline")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue10.FillVDLTarget(fieldTarget12, tt.NonOptional().Field(4).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
			return err
		}
	}
	var wireValue13 security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireValue13, m.GrantedBlessings); err != nil {
		return err
	}

	keyTarget14, fieldTarget15, err := fieldsTarget1.StartField("GrantedBlessings")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue13.FillVDLTarget(fieldTarget15, tt.NonOptional().Field(5).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget14, fieldTarget15); err != nil {
			return err
		}
	}
	keyTarget16, fieldTarget17, err := fieldsTarget1.StartField("TraceRequest")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.TraceRequest.FillVDLTarget(fieldTarget17, tt.NonOptional().Field(6).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget16, fieldTarget17); err != nil {
			return err
		}
	}
	keyTarget18, fieldTarget19, err := fieldsTarget1.StartField("Language")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget19.FromString(string(m.Language), tt.NonOptional().Field(7).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget18, fieldTarget19); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Request) MakeVDLTarget() vdl.Target {
	return &RequestTarget{Value: m}
}

type RequestTarget struct {
	Value                  *Request
	suffixTarget           vdl.StringTarget
	methodTarget           vdl.StringTarget
	numPosArgsTarget       vdl.Uint64Target
	endStreamArgsTarget    vdl.BoolTarget
	deadlineTarget         time.WireDeadlineTarget
	grantedBlessingsTarget security.WireBlessingsTarget
	traceRequestTarget     vtrace.RequestTarget
	languageTarget         vdl.StringTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *RequestTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Request)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *RequestTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Suffix":
		t.suffixTarget.Value = &t.Value.Suffix
		target, err := &t.suffixTarget, error(nil)
		return nil, target, err
	case "Method":
		t.methodTarget.Value = &t.Value.Method
		target, err := &t.methodTarget, error(nil)
		return nil, target, err
	case "NumPosArgs":
		t.numPosArgsTarget.Value = &t.Value.NumPosArgs
		target, err := &t.numPosArgsTarget, error(nil)
		return nil, target, err
	case "EndStreamArgs":
		t.endStreamArgsTarget.Value = &t.Value.EndStreamArgs
		target, err := &t.endStreamArgsTarget, error(nil)
		return nil, target, err
	case "Deadline":
		t.deadlineTarget.Value = &t.Value.Deadline
		target, err := &t.deadlineTarget, error(nil)
		return nil, target, err
	case "GrantedBlessings":
		t.grantedBlessingsTarget.Value = &t.Value.GrantedBlessings
		target, err := &t.grantedBlessingsTarget, error(nil)
		return nil, target, err
	case "TraceRequest":
		t.traceRequestTarget.Value = &t.Value.TraceRequest
		target, err := &t.traceRequestTarget, error(nil)
		return nil, target, err
	case "Language":
		t.languageTarget.Value = &t.Value.Language
		target, err := &t.languageTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/rpc.Request", name)
	}
}
func (t *RequestTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *RequestTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Response describes the response header sent by the server to the client.  A
// zero response header is sent before each streaming arg.  Thereafter a
// non-zero response header is sent at the end of the RPC call, right before
// the positional results.
type Response struct {
	// Error in processing the RPC at the server. Implies EndStreamResults.
	Error error
	// EndStreamResults is true iff no more streaming results will be sent; the
	// remainder of the stream consists of NumPosResults positional results.
	EndStreamResults bool
	// NumPosResults is the number of positional results, which immediately follow
	// on the response stream.  After these results, no further data will be sent
	// on the response stream.
	NumPosResults uint64
	// TraceResponse maintains the vtrace context between clients and servers.
	// In some cases trace data will be included in this response as well.
	TraceResponse vtrace.Response
	// AckBlessings is true if the server successfully recevied the client's
	// blessings and stored them in the server's blessings cache.
	AckBlessings bool
}

func (Response) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/rpc.Response"`
}) {
}

func (m *Response) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Error")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Error == nil {
			if err := fieldTarget3.FromNil(tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
		} else {
			var wireError4 vdl.WireError
			if err := verror.WireFromNative(&wireError4, m.Error); err != nil {
				return err
			}
			if err := wireError4.FillVDLTarget(fieldTarget3, vdl.ErrorType); err != nil {
				return err
			}

		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("EndStreamResults")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget6.FromBool(bool(m.EndStreamResults), tt.NonOptional().Field(1).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
			return err
		}
	}
	keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("NumPosResults")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget8.FromUint(uint64(m.NumPosResults), tt.NonOptional().Field(2).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
			return err
		}
	}
	keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("TraceResponse")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.TraceResponse.FillVDLTarget(fieldTarget10, tt.NonOptional().Field(3).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
			return err
		}
	}
	keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("AckBlessings")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget12.FromBool(bool(m.AckBlessings), tt.NonOptional().Field(4).Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Response) MakeVDLTarget() vdl.Target {
	return &ResponseTarget{Value: m}
}

type ResponseTarget struct {
	Value                  *Response
	errorTarget            verror.ErrorTarget
	endStreamResultsTarget vdl.BoolTarget
	numPosResultsTarget    vdl.Uint64Target
	traceResponseTarget    vtrace.ResponseTarget
	ackBlessingsTarget     vdl.BoolTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ResponseTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Response)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *ResponseTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Error":
		t.errorTarget.Value = &t.Value.Error
		target, err := &t.errorTarget, error(nil)
		return nil, target, err
	case "EndStreamResults":
		t.endStreamResultsTarget.Value = &t.Value.EndStreamResults
		target, err := &t.endStreamResultsTarget, error(nil)
		return nil, target, err
	case "NumPosResults":
		t.numPosResultsTarget.Value = &t.Value.NumPosResults
		target, err := &t.numPosResultsTarget, error(nil)
		return nil, target, err
	case "TraceResponse":
		t.traceResponseTarget.Value = &t.Value.TraceResponse
		target, err := &t.traceResponseTarget, error(nil)
		return nil, target, err
	case "AckBlessings":
		t.ackBlessingsTarget.Value = &t.Value.AckBlessings
		target, err := &t.ackBlessingsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/rpc.Response", name)
	}
}
func (t *ResponseTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ResponseTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

//////////////////////////////////////////////////
// Const definitions

// TODO(toddw): Rename GlobMethod to ReservedGlob.
const GlobMethod = "__Glob"
const ReservedSignature = "__Signature"
const ReservedMethodSignature = "__MethodSignature"

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
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))

	return struct{}{}
}