// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// Package access defines types and services for dynamic access control
// in Veyron.  Examples: "allow app to read this photo", "prevent user
// from modifying this file".
//
// Target Developers
//
// Developers creating functionality to share data or services between
// multiple users/devices/apps.
//
// Overview
//
// Veyron objects provide GetACL and SetACL methods.  An ACL (Access Control
// List) contains the set of blessings that grant principals access to the
// object. All methods on objects can have "tags" on them and the access
// control list used for the method is selected based on that tag (from a
// TaggedACLMap).
//
// An object can have multiple names, so GetACL and SetACL can be invoked on
// any of these names, but the object itself has a single ACL.
//
// SetACL completely replaces the TaggedACLMap. To perform an atomic
// read-modify-write of the ACL, use the etag parameter.
//   client := access.ObjectClient(name)
//   for {
//     acl, etag, err := client.GetACL()
//     if err != nil {
//       return err
//     }
//     acl[newTag] = ACL{In: []security.BlessingPattern{newPattern}}
//     // Use the same etag with the modified acl to ensure that no other client
//     // has modified the acl since GetACL returned.
//     if err := client.SetACL(acl, etag); err != nil {
//       if verror.Is(err, access.ErrBadEtag) {
//         // Another client replaced the ACL after our GetACL returned.
//         // Try again.
//         continue
//       }
//       return err
//     }
//   }
//
// Conventions
//
// Service implementors should follow the conventions below to be consistent
// with other parts of Veyron and with each other.
//
// All methods that create an object (e.g. Put, Mount, Link) should take an
// optional ACL parameter.  If the ACL is not specified, the new object, O,
// copies its ACL from the parent.  Subsequent changes to the parent ACL are
// not automatically propagated to O.  Instead, a client library must make
// recursive ACL changes.
//
// Resolve access is required on all components of a name, except the last one, in
// order to access the object referenced by that name.  For example, for
// principal P to access the name "a/b/c", P must have resolve access to "a"
// and "a/b".
//
// The Resolve tag means that a principal can traverse that component of the name to
// access the child.  It does not give the principal permission to list the
// children via Glob or a similar method.  For example, a server might have an
// object named "home" with a child for each user of the system.  If these
// users were allowed to list the contents of "home", they could discover the
// other users of the system.  That could be a privacy violation.  Without
// Resolve, every user of the system would need read access to "home" to
// access "home/<user>".  If the user called Glob("home/*"), it would then be
// up to the server to filter out the names that the user could not access.
// That could be a very expensive operation if there were a lot of children of
// "home".  Resolve protects these servers against potential denial of
// service attacks on these large, shared directories.
//
// Groups and blessings allow for sweeping access changes.  A group is
// suitable for saying that the same set of principals have access to a set of
// unrelated resources (e.g. docs, VMs, images).  See the Group API for a
// complete description.  A blessing is useful for controlling access to objects
// that are always accessed together.  For example, a document may have
// embedded images and comments, each with a unique name.  When accessing a
// document, the server would generate a blessing that the client would use to
// fetch the images and comments; the images and comments would have this
// blessed identity in their ACLs.  Changes to the document’s ACL are
// therefore “propagated” to the images and comments.
//
// Some services will want a concept of implicit access control.  They are
// free to implement this as is best for their service.  However, GetACL
// should respond with the correct ACL.  For example, a corporate file server
// would allow all employees to create their own directory and have full
// control within that directory.  Employees should not be allowed to modify
// other employee directories.  In other words, within the directory "home",
// employee E should be allowed to modify only "home/E".  The file server
// doesn't know the list of all employees a priori, so it uses an
// implementation-specific rule to map employee identities to their home
// directory.
package access

