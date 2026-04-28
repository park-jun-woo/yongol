//ff:func feature=validate type=util control=iteration dimension=1 topic=manifest-infra
//ff:what sortedClaimFields — auth.Claims 의 field 이름을 정렬된 슬라이스로 반환

package manifest_ddl

import (
	"sort"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// sortedClaimFields returns every key of auth.Claims sorted lexicographically.
// XDN-03 / XDN-04 iterate in this order so diagnostic output is stable across
// runs (Go map iteration is randomised).
func sortedClaimFields(claims map[string]pmanifest.ClaimDef) []string {
	fields := make([]string, 0, len(claims))
	for f := range claims {
		fields = append(fields, f)
	}
	sort.Strings(fields)
	return fields
}
