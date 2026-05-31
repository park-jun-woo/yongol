//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

func firstTag(frag string) string {
	i := 1 // skip '<'
	j := i
	for j < len(frag) {
		c := frag[j]
		if c == ' ' || c == '>' || c == '/' {
			break
		}
		j++
	}
	return frag[i:j]
}
