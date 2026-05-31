//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"testing"
)

func TestIsLiteral(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"true", true},
		{"false", true},
		{"nil", true},
		{"42", true},
		{"-1", true},
		{"3.14", true},
		{"-0.5", true},
		{"", false},
		{"-", false},
		{"foo", false},
		{"1.2.3", false}, // second dot → not numeric
		{"course.ID", false},
	}
	for _, tt := range tests {
		if got := IsLiteral(tt.in); got != tt.want {
			t.Errorf("IsLiteral(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
