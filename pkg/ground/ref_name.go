//ff:func feature=rule type=util control=sequence
//ff:what refName — "#/components/schemas/X" 의 마지막 path 세그먼트 반환
package ground

import "strings"

// refName returns the last path segment of "#/components/schemas/X" → "X".
func refName(ref string) string {
	i := strings.LastIndex(ref, "/")
	if i < 0 {
		return ref
	}
	return ref[i+1:]
}
