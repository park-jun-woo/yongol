//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm32LinkParamsUnsatisfied — 구문·세그먼트·필수 충족·item/route 소스 발화/침묵 매트릭스 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM32LinkParamsUnsatisfied(t *testing.T) {
	detail := stml.PageSpec{
		Name:     "building-detail",
		FileName: "building-detail.html",
		Route:    "/buildings/:BuildingID/:PhotoID?",
	}
	twoRequired := stml.PageSpec{
		Name:     "unit-info",
		FileName: "unit-info.html",
		Route:    "/unit-info/:BuildingID/:UnitID",
	}
	settings := stml.PageSpec{Name: "settings", FileName: "settings.html"}
	raif := map[string]map[string]map[string]bool{
		"ListBuildings": {"buildings": {"id": true, "name": true}},
	}
	rowLink := func(params string) string {
		return `<main><section data-fetch="ListBuildings">
  <ul data-each="buildings">
    <li data-link="building-detail"` + params + `><span data-bind="name"></span></li>
  </ul>
</section></main>`
	}

	cases := []TestTM32LinkParamsCase{
		// Full mapping of the single required segment → silent
		// (the unmapped optional :PhotoID? is legal).
		{
			name:      "row_link_satisfied",
			html:      rowLink(` data-link-params="item.id -> BuildingID"`),
			targets:   []stml.PageSpec{detail},
			raif:      raif,
			wantCount: 0,
		},
		// Elided form against a single required segment → silent.
		{
			name:      "elided_single_required",
			html:      rowLink(` data-link-params="item.id"`),
			targets:   []stml.PageSpec{detail},
			raif:      raif,
			wantCount: 0,
		},
		// No params against a target without required segments → silent.
		{
			name:      "no_params_no_required",
			html:      `<main><a data-link="settings">설정</a></main>`,
			targets:   []stml.PageSpec{settings},
			wantCount: 0,
		},
		// Syntax violation → 1 ERROR (re-parse).
		{
			name:      "syntax_error",
			html:      rowLink(` data-link-params="id -> BuildingID"`),
			targets:   []stml.PageSpec{detail},
			raif:      raif,
			wantCount: 1,
			wantIn:    "data-link-params",
		},
		// Required segment left unmapped → ERROR with the resolved pattern.
		{
			name:      "required_unmapped",
			html:      rowLink(``),
			targets:   []stml.PageSpec{detail},
			raif:      raif,
			wantCount: 1,
			wantIn:    "/buildings/:BuildingID/:PhotoID?",
		},
		// SegmentName not in the target route → ERROR (+ required unmapped).
		{
			name:      "unknown_segment",
			html:      rowLink(` data-link-params="item.id -> RoomID"`),
			targets:   []stml.PageSpec{detail},
			raif:      raif,
			wantCount: 2,
			wantIn:    `segment "RoomID" is not in target page`,
		},
		// item.* source outside any data-each → ERROR.
		{
			name:      "item_outside_each",
			html:      `<main><a data-link="building-detail" data-link-params="item.id -> BuildingID">상세</a></main>`,
			targets:   []stml.PageSpec{detail},
			wantCount: 1,
			wantIn:    "only valid inside a data-each",
		},
		// item field not in the each item schema → ERROR.
		{
			name:      "item_field_missing",
			html:      rowLink(` data-link-params="item.nope -> BuildingID"`),
			targets:   []stml.PageSpec{detail},
			raif:      raif,
			wantCount: 1,
			wantIn:    "item schema",
		},
		// Unresolved item schema stays silent (TM-01/TM-07 territory).
		{
			name:      "unresolved_schema_silent",
			html:      rowLink(` data-link-params="item.nope -> BuildingID"`),
			targets:   []stml.PageSpec{detail},
			raif:      map[string]map[string]map[string]bool{},
			wantCount: 0,
		},
		// route.* source missing from this page's own resolved route → ERROR.
		{
			name:      "route_source_missing",
			html:      `<main><a data-link="building-detail" data-link-params="route.BuildingID -> BuildingID">상세</a></main>`,
			targets:   []stml.PageSpec{detail},
			wantCount: 1,
			wantIn:    "not a segment of this page's resolved route",
		},
		// route.* source present in this page's own route (data-route) → silent.
		{
			name:      "route_source_ok",
			html:      `<main data-route="/page/:BuildingID"><a data-link="building-detail" data-link-params="route.BuildingID -> BuildingID">상세</a></main>`,
			targets:   []stml.PageSpec{detail},
			wantCount: 0,
		},
		// Elided form against two required segments → ambiguity ERROR
		// (+ both required segments unmapped).
		{
			name: "elided_ambiguous",
			html: `<main><section data-fetch="ListBuildings">
  <ul data-each="buildings">
    <li data-link="unit-info" data-link-params="item.id"><span data-bind="name"></span></li>
  </ul>
</section></main>`,
			targets:   []stml.PageSpec{twoRequired},
			raif:      raif,
			wantCount: 3,
			wantIn:    "not exactly one",
		},
		// Missing target page stays silent — TM-31 owns it.
		{
			name:      "missing_target_silent",
			html:      `<main><a data-link="nope" data-link-params="item.id -> X">상세</a></main>`,
			targets:   []stml.PageSpec{detail},
			wantCount: 0,
		},
		// A page without any links short-circuits.
		{
			name:      "no_links",
			html:      `<main><h1>hello</h1></main>`,
			targets:   []stml.PageSpec{detail},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTM32LinkParams(t, c)
		})
	}
}
