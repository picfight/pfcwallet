// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import (
	"errors"

	"github.com/picfight/pfcd/pfcjson"
)

// TODO(jrick): There are several error paths which 'replace' various errors
// with a more appropiate error from the pfcjson package.  Create a map of
// these replacements so they can be handled once after an RPC handler has
// returned and before the error is marshaled.

// Error types to simplify the reporting of specific categories of
// errors, and their *pfcjson.RPCError creation.
type (
	// DeserializationError describes a failed deserializaion due to bad
	// user input.  It corresponds to pfcjson.ErrRPCDeserialization.
	DeserializationError struct {
		error
	}

	// InvalidParameterError describes an invalid parameter passed by
	// the user.  It corresponds to pfcjson.ErrRPCInvalidParameter.
	InvalidParameterError struct {
		error
	}

	// ParseError describes a failed parse due to bad user input.  It
	// corresponds to pfcjson.ErrRPCParse.
	ParseError struct {
		error
	}
)

// Errors variables that are defined once here to avoid duplication below.
var (
	ErrNeedPositiveAmount = InvalidParameterError{
		errors.New("amount must be positive"),
	}

	ErrNeedPositiveMinconf = InvalidParameterError{
		errors.New("minconf must be positive"),
	}

	ErrAddressNotInWallet = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWallet,
		Message: "address not found in wallet",
	}

	ErrAccountNameNotFound = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWalletInvalidAccountName,
		Message: "account name not found",
	}

	ErrUnloadedWallet = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWallet,
		Message: "Request requires a wallet but wallet has not loaded yet",
	}

	ErrWalletUnlockNeeded = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWalletUnlockNeeded,
		Message: "Enter the wallet passphrase with walletpassphrase first",
	}

	ErrNotImportedAccount = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWallet,
		Message: "imported addresses must belong to the imported account",
	}

	ErrNoTransactionInfo = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCNoTxInfo,
		Message: "No information for transaction",
	}

	ErrReservedAccountName = pfcjson.RPCError{
		Code:    pfcjson.ErrRPCInvalidParameter,
		Message: "Account name is reserved by RPC server",
	}
)
