//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestSafetyAddColumnNotNull — NOT NULL + default/backfill 없으면 MIG-002 ERROR
package migration

import (
	"testing"
)

func TestSafetyAddColumnNotNull(t *testing.T) {
	cases := []struct {
		name string
		op   AddColumn
		want bool // true means MIG-002 emitted
	}{
		{"nullable ok", AddColumn{Table: "u", Column: &Column{Name: "c", Nullable: true}}, false},
		{"has default", AddColumn{Table: "u", Column: &Column{Name: "c", Nullable: false, Default: "0"}}, false},
		{"has backfill", AddColumn{Table: "u", Column: &Column{Name: "c", Nullable: false}, Backfill: "0"}, false},
		{"unsafe", AddColumn{Table: "u", Column: &Column{Name: "c", Nullable: false}}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertSafetyAddColumnNotNull(t, c.op, c.want)
		})
	}
}
