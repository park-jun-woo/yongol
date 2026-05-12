//ff:func feature=stml-gen type=generator control=sequence
//ff:what 단일 필드 제약조건에서 zod 유효성 체인을 생성한다
package stml

import (
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// zodChain builds a zod validation chain for a single field constraint.
func zodChain(fc oapiparser.FieldConstraint) string {
	if len(fc.Enum) > 0 {
		return zodEnumChain(fc)
	}

	var base string
	if fc.Type == "array" {
		base = zodBaseTypeArray(fc.ItemType)
	} else {
		base = zodBaseType(fc.Type)
	}
	parts := []string{base}
	parts = appendStringValidations(parts, fc)
	parts = appendNumericValidations(parts, fc)

	if !fc.Required {
		parts = append(parts, ".optional()")
	}

	return strings.Join(parts, "")
}
