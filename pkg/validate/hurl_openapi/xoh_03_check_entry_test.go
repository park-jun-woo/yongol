//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh03CheckEntry — hurl entry body field ↔ OpenAPI request schema 비교 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh03CheckEntry(t *testing.T) {
	op := &openapi3.Operation{
		OperationID: "createUser",
		Responses:   &openapi3.Responses{},
		RequestBody: &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
				Properties: openapi3.Schemas{
					"name":  &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			}),
		},
	}
	opNoBody := &openapi3.Operation{
		OperationID: "getItems",
		Responses:   &openapi3.Responses{},
	}
	routes := []apiRoute{
		{Path: "/users", Method: "POST", Segments: []string{"users"}, Op: op},
		{Path: "/items", Method: "GET", Segments: []string{"items"}, Op: opNoBody},
	}

	cases := []struct {
		name      string
		entry     hurl.HurlEntry
		wantCount int
	}{
		{
			name:      "no_body_fields_skip",
			entry:     hurl.HurlEntry{Method: "POST", Path: "/users"},
			wantCount: 0,
		},
		{
			name:      "no_route_match_skip",
			entry:     hurl.HurlEntry{Method: "POST", Path: "/orders", BodyFields: []string{"name"}},
			wantCount: 0,
		},
		{
			name:      "valid_field_no_diag",
			entry:     hurl.HurlEntry{Method: "POST", Path: "/users", BodyFields: []string{"name", "email"}},
			wantCount: 0,
		},
		{
			name:      "invalid_field_produces_diag",
			entry:     hurl.HurlEntry{Method: "POST", Path: "/users", BodyFields: []string{"name", "phone"}, File: "t.hurl", Line: 5},
			wantCount: 1,
		},
		{
			name:      "all_invalid_fields",
			entry:     hurl.HurlEntry{Method: "POST", Path: "/users", BodyFields: []string{"age", "phone"}, File: "t.hurl", Line: 5},
			wantCount: 2,
		},
		{
			name:      "route_match_but_no_request_body_schema_skip",
			entry:     hurl.HurlEntry{Method: "GET", Path: "/items", BodyFields: []string{"name"}},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh03CheckEntry(c.entry, routes), c.wantCount, "[XOH-03]")
		})
	}
}
