//ff:type feature=projectconfig type=model
//ff:what 쿠키 인증 / CSRF 설정 구조체 (backend.auth.cookie, backend.auth.csrf)
package manifest

// CookieConfig mirrors the backend.auth.cookie: section. It carries the
// parameters for the server-issued session cookie. Active only when
// Auth.Mode == "cookie" or "hybrid"; otherwise ignored by the generator.
//
// Phase020 schema (reduced from Phase005): the generator enforces the
// 2026 standard (HttpOnly + Secure + __Host- prefix) unconditionally, so
// the manifest only exposes the three knobs that can vary per deployment.
// Callers upgrading from Phase005:
//   - Name      → AccessName (renamed; previous `session`-style values no
//                 longer apply because the generator enforces __Host-).
//   - Secure    → removed (always true; __Host- prefix mandates it).
//   - HTTPOnly  → removed (always true; JS access is never permitted for
//                 session tokens per OWASP ASVS 5.0).
//   - MaxAge    → removed (derived from auth.access_token_ttl /
//                 refresh_token_ttl, keeping cookie lifetime locked to
//                 token lifetime).
type CookieConfig struct {
	// AccessName overrides the Set-Cookie name of the access token.
	// Defaults to "__Host-access_token". Names that don't begin with
	// "__Host-" or "__Secure-" lose browser-enforced prefix guarantees
	// (Secure + Domain=none for __Host-). Kept as a knob for specialised
	// deployments (e.g. hosted at a path that __Host- rejects) but the
	// default is strongly recommended.
	AccessName string `yaml:"access_name"`
	// RefreshName overrides the Set-Cookie name of the refresh token.
	// Defaults to "__Host-refresh_token".
	RefreshName string `yaml:"refresh_name"`
	// SameSite controls the Set-Cookie SameSite attribute. One of
	// "Lax" (default, safe top-level navigation), "Strict" (blocks even
	// same-site cross-document requests, breaks some login flows), or
	// "None" (required for cross-origin SPA; must be paired with
	// Secure=true, which the __Host- prefix already enforces). An empty
	// value resolves to "Lax" at generation time.
	SameSite string `yaml:"same_site"`
}

// CsrfConfig mirrors the backend.auth.csrf: section. Double-submit cookie
// defense: server sets CookieName (JS-readable), client duplicates it into
// HeaderName on state-changing requests. Safe methods (GET/HEAD/OPTIONS)
// and ExemptPaths skip verification.
//
// env overrides:
//   BACKEND_AUTH_CSRF_ENABLED=false  — emergency disable.
type CsrfConfig struct {
	Enabled     bool     `yaml:"enabled"`
	CookieName  string   `yaml:"cookie_name"`   // default "XSRF-TOKEN"
	HeaderName  string   `yaml:"header_name"`   // default "X-XSRF-TOKEN"
	ExemptPaths []string `yaml:"exempt_paths"`  // prefix match
	// MaxAge is the CSRF cookie lifetime in seconds (default 86400).
	// Distinct from the session cookie MaxAge — this one is for the
	// JS-readable XSRF token, whose expiry can legitimately differ.
	MaxAge int `yaml:"max_age"`
}
