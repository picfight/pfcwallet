// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import (
	"fmt"

	"github.com/picfight/pfcd/pfcjson"
	"github.com/picfight/pfcwallet/errors"
)

func convertError(err error) *pfcjson.RPCError {
	if err, ok := err.(*pfcjson.RPCError); ok {
		return err
	}

	code := pfcjson.ErrRPCWallet
	if err, ok := err.(*errors.Error); ok {
		switch err.Kind {
		case errors.Bug:
			code = pfcjson.ErrRPCInternal.Code
		case errors.Encoding:
			code = pfcjson.ErrRPCInvalidParameter
		case errors.Locked:
			code = pfcjson.ErrRPCWalletUnlockNeeded
		case errors.Passphrase:
			code = pfcjson.ErrRPCWalletPassphraseIncorrect
		case errors.NoPeers:
			code = pfcjson.ErrRPCClientNotConnected
		case errors.InsufficientBalance:
			code = pfcjson.ErrRPCWalletInsufficientFunds
		}
	}
	return &pfcjson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcError(code pfcjson.RPCErrorCode, err error) *pfcjson.RPCError {
	return &pfcjson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcErrorf(code pfcjson.RPCErrorCode, format string, args ...interface{}) *pfcjson.RPCError {
	return &pfcjson.RPCError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// Errors variables that are defined once here to avoid duplication.
var (
	errUnloadedWallet = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWallet,
		Message: "request requires a wallet but wallet has not loaded yet",
	}

	errClientNotConnected = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCClientNotConnected,
		Message: "disconnected from network",
	}

	errAccountNotFound = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWalletInvalidAccountName,
		Message: "account not found",
	}

	errAddressNotInWallet = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWallet,
		Message: "address not found in wallet",
	}

	errNotImportedAccount = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWallet,
		Message: "imported addresses must belong to the imported account",
	}

	errNeedPositiveAmount = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCInvalidParameter,
		Message: "amount must be positive",
	}

	errWalletUnlockNeeded = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCWalletUnlockNeeded,
		Message: "enter the wallet passphrase with walletpassphrase first",
	}

	errReservedAccountName = &pfcjson.RPCError{
		Code:    pfcjson.ErrRPCInvalidParameter,
		Message: "account name is reserved by RPC server",
	}
)
