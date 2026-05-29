//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema number 필드 스키마 생성 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaNumber(t *testing.T) {
	min := 0.0
	max := 99.99
	fields := map[string]oapiparser.FieldConstraint{
		"price": {
			Type:    "number",
			Minimum: &min,
			Maximum: &max,
		},
	}
	code := generateZodSchema("SetPrice", fields)
	assertContains(t, code, "price: z.number().min(0).max(99.99).optional()")
}
