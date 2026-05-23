//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what isPasswordField — password/접두사/접미사/무관 필드명 검증

package openapi

import "testing"

func TestIsPasswordField(t *testing.T) {
	tests := []struct {
		name  string
		field string
		want  bool
	}{
		{"exact password", "password", true},
		{"upper Password", "Password", true},
		{"suffix old_password", "old_password", true},
		{"prefix password_hash", "password_hash", true},
		{"unrelated name", "username", false},
		{"empty string", "", false},
		{"PASSWORD all caps", "PASSWORD", true},
		{"myPassword camel", "myPassword", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPasswordField(tt.field)
			if got != tt.want {
				t.Errorf("isPasswordField(%q) = %v, want %v", tt.field, got, tt.want)
			}
		})
	}
}
