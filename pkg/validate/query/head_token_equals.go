//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what headTokenEquals — RawType 의 head (배열/파라미터 제거 후) 와 want 를 case-insensitive 비교

package query

import "strings"

// headTokenEquals upper-cases the head token of raw (stripping "[]" and
// "(...)") and compares case-insensitively to want. Shared by Q-12 ~
// Q-18 column filters.
func headTokenEquals(raw, want string) bool {
	t := strings.TrimSpace(raw)
	if strings.HasSuffix(t, "[]") {
		t = strings.TrimSpace(strings.TrimSuffix(t, "[]"))
	}
	if idx := strings.Index(t, "("); idx > 0 {
		t = strings.TrimSpace(t[:idx])
	}
	return strings.EqualFold(t, want)
}
