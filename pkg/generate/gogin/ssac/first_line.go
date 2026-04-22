//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what firstLine — 멀티라인 문자열에서 첫 비어있지 않은 줄 반환

package ssac

import "strings"

// firstLine trims to the first non-empty line (collapse multi-line descriptions).
func firstLine(s string) string {
	for _, part := range strings.Split(s, "\n") {
		if p := strings.TrimSpace(part); p != "" {
			return p
		}
	}
	return strings.TrimSpace(s)
}
