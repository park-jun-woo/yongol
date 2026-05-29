//ff:func feature=migration type=test control=sequence
//ff:what TestDropIndex_SQL — DROP INDEX 문장 렌더 확인
package migration

import "testing"

func TestDropIndex_SQL(t *testing.T) {
	if got, want := (DropIndex{Name: "idx_email"}).SQL(), "DROP INDEX idx_email;"; got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
