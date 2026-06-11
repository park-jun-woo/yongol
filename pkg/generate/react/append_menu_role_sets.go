//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what appendMenuRoleSets — 메뉴 트리를 재귀 순회하며 미등록 role 목록을 sets 에 추가

package react

// appendMenuRoleSets walks the menu items depth-first, recording each
// data-roles allowlist whose constant name (rolesConstName) has not been
// seen yet — the document-order dedup behind collectMenuRoleSets.
func appendMenuRoleSets(items []sitemapMenuItem, seen map[string]bool, sets *[][]string) {
	for _, it := range items {
		if name := rolesConstName(it.Roles); len(it.Roles) > 0 && !seen[name] {
			seen[name] = true
			*sets = append(*sets, it.Roles)
		}
		appendMenuRoleSets(it.Children, seen, sets)
	}
}
