//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema required+minLength 미지정 시 .min(1) 미부착 검증 (BUG-134)
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// BUG-134: required (key presence) must not imply nonempty (.min(1)).
// A required string without an explicit OpenAPI minLength accepts "" so that
// partial PUT updates remain possible.
func TestGenerateZodSchemaRequiredNoMinLength(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{
		"memo": {Type: "string", Required: true},
	}
	code := generateZodSchema("CreatePhoto", fields)
	assertContains(t, code, "memo: z.string(),")
	if strings.Contains(code, ".min(") {
		t.Errorf("required-without-minLength string must not carry .min(): %s", code)
	}
	// required field is still not optional (key presence preserved).
	if strings.Contains(code, ".optional()") {
		t.Errorf("required field must not be .optional(): %s", code)
	}
}
