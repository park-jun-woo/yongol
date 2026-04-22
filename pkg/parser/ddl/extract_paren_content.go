//ff:func feature=manifest type=util control=sequence
//ff:what extractParenContent — "(content)" 에서 content 추출
package ddl

import "strings"

func extractParenContent(s string) string {
	open := strings.Index(s, "(")
	close := strings.Index(s, ")")
	if open < 0 || close < 0 || close <= open {
		return ""
	}
	return strings.TrimSpace(s[open+1 : close])
}
