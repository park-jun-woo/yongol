//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm30CheckActionParams — each 외부/스키마 부재/미해석 침묵/비-item 스킵 분기 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM30CheckActionParams(t *testing.T) {
	a := stml.ActionBlock{
		OperationID: "DeletePhoto",
		Params: []stml.ParamBind{
			{Name: "buildingId", Source: "route.BuildingID"},
			{Name: "photoId", Source: "item.id"},
		},
	}
	fields := map[string]bool{"id": true}

	// inside each with the field present → silent (route.* skipped too)
	if got := tm30CheckActionParams(a, "p.html", fields, true); len(got) != 0 {
		t.Errorf("valid: got %v", got)
	}
	// inside each, field missing → 1 error
	missing := stml.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stml.ParamBind{{Name: "photoId", Source: "item.nope"}},
	}
	got := tm30CheckActionParams(missing, "p.html", fields, true)
	if len(got) != 1 || !strings.Contains(got[0].Message, "[TM-30]") {
		t.Fatalf("missing field: got %v", got)
	}
	// unresolved schema → silent
	if got := tm30CheckActionParams(missing, "p.html", nil, true); len(got) != 0 {
		t.Errorf("nil schema must stay silent: %v", got)
	}
	// outside each → error regardless of schema
	got = tm30CheckActionParams(a, "p.html", nil, false)
	if len(got) != 1 || !strings.Contains(got[0].Message, "only valid inside a data-each") {
		t.Fatalf("outside each: got %v", got)
	}
}
