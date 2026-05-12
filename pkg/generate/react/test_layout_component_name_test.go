//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what layoutComponentName 변환 테이블 검증

package react

import "testing"

func TestLayoutComponentName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"app", "AppLayout"},
		{"auth", "AuthLayout"},
		{"main-nav", "MainNavLayout"},
		{"admin-panel", "AdminPanelLayout"},
	}
	for _, tt := range tests {
		got := layoutComponentName(tt.in)
		if got != tt.want {
			t.Errorf("layoutComponentName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
