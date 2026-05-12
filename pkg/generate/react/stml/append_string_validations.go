//ff:func feature=stml-gen type=util control=sequence
//ff:what 문자열 타입 필드의 zod 유효성 규칙을 추가한다
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// appendStringValidations appends string-specific zod validations.
func appendStringValidations(parts []string, fc oapiparser.FieldConstraint) []string {
	if fc.Type != "string" && fc.Type != "" {
		return parts
	}
	if fc.Format == "email" {
		parts = append(parts, ".email()")
	}
	if fc.Required {
		parts = append(parts, ".min(1)")
	}
	if fc.MinLength != nil && *fc.MinLength > 0 {
		if !fc.Required || *fc.MinLength > 1 {
			parts = append(parts, fmt.Sprintf(".min(%d)", *fc.MinLength))
		}
	}
	if fc.MaxLength != nil {
		parts = append(parts, fmt.Sprintf(".max(%d)", *fc.MaxLength))
	}
	if fc.Pattern != "" {
		parts = append(parts, fmt.Sprintf(`.regex(new RegExp("%s"))`, fc.Pattern))
	}
	return parts
}
