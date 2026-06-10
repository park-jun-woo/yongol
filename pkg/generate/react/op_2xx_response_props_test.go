//ff:func feature=gen-react type=test control=sequence
//ff:what op2xxResponseProps — 2xx JSON 스키마 프로퍼티만 수집, 비2xx·비JSON·스키마부재·nil 제외 검증

package react

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOp2xxResponseProps(t *testing.T) {
	// nil op -> empty (non-nil) set
	if got := op2xxResponseProps(nil); got == nil || len(got) != 0 {
		t.Errorf("nil op = %v, want empty set", got)
	}

	// op with no Responses -> empty
	if got := op2xxResponseProps(&openapi3.Operation{}); len(got) != 0 {
		t.Errorf("no responses = %v, want empty", got)
	}

	objSchema := func(fields ...string) *openapi3.SchemaRef {
		props := openapi3.Schemas{}
		for _, f := range fields {
			props[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
		}
		return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, Properties: props}}
	}

	responses := openapi3.NewResponses()
	// 200 JSON -> collected
	responses.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{"application/json": &openapi3.MediaType{Schema: objSchema("access_token", "refresh_token")}},
	}})
	// non-2xx -> ignored
	responses.Set("400", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{"application/json": &openapi3.MediaType{Schema: objSchema("error")}},
	}})
	// 2xx non-JSON media type -> ignored
	responses.Set("204", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{"text/plain": &openapi3.MediaType{Schema: objSchema("plain")}},
	}})
	// 2xx JSON but schema-less media type -> ignored (no panic)
	responses.Set("201", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: openapi3.Content{"application/json": &openapi3.MediaType{}},
	}})
	// 2xx ResponseRef with nil Value -> ignored
	responses.Set("202", &openapi3.ResponseRef{})

	props := op2xxResponseProps(&openapi3.Operation{Responses: responses})

	if len(props) != 2 || !props["access_token"] || !props["refresh_token"] {
		t.Fatalf("props = %v, want only 200 JSON props", props)
	}
	if props["error"] || props["plain"] {
		t.Errorf("non-2xx or non-JSON props leaked: %v", props)
	}
}
