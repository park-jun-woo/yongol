//ff:type feature=projectconfig type=model
//ff:what SecurityHeadersConfig — backend.security_headers 섹션 모델 (HSTS/CSP/XFO/Referrer/Permissions)

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
//	BACKEND_SECURITY_HEADERS_ENABLED
//	BACKEND_SECURITY_HEADERS_PROFILE
//	BACKEND_SECURITY_HEADERS_HSTS_MAX_AGE
//	BACKEND_SECURITY_HEADERS_CSP_REPORT_ONLY
type SecurityHeadersConfig struct {
	Enabled           *bool               `yaml:"enabled,omitempty"`
	Profile           string              `yaml:"profile,omitempty"` // production | dev | api
	HSTS              *HSTSConfig         `yaml:"hsts,omitempty"`
	CSP               *CSPConfig          `yaml:"csp,omitempty"`
	XFrameOptions     string              `yaml:"x_frame_options,omitempty"`    // DENY | SAMEORIGIN
	ReferrerPolicy    string              `yaml:"referrer_policy,omitempty"`    // strict-origin-when-cross-origin (default)
	PermissionsPolicy map[string][]string `yaml:"permissions_policy,omitempty"` // camera: [] → camera=()
}
