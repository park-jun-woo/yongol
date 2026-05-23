//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what o05ResponseBodyRequired — nil doc + nil item skip + 4xx body 유무 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO05ResponseBodyRequired(t *testing.T) {
	t.Run("nil OpenAPIDoc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := o05ResponseBodyRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil path item skipped", func(t *testing.T) {
		paths := openapi3.NewPaths()
		paths.Set("/health", nil)
		doc := &openapi3.T{Paths: paths}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Paths: map[string]int{}, Operations: map[string]int{}},
		}
		diags := o05ResponseBodyRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("4xx without body raises diagnostic", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{
					Get: &openapi3.Operation{
						OperationID: "getUser",
						Responses:   resps,
					},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{"getUser": 10}, Paths: map[string]int{}},
		}
		diags := o05ResponseBodyRequired(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "O-5") {
			t.Errorf("Message missing O-5: %s", diags[0].Message)
		}
	})

	t.Run("4xx with JSON schema passes", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{Ref: "#/components/schemas/Error"},
					},
				},
			},
		})
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{
					Get: &openapi3.Operation{
						OperationID: "getUser",
						Responses:   resps,
					},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}, Paths: map[string]int{}},
		}
		diags := o05ResponseBodyRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
