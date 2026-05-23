//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo77ColumnTypeMismatch — nil doc/스키마 있는 경우 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo77ColumnTypeMismatch(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo77ColumnTypeMismatch(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching types no diagnostics", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"id":    {RawType: "BIGINT"},
					"email": {RawType: "TEXT"},
				}},
			},
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					Schemas: openapi3.Schemas{
						"User": &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: &openapi3.Types{"object"},
								Properties: openapi3.Schemas{
									"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}},
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
		diags := xdo77ColumnTypeMismatch(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
