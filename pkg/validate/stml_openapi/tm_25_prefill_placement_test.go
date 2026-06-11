//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM25PrefillPlacement — misplaced data-prefill 레코드가 TM-25 ERROR로 발화하는지 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM25PrefillPlacement(t *testing.T) {
	page := stml.PageSpec{
		FileName: "p.html",
		FlowAttrMisplaced: []stml.FlowAttrMisplaced{
			{Attr: "data-prefill", Tag: "div"},
		},
	}
	got := tm25FlowAttrPlacement(page)
	if countDiag(got, "[TM-25]") != 1 {
		t.Fatalf("expected 1 TM-25 diagnostic, got %+v", got)
	}
	if got[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR level, got %+v", got[0])
	}
	if !strings.Contains(got[0].Message, "data-prefill") || !strings.Contains(got[0].Message, "requires data-action") {
		t.Errorf("message = %q", got[0].Message)
	}
}
