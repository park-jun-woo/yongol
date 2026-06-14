//ff:func feature=validate type=util control=selection topic=openapi-ddl
//ff:what responseValueBaseVar — @response 필드 값 표현식의 base var(첫 점 앞) 추출, 리터럴이면 ""

package openapi_ddl

import "strings"

// responseValueBaseVar returns the base variable of a @response field value
// expression. The SSaC @response value grammar is closed (corpus-verified): a
// value is either ① a bare var, ② a single-level dotted "var.Field", or ③ a
// literal. The base var is the token before the first '.', or the whole token
// when bare. Literals (quoted / numeric / bool / null) yield "".
func responseValueBaseVar(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	switch c := value[0]; {
	case c == '"', c == '\'', c == '-', c >= '0' && c <= '9':
		return ""
	}
	if value == "true" || value == "false" || value == "null" || value == "nil" {
		return ""
	}
	if i := strings.IndexByte(value, '.'); i > 0 {
		return value[:i]
	}
	return value
}
