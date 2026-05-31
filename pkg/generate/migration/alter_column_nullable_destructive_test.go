//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnNullable_Destructive — NOT NULL 추가(To=false)만 파괴적
package migration

import (
	"testing"
)

func TestAlterColumnNullable_Destructive(t *testing.T) {
	if (AlterColumnNullable{To: true}).Destructive() {
		t.Error("To=true (drop NOT NULL) should be non-destructive")
	}
	if !(AlterColumnNullable{To: false}).Destructive() {
		t.Error("To=false (set NOT NULL) should be destructive")
	}
}
