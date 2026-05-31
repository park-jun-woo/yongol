//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestToCamelKey(t *testing.T) {
	tests := []struct{ in, want string }{
		{"HashPassword", "hashPassword"},
		{"hashPassword", "hashPassword"},
		{"", ""},
		{"X", "x"},
	}
	for _, tt := range tests {
		if got := toCamelKey(tt.in); got != tt.want {
			t.Errorf("toCamelKey(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
