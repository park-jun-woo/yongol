//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestUcFirst(t *testing.T) {
	tests := []struct{ in, want string }{
		{"hash", "Hash"},
		{"Hash", "Hash"},
		{"", ""},
		{"1abc", "1abc"},
	}
	for _, tt := range tests {
		if got := ucFirst(tt.in); got != tt.want {
			t.Errorf("ucFirst(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
