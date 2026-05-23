//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what buildOpIDLookup — METHOD /path → operationId 테이블 생성 검증

package hurl_statemachine

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildOpIDLookup(t *testing.T) {
	t.Run("nil_fs", func(t *testing.T) {
		got := buildOpIDLookup(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("nil_doc", func(t *testing.T) {
		got := buildOpIDLookup(&yongol.Fullstack{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("builds_lookup", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		doc.Paths.Set("/users", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "getUsers", Responses: &openapi3.Responses{}},
		})
		doc.Paths.Set("/users/{id}", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "getUser", Responses: &openapi3.Responses{}},
		})
		fs := &yongol.Fullstack{OpenAPIDoc: doc}
		got := buildOpIDLookup(fs)

		if got["GET /users"] != "getUsers" {
			t.Errorf("GET /users = %q, want getUsers", got["GET /users"])
		}
		if got["GET /users/:param"] != "getUser" {
			t.Errorf("GET /users/:param = %q, want getUser", got["GET /users/:param"])
		}
	})

	t.Run("skips_empty_operation_id", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		doc.Paths.Set("/health", &openapi3.PathItem{
			Get: &openapi3.Operation{Responses: &openapi3.Responses{}},
		})
		fs := &yongol.Fullstack{OpenAPIDoc: doc}
		got := buildOpIDLookup(fs)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})
}
