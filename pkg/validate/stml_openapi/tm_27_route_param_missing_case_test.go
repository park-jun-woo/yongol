//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what runTM27RouteParamMissing — TestTM27RouteParamMissing table-driven 개별 케이스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runTM27RouteParamMissing(t *testing.T, c TestTM27RouteParamMissingCase) {
	t.Helper()
	diags := tm27RouteParamMissing(c.page)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[TM-27]") {
			t.Errorf("expected [TM-27], got %q", d.Message)
		}
		if d.File != c.page.FileName {
			t.Errorf("expected file %q, got %q", c.page.FileName, d.File)
		}
	}
}
