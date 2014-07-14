// This file was auto-generated by the veyron vdl tool.
// Source: base.vdl

// Package test_base is a simple single-file test of vdl functionality.
package test_base

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_verror "veyron2/verror"
	_gen_wiretype "veyron2/wiretype"
)

type NamedBool bool

type NamedByte byte

type NamedUint16 uint16

type NamedUint32 uint32

type NamedUint64 uint64

type NamedInt16 int16

type NamedInt32 int32

type NamedInt64 int64

type NamedFloat32 float32

type NamedFloat64 float64

type NamedComplex64 complex64

type NamedComplex128 complex128

type NamedString string

//NamedEnum       enum{A;B;C}
type NamedArray [2]bool

type NamedList []uint32

//NamedSet        set[string]
type NamedMap map[string]float32

type NamedStruct struct {
	A bool
	B string
	C int32
}

type Scalars struct {
	A0  bool
	A1  byte
	A2  uint16
	A3  uint32
	A4  uint64
	A5  int16
	A6  int32
	A7  int64
	A8  float32
	A9  float64
	A10 complex64
	A11 complex128
	A12 string
	A13 error
	A14 _gen_vdlutil.Any
	A15 *_gen_vdl.Type
	B0  NamedBool
	B1  NamedByte
	B2  NamedUint16
	B3  NamedUint32
	B4  NamedUint64
	B5  NamedInt16
	B6  NamedInt32
	B7  NamedInt64
	B8  NamedFloat32
	B9  NamedFloat64
	B10 NamedComplex64
	B11 NamedComplex128
	B12 NamedString
}

type Composites struct {
	A0 Scalars
	A1 [2]Scalars
	A2 []Scalars
	//A3 set[Scalars]
	A4 map[string]Scalars
	A5 map[Scalars][]map[string]complex128
}

type CompComp struct {
	A0 Composites
	A1 [2]Composites
	A2 []Composites
	A3 map[string]Composites
	A4 map[Scalars][]map[string]Composites
}

// NestedArgs is defined before Args; that's allowed in regular Go, and also
// allowed in our vdl files.  The compiler will re-order dependent types to ease
// code generation in other languages.
type NestedArgs struct {
	Args Args
}

// Args will be reordered to show up before NestedArgs in the generated output.
type Args struct {
	A int32
	B int32
}

const (
	Cbool = true

	Cbyte = byte(1)

	Cint32 = int32(2)

	Cint64 = int64(3)

	Cuint32 = uint32(4)

	Cuint64 = uint64(5)

	Cfloat32 = float32(6)

	Cfloat64 = float64(7)

	Ccomplex64 = complex64(8 + 9i)

	Ccomplex128 = complex128(10 + 11i)

	Cstring = "foo"

	Cany = true

	True = true

	Foo = "foo"

	Five = int32(5)

	Six = uint64(6)

	SixSquared = uint64(36)

	FiveSquared = int32(25)
)

const ErrIDFoo = _gen_verror.ID("veyron2/vdl/test_base.ErrIDFoo")

const ErrIDBar = _gen_verror.ID("some/path.ErrIdOther")

// ServiceA is the interface the client binds and uses.
// ServiceA_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type ServiceA_ExcludingUniversal interface {
	MethodA1(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	MethodA2(ctx _gen_context.T, a int32, b string, opts ..._gen_ipc.CallOpt) (reply string, err error)
	MethodA3(ctx _gen_context.T, a int32, opts ..._gen_ipc.CallOpt) (reply ServiceAMethodA3Stream, err error)
	MethodA4(ctx _gen_context.T, a int32, opts ..._gen_ipc.CallOpt) (reply ServiceAMethodA4Stream, err error)
}
type ServiceA interface {
	_gen_ipc.UniversalServiceMethods
	ServiceA_ExcludingUniversal
}

// ServiceAService is the interface the server implements.
type ServiceAService interface {
	MethodA1(context _gen_ipc.ServerContext) (err error)
	MethodA2(context _gen_ipc.ServerContext, a int32, b string) (reply string, err error)
	MethodA3(context _gen_ipc.ServerContext, a int32, stream ServiceAServiceMethodA3Stream) (reply string, err error)
	MethodA4(context _gen_ipc.ServerContext, a int32, stream ServiceAServiceMethodA4Stream) (err error)
}

// ServiceAMethodA3Stream is the interface for streaming responses of the method
// MethodA3 in the service interface ServiceA.
type ServiceAMethodA3Stream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item Scalars, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (reply string, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the ServiceAMethodA3Stream interface that is not exported.
type implServiceAMethodA3Stream struct {
	clientCall _gen_ipc.Call
}

func (c *implServiceAMethodA3Stream) Recv() (item Scalars, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implServiceAMethodA3Stream) Finish() (reply string, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implServiceAMethodA3Stream) Cancel() {
	c.clientCall.Cancel()
}

// ServiceAServiceMethodA3Stream is the interface for streaming responses of the method
// MethodA3 in the service interface ServiceA.
type ServiceAServiceMethodA3Stream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item Scalars) error
}

