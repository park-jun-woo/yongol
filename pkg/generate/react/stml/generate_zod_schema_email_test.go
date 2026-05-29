//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema email 형식 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaEmail(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"email": {Type: "string", Format: "email", Required: true},
	}
	code := generateZodSchema("Register", fields)
	assertContains(t, code, "email: z.string().email().min(1)")
}
