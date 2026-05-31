//ff:func feature=orchestrator type=test control=sequence
//ff:what TestGoModCache — GOMODCACHE→GOPATH→HOME fallback 우선순위 검증
package yongol

import (
	"os"
	"testing"
)

func TestGoModCacheNoHome(t *testing.T) {
	t.Setenv("GOMODCACHE", "")
	t.Setenv("GOPATH", "")
	// On unix os.UserHomeDir reads $HOME; empty → error → final return "".
	t.Setenv("HOME", "")
	if _, err := os.UserHomeDir(); err == nil {
		t.Skip("UserHomeDir still resolves a home directory on this platform")
	}
	if got := goModCache(); got != "" {
		t.Fatalf("expected \"\" when no home resolvable, got %q", got)
	}
}
