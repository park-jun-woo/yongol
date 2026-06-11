//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm54PrefillSourceForAction — 소스 부재 ERROR / 미지 op·void 응답·prefill 없음 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM54PrefillSourceForAction(t *testing.T) {
	opMap := buildOperationMethodMap(tm54Doc())
	pageFetchOps := map[string]bool{"GetRule": true}

	// No data-prefill → silent.
	if d := tm54PrefillSourceForAction(stml.ActionBlock{OperationID: "UpdateRule"}, "p.html", pageFetchOps, opMap); d != nil {
		t.Errorf("no prefill should be silent, got %+v", d)
	}

	// Source not a same-page fetch → ERROR.
	bad := stml.ActionBlock{OperationID: "UpdateRule", Prefill: "NoSuch"}
	d := tm54PrefillSourceForAction(bad, "p.html", pageFetchOps, opMap)
	if len(d) != 1 || d[0].Level != diagnostic.LevelError {
		t.Fatalf("expected ERROR, got %+v", d)
	}

	// Source is a page fetch but unknown to OpenAPI → silent (TM-01 owns it).
	unknownOp := stml.ActionBlock{OperationID: "UpdateRule", Prefill: "GhostFetch"}
	if d := tm54PrefillSourceForAction(unknownOp, "p.html", map[string]bool{"GhostFetch": true}, opMap); d != nil {
		t.Errorf("unknown op should be silent, got %+v", d)
	}
}
