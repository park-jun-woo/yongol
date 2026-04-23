//ff:func feature=rule type=test control=iteration dimension=1
//ff:what populateSSaCSymbols — DDL 테이블 부재 시 Struct.* 키 미등록

package ground

import (
	"testing"
)

// TestPopulateSSaCSymbols_EmptyTables: no panics / no Struct.* keys when no
// DDL tables.
func TestPopulateSSaCSymbols_EmptyTables(t *testing.T) {
	g := newGround()
	populateSSaCSymbols(g, newMinimalFullstack())

	for k := range g.Types {
		if len(k) >= len("Struct.") && k[:len("Struct.")] == "Struct." {
			t.Errorf("unexpected Struct.* key %q when no DDL tables", k)
		}
	}
}
