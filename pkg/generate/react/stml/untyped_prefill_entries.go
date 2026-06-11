//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what 무타입 폼의 prefill values 항목 — 응답에 실재하는 필드명만 정렬·중복제거하여 data 참조로 방출
package stml

import (
	"fmt"
	"sort"
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// untypedPrefillEntries builds the value lines for an untyped form's prefill
// object. Only fields present in the prefill 2xx response (respFields) are
// emitted — an untyped form's `values` may be partial, so absent fields are
// dropped rather than coalesced. Returns "" when no field overlaps the response.
func untypedPrefillEntries(fields []stmlparser.FieldBind, dataVar string, respFields map[string]oapiparser.FieldTypeInfo) string {
	seen := make(map[string]bool, len(fields))
	names := make([]string, 0, len(fields))
	for _, f := range fields {
		if f.Name == "" || seen[f.Name] {
			continue
		}
		if _, ok := respFields[f.Name]; !ok {
			continue
		}
		seen[f.Name] = true
		names = append(names, f.Name)
	}
	sort.Strings(names)

	lines := make([]string, 0, len(names))
	for _, n := range names {
		lines = append(lines, fmt.Sprintf("          %s: %s.%s,", n, dataVar, n))
	}
	return strings.Join(lines, "\n")
}
