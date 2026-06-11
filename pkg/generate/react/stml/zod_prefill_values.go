//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what zod 폼의 prefill values 완전 객체 — 전 필드 방출, 응답 교집합만 data 참조·나머지는 타입별 빈 리터럴
package stml

import (
	"fmt"
	"sort"
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// zodPrefillValues builds the value lines for a zod form's prefill object. Every
// schema field is emitted (the `values` exact-type requirement), sorted the same
// way generateZodSchema sorts. A field present in the prefill 2xx response
// (respFields) is mapped to `<dataVar>.<field> ?? <empty>` (coalescing satisfies
// the exact type when the response field is optional/nullable); a field absent
// from the response gets the type-appropriate empty literal only, never a data
// reference (the strongly-typed Res<K> would not type-check it).
func zodPrefillValues(fields map[string]oapiparser.FieldConstraint, dataVar string, respFields map[string]oapiparser.FieldTypeInfo) string {
	names := make([]string, 0, len(fields))
	for n := range fields {
		names = append(names, n)
	}
	sort.Strings(names)

	lines := make([]string, 0, len(names))
	for _, n := range names {
		lit := prefillEmptyLiteral(fields[n])
		if _, ok := respFields[n]; ok {
			lines = append(lines, fmt.Sprintf("          %s: %s.%s ?? %s,", n, dataVar, n, lit))
			continue
		}
		lines = append(lines, fmt.Sprintf("          %s: %s,", n, lit))
	}
	return strings.Join(lines, "\n")
}
