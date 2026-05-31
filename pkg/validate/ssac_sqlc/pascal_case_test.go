//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestPascalCase(t *testing.T) {
	tests := []struct{ in, want string }{
		{"", ""},
		{"user_id", "UserId"},
		{"user", "User"},
		{"User", "User"},
		{"userName", "UserName"},
		{"audit_log_entry", "AuditLogEntry"},
		{"__leading", "Leading"},
	}
	for _, tt := range tests {
		if got := pascalCase(tt.in); got != tt.want {
			t.Errorf("pascalCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
