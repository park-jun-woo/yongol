//ff:func feature=gen-gogin type=generator control=sequence topic=security-headers
//ff:what blockSecurityHeaders — middleware.SecurityHeadersMiddleware 등록 + SecurityHeadersConfig 조립

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockSecurityHeaders emits the Security Headers middleware registration.
// Activated by default (Enabled defaults to true) — a manifest without
// security_headers block still ships with the production preset.
//
// Ordering: after request_id (Phase004) and error_envelope (Phase004) so
// any header the envelope writes is still visible, but before body_limit /
// rate_limit / cors / handlers so early rejections (413, 429) already carry
// the security headers. Phase004 absence is tolerated — the chain still
// works because collectActiveBlocks orders blockSecurityHeaders before
// blockBodyLimit regardless of the Phase004 blocks being present.
func blockSecurityHeaders(fs *yongol.Fullstack, modulePath string) MainBlock {
	if !hasSecurityHeaders(fs) {
		return MainBlock{
			Name:   "security-headers",
			Active: func(_ *yongol.Fullstack) bool { return false },
		}
	}

	cfg := resolveSecurityHeaders(fs)

	lines := []string{
		`shEnabled := envBool("BACKEND_SECURITY_HEADERS_ENABLED", true)`,
		fmt.Sprintf(`shProfile := envString("BACKEND_SECURITY_HEADERS_PROFILE", %q)`, cfg.Profile),
		fmt.Sprintf(`shHSTSMaxAge := envInt("BACKEND_SECURITY_HEADERS_HSTS_MAX_AGE", %d)`, cfg.HSTSMaxAge),
		fmt.Sprintf(`shCSPReportOnly := envBool("BACKEND_SECURITY_HEADERS_CSP_REPORT_ONLY", %v)`, cfg.CSPReportOnly),
		`if shEnabled {`,
		`	r.Use(middleware.SecurityHeadersMiddleware(buildSecurityHeadersCfg(shProfile, shHSTSMaxAge, shCSPReportOnly)))`,
		`}`,
	}

	helperFunc := fmt.Sprintf(`func buildSecurityHeadersCfg(profile string, hstsMaxAge int, cspReportOnly bool) middleware.SecurityHeadersConfig {
	return middleware.SecurityHeadersConfig{
		Enabled:           true,
		Profile:           profile,
		HSTSMaxAge:        hstsMaxAge,
		HSTSIncludeSubs:   %v,
		HSTSPreload:       %v,
		CSPEnabled:        %v,
		CSPReportOnly:     cspReportOnly,
		CSPDirectives:     %s,
		XFrameOptions:     %q,
		ReferrerPolicy:    %q,
		PermissionsPolicy: %s,
	}
}`, cfg.HSTSIncludeSubs, cfg.HSTSPreload, cfg.CSPEnabled,
		goStringListMap(cfg.CSPDirectives), cfg.XFrameOptions,
		cfg.ReferrerPolicy, goStringListMap(cfg.PermissionsPolicy))

	imports := []string{
		fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
	}

	return MainBlock{
		Name:    "security-headers",
		Imports: imports,
		Lines:   lines,
		Funcs:   []string{helperFunc},
	}
}
