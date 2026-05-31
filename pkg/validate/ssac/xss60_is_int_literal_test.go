//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what TestSSaCXss60Helpers — unit tests for the pure ssac validate helper functions
package ssac

import (
	"testing"
)

func TestXss60IsIntLiteral(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"0", true},
		{"42", true},
		{"-7", true},
		{"", false},
		{"-", false},
		{"3.14", false},
		{"abc", false},
		{"12a", false},
	}
	for _, tt := range tests {
		if got := xss60IsIntLiteral(tt.in); got != tt.want {
			t.Errorf("xss60IsIntLiteral(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
