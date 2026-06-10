//ff:func feature=gen-react type=util control=sequence
//ff:what resolveAPIClientPlan — prepared 인증 모드에서 api.ts 분기 계획(bearer refresh / cookie CSRF) 도출

package react

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveAPIClientPlan derives the api.ts emission plan from the prepared
// auth mode (prepared.AuthFor — the same derivation the backend emitters
// use, so the client can never disagree with the middleware it talks to).
// Bearer mode additionally resolves the 401 refresh plan; cookie/hybrid
// mode resolves the double-submit CSRF cookie/header names from
// backend.auth.csrf with the runtime defaults (XSRF-TOKEN / X-XSRF-TOKEN).
//
// The only error path is an ambiguous structural refresh-op inference
// (resolveRefreshPlan) — a generate-time ERROR asking for an explicit
// frontend.auth.refresh_op declaration.
func resolveAPIClientPlan(fs *yongol.Fullstack) (apiClientPlan, error) {
	var plan apiClientPlan
	auth := prepared.AuthFor(fs)
	if !auth.Present {
		return plan, nil
	}
	if auth.Mode == "bearer" {
		plan.bearer = true
		rp, err := resolveRefreshPlan(fs)
		if err != nil {
			return plan, err
		}
		plan.refresh = rp
		return plan, nil
	}
	// "cookie" and "hybrid": the browser session rides httpOnly cookies.
	plan.cookie = true
	plan.csrf = auth.CsrfRequired && (auth.Raw == nil || auth.Raw.Csrf == nil || auth.Raw.Csrf.Enabled)
	var csrfCfg *manifest.CsrfConfig
	if auth.Raw != nil {
		csrfCfg = auth.Raw.Csrf
	}
	plan.csrfCookieName = csrfCfg.ResolvedCookieName()
	plan.csrfHeaderName = csrfCfg.ResolvedHeaderName()
	return plan, nil
}
