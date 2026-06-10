//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM33CheckParam — 세그먼트 해석(명시·생략형·모호·미존재)과 소스 진단 병합 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM33CheckParam(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/contracts": postOpWithResp("CreateContract", map[string]*openapi3.SchemaRef{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		}),
	})
	opMap := buildOperationMethodMap(doc)
	a := stml.ActionBlock{OperationID: "CreateContract", Redirect: "contract-edit"}

	// Explicit segment resolves.
	seg, d := tm33CheckParam(stml.LinkParamBind{Source: "id", Segment: "ContractID"}, a, "f.html", "/contract-edit/:ContractID", []string{"ContractID"}, opMap)
	if seg != "ContractID" || len(d) != 0 {
		t.Errorf("explicit: seg=%q diags=%+v", seg, d)
	}

	// Elided form resolves against the single required segment.
	seg, d = tm33CheckParam(stml.LinkParamBind{Source: "id"}, a, "f.html", "/contract-edit/:ContractID", []string{"ContractID"}, opMap)
	if seg != "ContractID" || len(d) != 0 {
		t.Errorf("elided: seg=%q diags=%+v", seg, d)
	}

	// Ambiguous elision → ERROR, unresolved.
	seg, d = tm33CheckParam(stml.LinkParamBind{Source: "id"}, a, "f.html", "/unit-info/:BuildingID/:UnitID", []string{"BuildingID", "UnitID"}, opMap)
	if seg != "" || countDiag(d, "[TM-33]") != 1 {
		t.Errorf("ambiguous: seg=%q diags=%+v", seg, d)
	}

	// Segment not in the target route → ERROR, unresolved.
	seg, d = tm33CheckParam(stml.LinkParamBind{Source: "id", Segment: "ContractId"}, a, "f.html", "/contract-edit/:ContractID", []string{"ContractID"}, opMap)
	if seg != "" || countDiag(d, "[TM-33]") != 1 {
		t.Errorf("wrong segment: seg=%q diags=%+v", seg, d)
	}

	// Source diagnostics merge with the segment resolution: an unknown
	// respField still resolves its segment (one diag, mapped).
	seg, d = tm33CheckParam(stml.LinkParamBind{Source: "nope", Segment: "ContractID"}, a, "f.html", "/contract-edit/:ContractID", []string{"ContractID"}, opMap)
	if seg != "ContractID" || countDiag(d, "[TM-33]") != 1 {
		t.Errorf("bad source: seg=%q diags=%+v", seg, d)
	}
}
