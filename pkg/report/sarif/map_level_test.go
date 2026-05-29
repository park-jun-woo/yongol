//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestMapLevel — ERROR/WARNING/그외 → SARIF level 문자열 매핑 (table)
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestMapLevel covers all three branches of the severity switch.
func TestMapLevel(t *testing.T) {
	cases := []struct {
		in   diagnostic.Level
		want string
	}{
		{diagnostic.LevelError, "error"},
		{diagnostic.LevelWarning, "warning"},
		{diagnostic.Level("INFO"), "note"},
		{diagnostic.Level(""), "note"},
	}
	for _, c := range cases {
		if got := mapLevel(c.in); got != c.want {
			t.Errorf("mapLevel(%q): got %q, want %q", c.in, got, c.want)
		}
	}
}
