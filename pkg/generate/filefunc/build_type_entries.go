//ff:func feature=gen-filefunc type=util control=sequence
//ff:what buildTypeEntries — returns the fixed type category map for generated Go+Gin backend code
package filefunc

// buildTypeEntries returns the fixed type catalogue for generated Go+Gin
// backend code. Keys must stay in sync with //ff:type values emitted by the
// code generator (Phase002+).
func buildTypeEntries() map[string]string {
	return map[string]string{
		"handler":     "HTTP request handler (Gin handler)",
		"service":     "business logic function (SSaC @func)",
		"model":       "data transfer object (DTO)",
		"query":       "sqlc query wrapper",
		"middleware":  "Gin middleware",
		"config":      "environment variable and runtime configuration",
		"accessor":    "getter/setter/reexport",
		"util":        "utility function",
		"generator":   "code generation helper",
		"loader":      "initialization/loader",
		"command":     "entry-point function",
		"test":        "test function",
		"test-helper": "test helper",
	}
}
