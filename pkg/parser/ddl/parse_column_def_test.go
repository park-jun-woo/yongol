//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what parseColumnDef — 이름/타입/NOT NULL/default/varchar/check/주석 어노테이션 추출

package ddl

import "testing"

func TestParseColumnDef(t *testing.T) {
	t.Run("full column with default and varchar", func(t *testing.T) {
		tb := newTable()
		line := "email VARCHAR(255) NOT NULL DEFAULT 'x'"
		parseColumnDef(line, "EMAIL VARCHAR(255) NOT NULL DEFAULT 'X'", tb, false)
		c := tb.Columns["email"]
		if c.RawType != "VARCHAR(255)" || !c.NotNull || c.VarcharLen != 255 {
			t.Errorf("col = %+v", c)
		}
		if !c.HasDefault || c.DefaultLiteral != "x" {
			t.Errorf("default = %v %q", c.HasDefault, c.DefaultLiteral)
		}
		if len(tb.ColumnOrder) != 1 || tb.ColumnOrder[0] != "email" {
			t.Errorf("ColumnOrder = %v", tb.ColumnOrder)
		}
	})
	t.Run("comment annotations", func(t *testing.T) {
		tb := newTable()
		parseColumnDef("notes TEXT, -- @nullable", "NOTES TEXT,", tb, false)
		c := tb.Columns["notes"]
		if !c.NullableAnnot {
			t.Errorf("expected NullableAnnot")
		}
	})
	t.Run("sensitive annotation", func(t *testing.T) {
		tb := newTable()
		parseColumnDef("pw VARCHAR(60) NOT NULL -- @sensitive", "PW VARCHAR(60) NOT NULL", tb, false)
		if !tb.Columns["pw"].Sensitive {
			t.Errorf("expected Sensitive")
		}
	})
	t.Run("pending archived", func(t *testing.T) {
		tb := newTable()
		parseColumnDef("old_col BIGINT", "OLD_COL BIGINT", tb, true)
		if !tb.Columns["old_col"].Archived {
			t.Errorf("expected Archived from pendingArchived")
		}
	})
	t.Run("too few tokens is no-op", func(t *testing.T) {
		tb := newTable()
		parseColumnDef("lonely", "LONELY", tb, false)
		if len(tb.Columns) != 0 {
			t.Errorf("expected no column, got %v", tb.Columns)
		}
	})
	t.Run("comment-only line after strip is no-op", func(t *testing.T) {
		tb := newTable()
		parseColumnDef("-- @archived", "", tb, false)
		if len(tb.Columns) != 0 {
			t.Errorf("expected no column, got %v", tb.Columns)
		}
	})
}
