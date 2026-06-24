//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what Run — STML<->OpenAPI 교차 검증 진입점: 도메인 모드(XMO-11/12 집계)와 단일 사이트(전체 규칙) 분기

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run dispatches STML<->OpenAPI cross-validation. Single-site projects run the
// full rule set (runSingle, unchanged). Multi-domain projects carry their data
// under the per-domain fields, leaving the singular fs.OpenAPIDoc nil; runSingle
// would early-return on that, so domain mode routes to runDomained which keeps
// the domain-agnostic coverage rules XMO-11/12 (evaluated across all domains)
// and skips the single-site-only XMO-10.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.IsDomained() {
		return runDomained(fs)
	}
	return runSingle(fs)
}
