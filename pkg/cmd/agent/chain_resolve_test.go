//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestChainResolve — SSOT 미검출/파싱 진단 시 빈 반환 + 예제 specs 로 link 매칭 성공 경로 검증

package agent

import (
	"os"
	"path/filepath"
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

func TestChainResolveDetectError(t *testing.T) {
	// Pointing specsDir at a regular file (not a directory) makes DetectSSOTs
	// fail, exercising the early "" "" return at the top of chainResolve.
	file := filepath.Join(t.TempDir(), "not-a-dir.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	desc, path := chainResolve(file, "whatever", nil)
	if desc != "" || path != "" {
		t.Errorf("chainResolve on file = %q, %q; want empty, empty", desc, path)
	}
}

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
