//ff:func feature=agent type=test control=sequence
//ff:what TestResolveFeature — SSaC 파일의 op→desc/path 해석, 미스 및 비-SSaC 레이어 처리 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestResolveFeature(t *testing.T) {
	lookup := map[string]features.Feature{
		"Login": {Op: "Login", Desc: "log in", Path: "/auth/login"},
	}

	desc, path := resolveFeature("service/auth/Login.ssac", layerSSaC, lookup)
	if desc != "log in" || path != "/auth/login" {
		t.Errorf("found: desc=%q path=%q, want log in / /auth/login", desc, path)
	}

	desc, path = resolveFeature("service/auth/Unknown.ssac", layerSSaC, lookup)
	if desc != "Unknown (no desc)" || path != "" {
		t.Errorf("miss: desc=%q path=%q, want 'Unknown (no desc)' / ''", desc, path)
	}

	desc, path = resolveFeature("db/users.sql", layerDDL, lookup)
	if desc != "" || path != "" {
		t.Errorf("non-ssac: desc=%q path=%q, want both empty", desc, path)
	}
}
