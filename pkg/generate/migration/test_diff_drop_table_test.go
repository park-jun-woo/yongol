//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_DropTable — prev 에만 있는 테이블 → DropTable 생성
package migration

import "testing"

func TestDiff_DropTable(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	curr := NewSchema()
	ops := Diff(prev, curr, nil)
	foundDrop := false
	for _, op := range ops {
		if _, ok := op.(DropTable); ok {
			foundDrop = true
		}
	}
	if !foundDrop {
		t.Errorf("DropTable missing from ops: %+v", ops)
	}
}
