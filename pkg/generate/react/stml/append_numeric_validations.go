//ff:func feature=stml-gen type=util control=sequence
//ff:what 숫자 타입 필드의 zod 유효성 규칙을 추가한다
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// appendNumericValidations appends numeric-specific zod validations.
func appendNumericValidations(parts []string, fc oapiparser.FieldConstraint) []string {
	if fc.Type != "integer" && fc.Type != "number" {
		return parts
	}
	if fc.Minimum != nil {
		parts = append(parts, fmt.Sprintf(".min(%s)", formatFloat(*fc.Minimum)))
	}
	if fc.Maximum != nil {
		parts = append(parts, fmt.Sprintf(".max(%s)", formatFloat(*fc.Maximum)))
	}
	return parts
}
