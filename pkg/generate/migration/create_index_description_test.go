//ff:func feature=migration type=test control=sequence
//ff:what TestCreateIndex_Description — 설명에 인덱스 이름 포함
package migration

import "testing"

func TestCreateIndex_Description(t *testing.T) {
	op := CreateIndex{Index: &Index{Name: "idx_email"}}
	if got, want := op.Description(), "create index idx_email"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
