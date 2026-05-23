//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what checkOpResponseBodies — nil op/responses + 4xx body 유무 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestCheckOpResponseBodies(t *testing.T) {
	lines := &oapiparser.LineIndex{
		Operations: map[string]int{"getUser": 10},
	}

	t.Run("nil operation returns nil", func(t *testing.T) {
		diags := checkOpResponseBodies(nil, "/users", lines)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil responses returns nil", func(t *testing.T) {
		op := &openapi3.Operation{OperationID: "getUser"}
		diags := checkOpResponseBodies(op, "/users", lines)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("2xx response without body is not an error", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{
			Value: &openapi3.Response{},
		})
		op := &openapi3.Operation{
			OperationID: "getUser",
			Responses:   resps,
		}
		diags := checkOpResponseBodies(op, "/users", lines)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("4xx without body raises diagnostic", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{
			Value: &openapi3.Response{},
		})
		op := &openapi3.Operation{
			OperationID: "getUser",
			Responses:   resps,
		}
		diags := checkOpResponseBodies(op, "/users/{id}", lines)
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
						Schema: &openapi3.SchemaRef{
							Ref: "#/components/schemas/Error",
						},
					},
				},
			},
		})
		op := &openapi3.Operation{
			OperationID: "getUser",
			Responses:   resps,
		}
		diags := checkOpResponseBodies(op, "/users/{id}", lines)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
