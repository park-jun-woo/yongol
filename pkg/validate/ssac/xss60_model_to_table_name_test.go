//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what TestSSaCXss60Helpers — unit tests for the pure ssac validate helper functions
package ssac

import (
	"testing"
)

func TestXss60ModelToTableName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"User", "users"},
		{"RefreshToken", "refresh_tokens"},
		{"AuditLog", "audit_logs"},
	}
	for _, tt := range tests {
		if got := xss60ModelToTableName(tt.in); got != tt.want {
			t.Errorf("xss60ModelToTableName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
