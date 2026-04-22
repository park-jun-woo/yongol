//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what canResolvePathParams — path의 모든 {param}이 captures로 해소 가능한지 검사
package hurl

import "strings"

// canResolvePathParams returns true when every {param} in the path has a
// corresponding captured hurl variable. Derivation mirrors resolveParamVar:
//   1. direct match on snakeHurlName(param)
//   2. precedingResource + "_id" (for plain "id" params)
// If any segment is unresolvable, returns false → caller skips the step.
func canResolvePathParams(path string, captures map[string]bool) bool {
	if captures == nil {
		captures = map[string]bool{}
	}
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if !isPathParamResolvable(parts, i, part, captures) {
			return false
		}
	}
	return true
}
