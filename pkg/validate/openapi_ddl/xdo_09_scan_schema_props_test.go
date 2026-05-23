//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo09ScanSchemaProps — nil ref/테이블 미매칭/ghost property 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo09ScanSchemaProps(t *testing.T) {
	tables := map[string]map[string]bool{
		"users": {"id": true, "email": true},
	}
	lines := &oapiparser.LineIndex{
		SchemaProperties: map[string]map[string]int{},
		Schemas:          map[string]int{},
	}
	fs := &yongol.Fullstack{OpenAPILines: lines}

	t.Run("nil schemaRef returns nil", func(t *testing.T) {
		diags := xdo09ScanSchemaProps(fs, "User", nil, tables)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil Value returns nil", func(t *testing.T) {
		diags := xdo09ScanSchemaProps(fs, "User", &openapi3.SchemaRef{}, tables)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("schema does not map to any table returns nil", func(t *testing.T) {
		ref := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			},
		}
		diags := xdo09ScanSchemaProps(fs, "Product", ref, tables)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("all properties match DDL columns", func(t *testing.T) {
		ref := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			},
		}
		diags := xdo09ScanSchemaProps(fs, "User", ref, tables)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("ghost property raises diagnostic", func(t *testing.T) {
		ref := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id":    &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					"ghost": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			},
		}
		diags := xdo09ScanSchemaProps(fs, "User", ref, tables)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XDO-9") {
			t.Errorf("Message missing XDO-9: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "ghost") {
			t.Errorf("Message missing ghost property name: %s", diags[0].Message)
		}
	})
}
