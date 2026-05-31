//ff:func feature=agent type=test control=sequence
//ff:what TestChainResolve — SSOT 미검출/파싱 진단 시 빈 반환 + 예제 specs 로 link 매칭 성공 경로 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestChainResolveMatch(t *testing.T) {
	// Use the zenflow example specs (read-only fixture). The "Login" operation
	// chains to api/openapi.yaml, so a lookup keyed on "Login" with relPath
	// "api/openapi.yaml" hits the link-match branch and returns the feature's
	// Desc/Path.
	specsDir, err := filepath.Abs(filepath.Join("..", "..", "..", "examples", "zenflow", "opus4_7", "specs"))
	if err != nil {
		t.Fatalf("abs: %v", err)
	}
	if _, statErr := os.Stat(filepath.Join(specsDir, "manifest.yaml")); statErr != nil {
		t.Skipf("zenflow example specs unavailable: %v", statErr)
	}

	lookup := map[string]features.Feature{
		"Login": {Op: "Login", Desc: "user login", Path: "features/auth.yaml"},
	}
	desc, path := chainResolve(specsDir, "api/openapi.yaml", lookup)
	if desc != "user login" || path != "features/auth.yaml" {
		t.Errorf("chainResolve match = %q, %q; want %q, %q", desc, path, "user login", "features/auth.yaml")
	}
}
