//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what appendUniqueIndex — UNIQUE 컬럼에서 unique index 추가 / 컬럼 없으면 무시

package ddl

import "testing"

func TestAppendUniqueIndex(t *testing.T) {
	t.Run("single column", func(t *testing.T) {
		tb := &Table{}
		appendUniqueIndex("UNIQUE (email)", tb)
		if len(tb.Indexes) != 1 {
			t.Fatalf("Indexes len = %d, want 1", len(tb.Indexes))
		}
		idx := tb.Indexes[0]
		if !idx.IsUnique || idx.Name != "unique_email" || len(idx.Columns) != 1 || idx.Columns[0] != "email" {
			t.Errorf("index = %+v, want unique_email[email]", idx)
		}
	})
	t.Run("multi column", func(t *testing.T) {
		tb := &Table{}
		appendUniqueIndex("UNIQUE (org_id, name)", tb)
		if len(tb.Indexes) != 1 || tb.Indexes[0].Name != "unique_org_id_name" {
			t.Errorf("index name = %q, want unique_org_id_name", tb.Indexes[0].Name)
		}
	})
	t.Run("no columns ignored", func(t *testing.T) {
		tb := &Table{}
		appendUniqueIndex("UNIQUE ()", tb)
		if len(tb.Indexes) != 0 {
			t.Errorf("Indexes len = %d, want 0", len(tb.Indexes))
		}
	})
}
