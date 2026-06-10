//ff:func feature=validate type=helper control=iteration dimension=2 topic=stml-statemachine
//ff:what findRedirectTargetPage — data-redirect 값(정적 경로 또는 페이지명 참조)이 해석되는 첫 STML 페이지 반환

package stml_statemachine

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// findRedirectTargetPage resolves a data-redirect value to its target STML
// page, or nil when it resolves to no page (TM-26 in stml_openapi reports
// that case). A "/"-prefixed value is a static path matched against each
// page's route patterns (stml.RoutePaths); any other value is a page-name
// reference (page-flow Phase008) matched by page name — the more precise
// input, so TM-23's verdict only sharpens.
func findRedirectTargetPage(path string, pages []stml.PageSpec) *stml.PageSpec {
	if !strings.HasPrefix(path, "/") {
		for i := range pages {
			if pages[i].Name == path {
				return &pages[i]
			}
		}
		return nil
	}
	for i := range pages {
		for _, pattern := range stml.RoutePaths(pages[i]) {
			if stml.RouteMatchesPath(pattern, path) {
				return &pages[i]
			}
		}
	}
	return nil
}
