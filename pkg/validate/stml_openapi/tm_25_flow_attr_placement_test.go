//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestTM25FlowAttrPlacement — 파서가 기록한 흐름 속성 위치 위반이 ERROR로 발화하는지 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM25FlowAttrPlacement(t *testing.T) {
	page := stml.PageSpec{
		FileName: "p.html",
		FlowAttrMisplaced: []stml.FlowAttrMisplaced{
			{Attr: "data-capture", Tag: "div"},
			{Attr: "data-on-error", Tag: "p"},
		},
	}
	got := tm25FlowAttrPlacement(page)
	if countDiag(got, "[TM-25]") != 2 {
		t.Fatalf("expected 2 TM-25 diagnostics, got %+v", got)
	}
	for _, d := range got {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected ERROR level, got %+v", d)
		}
	}
	if !strings.Contains(got[0].Message, "data-capture") || !strings.Contains(got[1].Message, "outside any data-action") {
		t.Errorf("messages: %q / %q", got[0].Message, got[1].Message)
	}

	// No misplacements → no diagnostics.
	if d := tm25FlowAttrPlacement(stml.PageSpec{FileName: "p.html"}); len(d) != 0 {
		t.Errorf("clean page: expected 0 diagnostics, got %+v", d)
	}
}
