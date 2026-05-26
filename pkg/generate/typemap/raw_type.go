//ff:func feature=gen-typemap type=util control=sequence
//ff:what ParseRawType — DDL 원본 타입 문자열을 RawTypeInfo (head/param/array/multi-token) 로 정규화

package typemap

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// ParseRawType splits a DDL column type string into its constituent parts:
// the normalised head token, optional parameter, array marker, and
// multi-token flag.
func ParseRawType(raw string) RawTypeInfo {
	t := strings.TrimSpace(raw)
	info := RawTypeInfo{}
	if strings.HasSuffix(t, "[]") {
		info.IsArray = true
		t = strings.TrimSpace(strings.TrimSuffix(t, "[]"))
	}
	if idx := strings.Index(t, "("); idx > 0 {
		info.Param = strings.TrimSuffix(strings.TrimSpace(t[idx+1:]), ")")
		t = strings.TrimSpace(t[:idx])
	}
	upper := strings.ToUpper(t)
	if strings.Contains(upper, " ") {
		info.MultiToken = true
	}
	info.Head = ddl.NormalizePGTypeHead(upper)
	return info
}
