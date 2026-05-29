//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPI — components.securitySchemes 등록 회귀

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

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
