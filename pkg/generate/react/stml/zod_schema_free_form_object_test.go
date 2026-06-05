//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema — 자유형 object 필드가 z.record(z.unknown()) 로 생성되는지 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaFreeFormObject(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"meta": {Type: "object", MapValueType: "*", Required: true},
	}
	code := generateZodSchema("SaveMeta", fields)
	assertContains(t, code, "meta: z.record(z.unknown()),")
}
