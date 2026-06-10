//ff:func feature=gen-react type=test control=sequence
//ff:what resolveProtectedPages — security 보호 op 소비 페이지만 보호 판정 검증

package react

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveProtectedPages(t *testing.T) {
	sec := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearerAuth": {}}}
	optOut := openapi3.SecurityRequirements{}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/login", &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: "Login", Security: &optOut},
		}),
		openapi3.WithPath("/workflows", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "ListWorkflows", Security: &sec},
		}),
	)}
	fs := &yongol.Fullstack{
		Manifest:   &manifest.ProjectConfig{Backend: manifest.Backend{Auth: &manifest.Auth{Mode: "bearer"}}},
		OpenAPIDoc: doc,
		STMLPages: []stml.PageSpec{
			{
				FileName: "login.html",
				Actions:  []stml.ActionBlock{{OperationID: "Login"}},
			},
			{
				FileName: "workflows.html",
				Fetches:  []stml.FetchBlock{{OperationID: "ListWorkflows"}},
			},
			{FileName: "about.html"},
		},
	}

	got := resolveProtectedPages(fs)
	if got["login.html"] {
		t.Errorf("login.html should be public (op security opts out)")
	}
	if !got["workflows.html"] {
		t.Errorf("workflows.html should be protected (consumes secured op)")
	}
	if got["about.html"] {
		t.Errorf("about.html should be public (consumes no ops)")
	}

	if pages := resolveProtectedPages(&yongol.Fullstack{Manifest: fs.Manifest, STMLPages: fs.STMLPages}); pages != nil {
		t.Errorf("nil OpenAPI doc should yield nil, got %v", pages)
	}
	if pages := resolveProtectedPages(&yongol.Fullstack{OpenAPIDoc: doc, STMLPages: fs.STMLPages}); pages != nil {
		t.Errorf("backend.auth absent should yield nil, got %v", pages)
	}
}
