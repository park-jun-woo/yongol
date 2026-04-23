//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRunGoSuccess — `go env GOROOT` 은 0-exit 이므로 runGo 성공

package gogin

import (
	"os/exec"
	"testing"
)

// TestRunGoSuccess — `go env GOROOT` returns zero; runGo should complete
// without error.
func TestRunGoSuccess(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available; skipping")
	}
	if err := runGo(t.TempDir(), "env", "GOROOT"); err != nil {
		t.Fatalf("runGo env GOROOT: %v", err)
	}
}
