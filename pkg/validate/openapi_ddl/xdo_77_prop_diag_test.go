//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what TestXdo77PropDiag — nil/미존재 컬럼/타입 일치/불일치 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo77PropDiag(t *testing.T) {
	cols := map[string]string{
		"id":     "int64",
		"email":  "string",
		"active": "bool",
	}
	fs := &yongol.Fullstack{
		OpenAPILines: &oapiparser.LineIndex{
			SchemaProperties: map[string]map[string]int{},
			Schemas:          map[string]int{},
		},
	}

	t.Run("nil propRef returns false", func(t *testing.T) {
		_, ok := xdo77PropDiag(fs, "User", "users", "id", nil, cols)
		if ok {
			t.Error("expected false for nil propRef")
		}
	})

	t.Run("column not in DDL returns false", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
		_, ok := xdo77PropDiag(fs, "User", "users", "ghost", ref, cols)
		if ok {
			t.Error("expected false for missing column")
		}
	})

	t.Run("matching types returns false", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}}
		_, ok := xdo77PropDiag(fs, "User", "users", "id", ref, cols)
		if ok {
			t.Error("expected false for matching types")
		}
	})

	t.Run("type mismatch returns diagnostic", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
		diag, ok := xdo77PropDiag(fs, "User", "users", "id", ref, cols)
		if !ok {
			t.Fatal("expected true for mismatch")
		}
		if !strings.Contains(diag.Message, "XDO-77") {
			t.Errorf("Message missing XDO-77: %s", diag.Message)
		}
	})

	t.Run("unknown DDL Go type returns false", func(t *testing.T) {
		unknownCols := map[string]string{"data": "unknownType"}
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
		_, ok := xdo77PropDiag(fs, "User", "users", "data", ref, unknownCols)
		if ok {
			t.Error("expected false for unknown DDL Go type")
		}
	})

	t.Run("empty type slice returns false", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{}}}
		_, ok := xdo77PropDiag(fs, "User", "users", "id", ref, cols)
		if ok {
			t.Error("expected false for empty type slice")
		}
	})

	t.Run("format mismatch with format in display", func(t *testing.T) {
		// id is int64 which expects integer/int64, provide integer/int32
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int32"}}
		diag, ok := xdo77PropDiag(fs, "User", "users", "id", ref, cols)
		if !ok {
			t.Fatal("expected true for format mismatch")
		}
		if !strings.Contains(diag.Message, "int32") {
			t.Errorf("Message missing format: %s", diag.Message)
		}
	})

	t.Run("boolean match returns false", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"boolean"}}}
		_, ok := xdo77PropDiag(fs, "User", "users", "active", ref, cols)
		if ok {
			t.Error("expected false for matching boolean")
		}
	})
}
