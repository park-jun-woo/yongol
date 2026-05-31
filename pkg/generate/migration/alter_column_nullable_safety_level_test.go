//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestAlterColumnNullable_SafetyLevel — NOT NULL 추가+backfill 없으면 Error
package migration

import (
	"testing"
)

func TestAlterColumnNullable_SafetyLevel(t *testing.T) {
	cases := []struct {
		name string
		op   AlterColumnNullable
		want SafetyLevel
	}{
		{"drop not null", AlterColumnNullable{To: true}, SafetySafe},
		{"set not null no backfill", AlterColumnNullable{To: false}, SafetyError},
		{"set not null with backfill", AlterColumnNullable{To: false, Backfill: "0"}, SafetySafe},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := c.op.SafetyLevel(); got != c.want {
				t.Errorf("SafetyLevel() = %v, want %v", got, c.want)
			}
		})
	}
}
