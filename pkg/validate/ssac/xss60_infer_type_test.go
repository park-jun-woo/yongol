//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what TestSSaCXss60Helpers — unit tests for the pure ssac validate helper functions
package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60InferType(t *testing.T) {
	fn := parsessac.ServiceFunc{}
	tests := []struct{ expr, want string }{
		{`"completed"`, "string"},
		{"0", "int64"},
		{"42", "int64"},
		{"", ""},
		{"unknownvar", ""}, // no dot, not literal
	}
	for _, tt := range tests {
		if got := xss60InferType(tt.expr, fn, nil); got != tt.want {
			t.Errorf("xss60InferType(%q) = %q, want %q", tt.expr, got, tt.want)
		}
	}
}
