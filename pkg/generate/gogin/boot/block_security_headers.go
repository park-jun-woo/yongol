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
		`	secHeadersCfg := middleware.SecurityHeadersConfig{`,
		`		Enabled:           true,`,
		`		Profile:           shProfile,`,
		`		HSTSMaxAge:        shHSTSMaxAge,`,
		fmt.Sprintf(`		HSTSIncludeSubs:   %v,`, cfg.HSTSIncludeSubs),
		fmt.Sprintf(`		HSTSPreload:       %v,`, cfg.HSTSPreload),
		fmt.Sprintf(`		CSPEnabled:        %v,`, cfg.CSPEnabled),
		`		CSPReportOnly:     shCSPReportOnly,`,
		fmt.Sprintf(`		CSPDirectives:     %s,`, goStringListMap(cfg.CSPDirectives)),
		fmt.Sprintf(`		XFrameOptions:     %q,`, cfg.XFrameOptions),
		fmt.Sprintf(`		ReferrerPolicy:    %q,`, cfg.ReferrerPolicy),
		fmt.Sprintf(`		PermissionsPolicy: %s,`, goStringListMap(cfg.PermissionsPolicy)),
		`	}`,
		`	r.Use(middleware.SecurityHeadersMiddleware(secHeadersCfg))`,
		`}`,
	}

	imports := []string{
		fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
	}

	return MainBlock{
		Name:    "security-headers",
		Imports: imports,
		Lines:   lines,
	}
}
