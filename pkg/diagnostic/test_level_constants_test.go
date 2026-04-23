//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what Lock-in test for Level constant values
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestLevelConstants locks in the string values of Level constants.
// Fails intentionally on value change — update this test when changing a constant.
func TestLevelConstants(t *testing.T) {
	cases := []struct {
		name string
		got  diagnostic.Level
		want string
	}{
		{"LevelError", diagnostic.LevelError, "ERROR"},
		{"LevelWarning", diagnostic.LevelWarning, "WARNING"},
	}

	for _, c := range cases {
		if string(c.got) != c.want {
			t.Errorf("%s value changed: want %q, got %q", c.name, c.want, string(c.got))
		}
	}
}
