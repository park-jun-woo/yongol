//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"testing"
)

func TestIsZeroExpr(t *testing.T) {
	tests := []struct {
		src  string
		want bool
	}{
		{"nil", true},
		{"false", true},
		{"0", true},
		{"0.0", true},
		{`""`, true},
		{"Resp{}", true},
		{"42", false},
		{`"hello"`, false},
		{"true", false},
		{"Resp{Status: 1}", false},
		{"someVar", false},
	}
	for _, tt := range tests {
		if got := isZeroExpr(parseExpr(t, tt.src)); got != tt.want {
			t.Errorf("isZeroExpr(%q) = %v, want %v", tt.src, got, tt.want)
		}
	}
}
