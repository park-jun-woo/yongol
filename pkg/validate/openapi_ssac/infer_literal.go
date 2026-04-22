//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what inferLiteral — 리터럴 값을 Go 타입명으로 판별 (string/bool/nil/int64/float64)

package openapi_ssac

import (
	"strconv"
	"strings"
)

// inferLiteral mirrors ssac_func.inferLiteralType semantics; kept local to
// avoid cross-package cycle.
func inferLiteral(value string) string {
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
