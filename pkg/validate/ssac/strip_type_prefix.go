//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what stripTypePrefix — "[]" 접두어와 패키지 prefix 제거

package ssac

import "strings"

// stripTypePrefix removes "[]" and any package prefix from a type name.
func stripTypePrefix(t string) string {
	t = strings.TrimPrefix(t, "[]")
	if i := strings.LastIndexByte(t, '.'); i >= 0 {
		t = t[i+1:]
	}
	return t
}
