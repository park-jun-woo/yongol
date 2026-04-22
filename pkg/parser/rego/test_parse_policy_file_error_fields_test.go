//ff:func feature=policy type=test control=sequence
//ff:what ParsePolicyFile — 존재하지 않는 파일에서 Diagnostic 필드 완결성 회귀

package rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParsePolicyFile_Missing_DiagnosticFields(t *testing.T) {
	p, diags := ParsePolicyFile("/nonexistent/policy.rego")
	if p != nil {
		t.Errorf("policy should be nil")
	}
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1", len(diags))
	}
	d := diags[0]
	if d.File != "/nonexistent/policy.rego" {
		t.Errorf("File = %q", d.File)
	}
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want LevelError", d.Level)
	}
	if d.Message == "" {
		t.Errorf("Message empty")
	}
}
