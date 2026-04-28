//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what isAuthPath — /auth/, /login, /signin, /register, /signup 경로 heuristic

package hurl_manifest

import "strings"

// isAuthPath matches the path conventions yongol recognises as auth
// endpoints. Matching the path (not the operationId) lets the rule work
// even when the OpenAPI document was not parsed.
func isAuthPath(p string) bool {
	low := strings.ToLower(p)
	if strings.Contains(low, "/auth/") {
		return true
	}
	for _, suffix := range []string{"/login", "/signin", "/register", "/signup"} {
		if strings.HasSuffix(low, suffix) || strings.Contains(low, suffix+"/") {
			return true
		}
	}
	return false
}
