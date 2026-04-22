//ff:func feature=gen-gogin type=util control=sequence
//ff:what extractHelperBodyLines — top-level func 선언 문자열에서 body 라인 추출 (signature/braces 제거)

package boot

import "strings"

// extractHelperBodyLines strips the `func name(...) T {` signature and trailing `}`
// from a helper declaration so ffannot.DetectControl can inspect the depth-0 lines.
// Returns an empty slice if the input lacks the expected braces.
func extractHelperBodyLines(decl string) []string {
	open := strings.Index(decl, "{")
	if open < 0 {
		return nil
	}
	inner := decl[open+1:]
	close := strings.LastIndex(inner, "}")
	if close < 0 {
		return nil
	}
	inner = inner[:close]
	return strings.Split(inner, "\n")
}
