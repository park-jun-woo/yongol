//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM33CheckParamSource — respField 2xx 스키마 검사·route.* 면제·미지 op 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM33CheckParamSource(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/contracts": postOpWithResp("CreateContract", map[string]*openapi3.SchemaRef{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		}),
		// 201 (Created) create op exposing the new resource id — the canonical
		// "redirect to detail by id" flow (BUG-128 / Phase039).
		"/contracts201": postOpStatusResp("CreateContract201", 201, map[string]*openapi3.SchemaRef{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		}),
		// 204 (No Content) op — no body to source a redirect param from.
		"/ping204": postOpStatusResp("Ping204", 204, nil),
	})
	opMap := buildOperationMethodMap(doc)
	a := stml.ActionBlock{OperationID: "CreateContract", Redirect: "contract-edit"}

	// respField in the 2xx schema → silent.
	if d := tm33CheckParamSource(stml.LinkParamBind{Source: "id", Segment: "ContractID"}, a, "f.html", opMap); len(d) != 0 {
		t.Errorf("known field: expected 0 diagnostics, got %+v", d)
	}

	// respField missing from the 2xx schema → 1 ERROR.
	if d := tm33CheckParamSource(stml.LinkParamBind{Source: "contract_id", Segment: "ContractID"}, a, "f.html", opMap); countDiag(d, "[TM-33]") != 1 {
		t.Errorf("missing field: expected 1 TM-33, got %+v", d)
	}

	// route.* source is exempt from the schema check.
	if d := tm33CheckParamSource(stml.LinkParamBind{Source: "route.BuildingID", Segment: "BuildingID"}, a, "f.html", opMap); len(d) != 0 {
		t.Errorf("route source: expected 0 diagnostics, got %+v", d)
	}

	// Unknown operationId → silent (TM-02 reports it).
	unknown := stml.ActionBlock{OperationID: "Nope", Redirect: "contract-edit"}
	if d := tm33CheckParamSource(stml.LinkParamBind{Source: "id", Segment: "ContractID"}, unknown, "f.html", opMap); len(d) != 0 {
		t.Errorf("unknown op: expected 0 diagnostics, got %+v", d)
	}

	// 201 success body field → silent (Phase039).
	a201 := stml.ActionBlock{OperationID: "CreateContract201", Redirect: "contract-edit"}
	if d := tm33CheckParamSource(stml.LinkParamBind{Source: "id", Segment: "ContractID"}, a201, "f.html", opMap); len(d) != 0 {
		t.Errorf("201 field source: expected 0 diagnostics, got %+v", d)
	}

	// 204 no-body op → param source unsatisfiable → 1 ERROR.
	a204 := stml.ActionBlock{OperationID: "Ping204", Redirect: "contract-edit"}
	if d := tm33CheckParamSource(stml.LinkParamBind{Source: "id", Segment: "ContractID"}, a204, "f.html", opMap); countDiag(d, "[TM-33]") != 1 {
		t.Errorf("204 field source: expected 1 TM-33, got %+v", d)
	}
}
