//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what xoe02SchemaDiag — string 표시 필드 보유 시 nil, 부재 시 XOE-02 WARNING 반환 검증
package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestXOE02SchemaDiag(t *testing.T) {
	lines := &oapiparser.LineIndex{
		Schemas:          map[string]int{"ErrorResponse": 10},
		SchemaProperties: map[string]map[string]int{"ErrorResponse": {"error": 11}},
	}

	withString := &openapi3.Schema{Properties: openapi3.Schemas{
		"error": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}}
	if d := xoe02SchemaDiag("ErrorResponse", withString, lines); d != nil {
		t.Errorf("string error schema should yield nil, got %+v", d)
	}

	missing := &openapi3.Schema{Properties: openapi3.Schemas{
		"code": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}}
	d := xoe02SchemaDiag("ErrorResponse", missing, lines)
	if d == nil {
		t.Fatalf("missing display field should yield a diagnostic")
	}
	if d.Level != diagnostic.LevelWarning || !strings.Contains(d.Message, "[XOE-02]") {
		t.Errorf("unexpected diagnostic: %+v", d)
	}
}
