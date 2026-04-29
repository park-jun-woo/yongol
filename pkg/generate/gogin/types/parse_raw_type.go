//ff:func feature=gen-gogin type=util control=sequence
//ff:what parseRawType — Column.RawType 토큰을 rawTypeInfo (head/param/array/multi-token) 로 분해

package types

import "strings"

// parseRawType splits Column.RawType into family head, parameter list,
// and array marker. The result is consumed by the dispatcher in types.go.
func parseRawType(raw string) rawTypeInfo {
	t := strings.TrimSpace(raw)
	info := rawTypeInfo{}
	if strings.HasSuffix(t, "[]") {
		info.IsArray = true
		t = strings.TrimSpace(strings.TrimSuffix(t, "[]"))
	}
	if idx := strings.Index(t, "("); idx > 0 {
		info.Param = strings.TrimSuffix(strings.TrimSpace(t[idx+1:]), ")")
		t = strings.TrimSpace(t[:idx])
	}
	upper := strings.ToUpper(t)
	info.Head = upper
	if strings.Contains(upper, " ") {
		info.MultiToken = true
	}
	return info
}
