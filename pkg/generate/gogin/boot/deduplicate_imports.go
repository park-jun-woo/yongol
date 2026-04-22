//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what deduplicateImports — 블록별 import 합산 + 중복 제거 + 정렬

package boot

import "sort"

// deduplicateImports collects imports from all blocks and returns a sorted,
// deduplicated list.
func deduplicateImports(blocks []MainBlock) []string {
	seen := make(map[string]bool)
	for _, b := range blocks {
		for _, imp := range b.Imports {
			seen[imp] = true
		}
	}
	var out []string
	for imp := range seen {
		out = append(out, imp)
	}
	sort.Strings(out)
	return out
}
