//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what TestManifestHelpers — unit tests for the pure manifest parser helper functions
package manifest

import (
	"testing"
)

func TestIntToStr(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "0"},
		{42, "42"},
		{-7, "-7"},
		{900, "900"},
		{-1000000, "-1000000"},
	}
	for _, tt := range tests {
		if got := intToStr(tt.in); got != tt.want {
			t.Errorf("intToStr(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
