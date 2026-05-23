//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what xoh04AssertPathInSchema — hurl Asserts jsonpath가 OpenAPI response schema에 도달 가능한지 검증

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh04AssertPathInSchema(t *testing.T) {
	t.Run("nil_fs", func(t *testing.T) {
		diags := xoh04AssertPathInSchema(nil)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("nil_doc", func(t *testing.T) {
		diags := xoh04AssertPathInSchema(&yongol.Fullstack{})
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("reachable_assert_no_diag", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		resp := openapi3.NewResponse()
		resp.Content = openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}},
			},
		}
		op := &openapi3.Operation{Responses: &openapi3.Responses{}}
		op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})
		doc.Paths.Set("/users", &openapi3.PathItem{Get: op})

		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			HurlEntries: []hurl.HurlEntry{{
				Method: "GET", Path: "/users", StatusCode: "200",
				Asserts: []hurl.HurlAssert{{JSONPath: "$.id", Line: 5}},
			}},
		}
		diags := xoh04AssertPathInSchema(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("unreachable_assert_produces_diag", func(t *testing.T) {
		doc := &openapi3.T{Paths: &openapi3.Paths{}}
		resp := openapi3.NewResponse()
		resp.Content = openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}},
			},
		}
		op := &openapi3.Operation{Responses: &openapi3.Responses{}}
		op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})
		doc.Paths.Set("/users", &openapi3.PathItem{Get: op})

		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			HurlEntries: []hurl.HurlEntry{{
				Method: "GET", Path: "/users", StatusCode: "200", File: "t.hurl",
				Asserts: []hurl.HurlAssert{{JSONPath: "$.email", Line: 5}},
			}},
		}
		diags := xoh04AssertPathInSchema(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XOH-04]") {
			t.Errorf("expected [XOH-04], got %q", diags[0].Message)
		}
	})
}
