//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestRenderTable(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
		},
		PrimaryKey: []string{"id"},
		Indexes:    []*Index{{Name: "idx_email", Columns: []string{"email"}, Unique: true}},
	}
	got := renderTable(tbl)
	if !strings.HasPrefix(got, "CREATE TABLE users (") {
		t.Errorf("missing CREATE TABLE header: %q", got)
	}
	if !strings.Contains(got, "\n    id INTEGER NOT NULL,") {
		t.Errorf("missing id column: %q", got)
	}
	if !strings.Contains(got, ",\n    PRIMARY KEY (id)") {
		t.Errorf("missing PK clause: %q", got)
	}
	if !strings.Contains(got, "\n);\n") {
		t.Errorf("missing close paren: %q", got)
	}
	if !strings.Contains(got, "CREATE UNIQUE INDEX idx_email ON users (email);") {
		t.Errorf("missing index line: %q", got)
	}
}