// Implementation of the ServiceAServiceMethodA3Stream interface that is not exported.
type implServiceAServiceMethodA3Stream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implServiceAServiceMethodA3Stream) Send(item Scalars) error {
	return s.serverCall.Send(item)
}

// ServiceAMethodA4Stream is the interface for streaming responses of the method
// MethodA4 in the service interface ServiceA.
type ServiceAMethodA4Stream interface {

	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item int32) error

	// CloseSend indicates to the server that no more items will be sent; server
	// Recv calls will receive io.EOF after all sent items.  Subsequent calls to
	// Send on the client will fail.  This is an optional call - it's used by
	// streaming clients that need the server to receive the io.EOF terminator.
	CloseSend() error

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item string, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the ServiceAMethodA4Stream interface that is not exported.
type implServiceAMethodA4Stream struct {
	clientCall _gen_ipc.Call
}

func (c *implServiceAMethodA4Stream) Send(item int32) error {
	return c.clientCall.Send(item)
}

func (c *implServiceAMethodA4Stream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implServiceAMethodA4Stream) Recv() (item string, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implServiceAMethodA4Stream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implServiceAMethodA4Stream) Cancel() {
	c.clientCall.Cancel()
}

// ServiceAServiceMethodA4Stream is the interface for streaming responses of the method
// MethodA4 in the service interface ServiceA.
type ServiceAServiceMethodA4Stream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item string) error

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item int32, err error)
}

// Implementation of the ServiceAServiceMethodA4Stream interface that is not exported.
type implServiceAServiceMethodA4Stream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implServiceAServiceMethodA4Stream) Send(item string) error {
	return s.serverCall.Send(item)
}

func (s *implServiceAServiceMethodA4Stream) Recv() (item int32, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// BindServiceA returns the client stub implementing the ServiceA
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindServiceA(name string, opts ..._gen_ipc.BindOpt) (ServiceA, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubServiceA{client: client, name: name}

	return stub, nil
}

// NewServerServiceA creates a new server stub.
//
// It takes a regular server implementing the ServiceAService
// interface, and returns a new server stub.
func NewServerServiceA(server ServiceAService) interface{} {
	return &ServerStubServiceA{
		service: server,
	}
}

// clientStubServiceA implements ServiceA.
type clientStubServiceA struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubServiceA) MethodA1(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "MethodA1", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceA) MethodA2(ctx _gen_context.T, a int32, b string, opts ..._gen_ipc.CallOpt) (reply string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "MethodA2", []interface{}{a, b}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceA) MethodA3(ctx _gen_context.T, a int32, opts ..._gen_ipc.CallOpt) (reply ServiceAMethodA3Stream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "MethodA3", []interface{}{a}, opts...); err != nil {
		return
	}
	reply = &implServiceAMethodA3Stream{clientCall: call}
	return
}

func (__gen_c *clientStubServiceA) MethodA4(ctx _gen_context.T, a int32, opts ..._gen_ipc.CallOpt) (reply ServiceAMethodA4Stream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "MethodA4", []interface{}{a}, opts...); err != nil {
		return
	}
	reply = &implServiceAMethodA4Stream{clientCall: call}
	return
}

func (__gen_c *clientStubServiceA) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceA) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceA) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubServiceA wraps a server that implements
// ServiceAService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubServiceA struct {
	service ServiceAService
}

