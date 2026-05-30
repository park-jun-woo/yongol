//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseOpenAPIIfPresent — 미탐지 / 로드 에러 / 정상 로드 분기 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseOpenAPIIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseOpenAPIIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.OpenAPIDoc != nil {
		t.Fatalf("expected no OpenAPIDoc when absent")
	}
}

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

func TestParseOpenAPIIfPresent_Valid(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "openapi.yaml")
	doc := `openapi: 3.0.3
info:
  title: t
  version: "0"
paths:
  /login:
    post:
      operationId: Login
      responses:
        '200':
          description: OK
`
	if err := os.WriteFile(path, []byte(doc), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindOpenAPI: {Kind: KindOpenAPI, Path: path, Presence: SSOTPopulated},
	}
	parseOpenAPIIfPresent(fs, has)
	if fs.OpenAPIDoc == nil {
		t.Fatalf("expected OpenAPIDoc to be set, diags=%+v", fs.ParseDiagnostics)
	}
	if fs.OpenAPILines == nil {
		t.Fatalf("expected OpenAPILines to be set")
	}
}
