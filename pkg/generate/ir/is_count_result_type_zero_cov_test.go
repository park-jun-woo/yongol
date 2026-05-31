//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"
)

func TestIsCountResultType_ZeroCov(t *testing.T) {
	for _, ty := range []string{"int64", "int32", "int", "uint64"} {
		if !isCountResultType(ty) {
			t.Errorf("%q should be count", ty)
		}
	}
	if isCountResultType("Course") {
		t.Error("Course should not be count")
	}
}
