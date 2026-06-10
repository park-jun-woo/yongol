//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM33RedirectParams — 정적+params 모순·구문·필수 세그먼트 누락·없는 respField·생략형 매트릭스 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM33RedirectParams(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/contracts": postOpWithResp("CreateContract", map[string]*openapi3.SchemaRef{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		}),
	})
	opMap := buildOperationMethodMap(doc)
	pages := []stml.PageSpec{
		{Name: "contract-edit", FileName: "contract-edit.html", Route: "/contract-edit/:ContractID"},
		{Name: "unit-info", FileName: "unit-info.html", Route: "/unit-info/:BuildingID/:UnitID"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	action := func(redirect, params string) stml.ActionBlock {
		return stml.ActionBlock{OperationID: "CreateContract", Redirect: redirect, RedirectParamsRaw: params}
	}

	// Well-formed page-name redirect → silent.
	if d := tm33RedirectParams(action("contract-edit", "id -> ContractID"), "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("ok: expected 0 diagnostics, got %+v", d)
	}

	// Elided form against a single required segment → silent.
	if d := tm33RedirectParams(action("contract-edit", "id"), "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("elided ok: expected 0 diagnostics, got %+v", d)
	}

	// route.* source is exempt from the response-schema check.
	if d := tm33RedirectParams(action("contract-edit", "route.ContractID -> ContractID"), "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("route source: expected 0 diagnostics, got %+v", d)
	}

	// Target without segments and no params → silent.
	if d := tm33RedirectParams(action("dashboard", ""), "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("no segments: expected 0 diagnostics, got %+v", d)
	}

	// Static path + params declared → contradiction ERROR.
	if got := tm33RedirectParams(action("/", "id -> ContractID"), "f.html", opMap, pages); countDiag(got, "[TM-33]") != 1 {
		t.Errorf("static contradiction: expected 1 TM-33, got %+v", got)
	}

	// Static path without params → silent.
	if d := tm33RedirectParams(action("/", ""), "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("static plain: expected 0 diagnostics, got %+v", d)
	}

	// Syntax violation (item.* source) → 1 ERROR, reported alone.
	if got := tm33RedirectParams(action("contract-edit", "item.id -> ContractID"), "f.html", opMap, pages); countDiag(got, "[TM-33]") != 1 {
		t.Errorf("syntax: expected 1 TM-33, got %+v", got)
	}

	// respField not in the 2xx response schema → ERROR (TM-20 infrastructure).
	if got := tm33RedirectParams(action("contract-edit", "contract_id -> ContractID"), "f.html", opMap, pages); countDiag(got, "[TM-33]") != 1 {
		t.Errorf("missing respField: expected 1 TM-33, got %+v", got)
	}

	// SegmentName not in the target route → ERROR (plus the required
	// segment left unmapped).
	if got := tm33RedirectParams(action("contract-edit", "id -> ContractId"), "f.html", opMap, pages); countDiag(got, "[TM-33]") != 2 {
		t.Errorf("wrong segment: expected 2 TM-33, got %+v", got)
	}

	// Required segment unmet with no params attribute at all → ERROR.
	if got := tm33RedirectParams(action("contract-edit", ""), "f.html", opMap, pages); countDiag(got, "[TM-33]") != 1 {
		t.Errorf("no params: expected 1 TM-33, got %+v", got)
	}

	// Elided form against two required segments → ambiguity ERROR (plus
	// both required segments left unmapped).
	if got := tm33RedirectParams(action("unit-info", "id"), "f.html", opMap, pages); countDiag(got, "[TM-33]") != 3 {
		t.Errorf("ambiguous elision: expected 3 TM-33, got %+v", got)
	}

	// Unknown target page name → silent (TM-26 owns it).
	if d := tm33RedirectParams(action("contract-edt", "id -> ContractID"), "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("unknown target: expected 0 diagnostics, got %+v", d)
	}

	// Unknown operationId → respField check silently skipped (TM-02 owns it).
	unknown := stml.ActionBlock{OperationID: "Nope", Redirect: "contract-edit", RedirectParamsRaw: "id -> ContractID"}
	if d := tm33RedirectParams(unknown, "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("unknown op: expected 0 diagnostics, got %+v", d)
	}

	// No redirect → silent.
	if d := tm33RedirectParams(stml.ActionBlock{OperationID: "CreateContract"}, "f.html", opMap, pages); len(d) != 0 {
		t.Errorf("no redirect: expected 0 diagnostics, got %+v", d)
	}
}
