//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what countPureLines — lines[start..end) 중 공백/주석 제외 라인 수 반환

package qcheck

import "strings"

// countPureLines returns the number of non-blank, non-line-comment lines in
// lines[start..end-1]. Indices outside the slice are silently clamped so
// callers don't need to guard against boundary cases.
func countPureLines(lines []string, start, end int) int {
	pure := 0
	for i := start; i < end-1 && i < len(lines); i++ {
		trim := strings.TrimSpace(lines[i])
		if trim == "" || strings.HasPrefix(trim, "//") {
			continue
		}
		pure++
	}
	return pure
}
