//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what appendMenuDynamicOps — 메뉴 트리를 재귀 순회하며 동적 그룹 op 를 문서 순서·중복 없이 누적

package react

// appendMenuDynamicOps walks the menu items depth-first, appending every
// dynamic group's data-fetch operationId in document order, skipping ones
// already seen — collectMenuDynamicOps' recursion body.
func appendMenuDynamicOps(items []sitemapMenuItem, seen map[string]bool, ops *[]string) {
	for _, it := range items {
		if it.Fetch != "" && !seen[it.Fetch] {
			seen[it.Fetch] = true
			*ops = append(*ops, it.Fetch)
		}
		appendMenuDynamicOps(it.Children, seen, ops)
	}
}
