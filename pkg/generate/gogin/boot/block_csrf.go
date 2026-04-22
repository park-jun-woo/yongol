//ff:func feature=gen-gogin type=generator control=sequence topic=csrf
//ff:what blockCsrf — middleware.Csrf 등록 (쿠키 인증 조건부, Phase005)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockCsrf emits the middleware.Csrf registration when manifest
// declares backend.auth.mode="cookie" or "hybrid" AND csrf.enabled.
// On the default bearer configuration the block is inert (empty Lines)
// and collectActiveBlocks drops it via the Active gate — no imports or
// code land in main.go, keeping bearer deployments unchanged.
//
// Phase005 dormant: hasCsrf returns false for default manifests today.
// Phase020 (CookieSessionAuth) will flip projects to mode=cookie, at
// which point this block lights up automatically.
//
// Placement: registered before blockBodyLimit / blockRegisterHandlers so
// the CSRF check fires early in the chain, ahead of body reads. In
// hybrid mode the Csrf middleware short-circuits on Bearer headers so
// API clients are unaffected.
func blockCsrf(fs *yongol.Fullstack, modulePath string) MainBlock {
	if !hasCsrf(fs) {
		return MainBlock{Name: "csrf", Active: hasCsrf}
	}
	a := fs.Manifest.Backend.Auth
	cookieName := "XSRF-TOKEN"
	headerName := "X-XSRF-TOKEN"
	var exempt []string
	maxAge := 86400
	// Phase020 — Secure defaults to true for the CSRF cookie. Previously
	// this flag was pulled from Cookie.Secure (removed in Phase020)
	// through a plain bool, which conflated "unset" with "false" and
	// silently shipped an insecure cookie on production manifests. The
	// new default is true without an override knob; deployments needing
	// HTTP-only dev testing should set BACKEND_AUTH_CSRF_ENABLED=false
	// instead of carving a per-attribute escape hatch.
	secure := true
	if a.Csrf != nil {
		if a.Csrf.CookieName != "" {
			cookieName = a.Csrf.CookieName
		}
		if a.Csrf.HeaderName != "" {
			headerName = a.Csrf.HeaderName
		}
		exempt = a.Csrf.ExemptPaths
		if a.Csrf.MaxAge > 0 {
			maxAge = a.Csrf.MaxAge
		}
	}
	hybridSkip := a.ResolvedMode() == "hybrid"

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
		Name:    "csrf",
		Active:  hasCsrf,
		Imports: []string{fmt.Sprintf(`"%s/internal/middleware"`, modulePath)},
		Lines:   lines,
	}
}
