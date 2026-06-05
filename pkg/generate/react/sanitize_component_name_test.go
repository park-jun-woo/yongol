//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what sanitizeComponentName 선행 숫자 Page 접두사 살균 테이블 검증

package react

import "testing"

func TestSanitizeComponentName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},           // 빈 문자열은 그대로
		{"Login", "Login"}, // 정상 식별자는 그대로
		{"workflowDetail", "workflowDetail"},
		{"_private", "_private"},     // 언더스코어 시작은 유효
		{"$ref", "$ref"},             // 달러 시작은 유효
		{"2faSetup", "Page2faSetup"}, // 숫자 시작 → Page 접두사
		{"3dView", "Page3dView"},
		{"404Error", "Page404Error"},
		{"0", "Page0"},   // 경계: 단일 숫자
		{"9z", "Page9z"}, // 경계: 최대 숫자
	}
	for _, tt := range tests {
		got := sanitizeComponentName(tt.in)
		if got != tt.want {
			t.Errorf("sanitizeComponentName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
