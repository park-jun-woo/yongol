//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema required + minLength>1 시 양쪽 .min() 모두 포함 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaMinLengthGreaterThan1(t *testing.T) {
	minLen := 8
	fields := map[string]oapiparser.FieldConstraint{
		"password": {Type: "string", Required: true, MinLength: &minLen},
	}
	code := generateZodSchema("Login", fields)
	assertContains(t, code, "password: z.string().min(1).min(8)")
}
