//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what isAuthLayout 인증 레이아웃 판별 테이블 검증

package react

import "testing"

func TestIsAuthLayout(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"auth", true},
		{"app", false},
		{"admin", false},
		{"", false},
	}
	for _, tt := range tests {
		got := isAuthLayout(tt.name)
		if got != tt.want {
			t.Errorf("isAuthLayout(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}
