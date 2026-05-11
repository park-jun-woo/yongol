//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what XMO-10 test — OpenAPI operationId가 STML에서 미소비되는 경우 WARNING 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_Positive(t *testing.T) {
	// STML references only ListUsers; GetUser is unconsumed.
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users":      getOp("ListUsers", nil, nil),
		"/users/{id}": getOp("GetUser", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[XMO-10]") {
		t.Errorf("expected XMO-10 diagnostic for GetUser, got %v", diags)
	}
	if countDiag(diags, "[XMO-10]") != 1 {
		t.Errorf("expected exactly 1 XMO-10, got %d", countDiag(diags, "[XMO-10]"))
	}
}

func TestXMO10_Unconsumed_Negative_AllConsumed(t *testing.T) {
	// STML references both ops — no XMO-10 expected.
	pages := []stml.PageSpec{{
		FileName: "users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
		Actions: []stml.ActionBlock{{
			OperationID: "CreateUser",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
		"/users/create": postOp("CreateUser", map[string]*openapi3.SchemaRef{
			"name": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 diagnostic, got %v", diags)
	}
}

func TestXMO10_Unconsumed_Negative_AuthEndpointExcluded(t *testing.T) {
	// Auth endpoint (security: []) is excluded from XMO-10.
	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetDashboard",
		}},
	}}
	// Login op has explicit empty security → auth endpoint → excluded
	loginOp := &openapi3.Operation{
		OperationID: "Login",
		Security:    &openapi3.SecurityRequirements{},
		Responses:   openapi3.NewResponses(),
	}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/dashboard":  getOp("GetDashboard", nil, nil),
		"/auth/login": {Post: loginOp},
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 for auth endpoint Login, got %v", diags)
	}
}

func TestXMO10_Unconsumed_NestedFetchConsumed(t *testing.T) {
	// Nested fetch's operationId should be counted as consumed.
	pages := []stml.PageSpec{{
		FileName: "detail.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetWorkflow",
			NestedFetches: []stml.FetchBlock{{
				OperationID: "ListWorkflowVersions",
			}},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/workflows/{id}":          getOp("GetWorkflow", nil, nil),
		"/workflows/{id}/versions": getOp("ListWorkflowVersions", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 for nested fetch, got %v", diags)
	}
}

func TestXMO10_Unconsumed_MultipleUnconsumed(t *testing.T) {
	// Multiple unconsumed ops → multiple XMO-10 diagnostics.
	pages := []stml.PageSpec{{
		FileName: "home.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetDashboard",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/dashboard": getOp("GetDashboard", nil, nil),
		"/users":     getOp("ListUsers", nil, nil),
		"/actions":   getOp("ListActions", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if countDiag(diags, "[XMO-10]") != 2 {
		t.Errorf("expected 2 XMO-10 diagnostics, got %d", countDiag(diags, "[XMO-10]"))
	}
}

func TestXMO10_Unconsumed_ActionConsumed(t *testing.T) {
	// data-action also counts as consumption.
	pages := []stml.PageSpec{{
		FileName: "form.html",
		Actions: []stml.ActionBlock{{
			OperationID: "CreateWorkflow",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/workflows": postOp("CreateWorkflow", map[string]*openapi3.SchemaRef{
			"name": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 for consumed action, got %v", diags)
	}
}
