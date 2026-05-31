//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"testing"
)

func TestUcFirst(t *testing.T) {
	tests := []struct{ in, want string }{
		{"hashPassword", "HashPassword"},
		{"user_id", "UserID"}, // strcase ToGoPascal applies initialism handling
		{"id", "ID"},
	}
	for _, tt := range tests {
		if got := ucFirst(tt.in); got != tt.want {
			t.Errorf("ucFirst(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
