//ff:func feature=cli type=test control=sequence
//ff:what validate parse-fail — malformed openapi.yaml 이 parse failed 로 짧게 끝남

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegrationValidate_ParseFail writes a malformed api/openapi.yaml
// into a tmpdir and expects validate to surface `parse failed` with exit 1.
// Confirms parse errors short-circuit before the Validate step.
func TestIntegrationValidate_ParseFail(t *testing.T) {
	tmp := t.TempDir()
	apiDir := filepath.Join(tmp, "api")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatalf("mkdir api: %v", err)
	}
	// Intentionally malformed YAML — a mapping value on the root key.
	bad := []byte("invalid: ::: yaml: :\n")
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), bad, 0644); err != nil {
		t.Fatalf("write bad openapi.yaml: %v", err)
	}
	stdout, _, err := runCmd(t, "validate", tmp)
	if err == nil {
		t.Fatalf("expected parse error, got nil\nstdout:\n%s", stdout)
	}
	if isUsageError(err) {
		t.Fatalf("parse failure must be exit 1, not usage error: %v", err)
	}
	if !strings.Contains(err.Error(), "parse failed") {
		t.Errorf("expected error to contain `parse failed`, got: %v", err)
	}
	if !strings.Contains(stdout, "Parse Errors") {
		t.Errorf("expected stdout to print `Parse Errors` banner, got:\n%s", stdout)
	}
}
