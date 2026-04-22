//ff:type feature=gen-gogin type=generator
//ff:what securityHeadersSource — GenerateSecurityHeaders 가 기록하는 security_headers.go 정적 소스

package middleware

// securityHeadersSource carries the verbatim Go source for
// internal/middleware/security_headers.go. No placeholders — the generated
// middleware is config-driven at boot time by SecurityHeadersConfig.
// Bootstrap (blockSecurityHeaders in main.go) builds the static header map
// + CSP value once and passes them to SecurityHeadersMiddleware, so per-
// request cost is a short range loop over the pre-baked header set.
//
// Provides:
//
//   - SecurityHeadersConfig          — runtime shape fed by main.go.
//   - SecurityHeadersMiddleware(cfg) — gin.HandlerFunc writing headers.
//   - BuildStaticSecurityHeaders(cfg)/ BuildCSPHeader(cfg) — exported for
//     tests and for bootstrap code.
//
// Profile semantics:
//
//   - production: HSTS + nosniff + XFO + CSP + Referrer-Policy +
//     Permissions-Policy.
//   - dev:        same as production minus HSTS. CSP reported via
//                 Content-Security-Policy-Report-Only.
//   - api:        same as production minus CSP.
const securityHeadersSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=security-headers
//` + `ff:what SecurityHeadersMiddleware — 브라우저 보안 헤더 6종(HSTS/nosniff/XFO/CSP/Referrer/Permissions) 자동 주입

package middleware

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// SecurityHeadersConfig is the runtime shape consumed by
// SecurityHeadersMiddleware. Populated by generated main.go from
// manifest.backend.security_headers with env overrides applied.
type SecurityHeadersConfig struct {
	Enabled           bool
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

// SecurityHeadersMiddleware returns a gin middleware that writes the
// configured security headers on every response. When cfg.Enabled is false
// the middleware is a no-op. Static headers and the CSP value are computed
// once at middleware construction time.
func SecurityHeadersMiddleware(cfg SecurityHeadersConfig) gin.HandlerFunc {
	if !cfg.Enabled {
		return func(c *gin.Context) { c.Next() }
	}
	staticHeaders := BuildStaticSecurityHeaders(cfg)
	cspKey, cspValue := BuildCSPHeader(cfg)
	return func(c *gin.Context) {
		h := c.Writer.Header()
		for k, v := range staticHeaders {
			h.Set(k, v)
		}
		if cspKey != "" {
			h.Set(cspKey, cspValue)
		}
		c.Next()
	}
}

// BuildStaticSecurityHeaders assembles the non-CSP header set for the
// configured profile. Keys are returned exactly as they are emitted on the
// wire. CSP is handled separately because the header name depends on
// ReportOnly.
func BuildStaticSecurityHeaders(cfg SecurityHeadersConfig) map[string]string {
	headers := map[string]string{}
	if !cfg.Enabled {
		return headers
	}
	profile := strings.ToLower(strings.TrimSpace(cfg.Profile))
	if profile == "" {
		profile = "production"
	}

	// HSTS — omitted in dev profile. MaxAge <= 0 also disables it so
	// operators can toggle via BACKEND_SECURITY_HEADERS_HSTS_MAX_AGE=0.
	if profile != "dev" && cfg.HSTSMaxAge > 0 {
		parts := []string{"max-age=" + strconv.Itoa(cfg.HSTSMaxAge)}
		if cfg.HSTSIncludeSubs {
			parts = append(parts, "includeSubDomains")
		}
		if cfg.HSTSPreload {
			parts = append(parts, "preload")
		}
		headers["Strict-Transport-Security"] = strings.Join(parts, "; ")
	}

	// X-Content-Type-Options — always on when middleware is enabled.
	headers["X-Content-Type-Options"] = "nosniff"

	// X-Frame-Options — defaults to DENY.
	xfo := strings.TrimSpace(cfg.XFrameOptions)
	if xfo == "" {
		xfo = "DENY"
	}
	headers["X-Frame-Options"] = xfo

	// Referrer-Policy — defaults to strict-origin-when-cross-origin.
	ref := strings.TrimSpace(cfg.ReferrerPolicy)
	if ref == "" {
		ref = "strict-origin-when-cross-origin"
	}
	headers["Referrer-Policy"] = ref

	// Permissions-Policy — empty map yields no header.
	if len(cfg.PermissionsPolicy) > 0 {
		headers["Permissions-Policy"] = buildPermissionsPolicy(cfg.PermissionsPolicy)
	}

	return headers
}

// BuildCSPHeader returns the (header-name, header-value) pair for the
// Content-Security-Policy header. Empty name indicates CSP must not be
// emitted (api profile, CSP disabled, empty directives).
func BuildCSPHeader(cfg SecurityHeadersConfig) (string, string) {
	if !cfg.Enabled || !cfg.CSPEnabled {
		return "", ""
	}
	profile := strings.ToLower(strings.TrimSpace(cfg.Profile))
	if profile == "api" {
		return "", ""
	}
	value := BuildCSPValue(cfg.CSPDirectives)
	if value == "" {
		return "", ""
	}
	// dev profile forces report-only so inline violations are logged but
	// never block SSR pages. Explicit ReportOnly also honoured.
	if profile == "dev" || cfg.CSPReportOnly {
		return "Content-Security-Policy-Report-Only", value
	}
	return "Content-Security-Policy", value
}

// BuildCSPValue renders a CSP directives map into the canonical
//   "directive1 source1 source2; directive2 source3"
// header string. Directives are emitted in sorted order so output is
// deterministic across processes.
func BuildCSPValue(directives map[string][]string) string {
	if len(directives) == 0 {
		return ""
	}
	keys := make([]string, 0, len(directives))
	for k := range directives {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		sources := directives[k]
		if len(sources) == 0 {
			parts = append(parts, k)
			continue
		}
		parts = append(parts, k+" "+strings.Join(sources, " "))
	}
	return strings.Join(parts, "; ")
}

// buildPermissionsPolicy renders the Permissions-Policy header value. Each
// entry "feature=(origin1 origin2)" or "feature=()" when the list is empty.
// Keys are emitted in sorted order for reproducibility.
func buildPermissionsPolicy(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		origins := m[k]
		if len(origins) == 0 {
			parts = append(parts, k+"=()")
			continue
		}
		parts = append(parts, k+"=("+strings.Join(origins, " ")+")")
	}
	return strings.Join(parts, ", ")
}
`
