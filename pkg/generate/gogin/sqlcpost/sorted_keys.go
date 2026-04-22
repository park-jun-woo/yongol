//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sortedKeys — map[string]string 의 키를 정렬된 []string로 반환

package sqlcpost

import "sort"

// sortedKeys returns m's keys in ascending order. Used by renderLogValueFile
// as a deterministic fallback when a DDL table provides no explicit
// ColumnOrder.
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
