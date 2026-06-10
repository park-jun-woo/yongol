//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what runTM30ItemSource — TestTM30ItemSource table-driven 개별 케이스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func runTM30ItemSource(t *testing.T, c TestTM30ItemSourceCase) {
	t.Helper()
	page, parseDiags := stml.ParseReader("page.html", strings.NewReader(c.html))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}
	diags := tm30ItemSource(page, c.raif)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[TM-30]") {
			t.Errorf("expected [TM-30], got %q", d.Message)
		}
		if d.File != "page.html" {
			t.Errorf("expected file page.html, got %q", d.File)
		}
	}
}
