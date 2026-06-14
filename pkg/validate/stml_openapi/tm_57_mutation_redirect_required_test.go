//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestTM57MutationRedirectRequired — mutation data-redirect 미선언 발화·선언/GET/capture/미존재 침묵

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM57MutationRedirectRequired(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/buildings", &openapi3.PathItem{
		Get:  &openapi3.Operation{OperationID: "ListBuildings"},
		Post: &openapi3.Operation{OperationID: "CreateBuilding"},
	})
	doc.Paths.Set("/buildings/{buildingId}", &openapi3.PathItem{
		Put:    &openapi3.Operation{OperationID: "UpdateBuilding"},
		Delete: &openapi3.Operation{OperationID: "DeleteBuilding"},
	})
	opMap := buildOperationMethodMap(doc)

	// Mutation (POST) without data-redirect → fires.
	fire := tm57MutationRedirectRequired(stml.ActionBlock{OperationID: "CreateBuilding"}, "p.html", opMap)
	if len(fire) != 1 {
		t.Fatalf("POST without redirect: expected 1 diagnostic, got %d: %+v", len(fire), fire)
	}

	// PUT/DELETE without data-redirect → fires.
	for _, op := range []string{"UpdateBuilding", "DeleteBuilding"} {
		d := tm57MutationRedirectRequired(stml.ActionBlock{OperationID: op}, "p.html", opMap)
		if len(d) != 1 {
			t.Errorf("%s without redirect: expected 1 diagnostic, got %d", op, len(d))
		}
	}

	// Mutation WITH data-redirect → silent.
	withRedirect := tm57MutationRedirectRequired(stml.ActionBlock{OperationID: "CreateBuilding", Redirect: "/buildings"}, "p.html", opMap)
	if len(withRedirect) != 0 {
		t.Errorf("POST with redirect: expected 0 diagnostics, got %d: %+v", len(withRedirect), withRedirect)
	}

	// GET data-action (non-mutating) → silent even without redirect.
	getAction := tm57MutationRedirectRequired(stml.ActionBlock{OperationID: "ListBuildings"}, "p.html", opMap)
	if len(getAction) != 0 {
		t.Errorf("GET action: expected 0 diagnostics, got %d: %+v", len(getAction), getAction)
	}

	// Bearer login capture action (data-capture) → exempt.
	capture := tm57MutationRedirectRequired(stml.ActionBlock{OperationID: "CreateBuilding", CaptureRaw: "access_token -> auth.token"}, "p.html", opMap)
	if len(capture) != 0 {
		t.Errorf("capture action: expected 0 diagnostics, got %d: %+v", len(capture), capture)
	}

	// Unknown operationId → silent (TM-02 reports it).
	unknown := tm57MutationRedirectRequired(stml.ActionBlock{OperationID: "NopeOp"}, "p.html", opMap)
	if len(unknown) != 0 {
		t.Errorf("unknown op: expected 0 diagnostics, got %d: %+v", len(unknown), unknown)
	}
}
