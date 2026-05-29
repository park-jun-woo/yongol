//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema array(string) 필드가 z.array(z.string()) 생성 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaArrayString(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"tags": {
			Type:     "array",
			ItemType: "string",
			Required: true,
		},
	}
	code := generateZodSchema("CreatePost", fields)
	assertContains(t, code, "tags: z.array(z.string()),")
	assertNotContains(t, code, "tags: z.string()")
}
