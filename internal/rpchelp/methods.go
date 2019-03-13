// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
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
	{"accountaddressindex", []interface{}{(*int)(nil)}},
	{"accountsyncaddressindex", nil},
	{"addmultisigaddress", returnsString},
	{"addticket", nil},
	{"consolidate", returnsString},
	{"createmultisig", []interface{}{(*pfcjson.CreateMultiSigResult)(nil)}},
	{"createnewaccount", nil},
	{"dumpprivkey", returnsString},
	{"exportwatchingwallet", returnsString},
	{"generatevote", []interface{}{(*pfcjson.GenerateVoteResult)(nil)}},
	{"getaccountaddress", returnsString},
	{"getaccount", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", []interface{}{(*pfcjson.GetBalanceResult)(nil)}},
	{"getbestblockhash", returnsString},
	{"getbestblock", []interface{}{(*pfcjson.GetBestBlockResult)(nil)}},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*pfcjson.InfoWalletResult)(nil)}},
	{"getmasterpubkey", []interface{}{(*string)(nil)}},
	{"getmultisigoutinfo", []interface{}{(*pfcjson.GetMultisigOutInfoResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"getstakeinfo", []interface{}{(*pfcjson.GetStakeInfoResult)(nil)}},
	{"getticketfee", returnsNumber},
	{"gettickets", []interface{}{(*pfcjson.GetTicketsResult)(nil)}},
	{"gettransaction", []interface{}{(*pfcjson.GetTransactionResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"getvotechoices", []interface{}{(*pfcjson.GetVoteChoicesResult)(nil)}},
	{"getwalletfee", returnsNumber},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"importscript", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"listlockunspent", []interface{}{(*[]pfcjson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]pfcjson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]pfcjson.ListReceivedByAddressResult)(nil)}},
	{"listscripts", []interface{}{(*pfcjson.ListScriptsResult)(nil)}},
	{"listsinceblock", []interface{}{(*pfcjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*pfcjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"purchaseticket", returnsString},
	{"redeemmultisigout", []interface{}{(*pfcjson.RedeemMultiSigOutResult)(nil)}},
	{"redeemmultisigouts", []interface{}{(*pfcjson.RedeemMultiSigOutResult)(nil)}},
	{"renameaccount", nil},
	{"rescanwallet", nil},
	{"revoketickets", nil},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendtomultisig", returnsString},
	{"setticketfee", returnsBool},
	{"settxfee", returnsBool},
	{"setvotechoice", nil},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*pfcjson.SignRawTransactionResult)(nil)}},
	{"signrawtransactions", []interface{}{(*pfcjson.SignRawTransactionsResult)(nil)}},
	{"stakepooluserinfo", []interface{}{(*pfcjson.StakePoolUserInfoResult)(nil)}},
	{"startautobuyer", nil},
	{"stopautobuyer", nil},
	{"sweepaccount", []interface{}{(*pfcjson.SweepAccountResult)(nil)}},
	{"ticketsforaddress", returnsBool},
	{"validateaddress", []interface{}{(*pfcjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"version", []interface{}{(*map[string]pfcjson.VersionResult)(nil)}},
	{"walletinfo", []interface{}{(*pfcjson.WalletInfoResult)(nil)}},
	{"walletislocked", returnsBool},
	{"walletlock", nil},
	{"walletpassphrasechange", nil},
	{"walletpassphrase", nil},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
