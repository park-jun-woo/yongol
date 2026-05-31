//ff:func feature=validate type=test control=iteration dimension=1 topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		s, sub string
		want   bool
	}{
		{"deleteworkflow", "workflow", true},
		{"abc", "", true},   // empty substr always contained
		{"a", "abc", false}, // substr longer than s
		{"hello", "ell", true},
		{"hello", "xyz", false},
		{"hello", "hello", true},
	}
	for _, tt := range tests {
		if got := contains(tt.s, tt.sub); got != tt.want {
			t.Errorf("contains(%q,%q) = %v, want %v", tt.s, tt.sub, got, tt.want)
		}
	}
}
