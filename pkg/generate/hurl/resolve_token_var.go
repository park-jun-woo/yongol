//ff:func feature=gen-hurl type=util control=sequence
//ff:what resolveTokenVar — operation에 맞는 token 변수명 선택 (role 우선, alpha fallback)
package hurl

// resolveTokenVar picks the token variable name for an operation.
// Priority:
//  1. role-specific token from roleMap (token_{role}) if captured
//  2. plain "token" (single-role mode) if captured
//  3. first captured token_* alphabetically
//  4. "" — no token available
func resolveTokenVar(operationID string, roleMap map[string]string, captures map[string]bool) string {
	if role, ok := roleMap[operationID]; ok {
		roleToken := "token_" + role
		if captures[roleToken] {
			return roleToken
		}
	}
	if captures["token"] {
		return "token"
	}
	return firstTokenCapture(captures)
}
