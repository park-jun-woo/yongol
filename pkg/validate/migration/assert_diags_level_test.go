//ff:func feature=validate type=test-helper control=iteration dimension=1
//ff:what assertDiagsLevel — diags 레벨·ruleID 검증 헬퍼

package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func assertDiagsLevel(t *testing.T, diags []diagnostic.Diagnostic, wantCount int, wantLevel diagnostic.Level, ruleID string) {
	t.Helper()
	if len(diags) != wantCount {
		t.Fatalf("expected %d diagnostics, got %d: %+v", wantCount, len(diags), diags)
	}
	for _, d := range diags {
		if d.Level != wantLevel {
			t.Errorf("Level = %v, want %v", d.Level, wantLevel)
		}
		if !strings.Contains(d.Message, ruleID) {
			t.Errorf("Message missing %s: %s", ruleID, d.Message)
		}
	}
}
