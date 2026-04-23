//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockAuthInit — auth.Configure + refresh_tokens DDL + RefreshStore 주입 (라우트 마운트 없음)

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// blockAuthInit emits the main.go block that wires ssac/pkg/auth into the
// running server. It replaces the deprecated blockAuthRefresh (Phase009):
//
//  1. auth.Configure — persists SecretEnv + AccessTTL + RefreshTTL so
//     subsequent IssueToken / VerifyToken calls read os.Getenv(SecretEnv).
//  2. refresh_tokens DDL bootstrap — idempotent CREATE TABLE + GIN index
//     from auth.RefreshTokensDDL. Harmless on verifier-only services: the
//     table just stays empty when no Login / Refresh handler writes to it.
//  3. &auth.RefreshStore{DB: conn} — the rotation store instance; surfaced
//     as srv.RefreshStore so SSaC @call handlers (auth.RefreshToken,
//     auth.RefreshRotate, auth.Logout) can reach it via the Server struct.
//
// What is intentionally NOT done here (unlike the removed blockAuthRefresh):
//
//   - No r.POST("/auth/refresh", ...) mount. The canonical path is now
//     openapi.yaml → service/auth/refresh.ssac → StrictServer codegen, so
//     the route lives in the OpenAPI contract like every other endpoint.
//   - No hardcoded FixedRateLimit. A service that declares /auth/refresh in
//     its OpenAPI can attach rate-limit via the normal per-op middleware
//     chain (Phase009 Open Question #1 — for now the guard heuristic in
//     blockRegisterHandlers covers auth endpoints, see collect_active_blocks).
//
// Active iff prepared.Auth.Present. verifier-only services with
// backend.auth but no refresh/logout endpoints still get the DDL +
// RefreshStore wiring so BearerAuth's VerifyToken keeps working.
func blockAuthInit(a prepared.Auth, modulePath string) MainBlock {
	if !a.Present {
		return MainBlock{Name: "auth-init", Active: authAlwaysInactive}
	}
	cfg := resolveAuthInitConfig(a)
	return MainBlock{
		Name: "auth-init",
		// Active left nil: baseCandidateBlocks gates on a.Present before
		// calling blockAuthInit.
		Imports: authInitImports(modulePath),
		Lines:   authInitLines(cfg),
	}
}
