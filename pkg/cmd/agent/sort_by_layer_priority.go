//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what sortByLayerPriority — fileGroup을 레이어 우선순위로 정렬

package agent

import "sort"

// sortByLayerPriority sorts file groups by the layer priority order.
func sortByLayerPriority(groups []fileGroup) {
	priority := make(map[layer]int)
	for i, l := range layerPriority {
		priority[l] = i
	}
	sort.SliceStable(groups, func(i, j int) bool {
		pi, ok := priority[groups[i].layer]
		if !ok {
			pi = 999
		}
		pj, ok := priority[groups[j].layer]
		if !ok {
			pj = 999
		}
		return pi < pj
	})
}
