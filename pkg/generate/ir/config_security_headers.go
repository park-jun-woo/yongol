//ff:type feature=gen-ir type=model
//ff:what SecurityHeadersConfig -- 브라우저 보안 헤더 미들웨어 설정

package ir

// SecurityHeadersConfig holds the resolved security headers middleware
// configuration. Maps to the six browser-facing headers: HSTS,
// X-Content-Type-Options, X-Frame-Options, CSP, Referrer-Policy,
// Permissions-Policy.
type SecurityHeadersConfig struct {
	// Profile selects a preset: "production", "dev", or "api".
	Profile string

	// HSTSMaxAge is the Strict-Transport-Security max-age in seconds.
	HSTSMaxAge int

	// HSTSIncludeSubs enables includeSubDomains on HSTS.
	HSTSIncludeSubs bool

	// HSTSPreload enables preload on HSTS.
	HSTSPreload bool

	// CSPEnabled is true when Content-Security-Policy should be emitted.
	CSPEnabled bool

	// CSPReportOnly sends CSP as report-only (dev mode).
	CSPReportOnly bool

	// CSPDirectives maps CSP directive names to their source lists.
	CSPDirectives map[string][]string

	// XFrameOptions is "DENY" or "SAMEORIGIN".
	XFrameOptions string

	// ReferrerPolicy is the Referrer-Policy header value.
	ReferrerPolicy string

	// PermissionsPolicy maps feature names to allowed origin lists.
	PermissionsPolicy map[string][]string
}
