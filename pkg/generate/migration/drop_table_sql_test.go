//ff:func feature=migration type=test control=sequence
//ff:what TestDropTable_SQL — DROP TABLE 문장 렌더 확인
package migration

import "testing"

func TestDropTable_SQL(t *testing.T) {
	if got, want := (DropTable{Name: "users"}).SQL(), "DROP TABLE users;"; got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
