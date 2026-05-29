//ff:func feature=migration type=test control=sequence
//ff:what TestCreateIndex_Destructive — 인덱스 생성은 비파괴적(false)
package migration

import "testing"

func TestCreateIndex_Destructive(t *testing.T) {
	if (CreateIndex{}).Destructive() {
		t.Error("CreateIndex.Destructive() = true, want false")
	}
}
