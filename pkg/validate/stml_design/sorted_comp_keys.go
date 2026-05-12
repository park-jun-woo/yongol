//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what sortedCompKeys — ComponentToken 맵 키를 정렬된 슬라이스로 반환
package stml_design

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// sortedCompKeys returns the keys of a ComponentToken map in sorted order.
func sortedCompKeys(m map[string]design.ComponentToken) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
