//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh08CheckEntry — hurl entry의 Captures jsonpath를 response schema로 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh08CheckEntry(t *testing.T) {
	resp := openapi3.NewResponse()
	resp.Content = openapi3.Content{
		"application/json": &openapi3.MediaType{
			Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"token": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			}},
		},
	}
	op := &openapi3.Operation{OperationID: "login", Responses: &openapi3.Responses{}}
	op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})

	opNoSchema := &openapi3.Operation{Responses: &openapi3.Responses{}}
	opNoSchema.Responses.Set("204", &openapi3.ResponseRef{Value: openapi3.NewResponse()})

	routes := []apiRoute{
		{Path: "/auth/login", Method: "POST", Segments: []string{"auth", "login"}, Op: op},
		{Path: "/users", Method: "DELETE", Segments: []string{"users"}, Op: opNoSchema},
	}

	cases := []struct {
		name      string
		entry     hurl.HurlEntry
		wantCount int
	}{
		{
			name:      "no_captures_skip",
			entry:     hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "200"},
			wantCount: 0,
		},
		{
			name: "no_route_skip",
			entry: hurl.HurlEntry{Method: "POST", Path: "/orders",
				Captures: []hurl.HurlCapture{{Var: "t", Source: "jsonpath", JSONPath: "$.id", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "no_response_schema_skip",
			entry: hurl.HurlEntry{Method: "DELETE", Path: "/users", StatusCode: "204",
				Captures: []hurl.HurlCapture{{Var: "t", Source: "jsonpath", JSONPath: "$.id", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "non_jsonpath_capture_skip",
			entry: hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "200",
				Captures: []hurl.HurlCapture{{Var: "csrf", Source: "header", Header: "X-CSRF-Token", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "reachable_capture_no_diag",
			entry: hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "200",
				Captures: []hurl.HurlCapture{{Var: "tok", Source: "jsonpath", JSONPath: "$.token", Line: 1}}},
			wantCount: 0,
		},
		{
			name: "unreachable_capture_produces_diag",
			entry: hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "t.hurl",
				Captures: []hurl.HurlCapture{{Var: "tok", Source: "jsonpath", JSONPath: "$.missing", Line: 5}}},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh08CheckEntry(c.entry, routes), c.wantCount, "[XOH-08]")
		})
	}
}
