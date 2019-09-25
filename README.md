pfcwallet
=========

pfcwallet is a daemon handling PicFight coin wallet functionality.  All interaction
with the wallet is performed over RPC.

pfcwallet provides two modes of operation to connect to the PicFight coin
network.  The first (and default) is to communicate with a single
trusted `pfcd` instance using JSON-RPC.  The second is a
privacy-preserving Simplified Payment Verification (SPV) mode (enabled
with the `--spv` flag) where the wallet connects either to specified
peers (with `--spvconnect`) or peers discovered from seeders and other
peers. Both modes can be switched between with just a restart of the
wallet. It is advised to avoid SPV mode for heavily-used wallets
which require downloading most blocks regardless.

Not all functionality is available when running in SPV mode.
Currently, the following features are disabled or unavailable to SPV wallets:

  * Voting

  * Revoking tickets before expiry

  * Determining exact number of live and missed tickets (as opposed to
    simply unspent).

Wallet clients interact with the wallet using one of two RPC servers:

  1. A legacy JSON-RPC server inspired by the Bitcoin Core rpc server

     The JSON-RPC server exists to ease the migration of wallet applications
     from Core, but complete compatibility is not guaranteed.  Some portions of
     the API (and especially accounts) have to work differently due to other
     design decisions (mostly due to BIP0044).  However, if you find a
     compatibility issue and feel that it could be reasonably supported, please
     report an issue.  This server is enabled by default as long as a username
     and password are provided.

  2. A gRPC server

     The gRPC server uses a new API built for pfcwallet, but the API is not
     stabilized.  This server is enabled by default and may be disabled with
     the config option `--nogrpc`.  If you don't mind applications breaking
     due to API changes, don't want to deal with issues of the legacy API, or
     need notifications for changes to the wallet, this is the RPC server to
     use. The gRPC server is documented [here](./rpc/documentation/README.md).

## Installation

### Build from source (all platforms)

Building or updating from source requires the following build dependencies:

- **Go**

  Installation instructions can be found here: https://golang.org/doc/install.
  It is recommended to add `$GOPATH/bin` to your `PATH` at this point.

To build and install from a checked-out repo, run `go install` in the repo's
root directory.  Some notes:

* Set the `GO111MODULE=on` environment variable if using Go 1.11 and building
  from within `GOPATH`.

* The `pfcwallet` executable will be installed to `$GOPATH/bin`.  `GOPATH`
  defaults to `$HOME/go` (or `%USERPROFILE%\go` on Windows) if unset.

```bash
set GO111MODULE=on
  go version
  go build -v ./...
  go install . ./cmd/...
```

## Testing

To run the tests locally:

```bash
./run_tests.sh 
```

## Getting Started

The following instructions detail how to get started with pfcwallet connecting
to a localhost pfcd.  Commands should be run in `cmd.exe` or PowerShell on
Windows, or any terminal emulator on *nix.

- Run the following command to start pfcd:

```bash
pfcd -u rpcuser -P rpcpass
```

- Run the following command to create a wallet:

```bash
pfcwallet -u rpcuser -P rpcpass --create
```

- Run the following command to start pfcwallet:

```bash
pfcwallet -u rpcuser -P rpcpass
```

If everything appears to be working, it is recommended at this point to
copy the sample pfcd and pfcwallet configurations and update with your
RPC username and password.

## Example run commands

Launch mining node:
```bash
pfcd
     --generate
     --miningaddr "JsCVh5SVDQovpW1dswaZNan2mfNWy6uRpPx"
     --listen=127.0.0.1:30000
     --rpclisten=127.0.0.1:30001
     --datadir=nodeA
     --rpccert=nodeA\rpc.cert
     --rpckey=nodeA\rpc.key     
     --txindex
     --addrindex
     --rpcuser=node.user
     --rpcpass=node.pass
```

Copy `nodeA\rpc.cert` to `wallet\pfcd-rpc.cert` 

Launch wallet, add `--create` flag for the first run:
```bash
pfcwallet
       --rpclisten=127.0.0.1:20002
       --rpcconnect=127.0.0.1:30001
       --appdata=wallet
       --cafile=wallet\pfcd-rpc.cert
       --rpckey=wallet\rpc.key
       --rpccert=wallet\rpc.cert
       --pfcdusername=node.user
       --pfcdpassword=node.pass
       --username=wallet.user
       --password=wallet.pass 
```
```bash
       --create 
```

Check balance:
```bash
pfcctl /u wallet.user
       /P wallet.pass
       /s 127.0.0.1:20002
       /c wallet\rpc.cert
       --wallet getbalance
```

Generate wallet address:
```bash
pfcctl /u wallet.user
       /P wallet.pass
       /s 127.0.0.1:20002
       /c wallet\rpc.cert
       --wallet getnewaddress
```

## Issue Tracker

The [integrated github issue tracker](https://github.com/picfight/pfcwallet/issues)
is used for this project.

## License

pfcwallet is licensed under the liberal ISC License.
