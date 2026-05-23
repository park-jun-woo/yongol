//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what xoh08CapturePathInSchema — hurl Captures jsonpath가 response schema에 존재하는지 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh08CapturePathInSchema(t *testing.T) {
	t.Run("nil_fs", func(t *testing.T) {
		if diags := xoh08CapturePathInSchema(nil); len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("nil_doc", func(t *testing.T) {
		if diags := xoh08CapturePathInSchema(&yongol.Fullstack{}); len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("with_entries_and_doc", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
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
		op := &openapi3.Operation{Responses: &openapi3.Responses{}}
		op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})
		doc.Paths.Set("/auth/login", &openapi3.PathItem{Post: op})

		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			HurlEntries: []hurl.HurlEntry{{
				Method: "POST", Path: "/auth/login", StatusCode: "200",
				Captures: []hurl.HurlCapture{{Var: "tok", Source: "jsonpath", JSONPath: "$.token", Line: 5}},
			}},
		}
		diags := xoh08CapturePathInSchema(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
