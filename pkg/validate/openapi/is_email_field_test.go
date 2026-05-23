//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what isEmailField — email/접두사/접미사/무관 필드명 검증

package openapi

import "testing"

func TestIsEmailField(t *testing.T) {
	tests := []struct {
		name  string
		field string
		want  bool
	}{
		{"exact email", "email", true},
		{"upper Email", "Email", true},
		{"suffix user_email", "user_email", true},
		{"prefix email_address", "email_address", true},
		{"unrelated name", "username", false},
		{"empty string", "", false},
		{"EMAIL all caps", "EMAIL", true},
		{"myEmail camel", "myEmail", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isEmailField(tt.field)
			if got != tt.want {
				t.Errorf("isEmailField(%q) = %v, want %v", tt.field, got, tt.want)
			}
		})
	}
}
