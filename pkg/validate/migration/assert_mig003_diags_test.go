//ff:func feature=validate type=test-helper control=iteration dimension=1
//ff:what assertMig003Diags — MIG-003 진단 카운트·레벨·경로 검증 헬퍼

package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func assertMig003Diags(t *testing.T, diags []diagnostic.Diagnostic, wantCount int, missing []string) {
	t.Helper()
	if len(diags) != wantCount {
		t.Fatalf("expected %d diagnostics, got %d: %+v", wantCount, len(diags), diags)
	}
	for i, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("[%d] Level = %v, want LevelError", i, d.Level)
		}
		if !strings.Contains(d.Message, "MIG-003") {
			t.Errorf("[%d] Message missing MIG-003: %s", i, d.Message)
		}
		if !strings.Contains(d.Message, missing[i]) {
			t.Errorf("[%d] Message missing path %q: %s", i, missing[i], d.Message)
		}
	}
}
