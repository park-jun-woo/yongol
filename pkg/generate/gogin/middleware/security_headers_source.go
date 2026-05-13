//ff:type feature=gen-gogin type=generator
//ff:what securityHeadersSources — security_headers 를 6파일로 분할하는 소스 템플릿 맵

package middleware

// securityHeadersConfigSource — SecurityHeadersConfig type.
const securityHeadersConfigSource = `//` + `ff:type feature=runtime-middleware type=model topic=security-headers
//` + `ff:what SecurityHeadersConfig — 보안 헤더 미들웨어 런타임 설정

package middleware

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
`

// securityHeadersMiddlewareSource — SecurityHeadersMiddleware func.
const securityHeadersMiddlewareSource = `//` + `ff:func feature=runtime-middleware type=middleware control=sequence topic=security-headers
//` + `ff:what SecurityHeadersMiddleware — 브라우저 보안 헤더 6종 자동 주입 미들웨어

package middleware

import "github.com/gin-gonic/gin"

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
`

// buildStaticSecurityHeadersSource — BuildStaticSecurityHeaders func.
const buildStaticSecurityHeadersSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=security-headers
//` + `ff:what BuildStaticSecurityHeaders — profile 기반 비 CSP 보안 헤더 맵 조립

package middleware

import (
	"strconv"
	"strings"
)

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
	headers["X-Content-Type-Options"] = "nosniff"
	xfo := strings.TrimSpace(cfg.XFrameOptions)
	if xfo == "" {
		xfo = "DENY"
	}
	headers["X-Frame-Options"] = xfo
	ref := strings.TrimSpace(cfg.ReferrerPolicy)
	if ref == "" {
		ref = "strict-origin-when-cross-origin"
	}
	headers["Referrer-Policy"] = ref
	if len(cfg.PermissionsPolicy) > 0 {
		headers["Permissions-Policy"] = buildPermissionsPolicy(cfg.PermissionsPolicy)
	}
	return headers
}
`

// buildCSPHeaderSource — BuildCSPHeader func.
const buildCSPHeaderSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=security-headers
//` + `ff:what BuildCSPHeader — CSP 헤더 이름 + 값 결정 (profile/ReportOnly 분기)

package middleware

import "strings"

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
	if profile == "dev" || cfg.CSPReportOnly {
		return "Content-Security-Policy-Report-Only", value
	}
	return "Content-Security-Policy", value
}
`

// buildCSPValueSource — BuildCSPValue func.
const buildCSPValueSource = `//` + `ff:func feature=runtime-middleware type=util control=iteration dimension=1 topic=security-headers
//` + `ff:what BuildCSPValue — CSP directives map 을 표준 헤더 문자열로 렌더

package middleware

import (
	"sort"
	"strings"
)

// BuildCSPValue renders a CSP directives map into the canonical
//
//	"directive1 source1 source2; directive2 source3"
//
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
`

// buildPermissionsPolicySource — buildPermissionsPolicy func.
const buildPermissionsPolicySource = `//` + `ff:func feature=runtime-middleware type=util control=iteration dimension=1 topic=security-headers
//` + `ff:what buildPermissionsPolicy — Permissions-Policy 헤더 값 렌더

package middleware

import (
	"sort"
	"strings"
)

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
