//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema nil/empty 필드 시 빈 문자열 반환 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaEmpty(t *testing.T) {
	code := generateZodSchema("SomeOp", nil)
	if code != "" {
		t.Errorf("expected empty string for nil fields, got %q", code)
	}
	code = generateZodSchema("SomeOp", map[string]oapiparser.FieldConstraint{})
	if code != "" {
		t.Errorf("expected empty string for empty fields, got %q", code)
	}
}
