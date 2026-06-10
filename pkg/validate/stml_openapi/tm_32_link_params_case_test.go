//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what runTM32LinkParams — TestTM32LinkParamsUnsatisfied table-driven 개별 케이스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func runTM32LinkParams(t *testing.T, c TestTM32LinkParamsCase) {
	t.Helper()
	page, parseDiags := stml.ParseReader("page.html", strings.NewReader(c.html))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}
	pages := append([]stml.PageSpec{page}, c.targets...)
	diags := tm32LinkParamsUnsatisfied(page, pages, c.raif)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	found := c.wantIn == ""
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[TM-32]") {
			t.Errorf("expected [TM-32], got %q", d.Message)
		}
		if d.File != "page.html" {
			t.Errorf("expected file page.html, got %q", d.File)
		}
		if c.wantIn != "" && strings.Contains(d.Message, c.wantIn) {
			found = true
		}
	}
	if !found {
		t.Errorf("no diagnostic contains %q: %+v", c.wantIn, diags)
	}
}
