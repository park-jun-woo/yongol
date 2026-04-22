//ff:func feature=gen-gogin type=util control=sequence
//ff:what countBraceDelta — 라인에서 '{' 개수에서 '}' 개수를 뺀 depth 변화량 반환 (주석 제외)

package ffannot

import "strings"

// countBraceDelta returns opens-minus-closes for a single line, ignoring any
// text after a // comment marker so braces inside comments don't shift depth.
func countBraceDelta(line string) int {
	code := line
	if idx := strings.Index(code, "//"); idx >= 0 {
		code = code[:idx]
	}
	return strings.Count(code, "{") - strings.Count(code, "}")
}
