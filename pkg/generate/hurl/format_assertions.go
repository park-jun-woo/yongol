//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what formatAssertions — assertion 목록을 hurl [Asserts] 블록 텍스트로 변환
package hurl

import "strings"

// formatAssertions renders assertions as hurl [Asserts] block text.
func formatAssertions(asserts []string) string {
	if len(asserts) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("[Asserts]\n")
	for _, a := range asserts {
		b.WriteString(a + "\n")
	}
	return b.String()
}
