// Copyright (c) 2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build !generate

package rpchelp

import "github.com/picfight/pfcd/pfcjson"

// Common return types.
var (
	returnsBool        = []interface{}{(*bool)(nil)}
	returnsNumber      = []interface{}{(*float64)(nil)}
	returnsString      = []interface{}{(*string)(nil)}
	returnsStringArray = []interface{}{(*[]string)(nil)}
	returnsLTRArray    = []interface{}{(*[]pfcjson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"addmultisigaddress", returnsString},
	{"createmultisig", []interface{}{(*pfcjson.CreateMultiSigResult)(nil)}},
	{"dumpprivkey", returnsString},
	{"getaccount", returnsString},
	{"getaccountaddress", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", append(returnsNumber, returnsNumber[0])},
	{"getbestblockhash", returnsString},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*pfcjson.InfoWalletResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"gettransaction", []interface{}{(*pfcjson.GetTransactionResult)(nil)}},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listlockunspent", []interface{}{(*[]pfcjson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]pfcjson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]pfcjson.ListReceivedByAddressResult)(nil)}},
	{"listsinceblock", []interface{}{(*pfcjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*pfcjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"settxfee", returnsBool},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*pfcjson.SignRawTransactionResult)(nil)}},
	{"validateaddress", []interface{}{(*pfcjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"walletlock", nil},
	{"walletpassphrase", nil},
	{"walletpassphrasechange", nil},
	{"createnewaccount", nil},
	{"exportwatchingwallet", returnsString},
	{"getbestblock", []interface{}{(*pfcjson.GetBestBlockResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"renameaccount", nil},
	{"walletislocked", returnsBool},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
