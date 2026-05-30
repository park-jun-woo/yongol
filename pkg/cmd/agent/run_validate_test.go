//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestRunValidate — DetectSSOTs 에러 / 파싱 진단 / 정상 validate 진단 수집 분기 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunValidateDetectError(t *testing.T) {
	// A file path (not a directory) makes DetectSSOTs fail.
	file := filepath.Join(t.TempDir(), "f.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := runValidate(file); err == nil {
		t.Fatal("expected detect error for non-directory path")
	}
}

func TestRunValidateParseDiagnostics(t *testing.T) {
	// Malformed openapi.yaml produces parse diagnostics, returned directly.
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "api"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata:\n  name: x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "api", "openapi.yaml"), []byte("invalid: [yaml: broken"), 0o644); err != nil {
		t.Fatal(err)
	}
	diags, err := runValidate(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(diags) == 0 {
		t.Error("expected parse diagnostics for malformed openapi.yaml")
	}
}

func TestRunValidateFullValidate(t *testing.T) {
	// An empty (but valid) specs directory parses without diagnostics and runs
	// the full validate pass, exercising the report-aggregation loop.
	dir := t.TempDir()
	if _, err := runValidate(dir); err != nil {
		t.Fatalf("unexpected error on empty specs: %v", err)
	}
}
