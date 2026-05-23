//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what collectOpenAPIRoutes — OpenAPI doc에서 정규화된 route 목록 생성 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectOpenAPIRoutes(t *testing.T) {
	t.Run("nil_doc_returns_empty", func(t *testing.T) {
		routes := collectOpenAPIRoutes(nil)
		if len(routes) != 0 {
			t.Fatalf("expected empty, got %d routes", len(routes))
		}
	})

	t.Run("nil_paths_returns_empty", func(t *testing.T) {
		doc := &openapi3.T{}
		routes := collectOpenAPIRoutes(doc)
		if len(routes) != 0 {
			t.Fatalf("expected empty, got %d routes", len(routes))
		}
	})

	t.Run("single_get_operation", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: &openapi3.Paths{},
		}
		doc.Paths.Set("/users", &openapi3.PathItem{
			Get: &openapi3.Operation{
				Responses: &openapi3.Responses{},
			},
		})
		routes := collectOpenAPIRoutes(doc)
		if len(routes) != 1 {
			t.Fatalf("expected 1 route, got %d", len(routes))
		}
		if routes[0].Method != "GET" {
			t.Errorf("method = %q, want GET", routes[0].Method)
		}
		if routes[0].Path != "/users" {
			t.Errorf("path = %q, want /users", routes[0].Path)
		}
	})

	t.Run("multiple_methods_on_path", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: &openapi3.Paths{},
		}
		doc.Paths.Set("/users", &openapi3.PathItem{
			Get: &openapi3.Operation{
				Responses: &openapi3.Responses{},
			},
			Post: &openapi3.Operation{
				Responses: &openapi3.Responses{},
			},
		})
		routes := collectOpenAPIRoutes(doc)
		if len(routes) != 2 {
			t.Fatalf("expected 2 routes, got %d", len(routes))
		}
		methods := map[string]bool{}
		for _, r := range routes {
			methods[r.Method] = true
		}
		if !methods["GET"] || !methods["POST"] {
			t.Errorf("expected GET and POST, got %v", methods)
		}
	})

	t.Run("parameterized_path_normalized", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: &openapi3.Paths{},
		}
		doc.Paths.Set("/users/{id}", &openapi3.PathItem{
			Get: &openapi3.Operation{
				Responses: &openapi3.Responses{},
			},
		})
		routes := collectOpenAPIRoutes(doc)
		if len(routes) != 1 {
			t.Fatalf("expected 1 route, got %d", len(routes))
		}
		if len(routes[0].Segments) != 2 || routes[0].Segments[1] != ":param" {
			t.Errorf("segments = %v, want [users :param]", routes[0].Segments)
		}
	})
}
