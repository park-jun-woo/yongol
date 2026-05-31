//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockNames — MainBlock 슬라이스에서 Name 집합 추출 (테스트 헬퍼)
package boot

func blockNames(blocks []MainBlock) map[string]bool {
	out := make(map[string]bool, len(blocks))
	for _, b := range blocks {
		out[b.Name] = true
	}
	return out
}
