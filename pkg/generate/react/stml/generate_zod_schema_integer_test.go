//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema integer 필드 스키마 생성 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaInteger(t *testing.T) {
	min := 1.0
	max := 100.0
	fields := map[string]oapiparser.FieldConstraint{
		"capacity": {
			Type:     "integer",
			Required: true,
			Minimum:  &min,
			Maximum:  &max,
		},
	}
	code := generateZodSchema("UpdateRoom", fields)
	assertContains(t, code, "capacity: z.number().int().min(1).max(100)")
}
