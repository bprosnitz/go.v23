// This file was auto-generated by the veyron vdl tool.
// Source: common.vdl

package verror

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
)

var (
	// Unknown means the error has no known ID.  A more specific error should
	// always be used, if possible.  Unknown is typically only used when
	// automatically converting errors that do not contain an ID.
	ErrUnknown = Register("v.io/v23/verror.Unknown", NoRetry, "{1:}{2:} Error{:_}")
	// Internal means an internal error has occurred.  A more specific error
	// should always be used, if possible.
	ErrInternal = Register("v.io/v23/verror.Internal", NoRetry, "{1:}{2:} Internal error{:_}")
	// NotImplemented means that the request type is valid but that the method to
	// handle the request has not been implemented.
	ErrNotImplemented = Register("v.io/v23/verror.NotImplemented", NoRetry, "{1:}{2:} Not implemented{:_}")
	// EOF means the end-of-file has been reached; more generally, no more input
	// data is available.
	ErrEOF = Register("v.io/v23/verror.EOF", NoRetry, "{1:}{2:} EOF{:_}")
	// BadArg means the arguments to an operation are invalid or incorrectly
	// formatted.
	ErrBadArg = Register("v.io/v23/verror.BadArg", NoRetry, "{1:}{2:} Bad argument{:_}")
	// BadState means an operation was attempted on an object while the object was
	// in an incompatible state.
	ErrBadState = Register("v.io/v23/verror.BadState", NoRetry, "{1:}{2:} Invalid state{:_}")
	// BadEtag means the etag presented by the client was out of date or otherwise
	// invalid, likely because some other request caused the etag at the server to
	// change. The client should get a fresh etag and try again.
	// TODO(sadovsky): Rename "etag" to something else. HTTP etags are content
	// hashes, used to implement client-side response caching. We use etags for
	// for atomicity in read-modify-write operations, and generally recommend for
	// them to be implemented as (possibly lightly obfuscated) sequence numbers.
	ErrBadEtag = Register("v.io/v23/verror.BadEtag", NoRetry, "{1:}{2:} Etag is out of date")
	// Exist means that the requested item already exists; typically returned when
	// an attempt to create an item fails because it already exists.
	ErrExist = Register("v.io/v23/verror.Exist", NoRetry, "{1:}{2:} Already exists{:_}")
	// NoExist means that the requested item does not exist; typically returned
	// when an attempt to lookup an item fails because it does not exist.
	ErrNoExist = Register("v.io/v23/verror.NoExist", NoRetry, "{1:}{2:} Does not exist{:_}")
	// NoExistOrNoAccess means that either the requested item does not exist, or
	// is inaccessible.  Typically returned when the distinction between existence
	// and inaccessiblity needs to remain hidden, as a privacy feature.
	ErrNoExistOrNoAccess = Register("v.io/v23/verror.NoExistOrNoAccess", NoRetry, "{1:}{2:} Does not exist or access denied{:_}")
	// The following errors can occur during the process of establishing
	// an RPC connection.
	// NoExist (see above) is returned if the name of the server fails to
	// resolve any addresses.
	// NoServers is returned when the servers returned for the supplied name
	// are somehow unusable or unreachable by the client.
	// NoAccess is returned when a server does not authorize a client.
	// NotTrusted is returned when a client does not trust a server.
	//
	// TODO(toddw): These errors and descriptions were added by Cos; consider
	// moving the IPC-related ones into the ipc package.
	ErrNoServers        = Register("v.io/v23/verror.NoServers", RetryRefetch, "{1:}{2:} No usable servers found{:_}")
	ErrNoAccess         = Register("v.io/v23/verror.NoAccess", RetryRefetch, "{1:}{2:} Access denied{:_}")
	ErrNotTrusted       = Register("v.io/v23/verror.NotTrusted", RetryRefetch, "{1:}{2:} Client does not trust server{:_}")
	ErrNoServersAndAuth = Register("v.io/v23/verror.NoServersAndAuth", RetryRefetch, "{1:}{2:} Has no usable servers and is either not trusted or access was denied{:_}")
	// Aborted means that an operation was not completed because it was aborted by
	// the receiver.  A more specific error should be used if it would help the
	// caller decide how to proceed.
	ErrAborted = Register("v.io/v23/verror.Aborted", NoRetry, "{1:}{2:} Aborted{:_}")
	// BadProtocol means that an operation was not completed because of a protocol
	// or codec error.
	ErrBadProtocol = Register("v.io/v23/verror.BadProtocol", NoRetry, "{1:}{2:} Bad protocol or type{:_}")
	// Canceled means that an operation was not completed because it was
	// explicitly cancelled by the caller.
	ErrCanceled = Register("v.io/v23/verror.Canceled", NoRetry, "{1:}{2:} Canceled{:_}")
	// Timeout means that an operation was not completed before the time deadline
	// for the operation.
	ErrTimeout = Register("v.io/v23/verror.Timeout", NoRetry, "{1:}{2:} Timeout{:_}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknown.ID), "{1:}{2:} Error{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInternal.ID), "{1:}{2:} Internal error{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNotImplemented.ID), "{1:}{2:} Not implemented{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrEOF.ID), "{1:}{2:} EOF{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadArg.ID), "{1:}{2:} Bad argument{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadState.ID), "{1:}{2:} Invalid state{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadEtag.ID), "{1:}{2:} Etag is out of date")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExist.ID), "{1:}{2:} Already exists{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoExist.ID), "{1:}{2:} Does not exist{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoExistOrNoAccess.ID), "{1:}{2:} Does not exist or access denied{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoServers.ID), "{1:}{2:} No usable servers found{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoAccess.ID), "{1:}{2:} Access denied{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNotTrusted.ID), "{1:}{2:} Client does not trust server{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoServersAndAuth.ID), "{1:}{2:} Has no usable servers and is either not trusted or access was denied{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAborted.ID), "{1:}{2:} Aborted{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadProtocol.ID), "{1:}{2:} Bad protocol or type{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCanceled.ID), "{1:}{2:} Canceled{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrTimeout.ID), "{1:}{2:} Timeout{:_}")
}

// NewErrUnknown returns an error with the ErrUnknown ID.
func NewErrUnknown(ctx *context.T) error {
	return New(ErrUnknown, ctx)
}

// NewErrInternal returns an error with the ErrInternal ID.
func NewErrInternal(ctx *context.T) error {
	return New(ErrInternal, ctx)
}

// NewErrNotImplemented returns an error with the ErrNotImplemented ID.
func NewErrNotImplemented(ctx *context.T) error {
	return New(ErrNotImplemented, ctx)
}

// NewErrEOF returns an error with the ErrEOF ID.
func NewErrEOF(ctx *context.T) error {
	return New(ErrEOF, ctx)
}

// NewErrBadArg returns an error with the ErrBadArg ID.
func NewErrBadArg(ctx *context.T) error {
	return New(ErrBadArg, ctx)
}

// NewErrBadState returns an error with the ErrBadState ID.
func NewErrBadState(ctx *context.T) error {
	return New(ErrBadState, ctx)
}

// NewErrBadEtag returns an error with the ErrBadEtag ID.
func NewErrBadEtag(ctx *context.T) error {
	return New(ErrBadEtag, ctx)
}

// NewErrExist returns an error with the ErrExist ID.
func NewErrExist(ctx *context.T) error {
	return New(ErrExist, ctx)
}

// NewErrNoExist returns an error with the ErrNoExist ID.
func NewErrNoExist(ctx *context.T) error {
	return New(ErrNoExist, ctx)
}

// NewErrNoExistOrNoAccess returns an error with the ErrNoExistOrNoAccess ID.
func NewErrNoExistOrNoAccess(ctx *context.T) error {
	return New(ErrNoExistOrNoAccess, ctx)
}

// NewErrNoServers returns an error with the ErrNoServers ID.
func NewErrNoServers(ctx *context.T) error {
	return New(ErrNoServers, ctx)
}

// NewErrNoAccess returns an error with the ErrNoAccess ID.
func NewErrNoAccess(ctx *context.T) error {
	return New(ErrNoAccess, ctx)
}

// NewErrNotTrusted returns an error with the ErrNotTrusted ID.
func NewErrNotTrusted(ctx *context.T) error {
	return New(ErrNotTrusted, ctx)
}

// NewErrNoServersAndAuth returns an error with the ErrNoServersAndAuth ID.
func NewErrNoServersAndAuth(ctx *context.T) error {
	return New(ErrNoServersAndAuth, ctx)
}

// NewErrAborted returns an error with the ErrAborted ID.
func NewErrAborted(ctx *context.T) error {
	return New(ErrAborted, ctx)
}

// NewErrBadProtocol returns an error with the ErrBadProtocol ID.
func NewErrBadProtocol(ctx *context.T) error {
	return New(ErrBadProtocol, ctx)
}

// NewErrCanceled returns an error with the ErrCanceled ID.
func NewErrCanceled(ctx *context.T) error {
	return New(ErrCanceled, ctx)
}

// NewErrTimeout returns an error with the ErrTimeout ID.
func NewErrTimeout(ctx *context.T) error {
	return New(ErrTimeout, ctx)
}
