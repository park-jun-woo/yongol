//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what TestXdo77PropDiag_UUID — UUID 컬럼의 format 일치/불일치/잘못된 타입 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo77PropDiag_UUID(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPILines: &oapiparser.LineIndex{
			SchemaProperties: map[string]map[string]int{},
			Schemas:          map[string]int{},
		},
	}
	uuidCols := map[string]string{
		"id":      "uuid",
		"user_id": "uuid",
		"email":   "string",
	}

	t.Run("uuid with format uuid returns false (pass)", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"string"}, Format: "uuid",
		}}
		_, ok := xdo77PropDiag(fs, "Order", "orders", "id", ref, uuidCols)
		if ok {
			t.Error("expected false for matching uuid format")
		}
	})

	t.Run("uuid without format returns diagnostic", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"string"},
		}}
		diag, ok := xdo77PropDiag(fs, "Order", "orders", "id", ref, uuidCols)
		if !ok {
			t.Fatal("expected true for uuid without format")
		}
		if !strings.Contains(diag.Message, "XDO-77") {
			t.Errorf("Message missing XDO-77: %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "UUID") {
			t.Errorf("Message missing UUID: %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "without format: uuid") {
			t.Errorf("Message missing hint text: %s", diag.Message)
		}
		if !strings.Contains(diag.Advice, "format: uuid") {
			t.Errorf("Advice missing format: uuid hint: %s", diag.Advice)
		}
	})

	t.Run("uuid with wrong format returns diagnostic", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"string"}, Format: "byte",
		}}
		diag, ok := xdo77PropDiag(fs, "Order", "orders", "user_id", ref, uuidCols)
		if !ok {
			t.Fatal("expected true for uuid with wrong format")
		}
		if !strings.Contains(diag.Message, "UUID") {
			t.Errorf("Message missing UUID: %s", diag.Message)
		}
	})

	t.Run("uuid with wrong type returns generic diagnostic", func(t *testing.T) {
		ref := &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"integer"}, Format: "int64",
		}}
		diag, ok := xdo77PropDiag(fs, "Order", "orders", "id", ref, uuidCols)
		if !ok {
			t.Fatal("expected true for uuid with wrong type")
		}
		// type mismatch: should use the generic message, not UUID-specific
		if !strings.Contains(diag.Message, "mismatch") {
			t.Errorf("expected generic mismatch message: %s", diag.Message)
		}
	})
}
