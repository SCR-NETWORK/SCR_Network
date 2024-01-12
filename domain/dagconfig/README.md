dagconfig

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/SCR-NETWORK/SCR_Network/dagconfig)

Package dagconfig defines DAG configuration parameters for the standard
SCR-NETWORK networks and provides the ability for callers to define their own custom
SCR-NETWORK networks.
SCR_Network networks and provides the ability for callers to define their own custom
SCR_Network networks.

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SCR-NETWORK/SCR_Network/util"
	"github.com/SCR-NETWORK/SCR_Network/domain/dagconfig"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet SCR network")

// By default (without --testnet), use mainnet.
var dagParams = &dagconfig.MainnetParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		dagParams = &dagconfig.TestnetParams
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKey := make([]byte, 32)
	addr, err := util.NewAddressPubKey(pubKey, dagParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```
