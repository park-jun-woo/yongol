//ff:type feature=generate type=model
//ff:what Auth — 인증 설정 파생 상태 (Mode 기본값 해석 완료)

package prepared

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// Auth carries the derived auth configuration used by codegen. Mode is
// always one of "cookie", "bearer", or "hybrid" — the Phase020
// empty-string default ("cookie") is already applied by prepared.AuthFor
// so emitters never need to reach for manifest.Auth.ResolvedMode()
// again.
//
// Present when manifest.backend.auth is declared; Present==false means
// the generator should emit no auth-dependent wiring.
//
// CsrfRequired is the derived "is the build-time resolved mode
// cookie/hybrid?" flag (Phase001 debug01 / BUG-013). It is true only when
// the resolved Mode is "cookie" or "hybrid".
//
// BUG-116 / Phase-B1 — this flag no longer gates backend CSRF *emission*.
// Because the generated authMode() switch lets BACKEND_AUTH_MODE flip a
// bearer build to cookie/hybrid at runtime, the gogin emitters now mount
// CSRF whenever auth is Present and gate it at runtime inside the
// middleware. CsrfRequired survives as the build-time-default signal still
// consumed by the react client plan (resolve_api_client_plan.go) and is
// the natural anchor for the Phase-B2 validate guard.
type Auth struct {
	Present      bool
	Mode         string
	CsrfRequired bool
	// Raw keeps the underlying manifest.Auth pointer for fields that
	// have not yet been migrated into Auth (claims, TTLs, cookie
	// config, csrf, ...). Callers SHOULD prefer explicit fields once
	// available; Raw exists to keep the Phase001 migration incremental.
	Raw *manifest.Auth
}