import (
	"v.io/core/veyron2/security"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__veyron2 "v.io/core/veyron2"
	__context "v.io/core/veyron2/context"
	__ipc "v.io/core/veyron2/ipc"
	__vdl "v.io/core/veyron2/vdl"
	__vdlutil "v.io/core/veyron2/vdl/vdlutil"
	__verror "v.io/core/veyron2/verror"
	__wiretype "v.io/core/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// ACL represents an Access Control List - a set of blessings that should be
// granted access.
type ACL struct {
	// In denotes the set of blessings (represented as BlessingPatterns) that
	// should be granted access, unless blacklisted by an entry in NotIn.
	//
	// For example:
	//   In: {"alice/family/..."}
	// grants access to a principal that presents at least one of "alice",
	// "alice/family", "alice/family/friend" etc. as a blessing.
	In []security.BlessingPattern
	// NotIn denotes the set of blessings (and their delegates) that
	// have been explicitly blacklisted from the In set.
	//
	// For example:
	//   In: {"alice/friend/..."}, NotIn: {"alice/friend/bob"}
	// grants access to principals that present "alice", "alice/friend",
	// "alice/friend/carol" etc. but NOT to a principal that presents
	// "alice/friend/bob" or "alice/friend/bob/spouse" etc.
	NotIn []string
}

func (ACL) __VDLReflect(struct {
	Name string "v.io/core/veyron2/services/security/access.ACL"
}) {
}

// TaggedACLMap maps string tags to access control lists specifying the
// blessings required to invoke methods with that tag.
//
// These tags are meant to add a layer of interposition between the set of
// users (blessings, specifically) and the set of methods, much like "Roles" do
// in Role Based Access Control.
// (http://en.wikipedia.org/wiki/Role-based_access_control)
type TaggedACLMap map[string]ACL

func (TaggedACLMap) __VDLReflect(struct {
	Name string "v.io/core/veyron2/services/security/access.TaggedACLMap"
}) {
}

// Tag is used to associate methods with an ACL in a TaggedACLMap.
//
// While services can define their own tag type and values, many
// services should be able to use the type and values defined in
// this package.
type Tag string

func (Tag) __VDLReflect(struct {
	Name string "v.io/core/veyron2/services/security/access.Tag"
}) {
}

func init() {
	__vdl.Register(ACL{})
	__vdl.Register(TaggedACLMap(nil))
	__vdl.Register(Tag(""))
}

const Admin = Tag("Admin") // Operations that require privileged access for object administration.

const Debug = Tag("Debug") // Operations that return debugging information (e.g., logs, statistics etc.) about the object.

const Read = Tag("Read") // Operations that do not mutate the state of the object.

const Write = Tag("Write") // Operations that mutate the state of the object.

const Resolve = Tag("Resolve") // Operations involving namespace navigation.

// The etag passed to SetACL is invalid.  Likely, another client set
// the ACL already and invalidated the etag.  Use GetACL to fetch a
// fresh etag.
const ErrBadEtag = __verror.ID("v.io/core/veyron2/services/security/access.ErrBadEtag")

// The ACL is too big.  Use groups to represent large sets of principals.
const ErrTooBig = __verror.ID("v.io/core/veyron2/services/security/access.ErrTooBig")

// ObjectClientMethods is the client interface
// containing Object methods.
//
// Object provides access control for Veyron objects.
//
// Veyron services implementing dynamic access control would typically
// embed this interface and tag additional methods defined by the service
// with one of Admin, Read, Write, Resolve etc. For example,
// the VDL definition of the object would be:
//
//   package mypackage
//
//   import "v.io/core/veyron2/security/access"
//
//   type MyObject interface {
//     access.Object
//     MyRead()  (string, error) {access.Read}
//     MyWrite(string) error     {access.Write}
//   }
//
// If the set of pre-defined tags is insufficient, services may define their
// own tag type and annotate all methods with this new type.
// Instead of embedding this Object interface, define SetACL and GetACL in
// their own interface. Authorization policies will typically respect
// annotations of a single type. For example, the VDL definition of an object
// would be:
//
//  package mypackage
//
//  import "v.io/core/veyron2/security/access"
//
//  type MyTag string
//
//  const (
//    Blue = MyTag("Blue")
//    Red  = MyTag("Red")
//  )
//
//  type MyObject interface {
//    MyMethod() (string, error) {Blue}
//
//    // Allow clients to change access via the access.Object interface:
//    SetACL(acl access.TaggedACLMap, etag string) error         {Red}
//    GetACL() (acl access.TaggedACLMap, etag string, err error) {Blue}
//  }
type ObjectClientMethods interface {
	// SetACL replaces the current ACL for an object.  etag allows for optional,
	// optimistic concurrency control.  If non-empty, etag's value must come
	// from GetACL.  If any client has successfully called SetACL in the
	// meantime, the etag will be stale and SetACL will fail.
	//
	// ACL objects are expected to be small.  It is up to the implementation to
	// define the exact limit, though it should probably be around 100KB.  Large
	// lists of principals should use the Group API or blessings.
	//
	// There is some ambiguity when calling SetACL on a mount point.  Does it
	// affect the mount itself or does it affect the service endpoint that the
	// mount points to?  The chosen behavior is that it affects the service
	// endpoint.  To modify the mount point's ACL, use ResolveToMountTable
	// to get an endpoint and call SetACL on that.  This means that clients
	// must know when a name refers to a mount point to change its ACL.
	SetACL(ctx __context.T, acl TaggedACLMap, etag string, opts ...__ipc.CallOpt) error
	// GetACL returns the complete, current ACL for an object.  The returned etag
	// can be passed to a subsequent call to SetACL for optimistic concurrency
	// control. A successful call to SetACL will invalidate etag, and the client
	// must call GetACL again to get the current etag.
	GetACL(__context.T, ...__ipc.CallOpt) (acl TaggedACLMap, etag string, err error)
}

// ObjectClientStub adds universal methods to ObjectClientMethods.
type ObjectClientStub interface {
	ObjectClientMethods
	__ipc.UniversalServiceMethods
}

// ObjectClient returns a client stub for Object.
func ObjectClient(name string, opts ...__ipc.BindOpt) ObjectClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implObjectClientStub{name, client}
}

