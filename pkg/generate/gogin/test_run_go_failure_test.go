//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRunGoFailure — 존재하지 않는 subcommand 면 runGo 가 에러를 반환

package gogin

import (
	"os/exec"
	"strings"
	"testing"
)

// TestRunGoFailure — an unknown subcommand produces a non-zero exit and
// runGo must wrap the error with the invoked command and stderr.
func TestRunGoFailure(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available; skipping")
	}
	err := runGo(t.TempDir(), "this-subcommand-does-not-exist")
	if err == nil {
		t.Fatal("runGo returned nil for invalid subcommand")
	}
	if !strings.Contains(err.Error(), "go this-subcommand-does-not-exist failed") {
		t.Fatalf("error missing command trace: %q", err.Error())
	}
}
