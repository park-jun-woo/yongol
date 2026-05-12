//ff:func feature=validate type=util control=selection topic=ssac-statemachine
//ff:what inferStateLiteralType — 리터럴 값의 Go 타입 추론 (string/bool/nil/int/float64)

package ssac_statemachine

import "strings"

// inferStateLiteralType returns the Go type of a literal value.
func inferStateLiteralType(value string) string {
	switch {
	case len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"':
		return "string"
	case value == "true" || value == "false":
		return "bool"
	case value == "nil":
		return "nil"
	case len(value) > 0 && (value[0] >= '0' && value[0] <= '9' || value[0] == '-'):
		if strings.Contains(value, ".") {
			return "float64"
		}
		return "int"
	default:
		return ""
	}
}
