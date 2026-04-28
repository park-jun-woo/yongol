//ff:func feature=crosscheck type=test-helper control=iteration dimension=1 topic=scenario-check
//ff:what hasHeader — HurlHeader 슬라이스에서 이름이 일치하는 헤더 존재 여부

package hurl

// hasHeader reports whether hs contains at least one header with the
// given name (case-sensitive).
func hasHeader(hs []HurlHeader, name string) bool {
	for _, h := range hs {
		if h.Name == name {
			return true
		}
	}
	return false
}
