//ff:func feature=migration type=test control=selection dimension=4
//ff:what TestAddColumn_SafetyLevel — NOT NULL + default/backfill 조합별 안전등급
package migration

import "testing"

func TestAddColumn_SafetyLevel(t *testing.T) {
	cases := []struct {
		name string
		op   AddColumn
		want SafetyLevel
	}{
		{"nullable", AddColumn{Column: &Column{Nullable: true}}, SafetySafe},
		{"not-null no default no backfill", AddColumn{Column: &Column{Nullable: false}}, SafetyError},
		{"not-null with default", AddColumn{Column: &Column{Nullable: false, Default: "0"}}, SafetySafe},
		{"not-null with backfill", AddColumn{Column: &Column{Nullable: false}, Backfill: "0"}, SafetySafe},
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
