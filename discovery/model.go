// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package discovery defines types and interfaces for discovering services.
//
// TODO(jhahn): This is a work in progress and can change without notice.
package discovery

import (
	"v.io/v23/context"
	"v.io/v23/security"
)

// T is the interface for discovery operations; it is the client side library
// for the discovery service.
type T interface {
	// Advertise advertises the service to be discovered by "Scan" operations.
	// visibility is used to limit the principals that can see the advertisement. An
	// empty set means that there are no restrictions on visibility (i.e, equivalent
	// to []security.BlessingPattern{security.AllPrincipals}). Advertising will continue
	// until the context is canceled or exceeds its deadline and the returned channel
	// will be closed when it stops.
	//
	// If service.InstanceId is not specified, a random unique identifier will be assigned
	// to it. Any change to service will not be applied after advertising starts.
	//
	// It is an error to have simultaneously active advertisements for two identical
	// instances (service.InstanceId).
	Advertise(ctx *context.T, service *Service, visibility []security.BlessingPattern) (<-chan struct{}, error)

	// Scan scans services that match the query and returns the channel on which
	// new discovered services can be read. Scanning will continue until the context
	// is canceled or exceeds its deadline.
	//
	// Scan excludes the services that are advertised from the same T instance.
	//
	// The query is a WHERE expression of a syncQL query against advertised services, where
	// keys are InstanceIds and values are Services.
	//
	// Examples
	//
	//    v.InterfaceName = "v.io/i"
	//    v.InterfaceName = "v.io/i" AND v.Attrs["a"] = "v"
	//    v.Attrs["a"] = "v1" OR v.Attrs["a"] = "v2"
	//
	// SyncQL tutorial at:
	//    https://github.com/vanadium/docs/blob/master/tutorials/syncql-tutorial.md
	Scan(ctx *context.T, query string) (<-chan Update, error)
}
