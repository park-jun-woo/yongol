//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what formatCaptures — capture 목록을 hurl [Captures] 블록 텍스트로 변환
package hurl

import (
	"fmt"
	"strings"
)

// formatCaptures renders captures as hurl [Captures] block text.
func formatCaptures(caps []capture) string {
	if len(caps) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("[Captures]\n")
	for _, c := range caps {
		b.WriteString(fmt.Sprintf("%s: jsonpath \"%s\"\n", c.VarName, c.JSONPath))
	}
	return b.String()
}
