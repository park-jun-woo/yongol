//ff:func feature=validate type=rule control=selection topic=hurl-manifest
//ff:what isMutating — HTTP method 가 POST/PUT/PATCH/DELETE 인지 판정

package hurl_manifest

import "strings"

// isMutating reports whether an HTTP method causes server state to
// change — these are the methods covered by the CSRF middleware.
func isMutating(method string) bool {
	switch strings.ToUpper(method) {
	case "POST", "PUT", "PATCH", "DELETE":
		return true
	}
	return false
}
