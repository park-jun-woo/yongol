//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm54FieldCoverage — 응답에 없는 폼 필드만 WARNING, 빈 이름·커버된 필드 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM54FieldCoverage(t *testing.T) {
	a := stml.ActionBlock{
		OperationID: "UpdateRule",
		Prefill:     "GetRule",
		Fields:      []stml.FieldBind{{Name: "sheet_name"}, {Name: "note"}, {Name: ""}},
	}
	resp := map[string]responseFieldInfo{"sheet_name": {}}
	diags := tm54FieldCoverage(a, "p.html", resp)
	if len(diags) != 1 || diags[0].Level != diagnostic.LevelWarning {
		t.Fatalf("expected 1 WARNING for 'note', got %+v", diags)
	}

	// All covered → silent.
	covered := map[string]responseFieldInfo{"sheet_name": {}, "note": {}}
	if d := tm54FieldCoverage(a, "p.html", covered); len(d) != 0 {
		t.Errorf("all covered should be silent, got %+v", d)
	}
}
