//ff:func feature=validate type=test-helper control=sequence
//ff:what runFinalizeCase — TestFinalize 개별 케이스 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runFinalizeCase(t *testing.T, stepName string, diags []diagnostic.Diagnostic, wantStatus Status) {
	t.Helper()
	got := finalize(stepName, diags)
	if got.Status != wantStatus {
		t.Errorf("finalize(%q, ...).Status = %v, want %v", stepName, got.Status, wantStatus)
	}
	if got.Name != stepName {
		t.Errorf("finalize(%q, ...).Name = %q, want %q", stepName, got.Name, stepName)
	}
	if len(got.Diagnostics) != len(diags) {
		t.Errorf("finalize(%q, ...).Diagnostics len = %d, want %d", stepName, len(got.Diagnostics), len(diags))
	}
}
