//ff:func feature=ssacmeta type=util control=iteration dimension=1
//ff:what lookupPath — dot-path 로 map[string]any 를 내려가며 값을 조회

package ssacmeta

import "strings"

// lookupPath descends into a nested map[string]any using a dot-separated
// path and returns the terminal value. Returns (nil, false) the moment a
// path component is missing or a non-map is encountered.
func lookupPath(m map[string]any, path string) (any, bool) {
	parts := strings.Split(path, ".")
	var cur any = m
	for _, p := range parts {
		mm, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = mm[p]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}
