//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema — integer 값 맵 필드가 z.record(z.number().int()) 로 생성되는지 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaMapInteger(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"counts": {Type: "object", MapValueType: "integer", Required: true},
	}
	code := generateZodSchema("SaveCounts", fields)
	assertContains(t, code, "counts: z.record(z.number().int()),")
}
