//ff:func feature=agent type=test control=sequence
//ff:what TestRunValidate — DetectSSOTs 에러 / 파싱 진단 / 정상 validate 진단 수집 분기 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

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
