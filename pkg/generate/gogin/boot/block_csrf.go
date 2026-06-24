//ff:func feature=gen-gogin type=generator control=sequence topic=csrf
//ff:what blockCsrf — middleware.Csrf 등록 (auth 선언 시, BUG-116: bearer 빌드도 런타임 게이트로 방출)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockCsrf emits the middleware.Csrf registration whenever auth is
// declared and CSRF is not explicitly disabled (hasCsrf). When auth is
// absent the block is inert (empty Lines) and collectActiveBlocks drops it.
//
// BUG-116 / Phase-B1 — the block is registered regardless of the build-time
// resolved mode (bearer included). The emitted middleware no-ops at runtime
// in bearer mode (csrfRuntimeActive), so bearer deployments are byte-for-
// byte unchanged in behavior, while a BACKEND_AUTH_MODE=cookie/hybrid
// override on the same binary now actually reaches a live CSRF check.
//
// Placement: registered before blockBodyLimit / blockRegisterHandlers so
// the CSRF check fires early in the chain, ahead of body reads. In
// hybrid mode the Csrf middleware short-circuits on Bearer headers so
// API clients are unaffected.
func blockCsrf(fs *yongol.Fullstack, a prepared.Auth, modulePath string) MainBlock {
	if !hasCsrf(a) {
		return MainBlock{Name: "csrf", Active: csrfAlwaysInactive}
	}
	cookieName, headerName, exempt, maxAge, secure := csrfCookieSettings(a.Raw.Csrf)
	// Phase008 §4 — in domain mode, exempt bearer-only domain prefixes so the
	// global Csrf check applies to cookie/hybrid domain paths only. isExemptPath
	// is prefix-matched, so /api/admin (bearer) skips while /api (cookie) stays
	// verified.
	exempt = append(exempt, bearerDomainPrefixes(fs)...)
	hybridSkip := a.Mode == "hybrid"

	lines := []string{
		fmt.Sprintf(`csrfEnabled := envBool("BACKEND_AUTH_CSRF_ENABLED", %v)`, true),
		`if csrfEnabled {`,
		`	r.Use(middleware.Csrf(middleware.CsrfConfig{`,
		fmt.Sprintf(`		CookieName:       %q,`, cookieName),
		fmt.Sprintf(`		HeaderName:       %q,`, headerName),
		fmt.Sprintf(`		ExemptPaths:      %s,`, goStringSlice(exempt)),
		fmt.Sprintf(`		MaxAge:           %d,`, maxAge),
		fmt.Sprintf(`		Secure:           %v,`, secure),
		fmt.Sprintf(`		HybridBearerSkip: %v,`, hybridSkip),
		`	}))`,
		`}`,
	}

	return MainBlock{
		Name: "csrf",
		// Active left nil: caller already validated hasCsrf via
		// prepared.Auth before calling this function.
		Imports: []string{fmt.Sprintf(`"%s/internal/middleware"`, modulePath)},
		Lines:   lines,
	}
}
