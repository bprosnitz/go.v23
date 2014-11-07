// This file was auto-generated by the veyron vdl tool.
// Source: advanced.vdl

package arith

import (
	"veyron.io/veyron/veyron2/vdl/testdata/arith/exp"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// TrigonometryClientMethods is the client interface
// containing Trigonometry methods.
//
// Trigonometry is an interface that specifies a couple trigonometric functions.
type TrigonometryClientMethods interface {
	Sine(ctx __context.T, angle float64, opts ...__ipc.CallOpt) (float64, error)
	Cosine(ctx __context.T, angle float64, opts ...__ipc.CallOpt) (float64, error)
}

// TrigonometryClientStub adds universal methods to TrigonometryClientMethods.
type TrigonometryClientStub interface {
	TrigonometryClientMethods
	__ipc.UniversalServiceMethods
}

// TrigonometryClient returns a client stub for Trigonometry.
func TrigonometryClient(name string, opts ...__ipc.BindOpt) TrigonometryClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implTrigonometryClientStub{name, client}
}

type implTrigonometryClientStub struct {
	name   string
	client __ipc.Client
}

func (c implTrigonometryClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implTrigonometryClientStub) Sine(ctx __context.T, i0 float64, opts ...__ipc.CallOpt) (o0 float64, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Sine", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTrigonometryClientStub) Cosine(ctx __context.T, i0 float64, opts ...__ipc.CallOpt) (o0 float64, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Cosine", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTrigonometryClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implTrigonometryClientStub) GetMethodTags(ctx __context.T, method string, opts ...__ipc.CallOpt) (o0 []interface{}, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// TrigonometryServerMethods is the interface a server writer
// implements for Trigonometry.
//
// Trigonometry is an interface that specifies a couple trigonometric functions.
type TrigonometryServerMethods interface {
	Sine(ctx __ipc.ServerContext, angle float64) (float64, error)
	Cosine(ctx __ipc.ServerContext, angle float64) (float64, error)
}

// TrigonometryServerStubMethods is the server interface containing
// Trigonometry methods, as expected by ipc.Server.  The difference between
// this interface and TrigonometryServerMethods is that the first context
// argument for each method is always ipc.ServerCall here, while it is either
// ipc.ServerContext or a typed streaming context there.
type TrigonometryServerStubMethods interface {
	Sine(call __ipc.ServerCall, angle float64) (float64, error)
	Cosine(call __ipc.ServerCall, angle float64) (float64, error)
}

// TrigonometryServerStub adds universal methods to TrigonometryServerStubMethods.
type TrigonometryServerStub interface {
	TrigonometryServerStubMethods
	// GetMethodTags will be replaced with DescribeInterfaces.
	GetMethodTags(call __ipc.ServerCall, method string) ([]interface{}, error)
	// Signature will be replaced with DescribeInterfaces.
	Signature(call __ipc.ServerCall) (__ipc.ServiceSignature, error)
}

// TrigonometryServer returns a server stub for Trigonometry.
// It converts an implementation of TrigonometryServerMethods into
// an object that may be used by ipc.Server.
func TrigonometryServer(impl TrigonometryServerMethods) TrigonometryServerStub {
	stub := implTrigonometryServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implTrigonometryServerStub struct {
	impl TrigonometryServerMethods
	gs   *__ipc.GlobState
}

func (s implTrigonometryServerStub) Sine(call __ipc.ServerCall, i0 float64) (float64, error) {
	return s.impl.Sine(call, i0)
}

func (s implTrigonometryServerStub) Cosine(call __ipc.ServerCall, i0 float64) (float64, error) {
	return s.impl.Cosine(call, i0)
}

func (s implTrigonometryServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implTrigonometryServerStub) GetMethodTags(call __ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(toddw): Replace with new DescribeInterfaces implementation.
	switch method {
	case "Sine":
		return []interface{}{}, nil
	case "Cosine":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (s implTrigonometryServerStub) Signature(call __ipc.ServerCall) (__ipc.ServiceSignature, error) {
	// TODO(toddw) Replace with new DescribeInterfaces implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Cosine"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "angle", Type: 26},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 26},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Sine"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "angle", Type: 26},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 26},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

// AdvancedMathClientMethods is the client interface
// containing AdvancedMath methods.
//
// AdvancedMath is an interface for more advanced math than arith.  It embeds
// interfaces defined both in the same file and in an external package; and in
// turn it is embedded by arith.Calculator (which is in the same package but
// different file) to verify that embedding works in all these scenarios.
type AdvancedMathClientMethods interface {
	// Trigonometry is an interface that specifies a couple trigonometric functions.
	TrigonometryClientMethods
	exp.ExpClientMethods
}

// AdvancedMathClientStub adds universal methods to AdvancedMathClientMethods.
type AdvancedMathClientStub interface {
	AdvancedMathClientMethods
	__ipc.UniversalServiceMethods
}

// AdvancedMathClient returns a client stub for AdvancedMath.
func AdvancedMathClient(name string, opts ...__ipc.BindOpt) AdvancedMathClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implAdvancedMathClientStub{name, client, TrigonometryClient(name, client), exp.ExpClient(name, client)}
}

type implAdvancedMathClientStub struct {
	name   string
	client __ipc.Client

	TrigonometryClientStub
	exp.ExpClientStub
}

func (c implAdvancedMathClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implAdvancedMathClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implAdvancedMathClientStub) GetMethodTags(ctx __context.T, method string, opts ...__ipc.CallOpt) (o0 []interface{}, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// AdvancedMathServerMethods is the interface a server writer
// implements for AdvancedMath.
//
// AdvancedMath is an interface for more advanced math than arith.  It embeds
// interfaces defined both in the same file and in an external package; and in
// turn it is embedded by arith.Calculator (which is in the same package but
// different file) to verify that embedding works in all these scenarios.
type AdvancedMathServerMethods interface {
	// Trigonometry is an interface that specifies a couple trigonometric functions.
	TrigonometryServerMethods
	exp.ExpServerMethods
}

// AdvancedMathServerStubMethods is the server interface containing
// AdvancedMath methods, as expected by ipc.Server.  The difference between
// this interface and AdvancedMathServerMethods is that the first context
// argument for each method is always ipc.ServerCall here, while it is either
// ipc.ServerContext or a typed streaming context there.
type AdvancedMathServerStubMethods interface {
	// Trigonometry is an interface that specifies a couple trigonometric functions.
	TrigonometryServerStubMethods
	exp.ExpServerStubMethods
}

// AdvancedMathServerStub adds universal methods to AdvancedMathServerStubMethods.
type AdvancedMathServerStub interface {
	AdvancedMathServerStubMethods
	// GetMethodTags will be replaced with DescribeInterfaces.
	GetMethodTags(call __ipc.ServerCall, method string) ([]interface{}, error)
	// Signature will be replaced with DescribeInterfaces.
	Signature(call __ipc.ServerCall) (__ipc.ServiceSignature, error)
}

// AdvancedMathServer returns a server stub for AdvancedMath.
// It converts an implementation of AdvancedMathServerMethods into
// an object that may be used by ipc.Server.
func AdvancedMathServer(impl AdvancedMathServerMethods) AdvancedMathServerStub {
	stub := implAdvancedMathServerStub{
		impl: impl,
		TrigonometryServerStub: TrigonometryServer(impl),
		ExpServerStub:          exp.ExpServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implAdvancedMathServerStub struct {
	impl AdvancedMathServerMethods
	gs   *__ipc.GlobState

	TrigonometryServerStub
	exp.ExpServerStub
}

func (s implAdvancedMathServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implAdvancedMathServerStub) GetMethodTags(call __ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(toddw): Replace with new DescribeInterfaces implementation.
	if resp, err := s.TrigonometryServerStub.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	if resp, err := s.ExpServerStub.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	return nil, nil
}

func (s implAdvancedMathServerStub) Signature(call __ipc.ServerCall) (__ipc.ServiceSignature, error) {
	// TODO(toddw) Replace with new DescribeInterfaces implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}

	result.TypeDefs = []__vdlutil.Any{}
	var ss __ipc.ServiceSignature
	var firstAdded int
	ss, _ = s.TrigonometryServerStub.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= __wiretype.TypeIDFirst {
				v.InArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= __wiretype.TypeIDFirst {
				v.OutArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= __wiretype.TypeIDFirst {
			v.InStream += __wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= __wiretype.TypeIDFirst {
			v.OutStream += __wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case __wiretype.SliceType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.ArrayType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.MapType:
			if wt.Key >= __wiretype.TypeIDFirst {
				wt.Key += __wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= __wiretype.TypeIDFirst {
					wt.Fields[i].Type += __wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}
	ss, _ = s.ExpServerStub.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= __wiretype.TypeIDFirst {
				v.InArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= __wiretype.TypeIDFirst {
				v.OutArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= __wiretype.TypeIDFirst {
			v.InStream += __wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= __wiretype.TypeIDFirst {
			v.OutStream += __wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case __wiretype.SliceType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.ArrayType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.MapType:
			if wt.Key >= __wiretype.TypeIDFirst {
				wt.Key += __wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= __wiretype.TypeIDFirst {
					wt.Fields[i].Type += __wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}
