//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what runTM29ActionOnErrorMissing — TestTM29ActionOnErrorMissing table-driven 개별 케이스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runTM29ActionOnErrorMissing(t *testing.T, c TestTM29ActionOnErrorMissingCase, opMap map[string]operationEntry) {
	t.Helper()
	diags := tm29ActionOnErrorMissing(c.action, "page.html", opMap)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[TM-29]") {
			t.Errorf("expected [TM-29], got %q", d.Message)
		}
		if d.File != "page.html" {
			t.Errorf("expected file %q, got %q", "page.html", d.File)
		}
		if d.OperationID != c.action.OperationID {
			t.Errorf("expected operationId %q, got %q", c.action.OperationID, d.OperationID)
		}
	}
}
