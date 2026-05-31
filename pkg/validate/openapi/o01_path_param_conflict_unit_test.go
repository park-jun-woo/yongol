//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what o01PathParamConflict — nil doc + 중복 없음 + 중복 발견 검증
package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO01PathParamConflict_Unit(t *testing.T) {
	t.Run("nil OpenAPIDoc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := o01PathParamConflict(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no conflict returns nil", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{user_id}/posts/{post_id}", &openapi3.PathItem{}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{}},
		}
		diags := o01PathParamConflict(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("duplicate param raises diagnostic", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}/posts/{id}", &openapi3.PathItem{}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{"/users/{id}/posts/{id}": 5}},
		}
		diags := o01PathParamConflict(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "O-1") {
			t.Errorf("Message missing O-1: %s", diags[0].Message)
		}
	})
}