func (__gen_s *ServerStubServiceA) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "MethodA1":
		return []interface{}{}, nil
	case "MethodA2":
		return []interface{}{}, nil
	case "MethodA3":
		return []interface{}{"tag", uint64(6)}, nil
	case "MethodA4":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubServiceA) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["MethodA1"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["MethodA2"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "a", Type: 36},
			{Name: "b", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "s", Type: 3},
			{Name: "err", Type: 65},
		},
	}
	result.Methods["MethodA3"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "a", Type: 36},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "s", Type: 3},
			{Name: "err", Type: 65},
		},

		OutStream: 82,
	}
	result.Methods["MethodA4"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "a", Type: 36},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
		InStream:  36,
		OutStream: 3,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x7, Name: "TypeID", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x2, Name: "veyron2/vdl/test_base.NamedBool", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron2/vdl/test_base.NamedByte", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x33, Name: "veyron2/vdl/test_base.NamedUint16", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x34, Name: "veyron2/vdl/test_base.NamedUint32", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x35, Name: "veyron2/vdl/test_base.NamedUint64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x23, Name: "veyron2/vdl/test_base.NamedInt16", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x24, Name: "veyron2/vdl/test_base.NamedInt32", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x25, Name: "veyron2/vdl/test_base.NamedInt64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x19, Name: "veyron2/vdl/test_base.NamedFloat32", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1a, Name: "veyron2/vdl/test_base.NamedFloat64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x38, Name: "veyron2/vdl/test_base.NamedComplex64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x39, Name: "veyron2/vdl/test_base.NamedComplex128", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron2/vdl/test_base.NamedString", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x2, Name: "A0"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "A1"},
				_gen_wiretype.FieldType{Type: 0x33, Name: "A2"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "A3"},
				_gen_wiretype.FieldType{Type: 0x35, Name: "A4"},
				_gen_wiretype.FieldType{Type: 0x23, Name: "A5"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "A6"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "A7"},
				_gen_wiretype.FieldType{Type: 0x19, Name: "A8"},
				_gen_wiretype.FieldType{Type: 0x1a, Name: "A9"},
				_gen_wiretype.FieldType{Type: 0x38, Name: "A10"},
				_gen_wiretype.FieldType{Type: 0x39, Name: "A11"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "A12"},
				_gen_wiretype.FieldType{Type: 0x41, Name: "A13"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "A14"},
				_gen_wiretype.FieldType{Type: 0x44, Name: "A15"},
				_gen_wiretype.FieldType{Type: 0x45, Name: "B0"},
				_gen_wiretype.FieldType{Type: 0x46, Name: "B1"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "B2"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "B3"},
				_gen_wiretype.FieldType{Type: 0x49, Name: "B4"},
				_gen_wiretype.FieldType{Type: 0x4a, Name: "B5"},
				_gen_wiretype.FieldType{Type: 0x4b, Name: "B6"},
				_gen_wiretype.FieldType{Type: 0x4c, Name: "B7"},
				_gen_wiretype.FieldType{Type: 0x4d, Name: "B8"},
				_gen_wiretype.FieldType{Type: 0x4e, Name: "B9"},
				_gen_wiretype.FieldType{Type: 0x4f, Name: "B10"},
				_gen_wiretype.FieldType{Type: 0x50, Name: "B11"},
				_gen_wiretype.FieldType{Type: 0x51, Name: "B12"},
			},
			"veyron2/vdl/test_base.Scalars", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubServiceA) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubServiceA) MethodA1(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.MethodA1(call)
	return
}

func (__gen_s *ServerStubServiceA) MethodA2(call _gen_ipc.ServerCall, a int32, b string) (reply string, err error) {
	reply, err = __gen_s.service.MethodA2(call, a, b)
	return
}

func (__gen_s *ServerStubServiceA) MethodA3(call _gen_ipc.ServerCall, a int32) (reply string, err error) {
	stream := &implServiceAServiceMethodA3Stream{serverCall: call}
	reply, err = __gen_s.service.MethodA3(call, a, stream)
	return
}

func (__gen_s *ServerStubServiceA) MethodA4(call _gen_ipc.ServerCall, a int32) (err error) {
	stream := &implServiceAServiceMethodA4Stream{serverCall: call}
	err = __gen_s.service.MethodA4(call, a, stream)
	return
}

// ServiceB is the interface the client binds and uses.
// ServiceB_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type ServiceB_ExcludingUniversal interface {
	ServiceA_ExcludingUniversal
	MethodB1(ctx _gen_context.T, a Scalars, b Composites, opts ..._gen_ipc.CallOpt) (reply CompComp, err error)
}
type ServiceB interface {
	_gen_ipc.UniversalServiceMethods
	ServiceB_ExcludingUniversal
}

// ServiceBService is the interface the server implements.
type ServiceBService interface {
	ServiceAService
	MethodB1(context _gen_ipc.ServerContext, a Scalars, b Composites) (reply CompComp, err error)
}

// BindServiceB returns the client stub implementing the ServiceB
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindServiceB(name string, opts ..._gen_ipc.BindOpt) (ServiceB, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubServiceB{client: client, name: name}
	stub.ServiceA_ExcludingUniversal, _ = BindServiceA(name, client)

	return stub, nil
}

// NewServerServiceB creates a new server stub.
//
// It takes a regular server implementing the ServiceBService
// interface, and returns a new server stub.
func NewServerServiceB(server ServiceBService) interface{} {
	return &ServerStubServiceB{
		ServerStubServiceA: *NewServerServiceA(server).(*ServerStubServiceA),
		service:            server,
	}
}

// clientStubServiceB implements ServiceB.
type clientStubServiceB struct {
	ServiceA_ExcludingUniversal

	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubServiceB) MethodB1(ctx _gen_context.T, a Scalars, b Composites, opts ..._gen_ipc.CallOpt) (reply CompComp, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "MethodB1", []interface{}{a, b}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceB) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceB) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubServiceB) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubServiceB wraps a server that implements
