//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-24 — cookie 모드인데 토큰 캡처 또는 토큰 키를 가진 frontend.auth 선언이 존재 (WARNING; auth.claims.*·role_field 전용 블록은 예외)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm24CookieModeCaptureConflict checks mode consistency (the mode side of
// XOH-07): when backend.auth.mode resolves to "cookie", tokens live in
// httpOnly cookies and there is nothing for the frontend to capture, so an
// auth.token/auth.refresh data-capture or a manifest frontend.auth block
// with token keys is a stale or contradictory declaration. Each offending
// declaration yields a WARNING. Two exemptions (plans/stml/sitemap
// Phase005) — the rule's blocking ground is "httpOnly cookies cannot be
// captured by JS", which does not apply to either:
//   - auth.claims.* captures read the login response *body*, not the
//     cookie, so they are first-class in cookie mode (the data-roles menu
//     filter depends on them; TM-47 even requires one).
//   - a role_field-only frontend.auth block (manifest.FrontendAuth.
//     RoleFieldOnly — the same predicate XON-60 consumes, so the two
//     rules' judgments cannot drift) declares no token contract at all.
func tm24CookieModeCaptureConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if backendAuthMode(fs) != "cookie" {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, c := range collectPageCaptures(fs.STMLPages) {
		if _, claims := stml.ClaimsSinkName(c.Bind.Sink); claims {
			continue // response-body claim, not a cookie-held token
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    c.File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[TM-24] backend.auth.mode resolves to \"cookie\" but %q captures %q into %q — httpOnly cookies cannot be captured by the frontend", c.File, c.Bind.RespField, c.Bind.Sink),
			Advice:  "Remove the data-capture declaration (cookie mode needs no token capture), or set backend.auth.mode to \"bearer\"",
		})
	}
	if fs.Manifest != nil && fs.Manifest.Frontend.Auth != nil && !fs.Manifest.Frontend.Auth.RoleFieldOnly() {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[TM-24] backend.auth.mode resolves to \"cookie\" but manifest.yaml declares frontend.auth with token keys — those keys only apply to bearer mode",
			Advice:  "Remove the token-related keys (token_field / refresh_field / refresh_op / store) from frontend.auth — role_field may stay for the data-roles menu wiring — or set backend.auth.mode to \"bearer\"",
		})
	}
	return diags
}
