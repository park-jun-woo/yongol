//ff:func feature=agent type=helper control=sequence
//ff:what domainFromPath — API path에서 도메인 이름 추출 (/workflows/{id} -> workflow)

package agent

import "strings"

// domainFromPath extracts the domain name from an API path.
// /workflows/{id} -> workflow
// /auth/login -> auth
// /payment-intents/{id} -> payment_intent
func domainFromPath(path string) string {
	fields := strings.Fields(path)
	if len(fields) >= 2 {
		path = fields[1]
	}
	path = strings.TrimPrefix(path, "/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) == 0 || parts[0] == "" {
		return "default"
	}

	segment := parts[0]
	// Convert kebab-case to snake_case
	segment = strings.ReplaceAll(segment, "-", "_")
	// Remove trailing 's' for plural to singular (simple heuristic)
	if strings.HasSuffix(segment, "s") && !strings.HasSuffix(segment, "ss") {
		segment = segment[:len(segment)-1]
	}
	return segment
}
