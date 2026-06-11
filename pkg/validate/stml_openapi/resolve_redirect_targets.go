//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what resolveRedirectTargets — data-redirect 값을 대상 페이지명들로 해석 (페이지명 참조 / "/" 인덱스 / 정적 경로 매칭)

package stml_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// resolveRedirectTargets resolves one data-redirect value to the page
// names it navigates to, for the reachability edge (c) of DESIGN §4.10 —
// the TM-26 target resolution reused for edge collection. A page-name
// reference resolves to that page; "/" resolves to the index pages
// (collectIndexPages); any other "/"-prefixed value is a static path
// matched against every page's resolved route patterns
// (stml.RouteMatchesPath over stml.RoutePaths). Unresolvable values yield
// no edge — TM-26 reports them, an edge into the void proves nothing.
func resolveRedirectTargets(redirect string, pages []stml.PageSpec, indexPages []string) []string {
	if redirect == "" {
		return nil
	}
	if redirect == "/" {
		return indexPages
	}
	if !strings.HasPrefix(redirect, "/") {
		if findPageByName(pages, redirect) == nil {
			return nil
		}
		return []string{redirect}
	}
	var out []string
	for _, p := range pages {
		for _, pattern := range stml.RoutePaths(p) {
			if stml.RouteMatchesPath(pattern, redirect) {
				out = append(out, p.Name)
				break
			}
		}
	}
	return out
}
