//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what matchMultiTokenHead — upper 의 선두 토큰들이 head 와 일치하면 소비 토큰 수, 아니면 0

package ddl

// matchMultiTokenHead returns the number of tokens consumed when the
// leading tokens of upper match head, ignoring a trailing parameter list
// "(...)" or comma on the final token. Returns 0 when no match.
func matchMultiTokenHead(upper []string, head []string) int {
	if len(upper) < len(head) {
		return 0
	}
	for i, want := range head {
		got := upper[i]
		if i == len(head)-1 {
			got = stripParamAndComma(got)
		}
		if got != want {
			return 0
		}
	}
	return len(head)
}
