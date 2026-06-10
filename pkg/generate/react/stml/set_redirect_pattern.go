//ff:func feature=stml-gen type=util control=sequence
//ff:what ActionBlock의 페이지명 data-redirect에 대상 페이지의 해석 라우트 패턴(RedirectPattern)을 설정한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// setRedirectPattern resolves a page-name data-redirect to the target
// page's resolved route pattern (stml.RoutePaths first pattern, supplied
// as GenerateOptions.RoutePatterns) and stores it in
// ActionBlock.RedirectPattern. It mirrors populateLinkTargets: the parser
// leaves the field empty, codegen resolves it (page-flow Phase008). Static
// "/"-prefixed redirects and a nil routePatterns map leave the field
// empty — renderRedirectNavigate then falls back to "/<page-name>".
func setRedirectPattern(a *stmlparser.ActionBlock, routePatterns map[string]string) {
	if a.Redirect == "" || strings.HasPrefix(a.Redirect, "/") || routePatterns == nil {
		return
	}
	a.RedirectPattern = routePatterns[a.Redirect]
}
