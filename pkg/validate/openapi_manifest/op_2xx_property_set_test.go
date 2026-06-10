//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what op2xxPropertySet — nil op/응답없음/2xx 수집/비2xx·비JSON·스키마없음·Value nil 스킵 검증

package openapi_manifest

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOp2xxPropertySet(t *testing.T) {
	t.Run("nil op returns empty", func(t *testing.T) {
		if got := op2xxPropertySet(nil); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("no responses returns empty", func(t *testing.T) {
		if got := op2xxPropertySet(&openapi3.Operation{}); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("collects 2xx json object properties", func(t *testing.T) {
		op := &openapi3.Operation{Responses: build2xxObjectResponses([]string{"token", "user"})}
		got := op2xxPropertySet(op)
		if !got["token"] || !got["user"] || len(got) != 2 {
			t.Errorf("expected {token,user}, got %v", got)
		}
	})

	t.Run("non-2xx code skipped", func(t *testing.T) {
		resp := openapi3.NewResponses()
		resp.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{
			Content: openapi3.Content{"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"err": {Value: &openapi3.Schema{}}},
				}},
			}},
		}})
		if got := op2xxPropertySet(&openapi3.Operation{Responses: resp}); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("nil response value skipped", func(t *testing.T) {
		resp := openapi3.NewResponses()
		resp.Set("200", &openapi3.ResponseRef{Value: nil})
		if got := op2xxPropertySet(&openapi3.Operation{Responses: resp}); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("non-json content skipped", func(t *testing.T) {
		resp := openapi3.NewResponses()
		resp.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{
			Content: openapi3.Content{"text/plain": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{"x": {Value: &openapi3.Schema{}}},
				}},
			}},
		}})
		if got := op2xxPropertySet(&openapi3.Operation{Responses: resp}); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("nil schema skipped", func(t *testing.T) {
		resp := openapi3.NewResponses()
		resp.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{
			Content: openapi3.Content{"application/json": &openapi3.MediaType{Schema: nil}},
		}})
		if got := op2xxPropertySet(&openapi3.Operation{Responses: resp}); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})
}
