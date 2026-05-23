//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what appendUnique — 중복 없이 문자열 슬라이스에 추가

package agent

func appendUnique(slice []string, s string) []string {
	for _, v := range slice {
		if v == s {
			return slice
		}
	}
	return append(slice, s)
}
