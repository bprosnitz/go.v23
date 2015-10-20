// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: errors.vdl

package message

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var (
	ErrInvalidMsg         = verror.Register("v.io/v23/flow/message.InvalidMsg", verror.NoRetry, "{1:}{2:} message of type {3} and size {4} failed decoding at field {5}{:6}.")
	ErrInvalidSetupOption = verror.Register("v.io/v23/flow/message.InvalidSetupOption", verror.NoRetry, "{1:}{2:} setup option{:3} failed decoding at field{:4}.")
	ErrUnknownMsg         = verror.Register("v.io/v23/flow/message.UnknownMsg", verror.NoRetry, "{1:}{2:} unknown message type{:3}.")
	ErrMissingBlessings   = verror.Register("v.io/v23/flow/message.MissingBlessings", verror.NoRetry, "{1:}{2:} {3} message received with no blessings.")
	// TODO(mattr): Remove this after the transition is complete.
	ErrWrongProtocol = verror.Register("v.io/v23/flow/message.WrongProtocol", verror.NoRetry, "{1:}{2:} A message of the incorrect protocol version was detected.")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidMsg.ID), "{1:}{2:} message of type {3} and size {4} failed decoding at field {5}{:6}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidSetupOption.ID), "{1:}{2:} setup option{:3} failed decoding at field{:4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownMsg.ID), "{1:}{2:} unknown message type{:3}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrMissingBlessings.ID), "{1:}{2:} {3} message received with no blessings.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrWrongProtocol.ID), "{1:}{2:} A message of the incorrect protocol version was detected.")
}

// NewErrInvalidMsg returns an error with the ErrInvalidMsg ID.
func NewErrInvalidMsg(ctx *context.T, typ byte, size uint64, field uint64, err error) error {
	return verror.New(ErrInvalidMsg, ctx, typ, size, field, err)
}

// NewErrInvalidSetupOption returns an error with the ErrInvalidSetupOption ID.
func NewErrInvalidSetupOption(ctx *context.T, option uint64, field uint64) error {
	return verror.New(ErrInvalidSetupOption, ctx, option, field)
}

// NewErrUnknownMsg returns an error with the ErrUnknownMsg ID.
func NewErrUnknownMsg(ctx *context.T, typ byte) error {
	return verror.New(ErrUnknownMsg, ctx, typ)
}

// NewErrMissingBlessings returns an error with the ErrMissingBlessings ID.
func NewErrMissingBlessings(ctx *context.T, typ byte) error {
	return verror.New(ErrMissingBlessings, ctx, typ)
}

// NewErrWrongProtocol returns an error with the ErrWrongProtocol ID.
func NewErrWrongProtocol(ctx *context.T) error {
	return verror.New(ErrWrongProtocol, ctx)
}
