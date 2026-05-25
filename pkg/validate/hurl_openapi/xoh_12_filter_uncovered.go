//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what filterUncovered — declared 중 coveredSet 에 없는 코드 필터링

package hurl_openapi

func filterUncovered(declared []string, coveredSet map[string]bool) []string {
	var missing []string
	for _, code := range declared {
		if coveredSet == nil || !coveredSet[code] {
			missing = append(missing, code)
		}
	}
	return missing
}
