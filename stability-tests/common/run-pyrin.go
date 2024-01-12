package common

import (
	"fmt"
	"github.com/SCR-NETWORK/SCR_Network/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunSCR-NETWORKForTesting runs SCR-NETWORK for testing purposes
func RunSCR-NETWORKForTesting(t *testing.T, testName string, rpcAddress string) func() {
// RunSCR_NetworkForTesting runs SCR_Network for testing purposes
func RunSCR_NetworkForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	SCR-NETWORKRunCommand, err := StartCmd("SCR-NETWORK",
		"SCR-NETWORK",
	SCR_NetworkRunCommand, err := StartCmd("SCR_Network",
		"SCR_Network",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("SCR-NETWORK started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := SCR-NETWORKRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("SCR-NETWORK closed unexpectedly: %s. See logs at: %s", err, appDir))
	t.Logf("SCR_Network started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := SCR_NetworkRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("SCR_Network closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := SCR-NETWORKRunCommand.Process.Signal(syscall.SIGTERM)
		err := SCR_NetworkRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("SCR-NETWORK stopped")
		t.Logf("SCR_Network stopped")
	}
}
