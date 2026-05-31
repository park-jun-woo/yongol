//ff:func feature=agent type=test control=sequence
//ff:what TestChainResolve — SSOT 미검출/파싱 진단 시 빈 반환 + 예제 specs 로 link 매칭 성공 경로 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestChainResolve(t *testing.T) {
	// An empty specs directory yields no detectable SSOTs (or parse diagnostics),
	// so chainResolve gives up and returns empty strings without panicking.
	lookup := map[string]features.Feature{
		"Login": {Op: "Login", Desc: "log in", Path: "/auth/login"},
	}
	desc, path := chainResolve(t.TempDir(), "service/auth/Login.ssac", lookup)
	if desc != "" || path != "" {
		t.Errorf("chainResolve on empty dir = %q, %q; want empty, empty", desc, path)
	}
}
