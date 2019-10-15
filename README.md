pfcwallet
=========

```pfcwallet``` is a daemon handling the PicFight coin wallet functionality.
 
The wallet connects to a ```pfcd``` node via RPC. All interactions
with the wallet are also performed over RPC.

## Installing and updating

### Setup

Building or updating from source requires the following build dependencies:

- **Git**

  Installation instructions can be found at https://git-scm.com or
  https://gitforwindows.org.
  
- **Go 1.13**

  Installation instructions can be found here: https://golang.org/doc/install.
  It is recommended to add `$GOPATH/bin` to your `PATH` at this point.

* The `pfcd` and `pfcwallet` executables will be installed to `$GOPATH/bin`.  `GOPATH`
  defaults to `$HOME/go` (or `%USERPROFILE%\go` on Windows) if unset.
  
### Build from source (all platforms)

Tip: You can always verify your steps against the Travis. Simply consult with the
```.travis.yml``` and the ```run_tests.sh``` for the details.

### Example of obtaining and building from source on Windows:

Checkout:
```bash
go get github.com/picfight/pfcd
go get github.com/picfight/pfcwallet
```

Build and install pfcd:
```bash
cd %GOPATH%
cd src/github.com/picfight/pfcd

set GO111MODULE=on
go build ./...
go install . ./cmd/...
```

Build and install pfcwallet:
```bash
cd %GOPATH%
cd src/github.com/picfight/pfcwallet

set GO111MODULE=on
go build ./...
go install . ./cmd/...
```

### Running Tests

To run the tests locally:

```bash
cd %GOPATH%
cd src/github.com/picfight/pfcwallet

set GO111MODULE=on
go build ./...
go clean -testcache
go test ./...
```

or simply
```bash
./run_tests.sh 
```

## Getting Started

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
copy the sample pfcd and pfcwallet configurations (.conf) and update them with your
RPC username and password.

Then simply run: 
```bash
pfcd
pfcwallet
``` 

## Example run commands

Launch a customized node:
```bash
pfcd
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

Launch wallet (add `--create` flag-option for the first run):
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

Generate new wallet address:
```bash
pfcctl /u wallet.user
       /P wallet.pass
       /s 127.0.0.1:20002
       /c wallet\rpc.cert
       --wallet getnewaddress
```

## Stake-mining

Participating in Proof-of-Stake (PoS) validation requires a wallet to be
running 24/7. The wallet needs to be always online so that it can be called
to validate block; if the wallet is unavailable then the votes will be missed
and no block reward will be received.

Tip: Stake-mining wallet will use all of its funds to buy voting tickets.
Thus it is recommended to create a dedicated wallet (with a new seed) for stake
mining and transfer to it some of your funds from your main wallet. Min possible
ticket price is 2 coins + fees. So you need at least ~2.1 coins to buy a ticket and to
participate in the stake mining.

To enable stake-mining in pfcwallet you need:

- Edit your `pfcwallet.conf` setting `pass=`%your wallet password% and the `enablevoting=1` flag.

- Run `pfcwallet` with the following flag: `--enableticketbuyer`

At this point you should have a running `pfcd`-node connected to internet
and syncing with the external world, and `pfcwallet` connected to the `pfcd`
listening to the blockchain updates and validating blocks on request.

Check your setup using the following commands:
```bash
 pfcctl --wallet getstakeinfo
 pfcctl --wallet walletinfo
``` 

Check balance:
```bash
pfcctl --wallet getbalance
```

Generate a new wallet address:
```bash
pfcctl --wallet getnewaddress
```

## SPV 

pfcwallet provides two modes of operation to connect to the PicFight coin
network.  The first (and default) is to communicate with a single
trusted `pfcd` instance using JSON-RPC.

The second is a privacy-preserving Simplified Payment Verification (SPV) mode (enabled
with the `--spv` flag), where the wallet connects either to specified
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
     from the Core, but complete compatibility is not guaranteed.  Some portions of
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

## Issue Tracker

The [integrated github issue tracker](https://github.com/picfight/pfcwallet/issues)
is used for this project.

## License

pfcwallet is licensed under the liberal ISC License.