// ServiceBService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubServiceB struct {
	ServerStubServiceA

	service ServiceBService
}

func (__gen_s *ServerStubServiceB) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	if resp, err := __gen_s.ServerStubServiceA.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	switch method {
	case "MethodB1":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubServiceB) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["MethodB1"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "a", Type: 82},
			{Name: "b", Type: 89},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "c", Type: 95},
			{Name: "err", Type: 66},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x7, Name: "TypeID", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x2, Name: "veyron2/vdl/test_base.NamedBool", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "veyron2/vdl/test_base.NamedByte", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x33, Name: "veyron2/vdl/test_base.NamedUint16", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x34, Name: "veyron2/vdl/test_base.NamedUint32", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x35, Name: "veyron2/vdl/test_base.NamedUint64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x23, Name: "veyron2/vdl/test_base.NamedInt16", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x24, Name: "veyron2/vdl/test_base.NamedInt32", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x25, Name: "veyron2/vdl/test_base.NamedInt64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x19, Name: "veyron2/vdl/test_base.NamedFloat32", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1a, Name: "veyron2/vdl/test_base.NamedFloat64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x38, Name: "veyron2/vdl/test_base.NamedComplex64", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x39, Name: "veyron2/vdl/test_base.NamedComplex128", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron2/vdl/test_base.NamedString", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x2, Name: "A0"},
				_gen_wiretype.FieldType{Type: 0x41, Name: "A1"},
				_gen_wiretype.FieldType{Type: 0x33, Name: "A2"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "A3"},
				_gen_wiretype.FieldType{Type: 0x35, Name: "A4"},
				_gen_wiretype.FieldType{Type: 0x23, Name: "A5"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "A6"},
				_gen_wiretype.FieldType{Type: 0x25, Name: "A7"},
				_gen_wiretype.FieldType{Type: 0x19, Name: "A8"},
				_gen_wiretype.FieldType{Type: 0x1a, Name: "A9"},
				_gen_wiretype.FieldType{Type: 0x38, Name: "A10"},
				_gen_wiretype.FieldType{Type: 0x39, Name: "A11"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "A12"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "A13"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "A14"},
				_gen_wiretype.FieldType{Type: 0x44, Name: "A15"},
				_gen_wiretype.FieldType{Type: 0x45, Name: "B0"},
				_gen_wiretype.FieldType{Type: 0x46, Name: "B1"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "B2"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "B3"},
				_gen_wiretype.FieldType{Type: 0x49, Name: "B4"},
				_gen_wiretype.FieldType{Type: 0x4a, Name: "B5"},
				_gen_wiretype.FieldType{Type: 0x4b, Name: "B6"},
				_gen_wiretype.FieldType{Type: 0x4c, Name: "B7"},
				_gen_wiretype.FieldType{Type: 0x4d, Name: "B8"},
				_gen_wiretype.FieldType{Type: 0x4e, Name: "B9"},
				_gen_wiretype.FieldType{Type: 0x4f, Name: "B10"},
				_gen_wiretype.FieldType{Type: 0x50, Name: "B11"},
				_gen_wiretype.FieldType{Type: 0x51, Name: "B12"},
			},
			"veyron2/vdl/test_base.Scalars", []string(nil)},
		_gen_wiretype.ArrayType{Elem: 0x52, Len: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x52, Name: "", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x3, Elem: 0x52, Name: "", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x3, Elem: 0x39, Name: "", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x56, Name: "", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x52, Elem: 0x57, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x52, Name: "A0"},
				_gen_wiretype.FieldType{Type: 0x53, Name: "A1"},
				_gen_wiretype.FieldType{Type: 0x54, Name: "A2"},
				_gen_wiretype.FieldType{Type: 0x55, Name: "A4"},
				_gen_wiretype.FieldType{Type: 0x58, Name: "A5"},
			},
			"veyron2/vdl/test_base.Composites", []string(nil)},
		_gen_wiretype.ArrayType{Elem: 0x59, Len: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x59, Name: "", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x3, Elem: 0x59, Name: "", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x5c, Name: "", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x52, Elem: 0x5d, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x59, Name: "A0"},
				_gen_wiretype.FieldType{Type: 0x5a, Name: "A1"},
				_gen_wiretype.FieldType{Type: 0x5b, Name: "A2"},
				_gen_wiretype.FieldType{Type: 0x5c, Name: "A3"},
				_gen_wiretype.FieldType{Type: 0x5e, Name: "A4"},
			},
			"veyron2/vdl/test_base.CompComp", []string(nil)},
	}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = __gen_s.ServerStubServiceA.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					wt.Fields[i].Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}

func (__gen_s *ServerStubServiceB) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubServiceB) MethodB1(call _gen_ipc.ServerCall, a Scalars, b Composites) (reply CompComp, err error) {
	reply, err = __gen_s.service.MethodB1(call, a, b)
	return
}
