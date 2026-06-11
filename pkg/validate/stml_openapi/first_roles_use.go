//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what firstRolesUse — 사이트맵의 첫 data-roles 사용 노드의 위치 경로 (없으면 "")

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// firstRolesUse returns the position path of the first sitemap node (in
// document order) carrying a data-roles allowlist, or "" when none does —
// the TM-47 gate and the spot its messages point at.
func firstRolesUse(sm *stml.SitemapSpec) string {
	for _, e := range collectSitemapEntries(sm) {
		if len(e.Node.Roles) > 0 {
			return e.Path
		}
	}
	return ""
}
