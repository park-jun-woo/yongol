//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestCatalogHelpers — unit tests for the pure rule catalog helper functions
package catalog

import (
	"reflect"
	"testing"
)

func TestSplitRow(t *testing.T) {
	got := splitRow("| a | b | c |")
	want := []string{" a ", " b ", " c "}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitRow = %q, want %q", got, want)
	}
	// Without surrounding pipes.
	if got := splitRow("x|y"); !reflect.DeepEqual(got, []string{"x", "y"}) {
		t.Errorf("no-pipe-edge = %q", got)
	}
}
