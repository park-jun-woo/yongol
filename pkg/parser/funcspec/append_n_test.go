//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"reflect"
	"testing"
)

func TestAppendN(t *testing.T) {
	got := appendN(nil, "int", 3)
	want := []string{"int", "int", "int"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("appendN = %v, want %v", got, want)
	}
	// n=0 → unchanged.
	base := []string{"a"}
	if got := appendN(base, "x", 0); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("appendN n=0 = %v", got)
	}
}
