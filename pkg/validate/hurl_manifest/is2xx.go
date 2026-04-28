//ff:func feature=validate type=rule control=sequence topic=hurl-manifest
//ff:what is2xx — hurl status code 가 2xx 인지 판정 (미지정은 성공)

package hurl_manifest

// is2xx reports whether a hurl-asserted status code is 2xx. Empty means
// "not asserted" and is treated as success (hurl default behaviour).
func is2xx(code string) bool {
	if code == "" {
		return true
	}
	return len(code) == 3 && code[0] == '2'
}
