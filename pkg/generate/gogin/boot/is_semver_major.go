//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what isSemVerMajor — "v2", "v3" 같은 Go 모듈 버전 suffix 판정

package boot

import "unicode"

// isSemVerMajor reports whether s looks like "v2", "v3", … as used by Go
// module versioned path suffixes. Pure numeric suffix matters — "v3alpha"
// is a real package name.
func isSemVerMajor(s string) bool {
	if len(s) < 2 || (s[0] != 'v' && s[0] != 'V') {
		return false
	}
	for _, r := range s[1:] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
