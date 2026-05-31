//ff:func feature=crosscheck type=test control=iteration dimension=1 topic=scenario-check
//ff:what TestHurlHelpers — unit tests for the pure hurl parser helper functions
package hurl

import (
	"testing"
)

func TestTrimQuery(t *testing.T) {
	tests := []struct{ in, want string }{
		{"/users?page=2", "/users"},
		{"/users", "/users"},
		{"/a?b?c", "/a"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := trimQuery(tt.in); got != tt.want {
			t.Errorf("trimQuery(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
