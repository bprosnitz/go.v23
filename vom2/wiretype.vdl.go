// This file was auto-generated by the veyron vdl tool.
// Source: wiretype.vdl

package vom2

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__vdl "v.io/core/veyron2/vdl"
)

// WireNamed represents a type definition for named primitives.
type WireNamed struct {
	Name string
	Base TypeID
}

func (WireNamed) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireNamed"
}) {
}

// WireEnum represents an type definition for enum types.
type WireEnum struct {
	Name   string
	Labels []string
}

func (WireEnum) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireEnum"
}) {
}

// WireArray represents an type definition for array types.
type WireArray struct {
	Name string
	Elem TypeID
	Len  uint64
}

func (WireArray) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireArray"
}) {
}

// WireList represents a type definition for list types.
type WireList struct {
	Name string
	Elem TypeID
}

func (WireList) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireList"
}) {
}

// WireSet represents a type definition for set types.
type WireSet struct {
	Name string
	Key  TypeID
}

func (WireSet) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireSet"
}) {
}

// WireMap represents a type definition for map types.
type WireMap struct {
	Name string
	Key  TypeID
	Elem TypeID
}

func (WireMap) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireMap"
}) {
}

// WireField represents a field in a struct or union type.
type WireField struct {
	Name string
	Type TypeID
}

func (WireField) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireField"
}) {
}

// WireStruct represents a type definition for struct types.
type WireStruct struct {
	Name   string
	Fields []WireField
}

func (WireStruct) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireStruct"
}) {
}

// WireUnion represents a type definition for union types.
type WireUnion struct {
	Name   string
	Fields []WireField
}

func (WireUnion) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireUnion"
}) {
}

// WireOptional represents an type definition for optional types.
type WireOptional struct {
	Name string
	Elem TypeID
}

func (WireOptional) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.WireOptional"
}) {
}

// TypeID uniquely identifies a type definition within a vom stream.
type TypeID uint64

func (TypeID) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2.TypeID"
}) {
}

func init() {
	__vdl.Register(WireNamed{})
	__vdl.Register(WireEnum{})
	__vdl.Register(WireArray{})
	__vdl.Register(WireList{})
	__vdl.Register(WireSet{})
	__vdl.Register(WireMap{})
	__vdl.Register(WireField{})
	__vdl.Register(WireStruct{})
	__vdl.Register(WireUnion{})
	__vdl.Register(WireOptional{})
	__vdl.Register(TypeID(0))
}

// Primitives
const WireAnyID = TypeID(1)

const WireTypeID = TypeID(2)

const WireBoolID = TypeID(3)

const WireStringID = TypeID(4)

const WireByteID = TypeID(5)

const WireUint16ID = TypeID(6)

const WireUint32ID = TypeID(7)

const WireUint64ID = TypeID(8)

const WireInt16ID = TypeID(9)

const WireInt32ID = TypeID(10)

const WireInt64ID = TypeID(11)

const WireFloat32ID = TypeID(12)

const WireFloat64ID = TypeID(13)

const WireComplex64ID = TypeID(14)

const WireComplex128ID = TypeID(15)

// Composites only used in type definitions
const WireNamedID = TypeID(16)

const WireEnumID = TypeID(17)

const WireArrayID = TypeID(18)

const WireListID = TypeID(19)

const WireSetID = TypeID(20)

const WireMapID = TypeID(21)

const WireStructID = TypeID(22)

const WireFieldID = TypeID(23)

const WireFieldListID = TypeID(24)

const WireUnionID = TypeID(25)

const WireOptionalID = TypeID(29)

// Other commonly used composites
const WireByteListID = TypeID(26)

const WireStringListID = TypeID(27)

const WireTypeListID = TypeID(28)

// The first user-defined TypeID is 65.  Note that -64 is encoded as 1 byte,
// while -65 is encoded as 2 bytes.
const WireTypeFirstUserID = TypeID(65)
