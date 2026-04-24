//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what XOH-07 — cookie 인증 모드에서 mutating 요청은 X-CSRF-Token 헤더를 포함해야 함

package hurl_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh07CSRFOnMutation enforces XOH-07 (WARNING): when
// manifest.backend.auth.mode resolves to "cookie" (or "hybrid"), every
// mutating hurl request (POST/PUT/PATCH/DELETE) must carry an
// `X-CSRF-Token` header. Without it the generated middleware rejects
// the request at runtime (403), and the smoke test can never succeed.
//
// Auth-issuing endpoints (login / register) are exempt — the CSRF
// token is typically captured from their response headers.
func xoh07CSRFOnMutation(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	mode := fs.Manifest.Backend.Auth.ResolvedMode()
	if mode != "cookie" && mode != "hybrid" {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		if !isMutating(e.Method) {
			continue
		}
		if isAuthPath(e.Path) {
			continue
		}
		if hasCSRFHeader(e.Headers) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XOH-07] " + e.Method + " " + e.Path + " missing X-CSRF-Token (cookie auth mode requires CSRF)",
			Advice:  "Capture the CSRF token on login and add `X-CSRF-Token: {{csrf}}` to this request",
		})
	}
	return diags
}

// isMutating reports whether an HTTP method causes server state to
// change — these are the methods covered by the CSRF middleware.
func isMutating(method string) bool {
	switch strings.ToUpper(method) {
	case "POST", "PUT", "PATCH", "DELETE":
		return true
	}
	return false
}

// hasCSRFHeader returns true when any header is named X-CSRF-Token
// (case-insensitive).
func hasCSRFHeader(headers []hurl.HurlHeader) bool {
	for _, h := range headers {
		if strings.EqualFold(h.Name, "X-CSRF-Token") {
			return true
		}
	}
	return false
}
