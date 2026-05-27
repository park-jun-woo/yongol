//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what buildSortedPackages — 패키지 맵 → 정렬된 externalPackage 슬라이스 변환

package fastapi

import "sort"

// buildSortedPackages converts the collected map to a sorted slice.
func buildSortedPackages(pm map[string]map[string]bool) []externalPackage {
	result := make([]externalPackage, 0, len(pm))
	for pkg, ms := range pm {
		methods := make([]string, 0, len(ms))
		for m := range ms {
			methods = append(methods, m)
		}
		sort.Strings(methods)
		result = append(result, externalPackage{Name: pkg, Methods: methods})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}
