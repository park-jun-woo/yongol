//ff:func feature=migration type=test control=sequence
//ff:what TestDropIndex_Description — 설명에 인덱스 이름 포함
package migration

import "testing"

func TestDropIndex_Description(t *testing.T) {
	if got, want := (DropIndex{Name: "idx_email"}).Description(), "drop index idx_email"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
