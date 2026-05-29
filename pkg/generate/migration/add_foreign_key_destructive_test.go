//ff:func feature=migration type=test control=sequence
//ff:what TestAddForeignKey_Destructive — FK 추가는 비파괴적(false)
package migration

import "testing"

func TestAddForeignKey_Destructive(t *testing.T) {
	if (AddForeignKey{}).Destructive() {
		t.Error("AddForeignKey.Destructive() = true, want false")
	}
}
