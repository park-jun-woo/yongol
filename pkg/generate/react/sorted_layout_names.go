//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 레이아웃 이름을 정렬하되 빈 문자열(flat)을 마지막에 배치한다

package react

import "sort"

// sortedLayoutNames returns layout names in sorted order, with "" (flat) last.
func sortedLayoutNames(grouped map[string][]stmlRoute) []string {
	names := make([]string, 0, len(grouped))
	hasFlat := false
	for name := range grouped {
		if name == "" {
			hasFlat = true
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)
	if hasFlat {
		names = append(names, "")
	}
	return names
}
