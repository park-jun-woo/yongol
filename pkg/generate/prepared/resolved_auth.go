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
// CsrfRequired is the derived "must the generator emit CSRF wiring?"
// flag (Phase001 debug01 / BUG-013). It is true only when the
// resolved Mode is "cookie" or "hybrid"; bearer-only (including
// JWT-typed-but-mode-unspecified) projects are CSRF-immune because
// the Authorization header is not auto-sent cross-origin. Emitters
// MUST gate CSRF block/middleware emission on this field, not on the
// raw manifest.
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
