//ff:func feature=migration type=test control=sequence
//ff:what TestSafetyAddColumnNotNull — NOT NULL + default/backfill 없으면 MIG-002 ERROR
package migration

import "testing"

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
		c := c
		t.Run(c.name, func(t *testing.T) {
			issues := safetyAddColumnNotNull(c.op)
			if c.want {
				if len(issues) != 1 || issues[0].RuleID != "MIG-002" || issues[0].Level != SafetyError {
					t.Errorf("got %+v, want one MIG-002 error", issues)
				}
			} else if issues != nil {
				t.Errorf("want nil, got %+v", issues)
			}
		})
	}
}
