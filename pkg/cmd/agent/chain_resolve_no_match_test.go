//ff:func feature=agent type=test control=sequence
//ff:what TestChainResolve — SSOT 미검출/파싱 진단 시 빈 반환 + 예제 specs 로 link 매칭 성공 경로 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestChainResolveNoMatch(t *testing.T) {
	// Parse succeeds but no chained link matches the relPath, so chainResolve
	// exhausts the loop and returns empty strings (the final return).
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
	desc, path := chainResolve(specsDir, "no/such/file.txt", lookup)
	if desc != "" || path != "" {
		t.Errorf("chainResolve no-match = %q, %q; want empty, empty", desc, path)
	}
}
