//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestResolveFastAPIExpr — request.X → body.x 치환 + snake_case 변환 검증
package ssac

import (
	"testing"
)

func TestResolveFastAPIExpr(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"RequestEmail", "request.Email", "body.email"},
		{"RequestPassword", "request.Password", "body.password"},
		{"RequestCamelCase", "request.PasswordHash", "body.password_hash"},
		{"NoPrefix", "user.email", "user.email"},
		{"PlainVar", "email", "email"},
		{"EmptyString", "", ""},
		{"RequestOnly", "request.", "body."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveFastAPIExpr(tt.in)
			if got != tt.want {
				t.Errorf("resolveFastAPIExpr(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
