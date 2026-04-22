//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockAuthInit — auth.Configure + refresh_tokens DDL + RefreshStore 주입 (라우트 마운트 없음)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
// Active iff manifest.backend.auth is configured. verifier-only services
// with backend.auth but no refresh/logout endpoints still get the DDL +
// RefreshStore wiring so BearerAuth's VerifyToken keeps working.
func blockAuthInit(fs *yongol.Fullstack, modulePath string) MainBlock {
	secretEnv := "JWT_SECRET"
	accessTTL := "15m"
	refreshTTL := "168h"
	mode := "cookie"
	sameSite := "Lax"
	accessName := "__Host-access_token"
	refreshName := "__Host-refresh_token"
	if fs.Manifest != nil && fs.Manifest.Backend.Auth != nil {
		a := fs.Manifest.Backend.Auth
		if a.SecretEnv != "" {
			secretEnv = a.SecretEnv
		}
		if a.AccessTokenTTL != "" {
			accessTTL = a.AccessTokenTTL
		}
		if a.RefreshTokenTTL != "" {
			refreshTTL = a.RefreshTokenTTL
		}
		mode = a.ResolvedMode()
		if a.Cookie != nil {
			if a.Cookie.SameSite != "" {
				sameSite = a.Cookie.SameSite
			}
			if a.Cookie.AccessName != "" {
				accessName = a.Cookie.AccessName
			}
			if a.Cookie.RefreshName != "" {
				refreshName = a.Cookie.RefreshName
			}
		}
	}
	detectReuse := false
	if fs.Manifest != nil && fs.Manifest.Backend.Auth != nil {
		detectReuse = fs.Manifest.Backend.Auth.DetectReuseLogoutAll
	}

	return MainBlock{
		Name:   "auth-init",
		Active: hasAuth,
		Imports: []string{
			`"` + modulePath + `/internal/auth"`,
			`"net/http"`,
			`"time"`,
		},
		Lines: []string{
			`// Phase003 — Configure ssac/pkg/auth. SecretEnv stores the env var NAME;`,
			`// IssueToken/RefreshToken/VerifyToken read os.Getenv(SecretEnv) on every`,
			`// call so secret rotation does not require re-Configure.`,
			`accessTTL, err := time.ParseDuration(` + fmt.Sprintf("%q", accessTTL) + `)`,
			`if err != nil {`,
			`	slog.Error("parse access_token_ttl", "err", err)`,
			`	os.Exit(1)`,
			`}`,
			`refreshTTL, err := time.ParseDuration(` + fmt.Sprintf("%q", refreshTTL) + `)`,
			`if err != nil {`,
			`	slog.Error("parse refresh_token_ttl", "err", err)`,
			`	os.Exit(1)`,
			`}`,
			`// Phase020 — BACKEND_AUTH_MODE env overrides the manifest default`,
			`// so the same binary can serve web (cookie) and mobile (bearer)`,
			`// deployments from a shared image.`,
			`authMode := ` + fmt.Sprintf("%q", mode),
			`if v := os.Getenv("BACKEND_AUTH_MODE"); v != "" {`,
			`	switch v {`,
			`	case "bearer", "cookie", "hybrid":`,
			`		authMode = v`,
			`	}`,
			`}`,
			`// Phase020 — SameSite string → http.SameSite enum. Values outside`,
			`// {Lax, Strict, None} fall back to Lax which is the OWASP-recommended`,
			`// default for same-site SaaS.`,
			`var sameSite http.SameSite`,
			`switch ` + fmt.Sprintf("%q", sameSite) + ` {`,
			`case "Strict":`,
			`	sameSite = http.SameSiteStrictMode`,
			`case "None":`,
			`	sameSite = http.SameSiteNoneMode`,
			`default:`,
			`	sameSite = http.SameSiteLaxMode`,
			`}`,
			`auth.Configure(auth.Config{`,
			`	SecretEnv:  ` + fmt.Sprintf("%q", secretEnv) + `,`,
			`	AccessTTL:  accessTTL,`,
			`	RefreshTTL: refreshTTL,`,
			`	Mode:       authMode,`,
			`	CookieAttrs: auth.CookieAttrs{`,
			`		AccessName:  ` + fmt.Sprintf("%q", accessName) + `,`,
			`		RefreshName: ` + fmt.Sprintf("%q", refreshName) + `,`,
			`		SameSite:    sameSite,`,
			`		AccessTTL:   accessTTL,`,
			`		RefreshTTL:  refreshTTL,`,
			`	},`,
			`})`,
			`// Phase002 — bootstrap refresh_tokens schema (idempotent). Kept in`,
			`// main.go so a fresh DB is usable without running a separate`,
			`// migration tool; real deployments should instead run the DDL via`,
			`// their migration pipeline and drop this block.`,
			`if _, err := conn.ExecContext(ctx, auth.RefreshTokensDDL); err != nil {`,
			`	slog.Error("refresh_tokens DDL", "err", err)`,
			`	os.Exit(1)`,
			`}`,
			fmt.Sprintf(`refreshStore := &auth.RefreshStore{DB: conn, DetectReuseLogoutAll: %t}`, detectReuse),
			`// Phase004/Phase009 — inject the RefreshStore into the Server so SSaC`,
			`// handlers that call auth.RefreshToken / auth.RefreshRotate / auth.Logout`,
			`// can reach it via server.RefreshStore without threading the DB handle`,
			`// through every handler signature. Phase009 moved the auth-refresh`,
			`// route onto the canonical openapi + SSaC path, so this block does`,
			`// not mount any gin route — it only wires store + config.`,
			`srv.RefreshStore = refreshStore`,
		},
	}
}
