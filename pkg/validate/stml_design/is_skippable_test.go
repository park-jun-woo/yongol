//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestIsSkippable(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"[10px]", true},
		{"gray-500", true},
		{"full", true},
		{"none", true},
		{"semibold", true},
		{"4", true},
		{"0.5", true},
		{"primary", false},
		{"display", false},
	}
	for _, tt := range tests {
		if got := isSkippable(tt.in); got != tt.want {
			t.Errorf("isSkippable(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
