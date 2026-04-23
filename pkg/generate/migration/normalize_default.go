//ff:func feature=migration type=parser control=selection
//ff:what NormalizeDefault — DEFAULT 절 문자열을 canonical 표현으로 정규화
package migration

import "strings"

// NormalizeDefault rewrites a PostgreSQL DEFAULT expression into a
// canonical form so that equivalent expressions compare equal:
//
//   'foo'::text        -> 'foo'
//   'foo'              -> 'foo'
//   NOW()              -> CURRENT_TIMESTAMP
//   now()              -> CURRENT_TIMESTAMP
//   CURRENT_TIMESTAMP  -> CURRENT_TIMESTAMP
//   TRUE / true        -> TRUE
//   FALSE / false      -> FALSE
//   0::integer         -> 0
//   nextval('s'::regclass) -> nextval('s')
//
// Empty input returns "" (meaning "no default").
func NormalizeDefault(raw string) string {
	s := strings.TrimSpace(raw)
	if s == "" {
		return ""
	}
	s = stripDefaultCasts(s)

	upper := strings.ToUpper(s)
	switch {
	case upper == "NOW()" || upper == "CURRENT_TIMESTAMP":
		return "CURRENT_TIMESTAMP"
	case upper == "TRUE":
		return "TRUE"
	case upper == "FALSE":
		return "FALSE"
	case upper == "NULL":
		return "NULL"
	case strings.HasPrefix(upper, "NEXTVAL("):
		return normalizeNextval(s)
	}
	return s
}
