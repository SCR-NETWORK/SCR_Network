package standalone

import (
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/logger"
	"github.com/SCR-NETWORK/SCR_Network/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
