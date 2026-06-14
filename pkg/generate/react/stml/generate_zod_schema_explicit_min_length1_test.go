//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema 명시적 minLength:1 지정 시 .min(1) 부착 검증 (BUG-134)
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// BUG-134: an explicit minLength:1 is the sole source of nonempty (.min(1)).
func TestGenerateZodSchemaExplicitMinLength1(t *testing.T) {
	minLen := 1
	fields := map[string]oapiparser.FieldConstraint{
		"name": {Type: "string", Required: true, MinLength: &minLen},
	}
	code := generateZodSchema("Create", fields)
	assertContains(t, code, "name: z.string().min(1)")
}
