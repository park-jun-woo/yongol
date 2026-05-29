//ff:func feature=migration type=test control=sequence
//ff:what TestDropTable_Description — 설명에 테이블 이름 포함
package migration

import "testing"

func TestDropTable_Description(t *testing.T) {
	if got, want := (DropTable{Name: "users"}).Description(), "drop table users"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
