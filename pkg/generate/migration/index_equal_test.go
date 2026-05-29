//ff:func feature=migration type=test control=sequence
//ff:what TestIndexEqual — Unique/Where/Method(btree=="")/Columns 동등성 비교
package migration

import "testing"

func TestIndexEqual(t *testing.T) {
	base := &Index{Name: "ix", Columns: []string{"a"}, Unique: true, Method: "btree"}
	// "" and "btree" equivalent
	emptyMethod := &Index{Name: "ix", Columns: []string{"a"}, Unique: true, Method: ""}
	if !indexEqual(base, emptyMethod) {
		t.Error(`Method "" should equal "btree"`)
	}
	if indexEqual(base, nil) || indexEqual(nil, base) {
		t.Error("nil index never equal")
	}
	diffUnique := &Index{Name: "ix", Columns: []string{"a"}, Unique: false}
	if indexEqual(base, diffUnique) {
		t.Error("different Unique should not be equal")
	}
	diffCols := &Index{Name: "ix", Columns: []string{"b"}, Unique: true}
	if indexEqual(base, diffCols) {
		t.Error("different Columns should not be equal")
	}
}
