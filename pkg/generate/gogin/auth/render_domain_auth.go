//ff:func feature=gen-gogin type=generator control=sequence
//ff:what renderDomainAuthMiddleware — 도메인 strict 미들웨어 템플릿 placeholder 치환

package auth

import "strings"

// renderDomainAuthMiddleware fills a per-domain strict-middleware template
// (bearerAuthStrictDomainTemplate / cookieAuthStrictDomainTemplate) with the
// domain's api import, api package qualifier, module path, and PascalCase
// title. Single source of substitution so bearer and cookie emitters stay in
// lockstep with the placeholder set (__API_IMPORT__ / __MODULE__ / __API_PKG__
// / __TITLE__).
func renderDomainAuthMiddleware(tmpl, apiImport, modulePath, apiPkg, title string) string {
	r := strings.NewReplacer(
		"__API_IMPORT__", apiImport,
		"__MODULE__", modulePath,
		"__API_PKG__", apiPkg,
		"__TITLE__", title,
	)
	return r.Replace(tmpl)
}
