//ff:type feature=projectconfig type=model
//ff:what SecurityHeadersConfig 구조체 — backend.security_headers (HSTS/CSP/XFO/nosniff/Referrer/Permissions)

package manifest

// SecurityHeadersConfig mirrors the backend.security_headers: section of
// manifest.yaml. Phase007 wires six browser-facing security headers into the
// generated backend — HSTS, X-Content-Type-Options, X-Frame-Options, CSP,
// Referrer-Policy, Permissions-Policy. When the block is absent yongol
// applies production defaults (all headers emitted, CSP strict).
//
// Enabled is a *bool so the generator can distinguish "unset" (treat as true)
// from "explicitly false" (disable the middleware entirely).
//
// Profile selects a preset:
//
//   - "production" (default) — every header active. HSTS 1 year +
//     includeSubDomains. CSP strict. Referrer-Policy
//     strict-origin-when-cross-origin. Permissions-Policy denies camera /
//     microphone / geolocation.
//   - "dev" — HSTS omitted (local HTTPS not assumed), CSP emitted as
//     Content-Security-Policy-Report-Only so violations log without breaking
//     SSR inline styles during development.
//   - "api" — CSP omitted (JSON APIs do not render). Other five headers
//     remain active.
//
// Env overrides (resolved in generated main.go):
//
//   BACKEND_SECURITY_HEADERS_ENABLED
//   BACKEND_SECURITY_HEADERS_PROFILE
//   BACKEND_SECURITY_HEADERS_HSTS_MAX_AGE
//   BACKEND_SECURITY_HEADERS_CSP_REPORT_ONLY
type SecurityHeadersConfig struct {
	Enabled           *bool             `yaml:"enabled,omitempty"`
	Profile           string            `yaml:"profile,omitempty"` // production | dev | api
	HSTS              *HSTSConfig       `yaml:"hsts,omitempty"`
	CSP               *CSPConfig        `yaml:"csp,omitempty"`
	XFrameOptions     string            `yaml:"x_frame_options,omitempty"`  // DENY | SAMEORIGIN
	ReferrerPolicy    string            `yaml:"referrer_policy,omitempty"`  // strict-origin-when-cross-origin (default)
	PermissionsPolicy map[string][]string `yaml:"permissions_policy,omitempty"` // camera: [] → camera=()
}

// HSTSConfig controls the Strict-Transport-Security header. MaxAge is seconds.
// Preload requires MaxAge >= 31536000 and IncludeSubDomains=true to be
// accepted by browser preload lists; yongol does not enforce this — the
// SEC-302 validate rule warns when MaxAge is too short.
type HSTSConfig struct {
	MaxAge            int  `yaml:"max_age,omitempty"`
	IncludeSubDomains bool `yaml:"include_subdomains,omitempty"`
	Preload           bool `yaml:"preload,omitempty"`
}

// CSPConfig controls Content-Security-Policy. When Enabled is nil it is
// treated as true. ReportOnly swaps the emitted header name to
// Content-Security-Policy-Report-Only — directives still apply but the
// browser only logs violations instead of blocking resources.
//
// Directives maps CSP directive names (default-src, script-src, frame-ancestors,
// ...) to their source lists. Yongol concatenates the map into a single
// header string at codegen + runtime boot time; iteration order is made
// deterministic via sorted keys so generated output is reproducible.
type CSPConfig struct {
	Enabled    *bool               `yaml:"enabled,omitempty"`
	ReportOnly bool                `yaml:"report_only,omitempty"`
	Directives map[string][]string `yaml:"directives,omitempty"`
}
