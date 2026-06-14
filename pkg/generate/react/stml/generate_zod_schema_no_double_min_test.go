//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema minLength=1 시 .min(1) 단일 방출 검증 (중복 없음)
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaMinLengthNoDoubleMin(t *testing.T) {
	minLen := 1
	fields := map[string]oapiparser.FieldConstraint{
		"name": {Type: "string", Required: true, MinLength: &minLen},
	}
	code := generateZodSchema("Create", fields)
	count := strings.Count(code, ".min(1)")
	if count != 1 {
		t.Errorf("expected exactly 1 .min(1), got %d\n--- code ---\n%s", count, code)
	}
}
