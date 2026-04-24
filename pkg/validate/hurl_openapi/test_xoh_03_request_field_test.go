//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what XOH-03 positive/negative — request body 필드가 OpenAPI schema 에 존재

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh03_Negative_FieldMissing(t *testing.T) {
	op := withRequestBody("Register", map[string]*openapi3.Schema{
		"email":    {Type: &openapi3.Types{"string"}},
		"password": {Type: &openapi3.Types{"string"}},
	})
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/register": {"POST": op},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/register", StatusCode: "201",
				File: "t.hurl", Line: 1,
				BodyFields: []string{"emale"}, // typo
			},
		},
	}
	diags := xoh03RequestFieldInSchema(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "emale") || !strings.Contains(diags[0].Advice, "email") {
		t.Fatalf("advice/msg missing field hints: %+v", diags[0])
	}
}

func TestXoh03_Positive_AllFieldsPresent(t *testing.T) {
	op := withRequestBody("Register", map[string]*openapi3.Schema{
		"email":    {Type: &openapi3.Types{"string"}},
		"password": {Type: &openapi3.Types{"string"}},
	})
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/register": {"POST": op},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/register", StatusCode: "201",
				File: "t.hurl", Line: 1,
				BodyFields: []string{"email", "password"},
			},
		},
	}
	if diags := xoh03RequestFieldInSchema(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}

// withRequestBody attaches a JSON request-body schema with the given
// properties to a synthetic operation. Keeps fixtures inline.
func withRequestBody(opID string, props map[string]*openapi3.Schema) *openapi3.Operation {
	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{},
	}
	for k, v := range props {
		schema.Properties[k] = &openapi3.SchemaRef{Value: v}
	}
	op := &openapi3.Operation{
		OperationID: opID,
		Responses:   openapi3.NewResponses(),
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{Value: schema},
					},
				},
			},
		},
	}
	return op
}
