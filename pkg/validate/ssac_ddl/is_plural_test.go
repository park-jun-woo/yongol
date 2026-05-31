//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"
)

func TestIsPlural(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"Workflows", true},
		{"Workflow", false},
		{"users", true},
		{"user", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := isPlural(tt.in); got != tt.want {
			t.Errorf("isPlural(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
