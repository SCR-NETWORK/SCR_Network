package main

import (
	"fmt"
	"os"

	"github.com/SCR-NETWORK/SCR_Network/infrastructure/config"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/standalone"
)

func connectToNode() *standalone.Routes {
	cfg := activeConfig()

	SCR-NETWORKConfig := config.DefaultConfig()
	SCR-NETWORKConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(SCR-NETWORKConfig)
	SCR_NetworkConfig := config.DefaultConfig()
	SCR_NetworkConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(SCR_NetworkConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating minimalNetAdapter: %+v", err)
		os.Exit(1)
	}
	routes, err := minimalNetAdapter.Connect(cfg.NodeP2PAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to node: %+v", err)
		os.Exit(1)
	}
	return routes
}
