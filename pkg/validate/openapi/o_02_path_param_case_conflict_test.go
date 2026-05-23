//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what o02PathParamCaseConflict — nil doc + 일관 케이스 + 불일치 케이스 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO02PathParamCaseConflict(t *testing.T) {
	t.Run("nil OpenAPIDoc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := o02PathParamCaseConflict(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("consistent casing returns nil", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{}),
				openapi3.WithPath("/posts/{id}", &openapi3.PathItem{}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{}},
		}
		diags := o02PathParamCaseConflict(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("case conflict raises diagnostic", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{}),
				openapi3.WithPath("/accounts/{ID}", &openapi3.PathItem{}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{}},
		}
		diags := o02PathParamCaseConflict(fs)
		if len(diags) == 0 {
			t.Fatal("expected at least 1 diagnostic")
		}
		if !strings.Contains(diags[0].Message, "O-2") {
			t.Errorf("Message missing O-2: %s", diags[0].Message)
		}
	})
}
