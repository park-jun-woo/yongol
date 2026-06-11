//ff:func feature=gen-react type=util control=sequence
//ff:what tsSingleQuote — 문자열을 TS 단일따옴표 리터럴로 이스케이프 (백슬래시/따옴표/개행)

package react

import "strings"

// tsSingleQuote renders a Go string as a TypeScript single-quoted string
// literal, escaping backslashes, single quotes and newlines — sitemap
// labels are user-authored text and flow into generated TS sources
// (breadcrumbs.ts), so they must never break the literal.
func tsSingleQuote(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `'`, `\'`, "\n", `\n`)
	return "'" + r.Replace(s) + "'"
}
