//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what kebab-case → PascalCase 컴포넌트명 변환 검증
package stml

import "testing"

func TestToComponentName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"login-page", "LoginPage"},
		{"my-reservations-page", "MyReservationsPage"},
		{"room-edit-page", "RoomEditPage"},
	}
	for _, tt := range tests {
		got := toComponentName(tt.in)
		if got != tt.want { t.Errorf("toComponentName(%q) = %q, want %q", tt.in, got, tt.want) }
	}
}
