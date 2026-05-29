//ff:func feature=migration type=test control=sequence
//ff:what TestInsertSentinel_Destructive — sentinel INSERT 는 멱등이라 false
package migration

import "testing"

func TestInsertSentinel_Destructive(t *testing.T) {
	if (InsertSentinel{Table: "roles"}).Destructive() {
		t.Error("InsertSentinel.Destructive() = true, want false")
	}
}
