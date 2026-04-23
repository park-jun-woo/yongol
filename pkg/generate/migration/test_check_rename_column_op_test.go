//ff:func feature=migration type=test-helper control=selection
//ff:what checkRenameColumnOp — 한 Operation 이 RenameColumn(email→email_address) 인지 / DropColumn 인지 검사
package migration

import "testing"

// checkRenameColumnOp returns the updated foundRename flag after
// inspecting one op. It also fails the test on unexpected DropColumn.
func checkRenameColumnOp(t *testing.T, op Operation, foundRename bool, ops []Operation) bool {
	t.Helper()
	switch v := op.(type) {
	case RenameColumn:
		if v.From != "email" || v.To != "email_address" {
			t.Errorf("rename wrong: %+v", v)
		}
		return true
	case DropColumn:
		t.Errorf("unexpected DropColumn with rename hint: %+v", ops)
	}
	return foundRename
}
