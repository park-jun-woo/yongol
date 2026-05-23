//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xod10DDLToResponse — nil doc/스키마 매칭/컬럼 노출 여부 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXod10DDLToResponse_Unit(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xod10DDLToResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("all DDL columns present in schema no diagnostics", func(t *testing.T) {
		fs := buildXod10FS(
			[]ddl.Table{{Name: "users", Columns: map[string]ddl.Column{"id": {}, "email": {}}}},
			openapi3.Schemas{
				"User": &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{}},
						"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}},
			},
		)
		diags := xod10DDLToResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("DDL column missing from schema raises warning", func(t *testing.T) {
		fs := buildXod10FS(
			[]ddl.Table{{Name: "users", Columns: map[string]ddl.Column{"id": {}, "email": {}, "phone": {}}}},
			openapi3.Schemas{
				"User": &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{}},
						"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}},
			},
		)
		diags := xod10DDLToResponse(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("nil schemaRef.Value skipped", func(t *testing.T) {
		fs := buildXod10FS(
			[]ddl.Table{{Name: "users", Columns: map[string]ddl.Column{"id": {}}}},
			openapi3.Schemas{"User": &openapi3.SchemaRef{Value: nil}},
		)
		diags := xod10DDLToResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("table without matching schema skipped", func(t *testing.T) {
		fs := buildXod10FS(
			[]ddl.Table{{Name: "products", Columns: map[string]ddl.Column{"id": {}}}},
			openapi3.Schemas{
				"User": &openapi3.SchemaRef{Value: &openapi3.Schema{Properties: openapi3.Schemas{}}},
			},
		)
		diags := xod10DDLToResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
