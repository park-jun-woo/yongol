//ff:func feature=validate type=test-helper control=sequence
//ff:what assertEmitByRule — emitByRule 결과 카운트·메시지·레벨·phase 검증 헬퍼

package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func assertEmitByRule(t *testing.T, diags []diagnostic.Diagnostic, wantCount int, wantSub string, lvl diagnostic.Level) {
	t.Helper()
	if len(diags) != wantCount {
		t.Fatalf("expected %d diagnostics, got %d: %+v", wantCount, len(diags), diags)
	}
	if wantSub != "" && len(diags) > 0 {
		if !strings.Contains(diags[0].Message, wantSub) {
			t.Errorf("message %q missing %q", diags[0].Message, wantSub)
		}
		if diags[0].Level != lvl {
			t.Errorf("Level = %v, want %v", diags[0].Level, lvl)
		}
		if diags[0].Phase != diagnostic.PhaseValidate {
			t.Errorf("Phase = %v, want PhaseValidate", diags[0].Phase)
		}
	}
}
