//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPI — path/method/operationId 기본 등록 검증

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateOpenAPI_PathsAndOpIDs verifies path, method, and operationId sets
// are populated correctly.
func TestPopulateOpenAPI_PathsAndOpIDs(t *testing.T) {
	paths := openapi3.NewPaths(
		openapi3.WithPath("/workflows", &openapi3.PathItem{
			Get:  &openapi3.Operation{OperationID: "ListWorkflows"},
			Post: &openapi3.Operation{OperationID: "CreateWorkflow"},
		}),
		openapi3.WithPath("/workflows/{id}", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "GetWorkflow"},
		}),
	)
	fs := newMinimalFullstack(withOpenAPIDoc(&openapi3.T{Paths: paths}))
	g := newGround()

	populateOpenAPI(g, fs)

	opIDs := g.Lookup["OpenAPI.operationId"]
	if !opIDs["ListWorkflows"] || !opIDs["CreateWorkflow"] || !opIDs["GetWorkflow"] {
		t.Fatalf("operationId set incomplete: %v", opIDs)
	}
	p := g.Lookup["OpenAPI.path"]
	if !p["/workflows"] || !p["/workflows/{id}"] {
		t.Fatalf("path set incomplete: %v", p)
	}
	m1 := g.Lookup["OpenAPI.method./workflows"]
	if !m1["GET"] || !m1["POST"] {
		t.Fatalf("methods for /workflows = %v, want GET+POST", m1)
	}
}
