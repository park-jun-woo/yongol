//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema array(integer) 필드가 z.array(z.number().int()) 생성 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaArrayInteger(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"scores": {
			Type:     "array",
			ItemType: "integer",
			Required: false,
		},
	}
	code := generateZodSchema("UpdateResult", fields)
	assertContains(t, code, "scores: z.array(z.number().int()).optional(),")
}
