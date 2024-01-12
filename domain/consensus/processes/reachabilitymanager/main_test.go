package reachabilitymanager_test

import (
	"os"
	"testing"

	"github.com/SCR-NETWORK/SCR_Network/infrastructure/logger"
)

const logLevel = logger.LevelWarn

func TestMain(m *testing.M) {
	logger.SetLogLevels(logLevel)
	logger.InitLogStdout(logLevel)
	os.Exit(m.Run())
}
