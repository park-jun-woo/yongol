//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what isVersionSuffix — "v2", "v10" 같은 Go 모듈 버전 suffix 패턴 여부 판정
package splitter

// isVersionSuffix reports whether s matches the Go module major-version
// suffix pattern ("v" followed by one or more digits). Used by
// importName to skip versioned path segments when resolving the
// reference identifier — e.g. foo/v2 → foo.
func isVersionSuffix(s string) bool {
	if len(s) < 2 || s[0] != 'v' {
		return false
	}
	for _, r := range s[1:] {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
