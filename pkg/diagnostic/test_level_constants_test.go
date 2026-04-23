//ff:func feature=orchestrator type=test control=sequence
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

// TestLevel_DistinctValues verifies that LevelError and LevelWarning have distinct values.
func TestLevel_DistinctValues(t *testing.T) {
	if diagnostic.LevelError == diagnostic.LevelWarning {
		t.Errorf("LevelError and LevelWarning must differ, both=%q", diagnostic.LevelError)
	}
}

// TestLevel_TypeIsStringAlias verifies that Level is convertible to and from string.
func TestLevel_TypeIsStringAlias(t *testing.T) {
	l := diagnostic.LevelError
	if string(l) != "ERROR" {
		t.Errorf("string(LevelError): want %q, got %q", "ERROR", string(l))
	}

	custom := diagnostic.Level("INFO")
	if string(custom) != "INFO" {
		t.Errorf("Level(\"INFO\"): round-trip failed, got %q", string(custom))
	}
}
