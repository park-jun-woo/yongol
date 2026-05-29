//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo09GhostProperty — nil doc/nil components/empty schemas 조기 반환 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo09GhostProperty_Unit(t *testing.T) {
	t.Run("nil OpenAPIDoc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo09GhostProperty(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil Components returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{},
		}
		diags := xdo09GhostProperty(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil Schemas returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{},
			},
		}
		diags := xdo09GhostProperty(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty schemas returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					Schemas: openapi3.Schemas{},
				},
			},
		}
		diags := xdo09GhostProperty(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("schema with DDL tables scans properties", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{"id": {}, "email": {}}},
			},
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					Schemas: openapi3.Schemas{
						"User": &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: &openapi3.Types{"object"},
								Properties: openapi3.Schemas{
									"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
									"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
								},
							},
						},
					},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{
				SchemaProperties: map[string]map[string]int{},
				Schemas:          map[string]int{},
			},
		}
		diags := xdo09GhostProperty(fs)
		// All properties exist in DDL, so no ghost property diagnostics
		if len(diags) != 0 {
			t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
		}
	})
}

func TestXdo09GhostProperty(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
