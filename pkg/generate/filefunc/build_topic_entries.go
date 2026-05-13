//ff:func feature=gen-filefunc type=util control=sequence
//ff:what buildTopicEntries — returns the fixed topic map (no dynamic expansion in Phase001)
package filefunc

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildTopicEntries returns the topic catalogue. For Phase001 we emit a fixed
// baseline set that covers the semantic units produced by SSaC @call /
// @publish / @subscribe / @state. fs is accepted for future dynamic expansion
// but currently unused.
func buildTopicEntries(fs *yongol.Fullstack) map[string]string {
	_ = fs
	return map[string]string{
		"auth-check":           "@auth gate",
		"auth-refresh":         "refresh-token store (ssac/pkg/auth.RefreshStore)",
		"dos-guard":            "HTTP body size limit (DoS prevention)",
		"error-envelope":       "canonical JSON error envelope",
		"error-mapping":        "classify validation / domain / infra errors",
		"observability":        "slog / metric / trace",
		"pagination":           "apply pagination",
		"pointer-helper":       "ptr/deref generic helpers",
		"publish":              "queue publish",
		"rate-limit":           "fixed rate limit guard",
		"request-binding":      "collect request parameters and bind body",
		"request-id":           "ULID-based request id middleware",
		"response-serialize":   "compose response fields",
		"security-headers":     "browser security headers (HSTS/CSP/XFO/etc.)",
		"state-transition":     "execute @state transition",
		"subscribe":            "queue subscribe",
		"transaction-boundary": "BeginTx / Commit / Rollback",
	}
}
