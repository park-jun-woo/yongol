//ff:type feature=gen-ir type=model
//ff:what CSRFConfig -- double-submit cookie CSRF 방어 설정

package ir

// CSRFConfig holds the resolved CSRF middleware configuration.
type CSRFConfig struct {
	// CookieName is the JS-readable CSRF cookie name (default "XSRF-TOKEN").
	CookieName string

	// HeaderName is the header clients must set (default "X-XSRF-TOKEN").
	HeaderName string

	// ExemptPaths lists URL prefixes exempt from CSRF checks.
	ExemptPaths []string

	// MaxAge is the CSRF cookie lifetime in seconds.
	MaxAge int

	// Secure controls the Secure flag on the CSRF cookie.
	Secure bool

	// HybridBearerSkip is true when auth mode is "hybrid", causing the
	// middleware to skip CSRF checks for requests with Bearer tokens.
	HybridBearerSkip bool
}
