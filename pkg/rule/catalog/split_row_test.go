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
	// GFM escaped pipes (`\|`) stay inside the cell, restored to literal `|`.
	got = splitRow(`| a \|\| b | c |`)
	want = []string{" a || b ", " c "}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("escaped-pipe = %q, want %q", got, want)
	}
	got = splitRow(`| cookie\|hybrid | x |`)
	want = []string{" cookie|hybrid ", " x "}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("escaped-pipe-mid = %q, want %q", got, want)
	}
	// Non-pipe backslash sequences are kept verbatim; trailing backslash too.
	got = splitRow(`| a\b | c\ |`)
	want = []string{` a\b `, ` c\ `}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("backslash-verbatim = %q, want %q", got, want)
	}
}
