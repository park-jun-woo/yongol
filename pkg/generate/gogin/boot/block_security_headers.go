//ff:func feature=gen-gogin type=generator control=sequence topic=security-headers
//ff:what blockSecurityHeaders — middleware.SecurityHeadersMiddleware 등록 + SecurityHeadersConfig 조립

package boot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// Default header values applied when the manifest omits the security_headers
// block or an individual field. production profile enables every header.
const (
	defaultSHProfile        = "production"
	defaultHSTSMaxAge       = 31536000 // 1 year
	defaultHSTSIncludeSubs  = true
	defaultHSTSPreload      = false
	defaultXFrameOptions    = "DENY"
	defaultReferrerPolicy   = "strict-origin-when-cross-origin"
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
			Name: "security-headers",
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

// resolvedSecurityHeaders captures the generator-time view of
// manifest.backend.security_headers after defaults are applied. Runtime env
// overrides are layered on top in the generated main.go.
type resolvedSecurityHeaders struct {
	Profile           string
	HSTSMaxAge        int
	HSTSIncludeSubs   bool
	HSTSPreload       bool
	CSPEnabled        bool
	CSPReportOnly     bool
	CSPDirectives     map[string][]string
	XFrameOptions     string
	ReferrerPolicy    string
	PermissionsPolicy map[string][]string
}

// resolveSecurityHeaders collapses manifest config with production-profile
// defaults. A nil block yields the full production preset. Individual nil /
// zero sub-blocks fall back to their own defaults.
func resolveSecurityHeaders(fs *yongol.Fullstack) resolvedSecurityHeaders {
	out := resolvedSecurityHeaders{
		Profile:         defaultSHProfile,
		HSTSMaxAge:      defaultHSTSMaxAge,
		HSTSIncludeSubs: defaultHSTSIncludeSubs,
		HSTSPreload:     defaultHSTSPreload,
		CSPEnabled:      true,
		CSPReportOnly:   false,
		CSPDirectives: map[string][]string{
			"default-src":     {"'self'"},
			"frame-ancestors": {"'none'"},
			"base-uri":        {"'self'"},
		},
		XFrameOptions:  defaultXFrameOptions,
		ReferrerPolicy: defaultReferrerPolicy,
		PermissionsPolicy: map[string][]string{
			"camera":      {},
			"microphone":  {},
			"geolocation": {},
		},
	}

	if fs == nil || fs.Manifest == nil {
		return out
	}
	sh := fs.Manifest.Backend.SecurityHeaders
	if sh == nil {
		return out
	}

	if p := strings.TrimSpace(sh.Profile); p != "" {
		out.Profile = strings.ToLower(p)
	}
	if sh.HSTS != nil {
		if sh.HSTS.MaxAge > 0 {
			out.HSTSMaxAge = sh.HSTS.MaxAge
		}
		out.HSTSIncludeSubs = sh.HSTS.IncludeSubDomains
		out.HSTSPreload = sh.HSTS.Preload
	}
	if sh.CSP != nil {
		if sh.CSP.Enabled != nil {
			out.CSPEnabled = *sh.CSP.Enabled
		}
		out.CSPReportOnly = sh.CSP.ReportOnly
		if len(sh.CSP.Directives) > 0 {
			out.CSPDirectives = copyStringListMap(sh.CSP.Directives)
		}
	}
	if v := strings.TrimSpace(sh.XFrameOptions); v != "" {
		out.XFrameOptions = v
	}
	if v := strings.TrimSpace(sh.ReferrerPolicy); v != "" {
		out.ReferrerPolicy = v
	}
	if sh.PermissionsPolicy != nil {
		out.PermissionsPolicy = copyStringListMap(sh.PermissionsPolicy)
	}
	return out
}

// copyStringListMap deep-copies a map[string][]string so the generator
// never mutates the shared manifest instance.
func copyStringListMap(in map[string][]string) map[string][]string {
	if in == nil {
		return nil
	}
	out := make(map[string][]string, len(in))
	for k, v := range in {
		copied := make([]string, len(v))
		copy(copied, v)
		out[k] = copied
	}
	return out
}

// goStringListMap renders a map[string][]string as a Go literal. Keys are
// emitted in sorted order for deterministic codegen output. An empty or nil
// map yields `map[string][]string{}` so downstream struct assignment stays
// assignment-safe (no nil dereference when length-checking later).
func goStringListMap(m map[string][]string) string {
	if len(m) == 0 {
		return `map[string][]string{}`
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	b.WriteString("map[string][]string{\n")
	for _, k := range keys {
		b.WriteString(fmt.Sprintf("\t\t\t%q: %s,\n", k, goStringSlice(m[k])))
	}
	b.WriteString("\t\t}")
	return b.String()
}

// Keep manifest import referenced so a future refactor that drops all other
// uses still compiles.
var _ = manifest.SecurityHeadersConfig{}
