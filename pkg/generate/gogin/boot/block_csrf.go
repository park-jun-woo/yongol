//ff:func feature=gen-gogin type=generator control=sequence topic=csrf
//ff:what blockCsrf — middleware.Csrf 등록 (쿠키 인증 조건부, Phase005)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

// blockCsrf emits the middleware.Csrf registration when the resolved
// auth mode is "cookie" or "hybrid" AND csrf.enabled. On the default
// bearer configuration the block is inert (empty Lines) and
// collectActiveBlocks drops it via hasCsrf — no imports or code land
// in main.go, keeping bearer deployments unchanged.
//
// Phase005 dormant: hasCsrf returns false for default manifests today.
// Phase020 (CookieSessionAuth) will flip projects to mode=cookie, at
// which point this block lights up automatically.
//
// Placement: registered before blockBodyLimit / blockRegisterHandlers so
// the CSRF check fires early in the chain, ahead of body reads. In
// hybrid mode the Csrf middleware short-circuits on Bearer headers so
// API clients are unaffected.
func blockCsrf(a prepared.Auth, modulePath string) MainBlock {
	if !hasCsrf(a) {
		return MainBlock{Name: "csrf", Active: csrfAlwaysInactive}
	}
	cookieName, headerName, exempt, maxAge, secure := csrfCookieSettings(a.Raw.Csrf)
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
