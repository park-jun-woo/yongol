//ff:func feature=manifest type=test control=sequence
//ff:what extractVarcharLen — VARCHAR 포맷 아닌 TEXT 는 0 반환

package ddl

import "testing"

func TestExtractVarcharLen_NoMatch(t *testing.T) {
	if got := extractVarcharLen("TEXT"); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
