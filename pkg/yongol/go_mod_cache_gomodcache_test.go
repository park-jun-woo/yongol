//ff:func feature=orchestrator type=test control=sequence
//ff:what TestGoModCache — GOMODCACHE→GOPATH→HOME fallback 우선순위 검증
package yongol

import (
	"testing"
)

func TestGoModCacheGOMODCACHE(t *testing.T) {
	t.Setenv("GOMODCACHE", "/custom/modcache")
	t.Setenv("GOPATH", "/some/gopath")
	if got := goModCache(); got != "/custom/modcache" {
		t.Errorf("goModCache = %q, want /custom/modcache", got)
	}
}
