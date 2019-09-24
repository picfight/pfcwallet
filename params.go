// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"github.com/decred/dcrwallet/netparams"
	"github.com/picfight/pfcd/picfightcoin"
)

var activeNet = &PfcNetParams

// PfcNetParams contains parameters specific running dcrwallet and
// dcrd on the main network (wire.MainNet).
var PfcNetParams = netparams.Params{
	Params:            &picfightcoin.PicFightCoinNetParams,
	JSONRPCClientPort: "9109",
	JSONRPCServerPort: "9110",
	GRPCServerPort:    "9111",
}
