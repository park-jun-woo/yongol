//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what isHTTPMethod — YAML 키가 HTTP 메서드인지 확인

package agent

import "strings"

// isHTTPMethod checks if a trimmed line starts with a YAML key for an HTTP method.
func isHTTPMethod(trimmed string) bool {
	methods := []string{"get:", "post:", "put:", "delete:", "patch:", "head:", "options:"}
	lower := strings.ToLower(trimmed)
	for _, m := range methods {
		if lower == m || strings.HasPrefix(lower, m+" ") {
			return true
		}
	}
	return false
}
