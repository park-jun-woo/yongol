//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what computeDepth — FK 부모 체인 재귀 탐색으로 의존 깊이 계산
package hurl

// computeDepth recursively computes the FK dependency depth for a table.
func computeDepth(name string, parentMap map[string][]string, cache map[string]int) int {
	if d, ok := cache[name]; ok {
		return d
	}
	cache[name] = 0
	maxD := 0
	for _, p := range parentMap[name] {
		if d := computeDepth(p, parentMap, cache) + 1; d > maxD {
			maxD = d
		}
	}
	cache[name] = maxD
	return maxD
}
