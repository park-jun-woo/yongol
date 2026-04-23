//ff:func feature=validate type=util control=selection topic=ddl-structural
//ff:what serialReplacement — SERIAL 계열 토큰(소문자)에 대응하는 IDENTITY 치환 문자열 반환

package ddl

// serialReplacement maps the lowercase SERIAL token to its canonical
// IDENTITY-based replacement used in D-8 advice.
func serialReplacement(lower string) string {
	switch lower {
	case "bigserial":
		return "BIGINT GENERATED ALWAYS AS IDENTITY"
	case "serial":
		return "INTEGER GENERATED ALWAYS AS IDENTITY"
	case "smallserial":
		return "SMALLINT GENERATED ALWAYS AS IDENTITY"
	}
	return "BIGINT GENERATED ALWAYS AS IDENTITY"
}
