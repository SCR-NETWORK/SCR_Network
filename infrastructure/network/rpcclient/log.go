package rpcclient

import (
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/logger"
	"github.com/SCR-NETWORK/SCR_Network/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
