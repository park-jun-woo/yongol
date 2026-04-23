//ff:func feature=policy type=util control=sequence
//ff:what lineOfOffset — 문자열 내 byte offset 의 1-based 라인 번호 계산

package rego

import "strings"

// lineOfOffset returns the 1-based line number at byte offset off in string s.
// Returns 0 when off is out of range.
func lineOfOffset(s string, off int) int {
	if off < 0 || off > len(s) {
		return 0
	}
	return 1 + strings.Count(s[:off], "\n")
}
