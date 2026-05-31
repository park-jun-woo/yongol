//ff:func feature=orchestrator type=test control=sequence
//ff:what TestGoModCache — GOMODCACHE→GOPATH→HOME fallback 우선순위 검증
package yongol

import (
	"path/filepath"
	"testing"
)

func TestGoModCacheHomeFallback(t *testing.T) {
	t.Setenv("GOMODCACHE", "")
	t.Setenv("GOPATH", "")
	// HOME-based fallback: result must end in go/pkg/mod (assuming a home dir
	// is resolvable, which it is in the test environment).
	got := goModCache()
	suffix := filepath.Join("go", "pkg", "mod")
	if got == "" {
		t.Skip("no resolvable home directory in this environment")
	}
	if filepath.Base(got) != "mod" || !endsWith(got, suffix) {
		t.Errorf("goModCache = %q, want path ending in %q", got, suffix)
	}
}
