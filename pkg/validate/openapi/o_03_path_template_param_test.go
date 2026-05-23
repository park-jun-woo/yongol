//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what o03PathTemplateParam — nil doc + 일치 + 불일치 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO03PathTemplateParam(t *testing.T) {
	t.Run("nil OpenAPIDoc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := o03PathTemplateParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching template and params returns nil", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{
					Get: &openapi3.Operation{
						OperationID: "getUser",
						Parameters: openapi3.Parameters{
							&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "id"}},
						},
					},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}, Paths: map[string]int{}},
		}
		diags := o03PathTemplateParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("nil path item is skipped", func(t *testing.T) {
		paths := openapi3.NewPaths()
		paths.Set("/health", nil)
		paths.Set("/users/{id}", &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "getUser",
				Parameters: openapi3.Parameters{
					&openapi3.ParameterRef{Value: &openapi3.Parameter{In: "path", Name: "id"}},
				},
			},
		})
		doc := &openapi3.T{Paths: paths}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}, Paths: map[string]int{}},
		}
		diags := o03PathTemplateParam(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing param declaration raises diagnostic", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{
					Get: &openapi3.Operation{
						OperationID: "getUser",
						// no parameters declared
					},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}, Paths: map[string]int{}},
		}
		diags := o03PathTemplateParam(fs)
		if len(diags) == 0 {
			t.Fatal("expected at least 1 diagnostic")
		}
		if !strings.Contains(diags[0].Message, "O-3") {
			t.Errorf("Message missing O-3: %s", diags[0].Message)
		}
	})
}
