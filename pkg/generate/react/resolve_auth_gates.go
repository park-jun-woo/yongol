//ff:func feature=gen-react type=util control=sequence
//ff:what manifest에서 프론트 인증 게이트(hasAuth/bearerAuth/store 종류)를 해석한다

package react

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveAuthGates derives the frontend auth gates from the manifest
// (Phase003): hasAuth keys off backend.auth presence (the frontend session
// concern), not the Rego authz block. bearerAuth additionally requires the
// prepared mode to be "bearer" — it gates the session store and the Bearer
// interceptor together so neither can exist without the other (cookie mode
// emits neither). The mode comes from prepared.AuthFor, the same derivation
// the backend emitters use (Phase004 — including the BUG-014 jwt-without-
// mode → bearer rule), so frontend and backend can never disagree on the
// effective mode. authStore is the frontend.auth store kind
// (ResolvedStore() default applies when the block is absent).
func resolveAuthGates(fs *yongol.Fullstack) (hasAuth, bearerAuth bool, authStore string) {
	auth := prepared.AuthFor(fs)
	if !auth.Present {
		return false, false, ""
	}
	hasAuth = true
	bearerAuth = auth.Mode == "bearer"
	authStore = fs.Manifest.Frontend.Auth.ResolvedStore()
	return hasAuth, bearerAuth, authStore
}
