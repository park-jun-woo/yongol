//ff:func feature=gen-react type=util control=sequence
//ff:what manifest에서 프론트 인증 게이트(hasAuth/bearerAuth/store 종류)를 해석한다

package react

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveAuthGates derives the frontend auth gates from the manifest
// (Phase003): hasAuth keys off backend.auth presence (the frontend session
// concern), not the Rego authz block. bearerAuth additionally requires
// ResolvedMode() == "bearer" — it gates the session store and the Bearer
// interceptor together so neither can exist without the other (cookie mode
// emits neither). authStore is the frontend.auth store kind
// (ResolvedStore() default applies when the block is absent).
func resolveAuthGates(fs *yongol.Fullstack) (hasAuth, bearerAuth bool, authStore string) {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return false, false, ""
	}
	hasAuth = true
	bearerAuth = fs.Manifest.Backend.Auth.ResolvedMode() == "bearer"
	authStore = fs.Manifest.Frontend.Auth.ResolvedStore()
	return hasAuth, bearerAuth, authStore
}
