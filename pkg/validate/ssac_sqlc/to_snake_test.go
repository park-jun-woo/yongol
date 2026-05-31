//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestToSnake(t *testing.T) {
	tests := []struct{ in, want string }{
		{"UserId", "user_id"},
		{"AuditLog", "audit_log"},
		{"User", "user"},
	}
	for _, tt := range tests {
		if got := toSnake(tt.in); got != tt.want {
			t.Errorf("toSnake(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
