//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestCanonicalSQL(t *testing.T) {
	if CanonicalSQL(nil) != "" {
		t.Errorf("nil schema should render empty")
	}
	s := &Schema{Tables: map[string]*Table{
		"zebra": {Name: "zebra", Columns: []*Column{{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false}}},
		"apple": {Name: "apple", Columns: []*Column{{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false}}},
	}}
	got := CanonicalSQL(s)
	ai := strings.Index(got, "CREATE TABLE apple")
	zi := strings.Index(got, "CREATE TABLE zebra")
	if ai < 0 || zi < 0 {
		t.Fatalf("both tables should render: %q", got)
	}
	if ai > zi {
		t.Errorf("tables should be alphabetically sorted (apple before zebra): %q", got)
	}
}
