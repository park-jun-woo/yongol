//ff:func feature=gen-gogin type=util control=sequence
//ff:what parseRawType — Column.RawType 토큰을 rawTypeInfo (head/param/array/multi-token) 로 분해

package types

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// parseRawType splits Column.RawType into family head, parameter list,
// and array marker. The result is consumed by the dispatcher in types.go.
//
// Multi-word PostgreSQL type names (e.g. "DOUBLE PRECISION",
// "TIMESTAMP WITH TIME ZONE") are normalised to their canonical
// single-token alias ("FLOAT8", "TIMESTAMPTZ") via
// ddl.NormalizePGTypeHead so the downstream family matrices
// (floatHeads / pgtype_timestamp / stringHeads etc.) can stay keyed by
// the single-token form. MultiToken is retained as an informational
// flag for diagnostics; dispatch no longer routes on it.
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
	if strings.Contains(upper, " ") {
		info.MultiToken = true
	}
	info.Head = ddl.NormalizePGTypeHead(upper)
	return info
}
