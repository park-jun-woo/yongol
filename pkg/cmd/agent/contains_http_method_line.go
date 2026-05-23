//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what containsHTTPMethodLine — 텍스트에 HTTP 메서드 라인 포함 여부 확인

package agent

import "strings"

// containsHTTPMethodLine checks if the text contains a line starting with an HTTP method.
func containsHTTPMethodLine(s string) bool {
	methods := []string{"GET ", "POST ", "PUT ", "DELETE ", "PATCH ", "HEAD ", "OPTIONS "}
	for _, line := range strings.Split(s, "\n") {
		trimmed := strings.TrimSpace(line)
		for _, m := range methods {
			if strings.HasPrefix(trimmed, m) {
				return true
			}
		}
	}
	return false
}
