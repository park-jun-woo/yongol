//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema 선택적 enum 필드 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaEnumOptional(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"role": {Type: "string", Enum: []string{"admin", "member"}},
	}
	code := generateZodSchema("Assign", fields)
	assertContains(t, code, `role: z.enum(["admin", "member"]).optional()`)
}
