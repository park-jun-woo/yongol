//ff:func feature=funcspec type=test-helper control=iteration dimension=1
//ff:what keysOf — map[string][]Field 의 key 목록 반환 (테스트 진단 메시지용)

package funcspec

func keysOf(m map[string][]Field) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
