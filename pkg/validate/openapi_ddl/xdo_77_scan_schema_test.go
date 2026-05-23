//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo77ScanSchema — nil ref/테이블 미매칭/속성 스캔 검증

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo77ScanSchema(t *testing.T) {
	tableIndex := map[string]map[string]string{
		"users": {"id": "int64", "email": "string"},
	}
	fs := &yongol.Fullstack{
		OpenAPILines: &oapiparser.LineIndex{
			SchemaProperties: map[string]map[string]int{},
			Schemas:          map[string]int{},
		},
	}

	t.Run("nil schemaRef returns nil", func(t *testing.T) {
		diags := xdo77ScanSchema(fs, "User", nil, tableIndex)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("schema does not map to table returns nil", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{}}
		diags := xdo77ScanSchema(fs, "Product", ref, tableIndex)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching properties no diagnostics", func(t *testing.T) {
		ref := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}},
					"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			},
		}
		diags := xdo77ScanSchema(fs, "User", ref, tableIndex)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("mismatched property raises diagnostic", func(t *testing.T) {
		ref := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			},
		}
		diags := xdo77ScanSchema(fs, "User", ref, tableIndex)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})
}
