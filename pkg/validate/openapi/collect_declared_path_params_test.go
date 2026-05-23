//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what collectDeclaredPathParams — nil/empty/path+op-level 파라미터 합집합 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectDeclaredPathParams(t *testing.T) {
	t.Run("nil params returns empty set", func(t *testing.T) {
		got := collectDeclaredPathParams(nil, nil)
		if len(got) != 0 {
			t.Fatalf("expected empty, got %v", got)
		}
	})

	t.Run("path-level params collected", func(t *testing.T) {
		pathParams := openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "id"}},
			&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "query", Name: "page"}},
		}
		got := collectDeclaredPathParams(pathParams, nil)
		if !got["id"] {
			t.Error("expected id in set")
		}
		if got["page"] {
			t.Error("query param page should not be in set")
		}
	})

	t.Run("op-level params collected", func(t *testing.T) {
		opParams := openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "user_id"}},
		}
		got := collectDeclaredPathParams(nil, opParams)
		if !got["user_id"] {
			t.Error("expected user_id in set")
		}
	})

	t.Run("union of both levels", func(t *testing.T) {
		pathParams := openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "org_id"}},
		}
		opParams := openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "user_id"}},
		}
		got := collectDeclaredPathParams(pathParams, opParams)
		if len(got) != 2 {
			t.Fatalf("expected 2, got %d: %v", len(got), got)
		}
		if !got["org_id"] || !got["user_id"] {
			t.Errorf("missing expected params: %v", got)
		}
	})

	t.Run("nil ref.Value is skipped", func(t *testing.T) {
		params := openapi3.Parameters{
			nil,
			&openapi3.ParameterRef{Value: nil},
			&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "id"}},
		}
		got := collectDeclaredPathParams(params, nil)
		if len(got) != 1 || !got["id"] {
			t.Errorf("expected {id: true}, got %v", got)
		}
	})
}
