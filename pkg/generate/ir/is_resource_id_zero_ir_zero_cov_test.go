//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"
)

func TestIsResourceIDZeroIR_ZeroCov(t *testing.T) {
	for _, z := range []string{"", "  ", "0", `""`, "''", "nil", "NULL", "Null"} {
		if !isResourceIDZeroIR(z) {
			t.Errorf("%q should be zero", z)
		}
	}
	for _, nz := range []string{"project.ID", "5", "x"} {
		if isResourceIDZeroIR(nz) {
			t.Errorf("%q should be non-zero", nz)
		}
	}
}
