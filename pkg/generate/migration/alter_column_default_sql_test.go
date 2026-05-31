//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnDefault_SQL — To 빈값이면 DROP DEFAULT, 아니면 SET DEFAULT
package migration

import (
	"testing"
)

func TestAlterColumnDefault_SQL(t *testing.T) {
	if got, want := (AlterColumnDefault{Table: "t", Column: "c", To: ""}).SQL(), "ALTER TABLE t ALTER COLUMN c DROP DEFAULT;"; got != want {
		t.Errorf("drop: %q, want %q", got, want)
	}
	if got, want := (AlterColumnDefault{Table: "t", Column: "c", To: "0"}).SQL(), "ALTER TABLE t ALTER COLUMN c SET DEFAULT 0;"; got != want {
		t.Errorf("set: %q, want %q", got, want)
	}
}
