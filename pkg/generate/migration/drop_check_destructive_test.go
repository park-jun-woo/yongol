//ff:func feature=migration type=test control=sequence
//ff:what TestDropCheck_Destructive — CHECK 삭제는 비파괴적(false)
package migration

import "testing"

func TestDropCheck_Destructive(t *testing.T) {
	if (DropCheck{}).Destructive() {
		t.Error("DropCheck.Destructive() = true, want false")
	}
}
