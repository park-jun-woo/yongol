//ff:func feature=agent type=test control=sequence
//ff:what TestChainResolve — SSOT 미검출/파싱 진단 시 빈 반환 + 예제 specs 로 link 매칭 성공 경로 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestChainResolveParseDiagnostics(t *testing.T) {
	// A specs dir containing malformed YAML produces parse diagnostics, so
	// chainResolve returns empty before attempting any chain resolution.
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "api"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata:\n  name: broken\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "api", "openapi.yaml"), []byte("invalid: [yaml: broken"), 0o644); err != nil {
		t.Fatal(err)
	}
	desc, path := chainResolve(dir, "api/openapi.yaml", map[string]features.Feature{
		"Login": {Op: "Login", Desc: "x", Path: "y"},
	})
	if desc != "" || path != "" {
		t.Errorf("chainResolve with parse diagnostics = %q, %q; want empty, empty", desc, path)
	}
}
