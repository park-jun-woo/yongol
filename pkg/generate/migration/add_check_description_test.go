//ff:func feature=migration type=test control=sequence
//ff:what TestAddCheck_Description — 설명 문자열에 제약 이름 포함 확인
package migration

import "testing"

func TestAddCheck_Description(t *testing.T) {
	op := AddCheck{Table: "users", Check: &CheckConstraint{Name: "users_age_check"}}
	if got, want := op.Description(), "add check users_age_check"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
