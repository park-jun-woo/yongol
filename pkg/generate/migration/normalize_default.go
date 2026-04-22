//ff:func feature=migration type=parser control=selection
//ff:what NormalizeDefault — DEFAULT 절 문자열을 canonical 표현으로 정규화
package migration

import (
	"strings"
)

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

	// Strip ::type casts repeatedly (outermost only — we're not parsing
	// full SQL expressions).
	for {
		idx := strings.LastIndex(s, "::")
		if idx < 0 {
			break
		}
		// Only strip if the cast target looks like a bare identifier
		// (possibly quoted). Avoid touching casts inside inner parens.
		tail := strings.TrimSpace(s[idx+2:])
		if tail == "" {
			break
		}
		// If tail has unbalanced parens it's part of an inner expr — stop.
		if strings.Count(tail, "(") != strings.Count(tail, ")") {
			break
		}
		// Accept identifier or `"quoted"` as the cast target.
		if !looksLikeCastTarget(tail) {
			break
		}
		s = strings.TrimSpace(s[:idx])
	}

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
		// nextval('seqname'::regclass) — inner part may still have cast;
		// recursively normalize the argument.
		return normalizeNextval(s)
	}

	return s
}

func looksLikeCastTarget(s string) bool {
	if s == "" {
		return false
	}
	// Strip optional trailing array suffix `[]`.
	for strings.HasSuffix(s, "[]") {
		s = strings.TrimSpace(strings.TrimSuffix(s, "[]"))
	}
	// Handle `"quoted"`.
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return true
	}
	// Plain word made of letters / digits / underscore / spaces (for
	// "character varying" etc).
	for _, r := range s {
		if r == ' ' || r == '_' ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') {
			continue
		}
		return false
	}
	return true
}

func normalizeNextval(s string) string {
	// expect nextval(<inner>)
	open := strings.Index(s, "(")
	close := strings.LastIndex(s, ")")
	if open < 0 || close <= open {
		return s
	}
	inner := strings.TrimSpace(s[open+1 : close])
	// strip ::regclass and sibling casts
	if i := strings.LastIndex(inner, "::"); i >= 0 {
		tail := strings.TrimSpace(inner[i+2:])
		if looksLikeCastTarget(tail) {
			inner = strings.TrimSpace(inner[:i])
		}
	}
	return "nextval(" + inner + ")"
}
