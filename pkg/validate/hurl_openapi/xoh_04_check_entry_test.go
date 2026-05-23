//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh04CheckEntry — hurl entry의 Asserts jsonpath를 response schema로 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh04CheckEntry(t *testing.T) {
	resp := openapi3.NewResponse()
	resp.Content = openapi3.Content{
		"application/json": &openapi3.MediaType{
			Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					"name": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			}},
		},
	}
	op := &openapi3.Operation{OperationID: "getUser", Responses: &openapi3.Responses{}}
	op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})

	opNoSchema := &openapi3.Operation{OperationID: "deleteUser", Responses: &openapi3.Responses{}}
	opNoSchema.Responses.Set("204", &openapi3.ResponseRef{Value: openapi3.NewResponse()})

	routes := []apiRoute{
		{Path: "/users", Method: "GET", Segments: []string{"users"}, Op: op},
		{Path: "/users", Method: "DELETE", Segments: []string{"users"}, Op: opNoSchema},
	}

	cases := []struct {
		name      string
		entry     hurl.HurlEntry
		wantCount int
	}{
		{
			name:      "no_asserts_skip",
			entry:     hurl.HurlEntry{Method: "GET", Path: "/users", StatusCode: "200"},
			wantCount: 0,
		},
		{
			name: "no_route_skip",
			entry: hurl.HurlEntry{Method: "GET", Path: "/orders",
				Asserts: []hurl.HurlAssert{{JSONPath: "$.id", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "no_response_schema_skip",
			entry: hurl.HurlEntry{Method: "DELETE", Path: "/users", StatusCode: "204",
				Asserts: []hurl.HurlAssert{{JSONPath: "$.id", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "reachable_path_no_diag",
			entry: hurl.HurlEntry{Method: "GET", Path: "/users", StatusCode: "200",
				Asserts: []hurl.HurlAssert{{JSONPath: "$.id", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "unreachable_path_produces_diag",
			entry: hurl.HurlEntry{Method: "GET", Path: "/users", StatusCode: "200", File: "t.hurl",
				Asserts: []hurl.HurlAssert{{JSONPath: "$.email", Line: 5}}},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh04CheckEntry(c.entry, routes), c.wantCount, "[XOH-04]")
		})
	}
}
