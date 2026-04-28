//ff:func feature=cli-init type=util control=sequence
//ff:what writeProjectIDBoundary — emit an underscore when rune i crosses a camel/acronym boundary

package cliinit

import (
	"strings"
	"unicode"
)

// writeProjectIDBoundary writes an underscore separator into out when rune at
// index i in runes sits on a camelCase or acronym boundary relative to its
// neighbours. The helper is isolated so the caller remains a single range
// loop (depth 2) instead of nesting `if` chains (depth 3, violates Q1).
func writeProjectIDBoundary(out *strings.Builder, runes []rune, i int, r rune) {
	if i == 0 {
		return
	}
	prev := runes[i-1]
	// camelCase boundary: lower|digit → Upper
	if (unicode.IsLower(prev) || unicode.IsDigit(prev)) && unicode.IsUpper(r) {
		out.WriteByte('_')
		return
	}
	// acronym boundary: Upper Upper lower → insert before lower
	if i+1 < len(runes) && unicode.IsUpper(prev) && unicode.IsUpper(r) && unicode.IsLower(runes[i+1]) {
		out.WriteByte('_')
	}
}
