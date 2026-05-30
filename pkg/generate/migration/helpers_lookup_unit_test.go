//ff:func feature=migration type=test control=iteration dimension=1
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestCheckMap(t *testing.T) {
	a := &CheckConstraint{Name: "chk_a", Expression: "x > 0"}
	b := &CheckConstraint{Name: "chk_b", Expression: "y < 1"}
	m := checkMap([]*CheckConstraint{a, b})
	if len(m) != 2 {
		t.Fatalf("len = %d, want 2", len(m))
	}
	if m["chk_a"] != a || m["chk_b"] != b {
		t.Errorf("map did not key by name correctly")
	}
	if got := checkMap(nil); len(got) != 0 {
		t.Errorf("nil input -> len %d, want 0", len(got))
	}
}

func TestColumnMap(t *testing.T) {
	c := &Column{Name: "id"}
	m := columnMap([]*Column{c})
	if m["id"] != c {
		t.Errorf("columnMap did not key by name")
	}
	if got := columnMap(nil); len(got) != 0 {
		t.Errorf("nil -> len %d, want 0", len(got))
	}
}

func TestFKMap(t *testing.T) {
	fk := &ForeignKey{Name: "fk1"}
	m := fkMap([]*ForeignKey{fk})
	if m["fk1"] != fk {
		t.Errorf("fkMap did not key by name")
	}
	if got := fkMap(nil); len(got) != 0 {
		t.Errorf("nil -> len %d, want 0", len(got))
	}
}

func TestIndexMap(t *testing.T) {
	ix := &Index{Name: "ix1"}
	m := indexMap([]*Index{ix})
	if m["ix1"] != ix {
		t.Errorf("indexMap did not key by name")
	}
	if got := indexMap(nil); len(got) != 0 {
		t.Errorf("nil -> len %d, want 0", len(got))
	}
}

func TestRenamedColumnName(t *testing.T) {
	rules := []RenameColumnHint{
		{Table: "users", From: "old", To: "new"},
	}
	cases := []struct {
		name      string
		prev, cur string
		col       string
		want      string
	}{
		{"match on prev table", "users", "members", "old", "new"},
		{"match on new table", "people", "users", "old", "new"},
		{"no col match", "users", "users", "other", "other"},
		{"no table match", "orders", "orders", "old", "old"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := renamedColumnName(c.prev, c.cur, c.col, rules); got != c.want {
				t.Errorf("renamedColumnName = %q, want %q", got, c.want)
			}
		})
	}
	if got := renamedColumnName("t", "t", "c", nil); got != "c" {
		t.Errorf("nil rules -> %q, want c", got)
	}
}

func TestRenamedTableName(t *testing.T) {
	rules := []RenameTableHint{{From: "old_t", To: "new_t"}}
	if got := renamedTableName("old_t", rules); got != "new_t" {
		t.Errorf("match -> %q, want new_t", got)
	}
	if got := renamedTableName("other", rules); got != "other" {
		t.Errorf("no match -> %q, want other", got)
	}
	if got := renamedTableName("x", nil); got != "x" {
		t.Errorf("nil rules -> %q, want x", got)
	}
}

func TestSetFKAction(t *testing.T) {
	cases := []struct {
		action, val            string
		wantDelete, wantUpdate string
	}{
		{"DELETE", "CASCADE", "CASCADE", ""},
		{"UPDATE", "SET NULL", "", "SET NULL"},
		{"OTHER", "CASCADE", "", ""},
	}
	for _, c := range cases {
		c := c
		t.Run(c.action, func(t *testing.T) {
			fk := &ForeignKey{}
			setFKAction(fk, c.action, c.val)
			if fk.OnDelete != c.wantDelete || fk.OnUpdate != c.wantUpdate {
				t.Errorf("setFKAction(%q,%q) -> OnDelete=%q OnUpdate=%q, want %q/%q",
					c.action, c.val, fk.OnDelete, fk.OnUpdate, c.wantDelete, c.wantUpdate)
			}
		})
	}
}

func TestNewEmptyHints(t *testing.T) {
	h := newEmptyHints()
	if h.Casts == nil || h.Backfills == nil || h.DataMigrations == nil || h.AllowDestructive == nil {
		t.Fatalf("newEmptyHints left a nil map: %+v", h)
	}
	if len(h.RenameTables) != 0 || len(h.RenameColumns) != 0 {
		t.Errorf("rename slices should be empty initially")
	}
}

func TestNewSchema(t *testing.T) {
	s := NewSchema()
	if s == nil || s.Tables == nil {
		t.Fatalf("NewSchema returned nil or nil Tables: %+v", s)
	}
	if len(s.Tables) != 0 {
		t.Errorf("Tables should be empty, got %d", len(s.Tables))
	}
}

func TestCollectTypeTokens(t *testing.T) {
	cases := []struct {
		name     string
		in       []string
		wantType string
		wantRest []string
	}{
		{"empty", nil, "", nil},
		{"simple", []string{"INTEGER", "NOT", "NULL"}, "INTEGER", []string{"NOT", "NULL"}},
		{"varchar with len token", []string{"VARCHAR(255)", "NOT"}, "VARCHAR(255)", []string{"NOT"}},
		{"character varying", []string{"character", "varying", "NOT"}, "character varying", []string{"NOT"}},
		{"timestamp with time zone", []string{"timestamp", "with", "time", "zone"}, "timestamp with time zone", []string{}},
		{"timestamp without time zone", []string{"timestamp", "without", "time", "zone", "NOT"}, "timestamp without time zone", []string{"NOT"}},
		{"double precision", []string{"double", "precision"}, "double precision", []string{}},
		{"array", []string{"INTEGER", "[]"}, "INTEGER[]", []string{}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			gotType, gotRest := collectTypeTokens(c.in)
			if gotType != c.wantType {
				t.Errorf("type = %q, want %q", gotType, c.wantType)
			}
			if !reflect.DeepEqual(gotRest, c.wantRest) {
				t.Errorf("rest = %#v, want %#v", gotRest, c.wantRest)
			}
		})
	}
}

func TestConsumeMultiWordTypeTail(t *testing.T) {
	cases := []struct {
		name      string
		toks      []string
		i         int
		startPart []string
		wantParts []string
		wantIdx   int
	}{
		{"varying", []string{"character", "varying"}, 1, []string{"character"}, []string{"character", "varying"}, 2},
		{"precision", []string{"double", "precision"}, 1, []string{"double"}, []string{"double", "precision"}, 2},
		{"with time zone", []string{"timestamp", "with", "time", "zone"}, 1, []string{"timestamp"}, []string{"timestamp", "with", "time", "zone"}, 4},
		{"without time zone", []string{"timestamp", "without", "time", "zone"}, 1, []string{"timestamp"}, []string{"timestamp", "without", "time", "zone"}, 4},
		{"no tail match", []string{"timestamp", "NOT"}, 1, []string{"timestamp"}, []string{"timestamp"}, 1},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			parts := make([]string, len(c.startPart))
			copy(parts, c.startPart)
			gotIdx := consumeMultiWordTypeTail(nil, c.toks, c.i, &parts)
			if gotIdx != c.wantIdx {
				t.Errorf("idx = %d, want %d", gotIdx, c.wantIdx)
			}
			if !reflect.DeepEqual(parts, c.wantParts) {
				t.Errorf("parts = %#v, want %#v", parts, c.wantParts)
			}
		})
	}
}
