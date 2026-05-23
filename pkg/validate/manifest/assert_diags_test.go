//ff:func feature=validate type=test-helper control=sequence
//ff:what assertDiags — wantDiags/wantSub 패턴 공용 assertion 헬퍼

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func assertDiags(t *testing.T, diags []diagnostic.Diagnostic, wantDiags bool, wantSub string) {
	t.Helper()
	if wantDiags && len(diags) == 0 {
		t.Fatal("expected diagnostics but got none")
	}
	if !wantDiags && len(diags) != 0 {
		t.Fatalf("expected no diagnostics, got %d: %+v", len(diags), diags)
	}
	if wantDiags && wantSub != "" {
		if !strings.Contains(diags[0].Message, wantSub) {
			t.Errorf("diagnostic message %q missing substring %q", diags[0].Message, wantSub)
		}
	}
}
