//ff:func feature=migration type=test-helper control=iteration dimension=1
//ff:what mapKeys — 맵의 키를 슬라이스로 수집 (테스트 에러 메시지용)
package migration

func mapKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
