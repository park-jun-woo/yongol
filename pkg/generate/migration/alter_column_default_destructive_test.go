//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnDefault_Destructive — 기본값 변경은 비파괴적(false)
package migration

import "testing"

func TestAlterColumnDefault_Destructive(t *testing.T) {
	if (AlterColumnDefault{}).Destructive() {
		t.Error("AlterColumnDefault.Destructive() = true, want false")
	}
}
