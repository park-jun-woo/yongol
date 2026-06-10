//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm32CheckParam — 생략형 해석/모호 진단·세그먼트 존재 검사·정상 매핑 반환 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32CheckParam(t *testing.T) {
	ref := linkRefCtx{Link: &stml.LinkRef{TargetPage: "detail"}, InEach: true, ItemFields: map[string]bool{"id": true}}
	pattern := "/d/:AID/:BID?"
	own := map[string]bool{}

	// Explicit mapping to an existing segment.
	seg, diags := tm32CheckParam(stml.LinkParamBind{Source: "item.id", Segment: "AID"}, ref, "f.html", pattern, []string{"AID"}, own)
	if seg != "AID" || len(diags) != 0 {
		t.Errorf("explicit: seg=%q diags=%+v", seg, diags)
	}

	// Elided form against a single required segment.
	seg, diags = tm32CheckParam(stml.LinkParamBind{Source: "item.id"}, ref, "f.html", pattern, []string{"AID"}, own)
	if seg != "AID" || len(diags) != 0 {
		t.Errorf("elided: seg=%q diags=%+v", seg, diags)
	}

	// Elided form against two required segments → ambiguity ERROR.
	seg, diags = tm32CheckParam(stml.LinkParamBind{Source: "item.id"}, ref, "f.html", "/d/:AID/:CID", []string{"AID", "CID"}, own)
	if seg != "" || len(diags) != 1 || !strings.Contains(diags[0].Message, "not exactly one") {
		t.Errorf("ambiguous: seg=%q diags=%+v", seg, diags)
	}

	// Unknown segment name → ERROR.
	seg, diags = tm32CheckParam(stml.LinkParamBind{Source: "item.id", Segment: "ZID"}, ref, "f.html", pattern, []string{"AID"}, own)
	if seg != "" || len(diags) != 1 || !strings.Contains(diags[0].Message, `"ZID"`) {
		t.Errorf("unknown segment: seg=%q diags=%+v", seg, diags)
	}

	// Mapping to an optional segment is legal.
	seg, diags = tm32CheckParam(stml.LinkParamBind{Source: "item.id", Segment: "BID"}, ref, "f.html", pattern, []string{"AID"}, own)
	if seg != "BID" || len(diags) != 0 {
		t.Errorf("optional mapped: seg=%q diags=%+v", seg, diags)
	}
}
