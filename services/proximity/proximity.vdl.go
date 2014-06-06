// This file was auto-generated by the veyron vdl tool.
// Source: proximity.vdl

package proximity

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

// Device represents one neighborhood device.  It contains that device's
// MAC address, observed names, and the average distance to the device.
type Device struct {
	// MAC is remote device's MAC address, in one of the following formats
	// (as per http://golang.org/pkg/net/#ParseMAC):
	//   01:23:45:67:89:ab
	//   01:23:45:67:89:ab:cd:ef
	//   01-23-45-67-89-ab
	//   01-23-45-67-89-ab-cd-ef
	//   0123.4567.89ab
	//   0123.4567.89ab.cdef
	MAC string
	// Names represents all unique observed names of the remote device.
	Names []string
	// Distance represents the (estimated) distance to the neighborhood
	// device.  It can be parsed using distance.Parse method.
	Distance string
}

// Proximity maintains a list of devices in our close proximity, using scan
// readings from nearby devices.  It also continuously advertises a set of
// provided names, which will be visible at nearby devices and associated
// with this device.
// Proximity is the interface the client binds and uses.
// Proximity_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Proximity_ExcludingUniversal interface {
	// RegisterName adds a name that this device will be associated with;
	// a remote device will see all the unique names currently registered
	// with this device (see Names field in Device).
	RegisterName(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (err error)
	// UnregisterName removes a name previously associated with this device.
	// If the name doesn't exist, this method will return an error.
	// If the name has been registered multiple times, this method will
	// remove only one instance of that registration.
	UnregisterName(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (err error)
	// NearbyDevices returns the most up-to-date list of nearby devices,
	// sorted in increasing distance order.
	NearbyDevices(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []Device, err error)
}
type Proximity interface {
	_gen_ipc.UniversalServiceMethods
	Proximity_ExcludingUniversal
}

// ProximityService is the interface the server implements.
type ProximityService interface {

	// RegisterName adds a name that this device will be associated with;
	// a remote device will see all the unique names currently registered
	// with this device (see Names field in Device).
	RegisterName(context _gen_ipc.ServerContext, Name string) (err error)
	// UnregisterName removes a name previously associated with this device.
	// If the name doesn't exist, this method will return an error.
	// If the name has been registered multiple times, this method will
	// remove only one instance of that registration.
	UnregisterName(context _gen_ipc.ServerContext, Name string) (err error)
	// NearbyDevices returns the most up-to-date list of nearby devices,
	// sorted in increasing distance order.
	NearbyDevices(context _gen_ipc.ServerContext) (reply []Device, err error)
}

// BindProximity returns the client stub implementing the Proximity
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindProximity(name string, opts ..._gen_ipc.BindOpt) (Proximity, error) {
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
	stub := &clientStubProximity{client: client, name: name}

	return stub, nil
}

// NewServerProximity creates a new server stub.
//
// It takes a regular server implementing the ProximityService
// interface, and returns a new server stub.
func NewServerProximity(server ProximityService) interface{} {
	return &ServerStubProximity{
		service: server,
	}
}

// clientStubProximity implements Proximity.
type clientStubProximity struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubProximity) RegisterName(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "RegisterName", []interface{}{Name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubProximity) UnregisterName(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnregisterName", []interface{}{Name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubProximity) NearbyDevices(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []Device, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "NearbyDevices", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubProximity) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubProximity) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubProximity) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubProximity wraps a server that implements
// ProximityService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubProximity struct {
	service ProximityService
}

func (__gen_s *ServerStubProximity) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "RegisterName":
		return []interface{}{security.Label(2)}, nil
	case "UnregisterName":
		return []interface{}{security.Label(2)}, nil
	case "NearbyDevices":
		return []interface{}{security.Label(1)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubProximity) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["NearbyDevices"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 67},
			{Name: "", Type: 65},
		},
	}
	result.Methods["RegisterName"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Name", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["UnregisterName"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Name", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "MAC"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Names"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Distance"},
			},
			"veyron2/services/proximity.Device", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubProximity) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubProximity) RegisterName(call _gen_ipc.ServerCall, Name string) (err error) {
	err = __gen_s.service.RegisterName(call, Name)
	return
}

func (__gen_s *ServerStubProximity) UnregisterName(call _gen_ipc.ServerCall, Name string) (err error) {
	err = __gen_s.service.UnregisterName(call, Name)
	return
}

func (__gen_s *ServerStubProximity) NearbyDevices(call _gen_ipc.ServerCall) (reply []Device, err error) {
	reply, err = __gen_s.service.NearbyDevices(call)
	return
}
