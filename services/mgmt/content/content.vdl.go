// This file was auto-generated by the veyron vdl tool.
// Source: content.vdl

// Package content supports storing and serving arbitrary binary
// content, such as veyron application binaries.
//
// OVERVIEW: Content is expected to be organized using veyron's
// hierarchical namespace. The nodes of the hierarchy are expected to
// implement: 1) the MountTable interface, to enable extending the
// hierarchy, 2) the Repository interface, to enable content retrieval,
// and 3) the Glob interface, to enable content discovery.
package content

import (
	"veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

// Repository can be used to store and retrieve binaries.
// Repository is the interface the client binds and uses.
// Repository_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Repository_ExcludingUniversal interface {
	// Delete deletes the content.
	Delete(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Download opens a stream that can used for downloading the
	// content.
	Download(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply RepositoryDownloadStream, err error)
	// Upload opens a stream that can be used for uploading the content
	// and returns the name under which this content can be found.
	Upload(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply RepositoryUploadStream, err error)
}
type Repository interface {
	_gen_ipc.UniversalServiceMethods
	Repository_ExcludingUniversal
}

// RepositoryService is the interface the server implements.
type RepositoryService interface {

	// Delete deletes the content.
	Delete(context _gen_ipc.ServerContext) (err error)
	// Download opens a stream that can used for downloading the
	// content.
	Download(context _gen_ipc.ServerContext, stream RepositoryServiceDownloadStream) (err error)
	// Upload opens a stream that can be used for uploading the content
	// and returns the name under which this content can be found.
	Upload(context _gen_ipc.ServerContext, stream RepositoryServiceUploadStream) (reply string, err error)
}

// RepositoryDownloadStream is the interface for streaming responses of the method
// Download in the service interface Repository.
type RepositoryDownloadStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item []byte, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the RepositoryDownloadStream interface that is not exported.
type implRepositoryDownloadStream struct {
	clientCall _gen_ipc.Call
}

func (c *implRepositoryDownloadStream) Recv() (item []byte, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implRepositoryDownloadStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implRepositoryDownloadStream) Cancel() {
	c.clientCall.Cancel()
}

// RepositoryServiceDownloadStream is the interface for streaming responses of the method
// Download in the service interface Repository.
type RepositoryServiceDownloadStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item []byte) error
}

// Implementation of the RepositoryServiceDownloadStream interface that is not exported.
type implRepositoryServiceDownloadStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implRepositoryServiceDownloadStream) Send(item []byte) error {
	return s.serverCall.Send(item)
}

// RepositoryUploadStream is the interface for streaming responses of the method
// Upload in the service interface Repository.
type RepositoryUploadStream interface {

	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item []byte) error

	// CloseSend indicates to the server that no more items will be sent; server
	// Recv calls will receive io.EOF after all sent items.  Subsequent calls to
	// Send on the client will fail.  This is an optional call - it's used by
	// streaming clients that need the server to receive the io.EOF terminator.
	CloseSend() error

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (reply string, err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the RepositoryUploadStream interface that is not exported.
type implRepositoryUploadStream struct {
	clientCall _gen_ipc.Call
}

func (c *implRepositoryUploadStream) Send(item []byte) error {
	return c.clientCall.Send(item)
}

func (c *implRepositoryUploadStream) CloseSend() error {
	return c.clientCall.CloseSend()
}

func (c *implRepositoryUploadStream) Finish() (reply string, err error) {
	if ierr := c.clientCall.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implRepositoryUploadStream) Cancel() {
	c.clientCall.Cancel()
}

// RepositoryServiceUploadStream is the interface for streaming responses of the method
// Upload in the service interface Repository.
type RepositoryServiceUploadStream interface {

	// Recv fills itemptr with the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item []byte, err error)
}

// Implementation of the RepositoryServiceUploadStream interface that is not exported.
type implRepositoryServiceUploadStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implRepositoryServiceUploadStream) Recv() (item []byte, err error) {
	err = s.serverCall.Recv(&item)
	return
}

// BindRepository returns the client stub implementing the Repository
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindRepository(name string, opts ..._gen_ipc.BindOpt) (Repository, error) {
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
			return nil, _gen_vdl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdl.ErrTooManyOptionsToBind
	}
	stub := &clientStubRepository{client: client, name: name}

	return stub, nil
}

// NewServerRepository creates a new server stub.
//
// It takes a regular server implementing the RepositoryService
// interface, and returns a new server stub.
func NewServerRepository(server RepositoryService) interface{} {
	return &ServerStubRepository{
		service: server,
	}
}

// clientStubRepository implements Repository.
type clientStubRepository struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubRepository) Delete(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Delete", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRepository) Download(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply RepositoryDownloadStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Download", nil, opts...); err != nil {
		return
	}
	reply = &implRepositoryDownloadStream{clientCall: call}
	return
}

func (__gen_c *clientStubRepository) Upload(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply RepositoryUploadStream, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Upload", nil, opts...); err != nil {
		return
	}
	reply = &implRepositoryUploadStream{clientCall: call}
	return
}

func (__gen_c *clientStubRepository) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRepository) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubRepository) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubRepository wraps a server that implements
// RepositoryService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubRepository struct {
	service RepositoryService
}

func (__gen_s *ServerStubRepository) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Delete":
		return []interface{}{security.Label(2)}, nil
	case "Download":
		return []interface{}{security.Label(1)}, nil
	case "Upload":
		return []interface{}{security.Label(2)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubRepository) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Delete"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Download"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 67,
	}
	result.Methods["Upload"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 3},
			{Name: "", Type: 65},
		},
		InStream: 67,
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubRepository) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubRepository) Delete(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Delete(call)
	return
}

func (__gen_s *ServerStubRepository) Download(call _gen_ipc.ServerCall) (err error) {
	stream := &implRepositoryServiceDownloadStream{serverCall: call}
	err = __gen_s.service.Download(call, stream)
	return
}

func (__gen_s *ServerStubRepository) Upload(call _gen_ipc.ServerCall) (reply string, err error) {
	stream := &implRepositoryServiceUploadStream{serverCall: call}
	reply, err = __gen_s.service.Upload(call, stream)
	return
}
