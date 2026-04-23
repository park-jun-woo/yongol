//ff:func feature=validate type=util control=selection topic=tsx-openapi
//ff:what XOT-2 헬퍼 — body/data/payload/json 같은 transport wrapper key 여부 판정

package tsx_openapi

// isTransportKey returns true for argument keys that are wrappers around
// the request body rather than actual OpenAPI parameters. XOT-3 validates
// the body contents; XOT-2 must skip them to avoid double-counting.
func isTransportKey(key string) bool {
	switch key {
	case "body", "data", "payload", "json":
		return true
	}
	return false
}
