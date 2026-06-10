//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm32CheckParamSource — item.*(each 밖/스키마 부재/미해석 침묵)·route.*(자기 라우트) 소스 검사 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32CheckParamSource(t *testing.T) {
	own := map[string]bool{"BuildingID": true}

	// item.* outside data-each → ERROR.
	got := tm32CheckParamSource(stml.LinkParamBind{Source: "item.id"}, linkRefCtx{Link: &stml.LinkRef{}}, "f.html", own)
	if len(got) != 1 || !strings.Contains(got[0].Message, "only valid inside a data-each") {
		t.Errorf("outside each: %+v", got)
	}

	// item field missing from the item schema → ERROR.
	inEach := linkRefCtx{Link: &stml.LinkRef{}, InEach: true, ItemFields: map[string]bool{"id": true}}
	got = tm32CheckParamSource(stml.LinkParamBind{Source: "item.nope"}, inEach, "f.html", own)
	if len(got) != 1 || !strings.Contains(got[0].Message, "item schema") {
		t.Errorf("missing field: %+v", got)
	}

	// Valid item field → silent; unresolved schema → silent.
	if got = tm32CheckParamSource(stml.LinkParamBind{Source: "item.id"}, inEach, "f.html", own); len(got) != 0 {
		t.Errorf("valid item: %+v", got)
	}
	unresolved := linkRefCtx{Link: &stml.LinkRef{}, InEach: true}
	if got = tm32CheckParamSource(stml.LinkParamBind{Source: "item.id"}, unresolved, "f.html", own); len(got) != 0 {
		t.Errorf("unresolved schema: %+v", got)
	}

	// route.* present in the own route → silent; absent → ERROR.
	ctx := linkRefCtx{Link: &stml.LinkRef{}}
	if got = tm32CheckParamSource(stml.LinkParamBind{Source: "route.BuildingID"}, ctx, "f.html", own); len(got) != 0 {
		t.Errorf("valid route: %+v", got)
	}
	got = tm32CheckParamSource(stml.LinkParamBind{Source: "route.UnitID"}, ctx, "f.html", own)
	if len(got) != 1 || !strings.Contains(got[0].Message, "resolved route") {
		t.Errorf("missing route: %+v", got)
	}
}
