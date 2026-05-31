//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"
)

func TestNormalizeTypeName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"[]Reservation", "Reservation"},
		{"*User", "User"},
		{"User", "User"},
		{"[]*Order", "Order"}, // strips slice prefix, then pointer prefix
	}
	for _, tt := range tests {
		if got := normalizeTypeName(tt.in); got != tt.want {
			t.Errorf("normalizeTypeName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
