//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what responseSchemaForStatus — status별 JSON response schema 선택 로직 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestResponseSchemaForStatus(t *testing.T) {
	schema200 := &openapi3.Schema{Type: &openapi3.Types{"object"}}
	schema201 := &openapi3.Schema{Type: &openapi3.Types{"array"}}
	schemaDefault := &openapi3.Schema{Type: &openapi3.Types{"string"}}

	makeRoute := func(codes map[string]*openapi3.Schema) *apiRoute {
		op := &openapi3.Operation{Responses: &openapi3.Responses{}}
		for code, s := range codes {
			resp := openapi3.NewResponse()
			if s != nil {
				resp.Content = openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{Value: s},
					},
				}
			}
			op.Responses.Set(code, &openapi3.ResponseRef{Value: resp})
		}
		return &apiRoute{Op: op}
	}

	cases := []struct {
		name       string
		route      *apiRoute
		statusCode string
		wantNil    bool
		wantPtr    *openapi3.Schema
	}{
		{name: "nil_route", route: nil, statusCode: "200", wantNil: true},
		{
			name:       "exact_match",
			route:      makeRoute(map[string]*openapi3.Schema{"200": schema200}),
			statusCode: "200",
			wantPtr:    schema200,
		},
		{
			name:       "fallback_to_2xx_when_empty_status",
			route:      makeRoute(map[string]*openapi3.Schema{"201": schema201}),
			statusCode: "",
			wantPtr:    schema201,
		},
		{
			name:       "fallback_to_default",
			route:      makeRoute(map[string]*openapi3.Schema{"default": schemaDefault}),
			statusCode: "200",
			wantPtr:    schemaDefault,
		},
		{
			name:       "no_json_schema_returns_nil",
			route:      makeRoute(map[string]*openapi3.Schema{"200": nil}),
			statusCode: "200",
			wantNil:    true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSchemaPointerCase(t, responseSchemaForStatus(c.route, c.statusCode), c.wantNil, c.wantPtr)
		})
	}
}
