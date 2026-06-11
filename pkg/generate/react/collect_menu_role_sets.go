//ff:func feature=gen-react type=util control=sequence
//ff:what collectMenuRoleSets — 메뉴 트리의 고유 role 목록들을 문서 순서로 수집 (상수명 기준 중복 제거)

package react

// collectMenuRoleSets returns every distinct data-roles allowlist of the
// menu tree in document order, deduplicated by the constant name
// rolesConstName derives — each set becomes one module-level
// `const ROLES_... = [...]` in the layout TSX (plans/stml/sitemap
// Phase005).
func collectMenuRoleSets(items []sitemapMenuItem) [][]string {
	var sets [][]string
	appendMenuRoleSets(items, map[string]bool{}, &sets)
	return sets
}
