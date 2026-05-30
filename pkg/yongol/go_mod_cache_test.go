//ff:func feature=orchestrator type=test control=sequence
//ff:what TestGoModCache — GOMODCACHE→GOPATH→HOME fallback 우선순위 검증

package yongol

import (
	"path/filepath"
	"testing"
)

func TestGoModCacheGOMODCACHE(t *testing.T) {
	t.Setenv("GOMODCACHE", "/custom/modcache")
	t.Setenv("GOPATH", "/some/gopath")
	if got := goModCache(); got != "/custom/modcache" {
		t.Errorf("goModCache = %q, want /custom/modcache", got)
	}
}

func TestGoModCacheGOPATH(t *testing.T) {
	t.Setenv("GOMODCACHE", "")
	t.Setenv("GOPATH", "/some/gopath")
	want := filepath.Join("/some/gopath", "pkg", "mod")
	if got := goModCache(); got != want {
		t.Errorf("goModCache = %q, want %q", got, want)
	}
}

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

func endsWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
