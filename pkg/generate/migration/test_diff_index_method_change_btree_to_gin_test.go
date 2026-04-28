//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_IndexMethodChange_BtreeToGIN — btree→gin 변경 시 DROP+CREATE emit

package migration

import "testing"

// TestDiff_IndexMethodChange_BtreeToGIN verifies that changing an index
// method from btree → gin emits DROP + CREATE.
func TestDiff_IndexMethodChange_BtreeToGIN(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, c JSONB);
CREATE INDEX idx_t_c ON t (c);`)
	curr := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, c JSONB);
CREATE INDEX idx_t_c ON t USING GIN (c);`)
	ops := Diff(prev, curr, nil)
	var sawDrop, sawCreate bool
	for _, op := range ops {
		sawDrop = sawDrop || isDropIndex(op)
		sawCreate = sawCreate || inspectCreateIndexForMethod(t, op, "gin")
	}
	if !sawDrop || !sawCreate {
		t.Errorf("expected DROP + CREATE for method change, got: %+v", ops)
	}
}
