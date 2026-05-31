//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"testing"
)

func TestExprToString(t *testing.T) {
	tests := []struct{ src, want string }{
		{"int", "int"},
		{"pkg.Type", "pkg.Type"},
		{"*User", "*User"},
		{"[]string", "[]string"},
		{"map[string]int", "map[string]int"},
		{"[]*pkg.Type", "[]*pkg.Type"},
	}
	for _, tt := range tests {
		if got := exprToString(parseExpr(t, tt.src)); got != tt.want {
			t.Errorf("exprToString(%q) = %q, want %q", tt.src, got, tt.want)
		}
	}
	// Unknown expr falls back to interface{}.
	if got := exprToString(parseExpr(t, "func(){}")); got != "interface{}" {
		t.Errorf("func type → %q, want interface{}", got)
	}
}
