//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what TestSSaCXss60Helpers — unit tests for the pure ssac validate helper functions
package ssac

import (
	"testing"
)

func TestExpectedQueryFile(t *testing.T) {
	tests := []struct{ in, want string }{
		{"RefreshToken", "db/queries/refresh_tokens.sql"},
		{"User", "db/queries/users.sql"},
	}
	for _, tt := range tests {
		if got := expectedQueryFile(tt.in); got != tt.want {
			t.Errorf("expectedQueryFile(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
