//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what countLeadingSpaces — 문자열의 선행 공백 문자 수 반환

package agent

// countLeadingSpaces returns the number of leading space characters.
func countLeadingSpaces(s string) int {
	n := 0
	for _, c := range s {
		if c == ' ' {
			n++
		} else {
			break
		}
	}
	return n
}
