//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what XOH-04 positive/negative — assert jsonpath 가 response schema 에 존재

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh04_Negative_AssertPathMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": withJSONResponse("Login", "200", map[string]*openapi3.Schema{
				"access_token": {Type: &openapi3.Types{"string"}},
			})},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/login", StatusCode: "200",
				File: "t.hurl", Line: 1,
				Asserts: []hurl.HurlAssert{{JSONPath: "$.nonexistent", Line: 5}},
			},
		},
	}
	diags := xoh04AssertPathInSchema(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[XOH-04]") {
		t.Fatalf("want 1 XOH-04 diag, got %+v", diags)
	}
}

func TestXoh04_Positive_AssertPathPresent(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": withJSONResponse("Login", "200", map[string]*openapi3.Schema{
				"access_token": {Type: &openapi3.Types{"string"}},
			})},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/login", StatusCode: "200",
				File: "t.hurl", Line: 1,
				Asserts: []hurl.HurlAssert{{JSONPath: "$.access_token", Line: 5}},
			},
		},
	}
	if diags := xoh04AssertPathInSchema(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}

// withJSONResponse attaches a JSON response-body schema to a synthetic
// operation. Keeps fixtures inline.
func withJSONResponse(opID, status string, props map[string]*openapi3.Schema) *openapi3.Operation {
	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{},
	}
	for k, v := range props {
		schema.Properties[k] = &openapi3.SchemaRef{Value: v}
	}
	resp := &openapi3.Response{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: schema},
			},
		},
	}
	op := &openapi3.Operation{
		OperationID: opID,
		Responses:   openapi3.NewResponses(),
	}
	op.Responses.Set(status, &openapi3.ResponseRef{Value: resp})
	return op
}
