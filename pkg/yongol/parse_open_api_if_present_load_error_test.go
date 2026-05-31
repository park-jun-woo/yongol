//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseOpenAPIIfPresent — 미탐지 / 로드 에러 / 정상 로드 분기 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseOpenAPIIfPresent_LoadError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "openapi.yaml")
	// Invalid OpenAPI content → loader returns an error.
	if err := os.WriteFile(path, []byte(":\n  not: [valid"), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindOpenAPI: {Kind: KindOpenAPI, Path: path, Presence: SSOTPopulated},
	}
	parseOpenAPIIfPresent(fs, has)
	if fs.OpenAPIDoc != nil {
		t.Fatalf("expected nil doc on load error")
	}
	if len(fs.ParseDiagnostics) == 0 {
		t.Fatalf("expected a load-error diagnostic")
	}
}
