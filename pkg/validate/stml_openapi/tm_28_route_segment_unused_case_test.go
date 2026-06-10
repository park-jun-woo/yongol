//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what runTM28RouteSegmentUnused — TestTM28RouteSegmentUnused table-driven 개별 케이스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runTM28RouteSegmentUnused(t *testing.T, c TestTM28RouteSegmentUnusedCase) {
	t.Helper()
	diags := tm28RouteSegmentUnused(c.page)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[TM-28]") {
			t.Errorf("expected [TM-28], got %q", d.Message)
		}
		if d.File != c.page.FileName {
			t.Errorf("expected file %q, got %q", c.page.FileName, d.File)
		}
	}
}
