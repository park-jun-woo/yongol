//ff:func feature=validate type=test-helper control=sequence
//ff:what assertDiagCount — wantCount/wantSub 패턴 공용 assertion 헬퍼

package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func assertDiagCount(t *testing.T, diags []diagnostic.Diagnostic, wantCount int, wantSub string) {
	t.Helper()
	if len(diags) != wantCount {
		t.Fatalf("expected %d diagnostics, got %d: %+v", wantCount, len(diags), diags)
	}
	if wantSub != "" && len(diags) > 0 {
		if !strings.Contains(diags[0].Message, wantSub) {
			t.Errorf("message %q missing %q", diags[0].Message, wantSub)
		}
	}
}
