//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what findFieldFold — 필드 목록에서 case-insensitive로 일치하는 정본 표기를 찾는다

package ssac

import "strings"

// findFieldFold reports the canonical field name in fields that matches f
// case-insensitively (initialism casing divergence). Returns ("", false) when
// no fold-match exists.
func findFieldFold(fields []string, f string) (string, bool) {
	for _, x := range fields {
		if strings.EqualFold(x, f) {
			return x, true
		}
	}
	return "", false
}
