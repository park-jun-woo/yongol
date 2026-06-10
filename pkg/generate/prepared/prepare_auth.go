//ff:func feature=generate type=util control=sequence
//ff:what AuthFor — manifest.backend.auth 파생 (Mode 기본값 해석 + type=jwt → bearer 매핑)

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// AuthFor collapses the raw manifest.Auth into prepared.Auth with Mode
// already defaulted via manifest.Auth.ResolvedMode(). Every emitter
// that previously read a.Mode or a.ResolvedMode() now reads
// state.Auth.Mode, eliminating the BUG-009 class of inconsistency.
//
// Phase002 (BUG-014) — when manifest.auth.mode is empty but auth.type
// is "jwt", Mode resolves to "bearer" rather than the Phase020 default
// "cookie". The jwt type implies bearer transport; resolving to cookie
// here would emit CSRF middleware and cookie-first extractToken paths
// that a JWT-only project never uses.
//
// Phase001 debug01 (BUG-013) — CsrfRequired is derived from the
// resolved Mode: cookie/hybrid → true, bearer (and auth-absent) →
// false. Emitters gate all CSRF code on this flag so JWT-only
// projects emit no CSRF block, no csrf.go file, and no
// middleware.Csrf import.
//
// Exported since Phase004 (plans/stml/auth-flow): the react emitter
// consumes the same derived mode so the frontend client (bearer vs
// cookie/CSRF) can never diverge from the backend middleware it talks to.
func AuthFor(fs *yongol.Fullstack) Auth {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return Auth{}
	}
	a := fs.Manifest.Backend.Auth
	mode := a.ResolvedMode()
	if a.Mode == "" && a.Type == "jwt" {
		mode = "bearer"
	}
	csrfRequired := mode == "cookie" || mode == "hybrid"
	return Auth{
		Present:      true,
		Mode:         mode,
		CsrfRequired: csrfRequired,
		Raw:          a,
	}
}
