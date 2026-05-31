//ff:func feature=gen-ir type=test-helper control=selection
//ff:what assertResponseFieldSource — 변수 shadowing 해소 후 response 필드 Source 가 기대값인지 검증 헬퍼
package ir

import "testing"

// assertResponseFieldSource asserts the rewritten Source for a known response
// field name (user is renamed to user_result; token is unchanged).
func assertResponseFieldSource(t *testing.T, name, source string) {
	t.Helper()
	switch name {
	case "email":
		if source != "user_result.Email" {
			t.Errorf("email.Source = %q, want %q", source, "user_result.Email")
		}
	case "access_token":
		if source != "token.AccessToken" {
			t.Errorf("access_token.Source = %q, want %q", source, "token.AccessToken")
		}
	case "refresh_token":
		if source != "token.RefreshToken" {
			t.Errorf("refresh_token.Source = %q, want %q", source, "token.RefreshToken")
		}
	}
}
