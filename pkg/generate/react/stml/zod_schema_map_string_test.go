//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema — string 값 맵 필드가 z.record(z.string()) 로 생성되는지 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaMapString(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"data": {Type: "object", MapValueType: "string", Required: true},
	}
	code := generateZodSchema("GenerateFilledPDF", fields)
	assertContains(t, code, "data: z.record(z.string()),")
	assertNotContains(t, code, "data: z.string()")
}
