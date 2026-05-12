//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what dotted path의 중간 세그먼트에 ?. (optional chaining)을 삽입한다
package stml

import "strings"

// optionalChainPath converts a dotted field path so that every "." after the
// first segment becomes "?." for TypeScript optional chaining.
//
//	"title"                 → "title"            (single field — no change)
//	"workflow.title"        → "workflow?.title"
//	"summary.credits_balance" → "summary?.credits_balance"
//	"a.b.c"                 → "a?.b?.c"
func optionalChainPath(path string) string {
	parts := strings.Split(path, ".")
	if len(parts) <= 1 {
		return path
	}
	var b strings.Builder
	b.WriteString(parts[0])
	for _, p := range parts[1:] {
		b.WriteString("?.")
		b.WriteString(p)
	}
	return b.String()
}
