//ff:func feature=stml-gen type=util control=sequence
//ff:what tsSingleQuote — 문자열을 TS 단일따옴표 리터럴로 이스케이프 (백슬래시/따옴표/개행)

package stml

import "strings"

// tsSingleQuote renders a Go string as a TypeScript single-quoted string
// literal, escaping backslashes, single quotes and newlines — the
// document.title strings of plans/stml/sitemap Phase004 carry
// user-authored sitemap labels and manifest names, which must never break
// the generated literal.
func tsSingleQuote(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `'`, `\'`, "\n", `\n`)
	return "'" + r.Replace(s) + "'"
}
