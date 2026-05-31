//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestInferLiteralType(t *testing.T) {
	tests := []struct{ in, want string }{
		{`"foo"`, "string"},
		{"42", "int64"},
		{"3.14", "float64"},
		{"true", "bool"},
		{"false", "bool"},
		{"nil", "nil"},
		{"someVar", ""},
		{"course.ID", ""},
		{"", ""},
		{`  "spaced"  `, "string"},
	}
	for _, tt := range tests {
		if got := inferLiteralType(tt.in); got != tt.want {
			t.Errorf("inferLiteralType(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
