//ff:func feature=policy type=test control=sequence
//ff:what ParseDir / ParsePolicies — 디렉토리 열기 실패 시 Diagnostic 필드 완결성 회귀 + happy path

package rego

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseDir_Missing_DiagnosticFields(t *testing.T) {
	_, diags := ParseDir("/nonexistent/rego/dir")
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1", len(diags))
	}
	d := diags[0]
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want LevelError", d.Level)
	}
	if d.File != "/nonexistent/rego/dir" {
		t.Errorf("File = %q", d.File)
	}
}

func TestParsePolicies_Happy(t *testing.T) {
	dir := t.TempDir()
	content := `package authz

default allow := false

allow if {
    input.action == "Create"
    input.resource == "workflow"
}
`
	if err := os.WriteFile(filepath.Join(dir, "authz.rego"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	policies, diags := ParsePolicies(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(policies) != 1 {
		t.Fatalf("policies count = %d, want 1", len(policies))
	}
	if len(policies[0].Rules) != 1 {
		t.Errorf("Rules count = %d, want 1", len(policies[0].Rules))
	}
}

func TestParseDir_Happy(t *testing.T) {
	dir := t.TempDir()
	content := "package authz\n\ndefault allow := false\n"
	if err := os.WriteFile(filepath.Join(dir, "a.rego"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	mods, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(mods) != 1 {
		t.Fatalf("modules count = %d, want 1", len(mods))
	}
}
