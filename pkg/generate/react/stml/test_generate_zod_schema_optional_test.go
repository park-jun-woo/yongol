//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema 선택적 문자열 필드 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaOptionalField(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"bio": {Type: "string"},
	}
	code := generateZodSchema("UpdateProfile", fields)
	assertContains(t, code, "bio: z.string().optional()")
}
