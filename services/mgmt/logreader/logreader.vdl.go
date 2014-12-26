// This file was auto-generated by the veyron vdl tool.
// Source: logreader.vdl

// Package logreader contains the interface for reading log files remotely.
package logreader

import (
	"v.io/core/veyron2/services/mgmt/logreader/types"

	"v.io/core/veyron2/services/security/access"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "v.io/core/veyron2"
	__context "v.io/core/veyron2/context"
	__ipc "v.io/core/veyron2/ipc"
	__vdlutil "v.io/core/veyron2/vdl/vdlutil"
	__wiretype "v.io/core/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// LogFileClientMethods is the client interface
// containing LogFile methods.
//
// LogFile can be used to access log files remotely.
type LogFileClientMethods interface {
	// Size returns the number of bytes in the receiving object.
	Size(__context.T, ...__ipc.CallOpt) (int64, error)
	// ReadLog receives up to NumEntries log entries starting at the
	// StartPos offset (in bytes) in the receiving object. Each stream chunk
	// contains one log entry.
	//
	// If Follow is true, ReadLog will block and wait for more entries to
	// arrive when it reaches the end of the file.
	//
	// ReadLog returns the position where it stopped reading, i.e. the
	// position where the next entry starts. This value can be used as
	// StartPos for successive calls to ReadLog.
	//
	// The returned error will be EOF if and only if ReadLog reached the
	// end of the file and no log entries were returned.
	ReadLog(ctx __context.T, StartPos int64, NumEntries int32, Follow bool, opts ...__ipc.CallOpt) (LogFileReadLogCall, error)
}

// LogFileClientStub adds universal methods to LogFileClientMethods.
type LogFileClientStub interface {
	LogFileClientMethods
	__ipc.UniversalServiceMethods
}

// LogFileClient returns a client stub for LogFile.
func LogFileClient(name string, opts ...__ipc.BindOpt) LogFileClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implLogFileClientStub{name, client}
}

type implLogFileClientStub struct {
	name   string
	client __ipc.Client
}

func (c implLogFileClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implLogFileClientStub) Size(ctx __context.T, opts ...__ipc.CallOpt) (o0 int64, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Size", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implLogFileClientStub) ReadLog(ctx __context.T, i0 int64, i1 int32, i2 bool, opts ...__ipc.CallOpt) (ocall LogFileReadLogCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "ReadLog", []interface{}{i0, i1, i2}, opts...); err != nil {
		return
	}
	ocall = &implLogFileReadLogCall{Call: call}
	return
}

