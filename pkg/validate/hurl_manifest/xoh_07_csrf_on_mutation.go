//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what XOH-07 — cookie 인증 모드에서 mutating 요청은 manifest 해석 CSRF 헤더(기본 X-XSRF-TOKEN)를 포함해야 함

package hurl_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh07CSRFOnMutation enforces XOH-07 (WARNING): when
// manifest.backend.auth.mode resolves to "cookie" (or "hybrid"), every
// mutating hurl request (POST/PUT/PATCH/DELETE) must carry the
// manifest-resolved CSRF header (backend.auth.csrf.header_name, default
// `X-XSRF-TOKEN`). Without it the generated middleware rejects the
// request at runtime (403), and the smoke test can never succeed.
//
// Auth-issuing endpoints (login / register) are exempt — the CSRF
// cookie is typically captured after a preceding safe request.
func xoh07CSRFOnMutation(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	mode := fs.Manifest.Backend.Auth.ResolvedMode()
	if mode != "cookie" && mode != "hybrid" {
		return nil
	}
	csrf := fs.Manifest.Backend.Auth.Csrf
	headerName := csrf.ResolvedHeaderName()
	cookieName := csrf.ResolvedCookieName()
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		if !shouldCheckCSRF(e, headerName) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XOH-07] " + e.Method + " " + e.Path + " missing " + headerName + " (cookie auth mode requires CSRF)",
			Advice:  "Capture the CSRF cookie after a safe request (`csrf: cookie \"" + cookieName + "\"`) and add `" + headerName + ": {{csrf}}` to this request",
		})
	}
	return diags
}
