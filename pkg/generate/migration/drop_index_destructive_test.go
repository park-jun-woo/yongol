//ff:func feature=migration type=test control=sequence
//ff:what TestDropIndex_Destructive — 인덱스 삭제는 비파괴적(false)
package migration

import "testing"

func TestDropIndex_Destructive(t *testing.T) {
	if (DropIndex{Name: "ix"}).Destructive() {
		t.Error("DropIndex.Destructive() = true, want false")
	}
}
