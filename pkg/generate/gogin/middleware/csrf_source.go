//ff:type feature=gen-gogin type=generator topic=csrf
//ff:what csrfSourceTemplate — the csrf.go template written by GenerateCsrf (double-submit cookie)

package middleware

// csrfSourceTemplate is the verbatim Go source for
// internal/middleware/csrf.go. Double-submit cookie defense:
//
//  1. Server sets cfg.CookieName cookie (JS-readable) on safe requests.
//  2. Client duplicates cookie value into cfg.HeaderName header on
//     state-changing requests.
//  3. Server compares cookie vs header with constant-time equality.
//
// Token is 32 random bytes → base64url. Safe methods (GET/HEAD/OPTIONS)
// and cfg.ExemptPaths prefix matches skip verification. Hybrid mode
// (Authorization: Bearer header present) also skips — allows API clients
// to coexist with browser sessions on the same surface.
//
// The in-template fallbacks ("XSRF-TOKEN" / "X-XSRF-TOKEN" / 86400 in
// Csrf()) keep the emitted file self-contained; boot always injects the
// manifest-resolved values, so they are defense-in-depth only. They must
// stay byte-identical to the defaults in
// pkg/parser/manifest/csrf_config.go (ResolvedCookieName /
// ResolvedHeaderName / ResolvedMaxAge) — the single source consumed by
// the generator and validate (XOH-07).
//
// Error envelope keeps the Phase004 shape (error / message / request_id)
// so downstream JSON clients parse it uniformly.
//
// BUG-116 / Phase-B1 — the template is now a format string with a single
// %q verb carrying the build-time default auth mode. The emitted Csrf()
// gates every request on csrfRuntimeActive(), which reads BACKEND_AUTH_MODE
// (mirroring authMode() in auth_mode.go, duplicated here so csrf.go stays
// self-contained). In bearer mode the middleware is a no-op — neither
// issuing nor verifying tokens — so a manifest=bearer build ships this file
// yet behaves identically until BACKEND_AUTH_MODE flips it to cookie/hybrid.
const csrfSourceTemplate = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=csrf
//` + `ff:what Csrf — double-submit cookie CSRF defense (runtime-gated on auth transport mode)

package middleware

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CsrfConfig controls double-submit cookie behavior. CookieName must be
// JS-readable (HttpOnly=false); the SPA reads it and copies it into
// HeaderName on state-changing requests.
type CsrfConfig struct {
	CookieName  string
	HeaderName  string
	ExemptPaths []string
	MaxAge      int
	Secure      bool
	// HybridBearerSkip lets Authorization: Bearer ... requests bypass CSRF
	// when auth.mode=hybrid. API clients authenticate via header-bound
	// tokens (already CSRF-immune) while browser sessions still require
	// the double-submit check.
	HybridBearerSkip bool
}

// CsrfEnvelope mirrors the Phase004 error envelope shape.
type CsrfEnvelope struct {
	Error     string ` + "`" + `json:"error"` + "`" + `
	Message   string ` + "`" + `json:"message"` + "`" + `
	RequestID string ` + "`" + `json:"request_id,omitempty"` + "`" + `
}

// Csrf returns a gin.HandlerFunc enforcing double-submit cookie CSRF.
// Safe methods issue/refresh the cookie and pass through; state-changing
// methods compare cookie vs header in constant time.
func Csrf(cfg CsrfConfig) gin.HandlerFunc {
	if cfg.CookieName == "" {
		cfg.CookieName = "XSRF-TOKEN"
	}
	if cfg.HeaderName == "" {
		cfg.HeaderName = "X-XSRF-TOKEN"
	}
	if cfg.MaxAge == 0 {
		cfg.MaxAge = 86400
	}
	return func(c *gin.Context) {
		if !csrfRuntimeActive() {
			// bearer transport: Authorization header is not auto-sent
			// cross-origin, so CSRF does not apply — pass through without
			// issuing or verifying a token (BUG-116 runtime gate).
			c.Next()
			return
		}
		if (cfg.HybridBearerSkip || csrfAuthMode() == "hybrid") && hasBearerHeader(c) {
			c.Next()
			return
		}
		if isSafeMethod(c.Request.Method) || isExemptPath(c.Request.URL.Path, cfg.ExemptPaths) {
			if _, err := c.Cookie(cfg.CookieName); err != nil {
				tok, gerr := generateCsrfToken()
				if gerr == nil {
					c.SetCookie(cfg.CookieName, tok, cfg.MaxAge, "/", "", cfg.Secure, false)
				}
			}
			c.Next()
			return
		}
		cookieTok, _ := c.Cookie(cfg.CookieName)
		headerTok := c.GetHeader(cfg.HeaderName)
		if cookieTok == "" || headerTok == "" || !constantTimeEqual(cookieTok, headerTok) {
			c.AbortWithStatusJSON(http.StatusForbidden, CsrfEnvelope{
				Error:     "csrf_token_invalid",
				Message:   "CSRF token is invalid",
				RequestID: c.GetString("request_id"),
			})
			return
		}
		c.Next()
	}
}

// generateCsrfToken returns a 32-byte random token encoded as base64url
// (URL-safe, no padding) so it transports cleanly in both cookies and
// headers without escaping.
func generateCsrfToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// isSafeMethod returns true for HTTP methods that are idempotent and
// free of side effects (RFC 7231 §4.2.1). Safe methods skip CSRF check.
func isSafeMethod(m string) bool {
	switch strings.ToUpper(m) {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true
	}
	return false
}

// isExemptPath returns true when path begins with any ExemptPaths entry.
// Prefix matching — glob patterns are intentionally not supported to keep
// the match predictable for operators.
func isExemptPath(path string, exempt []string) bool {
	for _, p := range exempt {
		if p == "" {
			continue
		}
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

// hasBearerHeader detects an Authorization: Bearer <token> header used by
// API clients. Presence alone qualifies for the hybrid-mode skip; the
// bearer token itself is validated downstream by BearerAuthStrict.
func hasBearerHeader(c *gin.Context) bool {
	auth := c.GetHeader("Authorization")
	return strings.HasPrefix(strings.ToLower(auth), "bearer ")
}

// csrfAuthMode mirrors authMode() in auth_mode.go: the effective auth
// transport mode from BACKEND_AUTH_MODE, falling back to the build-time
// default when unset or outside the closed set. Duplicated here rather
// than calling authMode() so csrf.go stays self-contained.
func csrfAuthMode() string {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("BACKEND_AUTH_MODE")))
	switch v {
	case "bearer", "cookie", "hybrid":
		return v
	}
	return %q
}

// csrfRuntimeActive reports whether the live auth transport carries the
// session token in a cookie (mode cookie/hybrid) — the only case where
// CSRF defense applies. Bearer mode is CSRF-immune, so Csrf() no-ops.
func csrfRuntimeActive() bool {
	switch csrfAuthMode() {
	case "cookie", "hybrid":
		return true
	}
	return false
}

// constantTimeEqual compares two strings in constant time to prevent
// timing oracles on the token value.
func constantTimeEqual(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
`
