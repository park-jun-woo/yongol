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
		"request-binding":      "collect request parameters and bind body",
		"response-serialize":   "compose response fields",
		"transaction-boundary": "BeginTx / Commit / Rollback",
		"state-transition":     "execute @state transition",
		"auth-check":           "@auth gate",
		"pagination":           "apply pagination",
		"error-mapping":        "classify validation / domain / infra errors",
		"observability":        "slog / metric / trace",
		"publish":              "queue publish",
		"subscribe":            "queue subscribe",
		"pointer-helper":       "ptr/deref generic helpers",
	}
}
