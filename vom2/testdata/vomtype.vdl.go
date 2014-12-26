// This file was auto-generated by the veyron vdl tool.
// Source: vomtype.vdl

package testdata

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__fmt "fmt"
	__vdl "v.io/core/veyron2/vdl"
	__vdlutil "v.io/core/veyron2/vdl/vdlutil"
)

type NBool bool

func (NBool) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NBool"
}) {
}

type NString string

func (NString) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NString"
}) {
}

type NByteSlice []byte

func (NByteSlice) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NByteSlice"
}) {
}

type NByte byte

func (NByte) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NByte"
}) {
}

type NUint16 uint16

func (NUint16) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NUint16"
}) {
}

type NUint32 uint32

func (NUint32) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NUint32"
}) {
}

type NUint64 uint64

func (NUint64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NUint64"
}) {
}

type NInt16 int16

func (NInt16) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NInt16"
}) {
}

type NInt32 int32

func (NInt32) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NInt32"
}) {
}

type NInt64 int64

func (NInt64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NInt64"
}) {
}

type NFloat32 float32

func (NFloat32) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NFloat32"
}) {
}

type NFloat64 float64

func (NFloat64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NFloat64"
}) {
}

type NComplex64 complex64

func (NComplex64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NComplex64"
}) {
}

type NComplex128 complex128

func (NComplex128) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NComplex128"
}) {
}

type NArray2Uint64 [2]uint64

func (NArray2Uint64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NArray2Uint64"
}) {
}

type NListUint64 []uint64

func (NListUint64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NListUint64"
}) {
}

type NSetUint64 map[uint64]struct{}

func (NSetUint64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NSetUint64"
}) {
}

type NMapUint64String map[uint64]string

func (NMapUint64String) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NMapUint64String"
}) {
}

type NStruct struct {
	A bool
	B string
	C int64
}

func (NStruct) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NStruct"
}) {
}

type NEnum int

const (
	NEnumA NEnum = iota
	NEnumB
	NEnumC
)

// NEnumAll holds all labels for NEnum.
var NEnumAll = []NEnum{NEnumA, NEnumB, NEnumC}

// NEnumFromString creates a NEnum from a string label.
func NEnumFromString(label string) (x NEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *NEnum) Set(label string) error {
	switch label {
	case "A", "a":
		*x = NEnumA
		return nil
	case "B", "b":
		*x = NEnumB
		return nil
	case "C", "c":
		*x = NEnumC
		return nil
	}
	*x = -1
	return __fmt.Errorf("unknown label %q in testdata.NEnum", label)
}

// String returns the string label of x.
func (x NEnum) String() string {
	switch x {
	case NEnumA:
		return "A"
	case NEnumB:
		return "B"
	case NEnumC:
		return "C"
	}
	return ""
}

func (NEnum) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.NEnum"
	Enum struct{ A, B, C string }
}) {
}

type (
	// NUnion represents any single field of the NUnion union type.
	NUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the NUnion union type.
		__VDLReflect(__NUnionReflect)
	}
	// NUnionA represents field A of the NUnion union type.
	NUnionA struct{ Value bool }
	// NUnionB represents field B of the NUnion union type.
	NUnionB struct{ Value string }
	// NUnionC represents field C of the NUnion union type.
	NUnionC struct{ Value int64 }
	// __NUnionReflect describes the NUnion union type.
	__NUnionReflect struct {
		Name  string "v.io/core/veyron2/vom2/testdata.NUnion"
		Type  NUnion
		Union struct {
			A NUnionA
			B NUnionB
			C NUnionC
		}
	}
)

func (x NUnionA) Index() int                   { return 0 }
func (x NUnionA) Interface() interface{}       { return x.Value }
func (x NUnionA) Name() string                 { return "A" }
func (x NUnionA) __VDLReflect(__NUnionReflect) {}

func (x NUnionB) Index() int                   { return 1 }
func (x NUnionB) Interface() interface{}       { return x.Value }
func (x NUnionB) Name() string                 { return "B" }
func (x NUnionB) __VDLReflect(__NUnionReflect) {}

func (x NUnionC) Index() int                   { return 2 }
func (x NUnionC) Interface() interface{}       { return x.Value }
func (x NUnionC) Name() string                 { return "C" }
func (x NUnionC) __VDLReflect(__NUnionReflect) {}

// Nested Custom Types
type MBool NBool

func (MBool) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.MBool"
}) {
}

type MStruct struct {
	A bool
	B NBool
	C MBool
	D *NStruct
	E *__vdl.Type
	F __vdlutil.Any
}

func (MStruct) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.MStruct"
}) {
}

type MList []NListUint64

func (MList) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.MList"
}) {
}

type MMap map[NFloat32]NListUint64

func (MMap) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vom2/testdata.MMap"
}) {
}

func init() {
	__vdl.Register(NBool(false))
	__vdl.Register(NString(""))
	__vdl.Register(NByteSlice(""))
	__vdl.Register(NByte(0))
	__vdl.Register(NUint16(0))
	__vdl.Register(NUint32(0))
	__vdl.Register(NUint64(0))
	__vdl.Register(NInt16(0))
	__vdl.Register(NInt32(0))
	__vdl.Register(NInt64(0))
	__vdl.Register(NFloat32(0))
	__vdl.Register(NFloat64(0))
	__vdl.Register(NComplex64(0))
	__vdl.Register(NComplex128(0))
	__vdl.Register(NArray2Uint64{})
	__vdl.Register(NListUint64(nil))
	__vdl.Register(NSetUint64(nil))
	__vdl.Register(NMapUint64String(nil))
	__vdl.Register(NStruct{})
	__vdl.Register(NEnumA)
	__vdl.Register(NUnion(NUnionA{false}))
	__vdl.Register(MBool(false))
	__vdl.Register(MStruct{
		E: __vdl.TypeOf((*__vdlutil.Any)(nil)),
	})
	__vdl.Register(MList(nil))
	__vdl.Register(MMap(nil))
}
