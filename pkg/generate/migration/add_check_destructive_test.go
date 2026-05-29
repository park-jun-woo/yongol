//ff:func feature=migration type=test control=sequence
//ff:what TestAddCheck_Destructive — CHECK 추가는 비파괴적(false)
package migration

import "testing"

func TestAddCheck_Destructive(t *testing.T) {
	if (AddCheck{}).Destructive() {
		t.Error("AddCheck.Destructive() = true, want false")
	}
}
