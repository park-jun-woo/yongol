//ff:func feature=validate type=rule control=sequence topic=openapi-structural
//ff:what xoe02SchemaDiag — 단일 Error 스키마에 string 표시 필드가 없으면 XOE-02 진단을 반환(없으면 nil)

package openapi

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// xoe02SchemaDiag returns an XOE-02 WARNING for schema when neither "error"
// nor "message" exists as a string property, or nil when a usable string
// display field is present.
func xoe02SchemaDiag(name string, schema *openapi3.Schema, lines *oapiparser.LineIndex) *diagnostic.Diagnostic {
	if schemaPropIsString(schema, "error") || schemaPropIsString(schema, "message") {
		return nil
	}
	_, hasError := schema.Properties["error"]
	_, hasMessage := schema.Properties["message"]

	line, msg := xoe02Detail(name, hasError, hasMessage, lines)
	return &diagnostic.Diagnostic{
		File:    "api/openapi.yaml",
		Line:    line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: msg,
		Advice:  fmt.Sprintf("Declare a string \"error\" (or \"message\") property on schema %q", name),
	}
}
