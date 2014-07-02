// This file was auto-generated by the veyron vdl tool.
// Source: application.vdl

// Package application defines the type for describing an application.
package application

// Envelope is a collection of metadata that describes an application.
type Envelope struct {
	// Arguments is an array of command-line arguments to be used when
	// executing the binary.
	Args []string
	// Binary is an object name that identifies the application binary.
	Binary string
	// Environment is a map that stores the environment variable values
	// to be used when executing the binary.
	Env []string
}
