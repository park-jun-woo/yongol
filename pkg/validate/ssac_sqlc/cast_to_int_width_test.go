//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestCastToIntWidth(t *testing.T) {
	tests := []struct{ in, want string }{
		{"bigint", "int64"},
		{"int8", "int64"},
		{"int", "int32"},
		{"int4", "int32"},
		{"integer", "int32"},
		{"", "int32"},
		{"text", ""},
		{"numeric", ""},
	}
	for _, tt := range tests {
		if got := castToIntWidth(tt.in); got != tt.want {
			t.Errorf("castToIntWidth(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
