//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what runTM52Case — TestTM52DuplicateFieldAcrossForms table-driven 개별 케이스 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func runTM52Case(t *testing.T, c tm52Case) {
	t.Helper()
	page, parseDiags := stml.ParseReader("page.html", strings.NewReader(c.html))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}
	diags := tm52DuplicateFieldAcrossForms(page)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d: %+v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[TM-52]") {
			t.Errorf("expected [TM-52], got %q", d.Message)
		}
		if d.File != "page.html" {
			t.Errorf("expected file page.html, got %q", d.File)
		}
	}
	if c.wantField == "" {
		return
	}
	d := diags[0]
	if !strings.Contains(d.Message, c.wantField) {
		t.Errorf("expected field %q in message, got %q", c.wantField, d.Message)
	}
	if !strings.Contains(d.Message, "UpdateBuilding") || !strings.Contains(d.Message, "CreateBuilding") {
		t.Errorf("expected both operationIds in message, got %q", d.Message)
	}
}
