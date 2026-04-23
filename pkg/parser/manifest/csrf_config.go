//ff:type feature=projectconfig type=model
//ff:what CsrfConfig — backend.auth.csrf 섹션 모델 (double-submit 쿠키 방어)

package manifest

// CsrfConfig mirrors the backend.auth.csrf: section. Double-submit cookie
// defense: server sets CookieName (JS-readable), client duplicates it into
// HeaderName on state-changing requests. Safe methods (GET/HEAD/OPTIONS)
// and ExemptPaths skip verification.
//
// env overrides:
//   BACKEND_AUTH_CSRF_ENABLED=false  — emergency disable.
type CsrfConfig struct {
	Enabled     bool     `yaml:"enabled"`
	CookieName  string   `yaml:"cookie_name"`  // default "XSRF-TOKEN"
	HeaderName  string   `yaml:"header_name"`  // default "X-XSRF-TOKEN"
	ExemptPaths []string `yaml:"exempt_paths"` // prefix match
	// MaxAge is the CSRF cookie lifetime in seconds (default 86400).
	// Distinct from the session cookie MaxAge — this one is for the
	// JS-readable XSRF token, whose expiry can legitimately differ.
	MaxAge int `yaml:"max_age"`
}
