//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRunGoModTidyFailure — go.mod 누락 시 runGoModTidy 가 stderr 포함 에러 반환

package gogin

import (
	"os/exec"
	"strings"
	"testing"
)

// TestRunGoModTidyFailure verifies that runGoModTidy returns a wrapped error
// whose message mentions "go mod tidy failed" when go.mod is missing.
// go refuses to run `go mod tidy` without a go.mod, so this reliably
// produces a non-zero exit with stderr output that must be surfaced.
func TestRunGoModTidyFailure(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available; skipping")
	}
	dir := t.TempDir() // no go.mod inside

	err := runGoModTidy(dir)
	if err == nil {
		t.Fatal("runGoModTidy returned nil; expected error when go.mod is missing")
	}
	msg := err.Error()
	if !strings.Contains(msg, "go mod tidy failed") {
		t.Fatalf("error missing 'go mod tidy failed' prefix: %q", msg)
	}
	// stderr from `go` should be present (message contents vary by version,
	// but it always mentions go.mod / module somewhere). Keep the check
	// loose to avoid coupling to toolchain text.
	lower := strings.ToLower(msg)
	if !strings.Contains(lower, "go.mod") && !strings.Contains(lower, "module") {
		t.Fatalf("error message lacks stderr details: %q", msg)
	}
}
