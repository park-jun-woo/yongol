//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"testing"
)

func TestIsIntType(t *testing.T) {
	for _, ty := range []string{"int", "int8", "int16", "int32", "int64", "uint", "uint64", "byte", "rune"} {
		if !isIntType(ty) {
			t.Errorf("%q should be int type", ty)
		}
	}
	for _, ty := range []string{"string", "float64", "bool", ""} {
		if isIntType(ty) {
			t.Errorf("%q should not be int type", ty)
		}
	}
}
