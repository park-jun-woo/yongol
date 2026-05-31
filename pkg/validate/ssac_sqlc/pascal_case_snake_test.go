//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestPascalCaseSnake(t *testing.T) {
	tests := []struct{ in, want string }{
		{"a_b_c", "ABC"},
		{"user_id", "UserId"},
		{"_x_", "X"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := pascalCaseSnake(tt.in); got != tt.want {
			t.Errorf("pascalCaseSnake(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
