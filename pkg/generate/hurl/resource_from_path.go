//ff:func feature=gen-hurl type=util control=sequence
//ff:what resourceFromPath — URL path에서 리소스명을 snake_case로 추출
package hurl

import "strings"

// resourceFromPath extracts the leading resource segment from a URL path,
// strips a trailing "s" (simple plural), and normalizes to snake_case.
// e.g. "/audit-logs/{id}" → "audit_log".
func resourceFromPath(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return ""
	}
	return snakeHurlName(strings.TrimSuffix(parts[0], "s"))
}
