// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"github.com/picfight/pfcd/chaincfg"
	"github.com/picfight/pfcwallet/netparams"
)

var activeNet = &PfcNetParams

// PfcNetParams contains parameters specific running pfcwallet and
// pfcd on the main network.
var PfcNetParams = netparams.Params{
	Params:            &chaincfg.PicFightCoinNetParams,
	JSONRPCClientPort: "9109",
	JSONRPCServerPort: "9110",
	GRPCServerPort:    "9111",
}
