package main

import (
	"fmt"
	"os"

	"github.com/SCR-NETWORK/SCR_Network/infrastructure/config"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/standalone"
	"github.com/SCR-NETWORK/SCR_Network/stability-tests/common"
	"github.com/SCR-NETWORK/SCR_Network/util/panics"
	"github.com/SCR-NETWORK/SCR_Network/util/profiling"
)

func main() {
	defer panics.HandlePanic(log, "applicationLevelGarbage-main", nil)
	err := parseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %+v", err)
		os.Exit(1)
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())
	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	SCR-NETWORKConfig := config.DefaultConfig()
	SCR-NETWORKConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(SCR-NETWORKConfig)
	SCR_NetworkConfig := config.DefaultConfig()
	SCR_NetworkConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(SCR_NetworkConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating minimalNetAdapter: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	blocksChan, err := readBlocks()
	if err != nil {
		log.Errorf("Error reading blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	err = sendBlocks(cfg.NodeP2PAddress, minimalNetAdapter, blocksChan)
	if err != nil {
		log.Errorf("Error sending blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}
}
