//ff:func feature=manifest type=test control=sequence
//ff:what extractVarcharLen — VARCHAR(64) → 64

package ddl

import "testing"

func TestExtractVarcharLen_Match(t *testing.T) {
	if got := extractVarcharLen("VARCHAR(64)"); got != 64 {
		t.Errorf("got %d, want 64", got)
	}
}
