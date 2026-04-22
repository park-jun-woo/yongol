//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateOpenAPI — path/method/operationId/security 기본 등록 검증

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

// TestPopulateOpenAPI_NilDoc ensures the function short-circuits safely.
func TestPopulateOpenAPI_NilDoc(t *testing.T) {
	g := newGround()
	populateOpenAPI(g, newMinimalFullstack())

	if len(g.Lookup) != 0 {
		t.Errorf("expected empty Lookup when doc is nil, got %d keys", len(g.Lookup))
	}
}

// TestPopulateOpenAPI_SecuritySchemes verifies components.securitySchemes are
// registered.
func TestPopulateOpenAPI_SecuritySchemes(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(),
		Components: &openapi3.Components{
			SecuritySchemes: openapi3.SecuritySchemes{
				"bearerAuth": &openapi3.SecuritySchemeRef{
					Value: &openapi3.SecurityScheme{Type: "http", Scheme: "bearer"},
				},
			},
		},
	}
	fs := newMinimalFullstack(withOpenAPIDoc(doc))
	g := newGround()

	populateOpenAPI(g, fs)

	if !g.Lookup["OpenAPI.security"]["bearerAuth"] {
		t.Errorf("OpenAPI.security missing bearerAuth: %v", g.Lookup["OpenAPI.security"])
	}
}
