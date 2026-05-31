//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestIsPureNumeric(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"4", true},
		{"0.5", true},
		{"1/2", true},
		{"", true}, // empty has no non-numeric chars
		{"sm", false},
		{"4px", false},
		{"-4", false},
	}
	for _, tt := range tests {
		if got := isPureNumeric(tt.in); got != tt.want {
			t.Errorf("isPureNumeric(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
