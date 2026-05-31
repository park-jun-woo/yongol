//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"testing"
)

func TestParseTwoQuoted(t *testing.T) {
	first, second, rem := parseTwoQuoted(`"a" "b" tail`)
	if first != "a" || second != "b" || rem != "tail" {
		t.Errorf("got (%q,%q,%q)", first, second, rem)
	}
	// Only one quoted → second empty.
	f2, s2, r2 := parseTwoQuoted(`"only"`)
	if f2 != "only" || s2 != "" || r2 != "" {
		t.Errorf("one-quoted = (%q,%q,%q)", f2, s2, r2)
	}
}
