//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what nolintMatches — 단일 `//` 주석이 `nolint:<rule>` 지시를 포함하는지 판정

package contract

import "strings"

// nolintMatches parses a single `//` comment text and reports whether
// it carries the requested nolint rule. rule must be lowercase-trimmed;
// the caller (hasNolint) normalises it once per query.
//
// Accepted forms (case-insensitive on the raw comment):
//
//	// nolint:panic
//	// nolint:prv-12
//	// nolint:prv-12,prv-17
//	//nolint:prv-14 extra words
func nolintMatches(text, rule string) bool {
	lower := strings.ToLower(text)
	idx := strings.Index(lower, "nolint:")
	if idx < 0 {
		return false
	}
	rest := strings.TrimSpace(lower[idx+len("nolint:"):])
	if rest == "" {
		return false
	}
	// Stop at first whitespace — tokens follow the directive word.
	if sp := strings.IndexAny(rest, " \t"); sp >= 0 {
		rest = rest[:sp]
	}
	for _, tok := range strings.Split(rest, ",") {
		tok = strings.TrimSpace(tok)
		if tok == rule {
			return true
		}
	}
	return false
}
