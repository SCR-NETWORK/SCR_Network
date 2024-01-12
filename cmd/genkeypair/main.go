package main

import (
	"fmt"
	"github.com/SCR-NETWORK/SCR_Network/cmd/SCRwallet/libSCRwallet"
	"github.com/SCR-NETWORK/SCR_Network/cmd/SCRwallet/libSCRwallet"
	"github.com/SCR-NETWORK/SCR_Network/util"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(err)
	}

	privateKey, publicKey, err := libSCRwallet.CreateKeyPair(false)
	privateKey, publicKey, err := libSCRwallet.CreateKeyPair(false)
	if err != nil {
		panic(err)
	}

	addr, err := util.NewAddressPublicKey(publicKey, cfg.NetParams().Prefix)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Private key: %x\n", privateKey)
	fmt.Printf("Address: %s\n", addr)
}
