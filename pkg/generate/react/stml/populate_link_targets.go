//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 페이지의 모든 LinkRef에 대상 페이지의 해석된 라우트 패턴(TargetPattern)을 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// populateLinkTargets walks every LinkRef in the page and sets
// TargetPattern to the target page's resolved route pattern
// (stml.RoutePaths first pattern, supplied by the caller as
// GenerateOptions.RoutePatterns). It mirrors populateEachKeyFields:
// the parser leaves the field empty, codegen resolves it (page-flow
// Phase007). When routePatterns is nil, no patterns are set and the
// renderer falls back to "/<page-name>".
func populateLinkTargets(page *stmlparser.PageSpec, routePatterns map[string]string) {
	if routePatterns == nil {
		return
	}
	for i := range page.Fetches {
		setLinkTargetsInChildren(page.Fetches[i].Children, routePatterns)
	}
	setLinkTargetsInChildren(page.Children, routePatterns)
}
