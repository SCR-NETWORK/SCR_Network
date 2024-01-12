package main

import (
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/logger"
	"github.com/SCR-NETWORK/SCR_Network/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
