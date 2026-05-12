//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what contains — 문자열 포함 여부 확인 (strings import 회피용)
package domain_security

// contains checks if s contains substr.
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
