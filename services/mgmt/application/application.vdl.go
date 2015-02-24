// This file was auto-generated by the veyron vdl tool.
// Source: application.vdl

// Package application defines the type for describing an application.
package application

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security"
)

// Envelope is a collection of metadata that describes an application.
type Envelope struct {
	// Title is the publisher-assigned application title.  Application
	// installations with the same title are considered as belonging to the
	// same application by the application management system.
	//
	// A change in the title signals a new application.
	Title string
	// Args is an array of command-line arguments to be used when executing
	// the binary.
	Args []string
	// Binary identifies the application binary.
	Binary SignedFile
	// Publisher represents the set of blessings that have been bound to
	// the principal who published this binary.
	Publisher security.WireBlessings
	// Env is an array that stores the environment variable values to be
	// used when executing the binary.
	Env []string
	// Packages is the set of packages to install on the local filesystem
	// before executing the binary
	Packages Packages
}

func (Envelope) __VDLReflect(struct {
	Name string "v.io/v23/services/mgmt/application.Envelope"
}) {
}

// Packages represents a set of packages. The map key is the local
// file/directory name, relative to the instance's packages directory, where the
// package should be installed. For archives, this name represents a directory
// into which the archive is to be extracted, and for regular files it
// represents the name for the file.  The map value is the package
// specification.
//
// Each object's media type determines how to install it.
//
// For example, with key=pkg1,value=SignedFile{File:binaryrepo/configfiles} (an
// archive), the "configfiles" package will be installed under the "pkg1"
// directory. With key=pkg2,value=SignedFile{File:binaryrepo/binfile} (a
// binary), the "binfile" file will be installed as the "pkg2" file.
//
// The keys must be valid file/directory names, without path separators.
//
// Any number of packages may be specified.
type Packages map[string]SignedFile

func (Packages) __VDLReflect(struct {
	Name string "v.io/v23/services/mgmt/application.Packages"
}) {
}

// SignedFile represents a file accompanied by a signature of its contents.
type SignedFile struct {
	//  File is the object name of the file.
	File string
	// Signature represents a signature on the sha256 hash of the file
	// contents by the publisher principal.
	Signature security.Signature
}

func (SignedFile) __VDLReflect(struct {
	Name string "v.io/v23/services/mgmt/application.SignedFile"
}) {
}

func init() {
	vdl.Register((*Envelope)(nil))
	vdl.Register((*Packages)(nil))
	vdl.Register((*SignedFile)(nil))
}

// Device manager application envelopes must present this title.
const DeviceManagerTitle = "device manager"
