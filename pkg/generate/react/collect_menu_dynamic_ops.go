//ff:func feature=gen-react type=util control=sequence
//ff:what collectMenuDynamicOps — 메뉴 트리의 동적 그룹 data-fetch operationId 들을 문서 순서·중복 제거로 수집 (useQuery 방출 게이트)

package react

// collectMenuDynamicOps returns the data-fetch operationIds of every
// dynamic menu group in the tree (plans/stml/sitemap Phase007), document
// order, deduplicated — each becomes exactly one layout useQuery const,
// so two groups over the same operation share one query (and one cache
// entry, like two page fetches of the same op would). The walk lives in
// appendMenuDynamicOps (the collectMenuRoleSets split).
func collectMenuDynamicOps(items []sitemapMenuItem) []string {
	var ops []string
	appendMenuDynamicOps(items, map[string]bool{}, &ops)
	return ops
}
