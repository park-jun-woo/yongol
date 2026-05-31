//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestCallFuncName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"auth.VerifyPassword", "VerifyPassword"},
		{"plainname", ""},
		{".trailing", ""}, // idx <= 0
		{"pkg.", ""},      // idx >= len-1
	}
	for _, tt := range tests {
		if got := callFuncName(tt.in); got != tt.want {
			t.Errorf("callFuncName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
