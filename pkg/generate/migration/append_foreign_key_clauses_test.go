//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendForeignKeyClauses(t *testing.T) {
	var b strings.Builder
	appendForeignKeyClauses(&b, []*ForeignKey{
		{Name: "fk_x", Columns: []string{"a"}, RefTable: "other", RefColumns: []string{"id"}, OnDelete: "CASCADE", OnUpdate: "SET NULL"},
		{Name: "fk_y", Columns: []string{"b", "c"}, RefTable: "t2", RefColumns: []string{"d", "e"}},
	})
	got := b.String()
	if !strings.Contains(got, ",\n    CONSTRAINT fk_x FOREIGN KEY (a) REFERENCES other (id) ON DELETE CASCADE ON UPDATE SET NULL") {
		t.Errorf("fk_x clause missing/wrong: %q", got)
	}
	if !strings.Contains(got, ",\n    CONSTRAINT fk_y FOREIGN KEY (b, c) REFERENCES t2 (d, e)") {
		t.Errorf("fk_y clause missing/wrong: %q", got)
	}
	if strings.Contains(got, "fk_y") && strings.Contains(strings.Split(got, "fk_y")[1], "ON DELETE") {
		t.Errorf("fk_y should not emit ON DELETE when empty: %q", got)
	}
}