type implObjectClientStub struct {
	name   string
	client __ipc.Client
}

func (c implObjectClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implObjectClientStub) SetACL(ctx __context.T, i0 TaggedACLMap, i1 string, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "SetACL", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implObjectClientStub) GetACL(ctx __context.T, opts ...__ipc.CallOpt) (o0 TaggedACLMap, o1 string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetACL", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &o1, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implObjectClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// ObjectServerMethods is the interface a server writer
// implements for Object.
//
// Object provides access control for Veyron objects.
//
// Veyron services implementing dynamic access control would typically
// embed this interface and tag additional methods defined by the service
// with one of Admin, Read, Write, Resolve etc. For example,
// the VDL definition of the object would be:
//
//   package mypackage
//
//   import "v.io/core/veyron2/security/access"
//
//   type MyObject interface {
//     access.Object
//     MyRead()  (string, error) {access.Read}
//     MyWrite(string) error     {access.Write}
//   }
//
// If the set of pre-defined tags is insufficient, services may define their
// own tag type and annotate all methods with this new type.
// Instead of embedding this Object interface, define SetACL and GetACL in
// their own interface. Authorization policies will typically respect
// annotations of a single type. For example, the VDL definition of an object
// would be:
//
//  package mypackage
//
//  import "v.io/core/veyron2/security/access"
//
//  type MyTag string
//
//  const (
//    Blue = MyTag("Blue")
//    Red  = MyTag("Red")
//  )
//
//  type MyObject interface {
//    MyMethod() (string, error) {Blue}
//
//    // Allow clients to change access via the access.Object interface:
//    SetACL(acl access.TaggedACLMap, etag string) error         {Red}
//    GetACL() (acl access.TaggedACLMap, etag string, err error) {Blue}
//  }
type ObjectServerMethods interface {
	// SetACL replaces the current ACL for an object.  etag allows for optional,
	// optimistic concurrency control.  If non-empty, etag's value must come
	// from GetACL.  If any client has successfully called SetACL in the
	// meantime, the etag will be stale and SetACL will fail.
	//
	// ACL objects are expected to be small.  It is up to the implementation to
	// define the exact limit, though it should probably be around 100KB.  Large
	// lists of principals should use the Group API or blessings.
	//
	// There is some ambiguity when calling SetACL on a mount point.  Does it
	// affect the mount itself or does it affect the service endpoint that the
	// mount points to?  The chosen behavior is that it affects the service
	// endpoint.  To modify the mount point's ACL, use ResolveToMountTable
	// to get an endpoint and call SetACL on that.  This means that clients
	// must know when a name refers to a mount point to change its ACL.
	SetACL(ctx __ipc.ServerContext, acl TaggedACLMap, etag string) error
	// GetACL returns the complete, current ACL for an object.  The returned etag
	// can be passed to a subsequent call to SetACL for optimistic concurrency
	// control. A successful call to SetACL will invalidate etag, and the client
	// must call GetACL again to get the current etag.
	GetACL(__ipc.ServerContext) (acl TaggedACLMap, etag string, err error)
}

// ObjectServerStubMethods is the server interface containing
// Object methods, as expected by ipc.Server.
// There is no difference between this interface and ObjectServerMethods
// since there are no streaming methods.
type ObjectServerStubMethods ObjectServerMethods

// ObjectServerStub adds universal methods to ObjectServerStubMethods.
type ObjectServerStub interface {
	ObjectServerStubMethods
	// Describe the Object interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// ObjectServer returns a server stub for Object.
// It converts an implementation of ObjectServerMethods into
// an object that may be used by ipc.Server.
func ObjectServer(impl ObjectServerMethods) ObjectServerStub {
	stub := implObjectServerStub{
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

type implObjectServerStub struct {
	impl ObjectServerMethods
	gs   *__ipc.GlobState
}

func (s implObjectServerStub) SetACL(ctx __ipc.ServerContext, i0 TaggedACLMap, i1 string) error {
	return s.impl.SetACL(ctx, i0, i1)
}

func (s implObjectServerStub) GetACL(ctx __ipc.ServerContext) (TaggedACLMap, string, error) {
	return s.impl.GetACL(ctx)
}

func (s implObjectServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implObjectServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{ObjectDesc}
}

// ObjectDesc describes the Object interface.
var ObjectDesc __ipc.InterfaceDesc = descObject

// descObject hides the desc to keep godoc clean.
var descObject = __ipc.InterfaceDesc{
	Name:    "Object",
	PkgPath: "v.io/core/veyron2/services/security/access",
	Doc:     "// Object provides access control for Veyron objects.\n//\n// Veyron services implementing dynamic access control would typically\n// embed this interface and tag additional methods defined by the service\n// with one of Admin, Read, Write, Resolve etc. For example,\n// the VDL definition of the object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/core/veyron2/security/access\"\n//\n//   type MyObject interface {\n//     access.Object\n//     MyRead()  (string, error) {access.Read}\n//     MyWrite(string) error     {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n// Instead of embedding this Object interface, define SetACL and GetACL in\n// their own interface. Authorization policies will typically respect\n// annotations of a single type. For example, the VDL definition of an object\n// would be:\n//\n//  package mypackage\n//\n//  import \"v.io/core/veyron2/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetACL(acl access.TaggedACLMap, etag string) error         {Red}\n//    GetACL() (acl access.TaggedACLMap, etag string, err error) {Blue}\n//  }",
	Methods: []__ipc.MethodDesc{
		{
			Name: "SetACL",
			Doc:  "// SetACL replaces the current ACL for an object.  etag allows for optional,\n// optimistic concurrency control.  If non-empty, etag's value must come\n// from GetACL.  If any client has successfully called SetACL in the\n// meantime, the etag will be stale and SetACL will fail.\n//\n// ACL objects are expected to be small.  It is up to the implementation to\n// define the exact limit, though it should probably be around 100KB.  Large\n// lists of principals should use the Group API or blessings.\n//\n// There is some ambiguity when calling SetACL on a mount point.  Does it\n// affect the mount itself or does it affect the service endpoint that the\n// mount points to?  The chosen behavior is that it affects the service\n// endpoint.  To modify the mount point's ACL, use ResolveToMountTable\n// to get an endpoint and call SetACL on that.  This means that clients\n// must know when a name refers to a mount point to change its ACL.",
			InArgs: []__ipc.ArgDesc{
				{"acl", ``},  // TaggedACLMap
				{"etag", ``}, // string
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{Tag("Admin")},
		},
		{
			Name: "GetACL",
			Doc:  "// GetACL returns the complete, current ACL for an object.  The returned etag\n// can be passed to a subsequent call to SetACL for optimistic concurrency\n// control. A successful call to SetACL will invalidate etag, and the client\n// must call GetACL again to get the current etag.",
			OutArgs: []__ipc.ArgDesc{
				{"acl", ``},  // TaggedACLMap
				{"etag", ``}, // string
				{"err", ``},  // error
			},
			Tags: []__vdlutil.Any{Tag("Admin")},
		},
	},
}

func (s implObjectServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["GetACL"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "acl", Type: 68},
			{Name: "etag", Type: 3},
			{Name: "err", Type: 69},
		},
	}
	result.Methods["SetACL"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "acl", Type: 68},
			{Name: "etag", Type: 3},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 69},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x3, Name: "v.io/core/veyron2/security.BlessingPattern", Tags: []string(nil)}, __wiretype.SliceType{Elem: 0x41, Name: "", Tags: []string(nil)}, __wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x42, Name: "In"},
				__wiretype.FieldType{Type: 0x3d, Name: "NotIn"},
			},
			"v.io/core/veyron2/services/security/access.ACL", []string(nil)},
		__wiretype.MapType{Key: 0x3, Elem: 0x43, Name: "v.io/core/veyron2/services/security/access.TaggedACLMap", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}
