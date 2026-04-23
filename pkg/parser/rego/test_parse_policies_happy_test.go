//ff:func feature=policy type=test control=sequence
//ff:what ParsePolicies — 정상 rego 파일 1개 → Policy 1개 + Rules 1개

package rego

import (
	"os"
	"path/filepath"
	"testing"
)

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
