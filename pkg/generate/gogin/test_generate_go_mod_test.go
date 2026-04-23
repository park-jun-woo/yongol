//ff:func feature=gen-gogin type=test control=sequence
//ff:what test_generate_go_mod — runGoModTidy + truncateStderr 단위 테스트
package gogin

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestRunGoModTidySuccess verifies that runGoModTidy succeeds on a minimal
// valid go.mod (no require directives — nothing to resolve).
func TestRunGoModTidySuccess(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available; skipping")
	}
	dir := t.TempDir()
	modContent := "module example.com/phase010/success\n\ngo 1.22\n"
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := runGoModTidy(dir); err != nil {
		t.Fatalf("runGoModTidy unexpected error: %v", err)
	}
}

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

// TestTruncateStderrNoTruncate — string fits under the limit, returned as-is
// except for trailing whitespace trimming.
func TestTruncateStderrNoTruncate(t *testing.T) {
	got := truncateStderr("hello\n", 64)
	if got != "hello" {
		t.Fatalf("truncateStderr = %q, want %q", got, "hello")
	}
}

// TestTruncateStderrTruncates — string over the limit is cut and marked.
func TestTruncateStderrTruncates(t *testing.T) {
	in := strings.Repeat("x", 50)
	got := truncateStderr(in, 10)
	want := strings.Repeat("x", 10) + "...(truncated)"
	if got != want {
		t.Fatalf("truncateStderr = %q, want %q", got, want)
	}
}

// TestTruncateStderrZeroLimit — limit<=0 short-circuits and just trims.
func TestTruncateStderrZeroLimit(t *testing.T) {
	got := truncateStderr("abc  \n", 0)
	if got != "abc" {
		t.Fatalf("truncateStderr(limit=0) = %q, want %q", got, "abc")
	}
}

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
