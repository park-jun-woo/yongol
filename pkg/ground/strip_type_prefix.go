//ff:func feature=rule type=util control=sequence
//ff:what stripTypePrefix — 타입 이름에서 슬라이스/포인터 prefix 제거 ("[]Workflow" → "Workflow")
package ground

import "strings"

// stripTypePrefix removes slice and pointer prefixes from a type name.
// e.g. "[]Workflow" → "Workflow", "*User" → "User".
func stripTypePrefix(t string) string {
	t = strings.TrimPrefix(t, "[]")
	t = strings.TrimPrefix(t, "*")
	return t
}
