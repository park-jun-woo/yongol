//ff:func feature=migration type=test control=sequence
//ff:what TestFKEqual — RefTable/액션/컬럼 리스트 동등성 비교
package migration

import "testing"

func TestFKEqual(t *testing.T) {
	base := &ForeignKey{Columns: []string{"user_id"}, RefTable: "users", RefColumns: []string{"id"}, OnDelete: "CASCADE"}
	same := &ForeignKey{Columns: []string{"user_id"}, RefTable: "users", RefColumns: []string{"id"}, OnDelete: "CASCADE"}
	if !fkEqual(base, same) {
		t.Error("identical FKs should be equal")
	}
	if fkEqual(base, nil) || fkEqual(nil, base) {
		t.Error("nil FK should never be equal")
	}
	diffAction := &ForeignKey{Columns: []string{"user_id"}, RefTable: "users", RefColumns: []string{"id"}, OnDelete: "SET NULL"}
	if fkEqual(base, diffAction) {
		t.Error("different OnDelete should not be equal")
	}
	diffCols := &ForeignKey{Columns: []string{"uid"}, RefTable: "users", RefColumns: []string{"id"}, OnDelete: "CASCADE"}
	if fkEqual(base, diffCols) {
		t.Error("different Columns should not be equal")
	}
}
