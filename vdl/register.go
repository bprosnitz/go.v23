package vdl

import (
	"fmt"
	"reflect"
	"sync"
)

// Register registers a type, identified by a value for that type.  The type
// should be a type that will be sent over the wire.  Subtypes are recursively
// registered.  This creates a type name <-> reflect.Type bijective mapping.
//
// Type registration is only required for VDL conversion into interface{}
// values, so that values of the correct type may be generated.  Conversion into
// interface{} values for types that are not registered will fill in *vdl.Value
// into the interface{} value.
//
// Panics if wire is not a valid wire type, or if the name<->type mapping is not
// bijective.
//
// Register is not intended to be called by end users; calls are auto-generated
// for all types defined in *.vdl files.
func Register(wire interface{}) {
	if wire == nil {
		return
	}
	if err := registerRecursive(reflect.TypeOf(wire)); err != nil {
		panic(err)
	}
}

func registerRecursive(rt reflect.Type) error {
	// 1) Normalize and derive reflect information.
	rt = normalizeType(rt)
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	ri, err := deriveReflectInfo(rt)
	if err != nil {
		return err
	}
	// 2) Add reflect information to the registry.
	first, err := riReg.addReflectInfo(rt, ri)
	if err != nil {
		return err
	}
	if !first {
		// Break cycles for recursive types.
		return nil
	}
	// 3) Register subtypes, if this is the first time we've seen the type.
	//
	// Special-case to recurse on union fields.
	if len(ri.UnionFields) > 0 {
		for _, field := range ri.UnionFields {
			if err := registerRecursive(field.Type); err != nil {
				return err
			}
		}
		return nil
	}
	// Recurse on subtypes contained in regular composite types.
	switch wt := ri.Type; wt.Kind() {
	case reflect.Array, reflect.Slice, reflect.Ptr:
		return registerRecursive(wt.Elem())
	case reflect.Map:
		if err := registerRecursive(wt.Key()); err != nil {
			return err
		}
		return registerRecursive(wt.Elem())
	case reflect.Struct:
		for ix := 0; ix < wt.NumField(); ix++ {
			if err := registerRecursive(wt.Field(ix).Type); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

// riRegistry holds the reflectInfo registry.  Unlike rtRegistry (used for the
// rtCache), this information cannot be regenerated at will.  We expect a
// limited number of types to be used within a single address space.
type riRegistry struct {
	sync.RWMutex
	fromName map[string]*reflectInfo
	fromType map[reflect.Type]*reflectInfo
}

var riReg = &riRegistry{
	fromName: make(map[string]*reflectInfo),
	fromType: make(map[reflect.Type]*reflectInfo),
}

func (reg *riRegistry) addReflectInfo(rt reflect.Type, ri *reflectInfo) (bool, error) {
	reg.Lock()
	defer reg.Unlock()
	// Allow duplicates only if they're identical, meaning the same type was
	// registered multiple times.  Otherwise it indicates a non-bijective mapping.
	if riDup := reg.fromType[ri.Type]; riDup != nil {
		if !equalRI(ri, riDup) {
			return false, fmt.Errorf("vdl: Register(%v) duplicate type %v: %#v and %#v", rt, ri.Type, ri, riDup)
		}
		// We've seen the wire type before.
		return false, nil
	}
	reg.fromType[ri.Type] = ri
	if ri.Name != "" {
		if riDup := reg.fromName[ri.Name]; riDup != nil && !equalRI(ri, riDup) {
			return false, fmt.Errorf("vdl: Register(%v) duplicate name %q: %#v and %#v", rt, ri.Name, ri, riDup)
		}
		reg.fromName[ri.Name] = ri
	}
	return true, nil
}

// reflectInfoFromName returns the reflectInfo for the given vdl type name, or
// nil if Register has not been called for a type with the given name.
func reflectInfoFromName(name string) *reflectInfo {
	riReg.RLock()
	ri := riReg.fromName[name]
	riReg.RUnlock()
	return ri
}

// reflectInfo holds the reflection information for a type.  All fields are
// populated via reflection over the Type.
//
// The type may include a special __VDLReflect function to describe metadata.
// This is only required for enum and union vdl types, which don't have a
// canonical Go representation.  All other fields are optional.
//
//   type Foo struct{}
//   func (Foo) __VDLReflect(struct{
//     // Type represents the base type.  This is used by union to describe the
//     // union interface type, as opposed to the concrete struct field types.
//     Type Foo
//
//     // Name holds the vdl type name, including the package path, in a tag.
//     Name string "vdl/pkg.Foo"
//
//     // Only one of Enum or Union should be set; they're both shown here for
//     // explanatory purposes.
//
//     // Enum describes the labels for an enum type.
//     Enum struct { A, B string }
//
//     // Union describes the union field names, along with the concrete struct
//     // field types, which contain the actual field types.
//     Union struct {
//       A FieldA
//       B FieldB
//     }
//   }
type reflectInfo struct {
	// Type is the basis for all other information in this struct.
	Type reflect.Type

	// Name is the vdl type name including the vdl package path,
	// e.g. "v.io/v23/vdl.Foo".
	Name string

	// EnumLabels holds the labels of an enum; it is non-empty iff the Type
	// represents a vdl enum.
	EnumLabels []string

	// UnionFields holds the fields of a union; it is non-empty iff the Type
	// represents a vdl union.
	UnionFields []reflectField
}

// reflectField describes the reflection info for a Union field.
type reflectField struct {
	// Given a vdl type Foo union{A bool;B string}, we generate:
	//   type Foo interface{...}
	//   type FooA struct{ Value bool }
	//   type FooB struct{ Value string }
	Name    string       // Field name, e.g. "A", "B"
	Type    reflect.Type // Field type, e.g. bool, string
	RepType reflect.Type // Concrete type representing the field, e.g. FooA, FooB
}

func equalRI(a, b *reflectInfo) bool {
	// Since all information is derived from the Type, that's all we compare.
	return a.Type == b.Type
}

// isUnion returns true iff rt is a union vdl type; it runs a quicker form of
// deriveReflectInfo, only to check for union.  It's used while normalizing types,
// and is necessary since the generated union type is an interface, and must be
// distinguished from the any type.
func isUnion(rt reflect.Type) bool {
	if method, ok := rt.MethodByName("__VDLReflect"); ok {
		mtype := method.Type
		offsetIn := 1
		if rt.Kind() == reflect.Interface {
			offsetIn = 0
		}
		if mtype.NumIn() == 1+offsetIn && mtype.In(offsetIn).Kind() == reflect.Struct {
			rtReflect := mtype.In(offsetIn)
			if _, ok := rtReflect.FieldByName("Union"); ok {
				return true
			}
		}
	}
	return false
}

// deriveReflectInfo returns the reflectInfo corresponding to rt.
// REQUIRES: rt has been normalized, and pointers have been flattened.
func deriveReflectInfo(rt reflect.Type) (*reflectInfo, error) {
	// Set reasonable defaults for types that don't have the __VDLReflect method.
	ri := new(reflectInfo)
	ri.Type = rt
	if rt.PkgPath() != "" {
		ri.Name = rt.PkgPath() + "." + rt.Name()
	}
	// If rt is an non-interface type, methods include the receiver as the first
	// in-arg, otherwise they don't.
	offsetIn := 1
	if rt.Kind() == reflect.Interface {
		offsetIn = 0
	}
	// If rt has a __VDLReflect method, use it to extract metadata.
	if method, ok := rt.MethodByName("__VDLReflect"); ok {
		mtype := method.Type
		if mtype.NumOut() != 0 || mtype.NumIn() != 1+offsetIn || mtype.In(offsetIn).Kind() != reflect.Struct {
			return nil, fmt.Errorf("type %q invalid __VDLReflect (want __VDLReflect(struct{...}))", rt)
		}
		// rtReflect corresponds to the argument to __VDLReflect.
		rtReflect := mtype.In(offsetIn)
		if field, ok := rtReflect.FieldByName("Type"); ok {
			ri.Type = field.Type
			if wt := ri.Type; wt.PkgPath() != "" {
				ri.Name = wt.PkgPath() + "." + wt.Name()
			} else {
				ri.Name = ""
			}
		}
		if field, ok := rtReflect.FieldByName("Name"); ok {
			ri.Name = string(field.Tag)
		}
		if field, ok := rtReflect.FieldByName("Enum"); ok {
			if err := describeEnum(field.Type, rt, ri); err != nil {
				return nil, err
			}
		}
		if field, ok := rtReflect.FieldByName("Union"); ok {
			if err := describeUnion(field.Type, rt, ri); err != nil {
				return nil, err
			}
		}
		if len(ri.EnumLabels) > 0 && len(ri.UnionFields) > 0 {
			return nil, fmt.Errorf("type %q is both an enum and a union", rt)
		}
	}
	return ri, nil
}

// describeEnum fills in ri; we expect enumReflect has this format:
//   struct {A, B, C Foo}
//
// Here's the full type for vdl type Foo enum{A;B}
//   type Foo int
//   const (
//     FooA Foo = iota
//     FooB
//   )
//   func (Foo) __VDLReflect(struct{
//     Type Foo
//     Enum struct { A, B Foo }
//   }) {}
//   func (Foo) String() string {}
//   func (*Foo) Set(string) error {}
func describeEnum(enumReflect, rt reflect.Type, ri *reflectInfo) error {
	if rt != ri.Type || rt.Kind() == reflect.Interface {
		return fmt.Errorf("enum type %q invalid (mismatched type %q)", rt, ri.Type)
	}
	if enumReflect.Kind() != reflect.Struct || enumReflect.NumField() == 0 {
		return fmt.Errorf("enum type %q invalid (no labels)", rt)
	}
	for ix := 0; ix < enumReflect.NumField(); ix++ {
		ri.EnumLabels = append(ri.EnumLabels, enumReflect.Field(ix).Name)
	}
	if s, ok := rt.MethodByName("String"); !ok ||
		s.Type.NumIn() != 1 ||
		s.Type.NumOut() != 1 || s.Type.Out(0) != rtString {
		return fmt.Errorf("enum type %q must have method String() string", rt)
	}
	_, nonptr := rt.MethodByName("Set")
	if a, ok := reflect.PtrTo(rt).MethodByName("Set"); !ok || nonptr ||
		a.Type.NumIn() != 2 || a.Type.In(1) != rtString ||
		a.Type.NumOut() != 1 || a.Type.Out(0) != rtError {
		return fmt.Errorf("enum type %q must have pointer method Set(string) error", rt)
	}
	return nil
}

// describeUnion fills in ri; we expect unionReflect has this format:
//   struct {
//     A FooA
//     B FooB
//   }
//
// Here's the full type for vdl type Foo union{A bool; B string}
//   type (
//     // Foo is the union interface type, that can hold any field.
//     Foo interface {
//       Index() int
//       Name() string
//       __VDLReflect(__FooReflect)
//     }
//     // FooA and FooB are the concrete field types.
//     FooA struct { Value bool }
//     FooB struct { Value string }
//     // __FooReflect lets us re-construct the union type via reflection.
//     __FooReflect struct {
//       Type  Foo // Tells us the union interface type.
//       Union struct {
//         A FooA  // Tells us field 0 has name A and concrete type FooA.
//         B FooB  // Tells us field 1 has name B and concrete type FooB.
//       }
//     }
//   )
func describeUnion(unionReflect, rt reflect.Type, ri *reflectInfo) error {
	if ri.Type.Kind() != reflect.Interface {
		return fmt.Errorf("union type %q has non-interface type %q", rt, ri.Type)
	}
	if unionReflect.Kind() != reflect.Struct || unionReflect.NumField() == 0 {
		return fmt.Errorf("union type %q invalid (no fields)", rt)
	}
	for ix := 0; ix < unionReflect.NumField(); ix++ {
		f := unionReflect.Field(ix)
		if f.PkgPath != "" {
			return fmt.Errorf("union type %q field %q.%q must be exported", rt, f.PkgPath, f.Name)
		}
		// f.Type corresponds to FooA and FooB in __FooReflect above.
		if f.Type.Kind() != reflect.Struct || f.Type.NumField() != 1 || f.Type.Field(0).Name != "Value" {
			return fmt.Errorf("union type %q field %q has bad concrete field type %q", rt, f.Name, f.Type)
		}
		ri.UnionFields = append(ri.UnionFields, reflectField{
			Name:    f.Name,
			Type:    f.Type.Field(0).Type,
			RepType: f.Type,
		})
	}
	// Check for Name method on interface and all concrete field structs.
	if n, ok := ri.Type.MethodByName("Name"); !ok || n.Type.NumIn() != 0 ||
		n.Type.NumOut() != 1 || n.Type.Out(0) != rtString {
		return fmt.Errorf("union interface type %q must have method Name() string", ri.Type)
	}
	for _, f := range ri.UnionFields {
		if n, ok := f.RepType.MethodByName("Name"); !ok || n.Type.NumIn() != 1 ||
			n.Type.NumOut() != 1 || n.Type.Out(0) != rtString {
			return fmt.Errorf("union field %q type %q must have method Name() string", f.Name, f.RepType)
		}
	}
	return nil
}

// TypeToReflect returns the reflect.Type corresponding to t.  We look up
// named types in our registry, and build the unnamed types that we can via the
// Go reflect package.
func TypeToReflect(t *Type) reflect.Type {
	if t.Name() != "" {
		// Named types cannot be manufactured via Go reflect, so we lookup in our
		// registry instead.
		if ri := reflectInfoFromName(t.Name()); ri != nil {
			if ni := nativeInfoFromWire(ri.Type); ni != nil {
				return ni.NativeType
			}
			return ri.Type
		}
		return nil
	}
	// We can make some unnamed types via Go reflect.  Everything else drops
	// through and returns nil.
	switch t.Kind() {
	case Any, Array, Enum, Union:
		// We can't make unnamed versions of any of these types.
		return nil
	case Optional:
		if elem := TypeToReflect(t.Elem()); elem != nil {
			return reflect.PtrTo(elem)
		}
		return nil
	case List:
		if elem := TypeToReflect(t.Elem()); elem != nil {
			return reflect.SliceOf(elem)
		}
		return nil
	case Set:
		if key := TypeToReflect(t.Key()); key != nil {
			return reflect.MapOf(key, rtUnnamedEmptyStruct)
		}
		return nil
	case Map:
		if key, elem := TypeToReflect(t.Key()), TypeToReflect(t.Elem()); key != nil && elem != nil {
			return reflect.MapOf(key, elem)
		}
		return nil
	case Struct:
		if t.NumField() == 0 {
			return rtUnnamedEmptyStruct
		}
		return nil
	default:
		return rtFromKind[t.Kind()]
	}
}

var rtFromKind = [...]reflect.Type{
	Bool:       rtBool,
	Byte:       rtByte,
	Uint16:     rtUint16,
	Uint32:     rtUint32,
	Uint64:     rtUint64,
	Int16:      rtInt16,
	Int32:      rtInt32,
	Int64:      rtInt64,
	Float32:    rtFloat32,
	Float64:    rtFloat64,
	Complex64:  rtComplex64,
	Complex128: rtComplex128,
	String:     rtString,
	TypeObject: rtPtrToType,
}
