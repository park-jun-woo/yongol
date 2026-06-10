//ff:func feature=gen-react type=test control=sequence
//ff:what opSecurityIndex — operationId → 보호 여부 인덱스 / nil doc / 빈 opID 분기 검증

package react

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOpSecurityIndex(t *testing.T) {
	sec := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearerAuth": {}}}
	optOut := openapi3.SecurityRequirements{}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/login", &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: "Login", Security: &optOut},
		}),
		openapi3.WithPath("/workflows", &openapi3.PathItem{
			Get:  &openapi3.Operation{OperationID: "ListWorkflows", Security: &sec},
			Post: &openapi3.Operation{OperationID: "", Security: &sec}, // empty opID skipped
		}),
		openapi3.WithPath("/public", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "GetPublic"},
		}),
	)}

	idx := opSecurityIndex(doc)

	if got, ok := idx["ListWorkflows"]; !ok || !got {
		t.Errorf("ListWorkflows should be protected, got (%v, %v)", got, ok)
	}
	if got, ok := idx["Login"]; !ok || got {
		t.Errorf("Login should be present and unprotected (opt-out), got (%v, %v)", got, ok)
	}
	if got, ok := idx["GetPublic"]; !ok || got {
		t.Errorf("GetPublic should be present and unprotected, got (%v, %v)", got, ok)
	}
	if _, ok := idx[""]; ok {
		t.Errorf("empty operationId must be skipped")
	}
	if len(idx) != 3 {
		t.Errorf("expected 3 indexed ops, got %d: %+v", len(idx), idx)
	}

	if got := opSecurityIndex(nil); len(got) != 0 {
		t.Errorf("nil doc should yield empty index, got %+v", got)
	}
	if got := opSecurityIndex(&openapi3.T{}); len(got) != 0 {
		t.Errorf("nil Paths should yield empty index, got %+v", got)
	}
}
