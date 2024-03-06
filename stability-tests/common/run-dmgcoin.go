package common

import (
	"fmt"
	"github.com/dmgcoin/dmgcoin/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunDmgcoinForTesting runs dmgcoin for testing purposes
func RunDmgcoinForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	dmgcoinRunCommand, err := StartCmd("DMGCOIN",
		"dmgcoin",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Dmgcoin started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := dmgcoinRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Dmgcoin closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := dmgcoinRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Dmgcoin stopped")
	}
}
