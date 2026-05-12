//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema boolean 필드 스키마 생성 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaBoolean(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"active": {Type: "boolean", Required: true},
	}
	code := generateZodSchema("Toggle", fields)
	assertContains(t, code, "active: z.boolean()")
	assertNotContains(t, code, ".optional()")
}
