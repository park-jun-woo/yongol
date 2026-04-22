//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what inferLiteralType — seq.Inputs value 문자열에서 literal Go 타입을 추론

package ssac_func

import (
	"strconv"
	"strings"
)

// inferLiteralType classifies a raw @call input expression as a literal Go
// type, or returns "" when the expression references a variable/field
// (handled separately by resolveInputType). Rules:
//
//   "foo"     → "string" (quoted)
//   1, 42     → "int64"
//   1.5, 3.14 → "float64"
//   true/false → "bool"
//   nil       → "nil"
//   other     → "" (not a literal)
func inferLiteralType(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) && len(value) >= 2 {
		return "string"
	}
	if value == "true" || value == "false" {
		return "bool"
	}
	if value == "nil" {
		return "nil"
	}
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return "int64"
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float64"
	}
	return ""
}
