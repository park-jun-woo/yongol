//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-24 — cookie 모드인데 auth.* 캡처 또는 frontend.auth 선언이 존재 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm24CookieModeCaptureConflict checks mode consistency (the mode side of
// XOH-07): when backend.auth.mode resolves to "cookie", tokens live in
// httpOnly cookies and there is nothing for the frontend to capture, so an
// auth.* data-capture or a manifest frontend.auth block is a stale or
// contradictory declaration. Each offending declaration yields a WARNING.
func tm24CookieModeCaptureConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if backendAuthMode(fs) != "cookie" {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, c := range collectPageCaptures(fs.STMLPages) {
		diags = append(diags, diagnostic.Diagnostic{
			File:    c.File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[TM-24] backend.auth.mode resolves to \"cookie\" but %q captures %q into %q — httpOnly cookies cannot be captured by the frontend", c.File, c.Bind.RespField, c.Bind.Sink),
			Advice:  "Remove the data-capture declaration (cookie mode needs no token capture), or set backend.auth.mode to \"bearer\"",
		})
	}
	if fs.Manifest != nil && fs.Manifest.Frontend.Auth != nil {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[TM-24] backend.auth.mode resolves to \"cookie\" but manifest.yaml declares frontend.auth — the block only applies to bearer mode",
			Advice:  "Remove the frontend.auth block (cookie mode stores tokens in httpOnly cookies), or set backend.auth.mode to \"bearer\"",
		})
	}
	return diags
}
