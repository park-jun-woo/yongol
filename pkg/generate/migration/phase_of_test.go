//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestPhaseOf — Operation 타입을 정렬 단계(1..13) 로 매핑
package migration

import (
	"testing"
)

func TestPhaseOf(t *testing.T) {
	cases := []struct {
		name string
		op   Operation
		want int
	}{
		{"rename table", RenameTable{}, 1},
		{"drop fk", DropForeignKey{}, 2},
		{"drop index", DropIndex{}, 3},
		{"drop check", DropCheck{}, 4},
		{"drop column", DropColumn{}, 5},
		{"drop table", DropTable{}, 6},
		{"create table", CreateTable{Table: &Table{}}, 7},
		{"insert sentinel", InsertSentinel{}, 8},
		{"add column", AddColumn{Column: &Column{}}, 9},
		{"alter type", AlterColumnType{}, 10},
		{"add check", AddCheck{}, 11},
		{"create index", CreateIndex{Index: &Index{}}, 12},
		{"add fk", AddForeignKey{FK: &ForeignKey{}}, 13},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := phaseOf(c.op); got != c.want {
				t.Errorf("phaseOf(%T) = %d, want %d", c.op, got, c.want)
			}
		})
	}
}