func (c implLogFileClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// LogFileReadLogClientStream is the client stream for LogFile.ReadLog.
type LogFileReadLogClientStream interface {
	// RecvStream returns the receiver side of the LogFile.ReadLog client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() types.LogEntry
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// LogFileReadLogCall represents the call returned from LogFile.ReadLog.
type LogFileReadLogCall interface {
	LogFileReadLogClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (int64, error)
}

type implLogFileReadLogCall struct {
	__ipc.Call
	valRecv types.LogEntry
	errRecv error
}

func (c *implLogFileReadLogCall) RecvStream() interface {
	Advance() bool
	Value() types.LogEntry
	Err() error
} {
	return implLogFileReadLogCallRecv{c}
}

type implLogFileReadLogCallRecv struct {
	c *implLogFileReadLogCall
}

func (c implLogFileReadLogCallRecv) Advance() bool {
	c.c.valRecv = types.LogEntry{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implLogFileReadLogCallRecv) Value() types.LogEntry {
	return c.c.valRecv
}
func (c implLogFileReadLogCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implLogFileReadLogCall) Finish() (o0 int64, err error) {
	if ierr := c.Call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// LogFileServerMethods is the interface a server writer
// implements for LogFile.
//
// LogFile can be used to access log files remotely.
type LogFileServerMethods interface {
	// Size returns the number of bytes in the receiving object.
	Size(__ipc.ServerContext) (int64, error)
	// ReadLog receives up to NumEntries log entries starting at the
	// StartPos offset (in bytes) in the receiving object. Each stream chunk
	// contains one log entry.
	//
	// If Follow is true, ReadLog will block and wait for more entries to
	// arrive when it reaches the end of the file.
	//
	// ReadLog returns the position where it stopped reading, i.e. the
	// position where the next entry starts. This value can be used as
	// StartPos for successive calls to ReadLog.
	//
	// The returned error will be EOF if and only if ReadLog reached the
	// end of the file and no log entries were returned.
	ReadLog(ctx LogFileReadLogContext, StartPos int64, NumEntries int32, Follow bool) (int64, error)
}

// LogFileServerStubMethods is the server interface containing
// LogFile methods, as expected by ipc.Server.
// The only difference between this interface and LogFileServerMethods
// is the streaming methods.
type LogFileServerStubMethods interface {
	// Size returns the number of bytes in the receiving object.
	Size(__ipc.ServerContext) (int64, error)
	// ReadLog receives up to NumEntries log entries starting at the
	// StartPos offset (in bytes) in the receiving object. Each stream chunk
	// contains one log entry.
	//
	// If Follow is true, ReadLog will block and wait for more entries to
	// arrive when it reaches the end of the file.
	//
	// ReadLog returns the position where it stopped reading, i.e. the
	// position where the next entry starts. This value can be used as
	// StartPos for successive calls to ReadLog.
	//
	// The returned error will be EOF if and only if ReadLog reached the
	// end of the file and no log entries were returned.
	ReadLog(ctx *LogFileReadLogContextStub, StartPos int64, NumEntries int32, Follow bool) (int64, error)
}

// LogFileServerStub adds universal methods to LogFileServerStubMethods.
type LogFileServerStub interface {
	LogFileServerStubMethods
	// Describe the LogFile interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// LogFileServer returns a server stub for LogFile.
// It converts an implementation of LogFileServerMethods into
// an object that may be used by ipc.Server.
func LogFileServer(impl LogFileServerMethods) LogFileServerStub {
	stub := implLogFileServerStub{
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

type implLogFileServerStub struct {
	impl LogFileServerMethods
	gs   *__ipc.GlobState
}

func (s implLogFileServerStub) Size(ctx __ipc.ServerContext) (int64, error) {
	return s.impl.Size(ctx)
}

func (s implLogFileServerStub) ReadLog(ctx *LogFileReadLogContextStub, i0 int64, i1 int32, i2 bool) (int64, error) {
	return s.impl.ReadLog(ctx, i0, i1, i2)
}

func (s implLogFileServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implLogFileServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{LogFileDesc}
}

// LogFileDesc describes the LogFile interface.
var LogFileDesc __ipc.InterfaceDesc = descLogFile

// descLogFile hides the desc to keep godoc clean.
var descLogFile = __ipc.InterfaceDesc{
	Name:    "LogFile",
	PkgPath: "v.io/core/veyron2/services/mgmt/logreader",
	Doc:     "// LogFile can be used to access log files remotely.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Size",
			Doc:  "// Size returns the number of bytes in the receiving object.",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // int64
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{access.Tag("Debug")},
		},
		{
			Name: "ReadLog",
			Doc:  "// ReadLog receives up to NumEntries log entries starting at the\n// StartPos offset (in bytes) in the receiving object. Each stream chunk\n// contains one log entry.\n//\n// If Follow is true, ReadLog will block and wait for more entries to\n// arrive when it reaches the end of the file.\n//\n// ReadLog returns the position where it stopped reading, i.e. the\n// position where the next entry starts. This value can be used as\n// StartPos for successive calls to ReadLog.\n//\n// The returned error will be EOF if and only if ReadLog reached the\n// end of the file and no log entries were returned.",
			InArgs: []__ipc.ArgDesc{
				{"StartPos", ``},   // int64
				{"NumEntries", ``}, // int32
				{"Follow", ``},     // bool
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // int64
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{access.Tag("Debug")},
		},
	},
}

func (s implLogFileServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["ReadLog"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "StartPos", Type: 37},
			{Name: "NumEntries", Type: 36},
			{Name: "Follow", Type: 2},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 37},
			{Name: "", Type: 65},
		},

		OutStream: 66,
	}
	result.Methods["Size"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 37},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, __wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x25, Name: "Position"},
				__wiretype.FieldType{Type: 0x3, Name: "Line"},
			},
			"v.io/core/veyron2/services/mgmt/logreader/types.LogEntry", []string(nil)},
	}

	return result, nil
}

// LogFileReadLogServerStream is the server stream for LogFile.ReadLog.
type LogFileReadLogServerStream interface {
	// SendStream returns the send side of the LogFile.ReadLog server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item types.LogEntry) error
	}
}

// LogFileReadLogContext represents the context passed to LogFile.ReadLog.
type LogFileReadLogContext interface {
	__ipc.ServerContext
	LogFileReadLogServerStream
}

// LogFileReadLogContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements LogFileReadLogContext.
type LogFileReadLogContextStub struct {
	__ipc.ServerCall
}

// Init initializes LogFileReadLogContextStub from ipc.ServerCall.
func (s *LogFileReadLogContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the LogFile.ReadLog server stream.
func (s *LogFileReadLogContextStub) SendStream() interface {
	Send(item types.LogEntry) error
} {
	return implLogFileReadLogContextSend{s}
}

type implLogFileReadLogContextSend struct {
	s *LogFileReadLogContextStub
}

func (s implLogFileReadLogContextSend) Send(item types.LogEntry) error {
	return s.s.Send(item)
}
