//ff:func feature=validate type=test control=iteration dimension=1 topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"
)

func TestNormalizeCallKey(t *testing.T) {
	tests := []struct{ in, want string }{
		{"session.GetUser", "session.getUser"},
		{"cache.Set", "cache.set"},
		{"VerifyPassword", "verifyPassword"},
	}
	for _, tt := range tests {
		if got := normalizeCallKey(tt.in); got != tt.want {
			t.Errorf("normalizeCallKey(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
