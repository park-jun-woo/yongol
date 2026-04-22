//ff:func feature=validate type=util control=iteration dimension=1 topic=config-check
//ff:what sortedClaimFields — diagnostic 메시지용 claim field 정렬 목록 문자열

package ssac_manifest

import (
	"sort"
	"strings"
)

// sortedClaimFields returns a stable, comma-separated list of claim field
// names for diagnostic messages.
func sortedClaimFields(claims map[string]bool) string {
	names := make([]string, 0, len(claims))
	for k := range claims {
		names = append(names, k)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
