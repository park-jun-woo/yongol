//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sortedKeys — string-keyed map의 키를 정렬하여 반환

package state

import "sort"

// sortedKeys returns the keys of a string-keyed map in sorted order.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
