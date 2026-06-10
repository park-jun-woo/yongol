//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-38 — cookie 모드 + 값 없는 data-logout(클라이언트 단독 세션 종료 불가) / auth 없는 프로젝트의 data-logout (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm38LogoutMode judges every data-logout declaration against the
// project's effective auth mode (page-flow Phase010). The mode comes from
// prepared.AuthFor — the same derivation the layout emitter uses
// (including the BUG-014 jwt-without-mode → bearer rule), so the warning
// can never disagree with what generate actually emits. Two branches:
// (1) no backend.auth at all — there is no session to end, the
// declaration is meaningless and emission is skipped; (2) a non-bearer
// (cookie/hybrid) mode with a *valueless* data-logout — the session lives
// in an httpOnly cookie the client cannot clear, so only a server
// operation can end it; the emitted button would merely navigate away.
func tm38LogoutMode(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	auth := prepared.AuthFor(fs)
	var diags []diagnostic.Diagnostic
	for _, l := range fs.Layouts {
		if l.Logout == nil {
			continue
		}
		if !auth.Present {
			diags = append(diags, diagnostic.Diagnostic{
				File:    l.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[TM-38] data-logout in layout %q but the project declares no backend.auth — there is no session to end, so no logout UI is emitted", l.Name),
				Advice:  "Remove the data-logout declaration, or declare backend.auth in manifest.yaml",
			})
			continue
		}
		if auth.Mode != "bearer" && l.Logout.OperationID == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    l.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[TM-38] valueless data-logout in layout %q but backend.auth.mode resolves to %q — the session lives in an httpOnly cookie the client cannot clear", l.Name, auth.Mode),
				Advice:  "Declare the server logout operation: data-logout=\"<operationId>\" (only a server call can end a cookie session)",
			})
		}
	}
	return diags
}
