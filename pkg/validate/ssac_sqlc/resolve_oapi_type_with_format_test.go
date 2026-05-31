//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestResolveOAPITypeWithFormat(t *testing.T) {
	tests := []struct{ base, format, want string }{
		{"integer", "int64", "int64"},
		{"integer", "int32", "int32"},
		{"number", "double", "float64"},
		{"number", "float", "float32"},
		{"string", "uuid", "uuid"},
		{"string", "", "string"},         // empty format → base
		{"integer", "weird", "integer"},  // unknown format → base
		{"object", "anything", "object"}, // unknown base → base
	}
	for _, tt := range tests {
		if got := resolveOAPITypeWithFormat(tt.base, tt.format); got != tt.want {
			t.Errorf("resolveOAPITypeWithFormat(%q,%q) = %q, want %q", tt.base, tt.format, got, tt.want)
		}
	}
}
