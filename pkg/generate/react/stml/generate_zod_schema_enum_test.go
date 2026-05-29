//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema enum 필드 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaEnum(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"role": {Type: "string", Enum: []string{"admin", "member"}, Required: true},
	}
	code := generateZodSchema("Assign", fields)
	assertContains(t, code, `role: z.enum(["admin", "member"])`)
	assertNotContains(t, code, ".optional()")
}
