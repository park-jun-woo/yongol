//ff:func feature=stml-gen type=util control=selection
//ff:what FieldConstraint 타입별 빈 리터럴(prefill values 초기값/coalesce 기본) 반환
package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// prefillEmptyLiteral returns the empty TypeScript literal for a form field's
// type, used both as the value for a field absent from the prefill response and
// as the `??` coalesce fallback for a present-but-nullable field: string → ”,
// integer/number → 0, boolean → false, array → [], object → {}.
func prefillEmptyLiteral(fc oapiparser.FieldConstraint) string {
	switch fc.Type {
	case "integer", "number":
		return "0"
	case "boolean":
		return "false"
	case "array":
		return "[]"
	case "object":
		return "{}"
	default:
		return "''"
	}
}
