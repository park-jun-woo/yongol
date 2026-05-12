//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema 필드 알파벳 순서 보장 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaDeterministicOrder(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"z_field": {Type: "string", Required: true},
		"a_field": {Type: "string", Required: true},
		"m_field": {Type: "string", Required: true},
	}
	code := generateZodSchema("Test", fields)
	aIdx := strings.Index(code, "a_field")
	mIdx := strings.Index(code, "m_field")
	zIdx := strings.Index(code, "z_field")
	if aIdx > mIdx || mIdx > zIdx {
		t.Errorf("fields not in alphabetical order\n--- code ---\n%s", code)
	}
}
