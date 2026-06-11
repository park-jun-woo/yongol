//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what XOE-02 — 에러 표시 필드 정합 진단 (string error 통과 / 부재 WARNING / 비-string WARNING / nil doc)
package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXOE02ErrorDisplayField(t *testing.T) {
	str := func() *openapi3.SchemaRef {
		return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	fsWith := func(props openapi3.Schemas) *yongol.Fullstack {
		return &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					Schemas: openapi3.Schemas{
						"ErrorResponse": {Value: &openapi3.Schema{
							Type:       &openapi3.Types{"object"},
							Properties: props,
						}},
					},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{
				Schemas:          map[string]int{"ErrorResponse": 10},
				SchemaProperties: map[string]map[string]int{"ErrorResponse": {"error": 11}},
			},
		}
	}

	t.Run("string error passes", func(t *testing.T) {
		if d := xoe02ErrorDisplayField(fsWith(openapi3.Schemas{"error": str(), "code": str()})); len(d) != 0 {
			t.Errorf("string error schema should produce 0 diags, got %d", len(d))
		}
	})

	t.Run("no display field warns", func(t *testing.T) {
		d := xoe02ErrorDisplayField(fsWith(openapi3.Schemas{"code": str()}))
		if len(d) != 1 {
			t.Fatalf("missing display field should produce 1 diag, got %d", len(d))
		}
		if !strings.Contains(d[0].Message, "[XOE-02]") {
			t.Errorf("diag message missing XOE-02 tag: %q", d[0].Message)
		}
	})

	t.Run("non-string error warns", func(t *testing.T) {
		d := xoe02ErrorDisplayField(fsWith(openapi3.Schemas{
			"error": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		}))
		if len(d) != 1 {
			t.Fatalf("non-string error should produce 1 diag, got %d", len(d))
		}
		if !strings.Contains(d[0].Message, "string 타입이 아니") {
			t.Errorf("diag message should flag non-string type: %q", d[0].Message)
		}
	})

	t.Run("nil doc no diag", func(t *testing.T) {
		if d := xoe02ErrorDisplayField(&yongol.Fullstack{}); d != nil {
			t.Errorf("nil doc should produce nil diags, got %v", d)
		}
	})
}
