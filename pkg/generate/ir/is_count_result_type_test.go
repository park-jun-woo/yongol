//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what convertPublish/resolveExposeInternal/isCountResultType/ddlTableSingularIR/DDLTableSingularIR/findDDLTable
package ir

import (
	"testing"
)

func TestIsCountResultType(t *testing.T) {
	for _, ty := range []string{"int64", "int32", "int", "uint64"} {
		if !isCountResultType(ty) {
			t.Errorf("%q should be count type", ty)
		}
	}
	for _, ty := range []string{"string", "bool", "float64", "User"} {
		if isCountResultType(ty) {
			t.Errorf("%q should not be count type", ty)
		}
	}
}
