//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema pattern regex 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaPattern(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"phone": {Type: "string", Required: true, Pattern: `^\d{3}-\d{4}-\d{4}$`},
	}
	code := generateZodSchema("UpdatePhone", fields)
	assertContains(t, code, `phone: z.string().min(1).regex(new RegExp("^\d{3}-\d{4}-\d{4}$"))`)
}
