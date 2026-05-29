//ff:func feature=migration type=test control=sequence
//ff:what TestDropForeignKey_Destructive — FK 삭제는 비파괴적(false)
package migration

import "testing"

func TestDropForeignKey_Destructive(t *testing.T) {
	if (DropForeignKey{}).Destructive() {
		t.Error("DropForeignKey.Destructive() = true, want false")
	}
}
