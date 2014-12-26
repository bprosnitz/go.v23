// This file was auto-generated by the veyron vdl tool.
// Source: signature.vdl

package ipc

import (
	"v.io/core/veyron2/wiretype"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__vdl "v.io/core/veyron2/vdl"
	__vdlutil "v.io/core/veyron2/vdl/vdlutil"
)

// ServiceSignature represents the signature of the service. This includes type information needed
// to resolve the method argument types.
// TODO(bprosnitz) Rename this and move it to wiretype.
type ServiceSignature struct {
	TypeDefs []__vdlutil.Any // A slice of wiretype structures form the type definition.
	Methods  map[string]MethodSignature
}

func (ServiceSignature) __VDLReflect(struct {
	Name string "v.io/core/veyron2/ipc.ServiceSignature"
}) {
}

// MethodSignature represents the structure for passing around method
// signatures. This is usually sent in a ServiceSignature.
type MethodSignature struct {
	InArgs    []MethodArgument // Positional Argument information.
	OutArgs   []MethodArgument
	InStream  wiretype.TypeID // Type of streaming arguments (or TypeIDInvalid if none). The type IDs here use the definitions in ServiceSigature.TypeDefs.
	OutStream wiretype.TypeID
}

func (MethodSignature) __VDLReflect(struct {
	Name string "v.io/core/veyron2/ipc.MethodSignature"
}) {
}

// MethodArgument represents the argument to a method in a method signature.
type MethodArgument struct {
	Name string // Argument name
	Type wiretype.TypeID
}

func (MethodArgument) __VDLReflect(struct {
	Name string "v.io/core/veyron2/ipc.MethodArgument"
}) {
}

func init() {
	__vdl.Register(ServiceSignature{})
	__vdl.Register(MethodSignature{})
	__vdl.Register(MethodArgument{})
}
