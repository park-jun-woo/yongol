//ff:func feature=orchestrator type=test control=sequence
//ff:what TestGoModCache — GOMODCACHE→GOPATH→HOME fallback 우선순위 검증
package yongol

import (
	"path/filepath"
	"testing"
)

func TestGoModCacheGOPATH(t *testing.T) {
	t.Setenv("GOMODCACHE", "")
	t.Setenv("GOPATH", "/some/gopath")
	want := filepath.Join("/some/gopath", "pkg", "mod")
	if got := goModCache(); got != want {
		t.Errorf("goModCache = %q, want %q", got, want)
	}
}
