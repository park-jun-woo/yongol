//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package migration

import (
	"testing"
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// bnSchemas builds a prev/curr pair used by several by-name diff tests.
func bnSchemas() (prev, curr *Table) {
	prev = &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "VARCHAR", Length: 255}, Nullable: false},
			{Name: "legacy", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
		},
		Indexes: []*Index{
			{Name: "idx_old", Columns: []string{"email"}},
		},
		ForeignKeys: []*ForeignKey{
			{Name: "fk_old", Columns: []string{"org_id"}, RefTable: "orgs", RefColumns: []string{"id"}},
		},
		Checks: []*CheckConstraint{
			{Name: "users_chk_old", Expression: "id > 0"},
		},
	}
	curr = &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "VARCHAR", Length: 320}, Nullable: false, Default: "''"},
			{Name: "age", Type: CanonicalType{Base: "INTEGER"}, Nullable: false},
		},
		Indexes: []*Index{
			{Name: "idx_new", Columns: []string{"age"}},
		},
		ForeignKeys: []*ForeignKey{
			{Name: "fk_new", Columns: []string{"org_id"}, RefTable: "orgs", RefColumns: []string{"id"}, OnDelete: "CASCADE"},
		},
		Checks: []*CheckConstraint{
			{Name: "users_chk_new", Expression: "age >= 0"},
		},
	}
	return prev, curr
}

func TestApplyColumnAttrs_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "c", Nullable: true}
	applyColumnAttrs(tbl, col, []string{"NOT", "NULL"})
	if col.Nullable {
		t.Errorf("NOT NULL not applied")
	}
}

func TestApplyHint_ZeroCov(t *testing.T) {
	h := newEmptyHints()
	h.AllowDestructive["t"] = true
	out := applyHint(DropTable{Name: "t"}, h)
	if dt, ok := out.(DropTable); !ok || !dt.AllowDestructive {
		t.Errorf("DropTable allow-destructive not applied: %#v", out)
	}
	// non-hint-aware op returned unchanged
	if got := applyHint(CreateTable{Table: &Table{Name: "x"}}, h); got == nil {
		t.Errorf("nil returned")
	}
}

func TestApplyIdentityAttr_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "id", Nullable: true}
	rest := []string{"GENERATED", "ALWAYS", "AS", "IDENTITY"}
	if c := applyIdentityAttr(tbl, col, rest, 0); c != 4 {
		t.Errorf("consumed=%d want 4", c)
	}
	if col.Identity == nil || !col.Identity.Always || col.Nullable {
		t.Errorf("identity not set: %#v", col)
	}
	// non-GENERATED returns 0
	if c := applyIdentityAttr(tbl, col, []string{"FOO"}, 0); c != 0 {
		t.Errorf("expected 0 for non-GENERATED")
	}
}

func TestApplyInlineCheck_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "c"}
	c := applyInlineCheck(tbl, col, []string{"CHECK", "(c > 0)"}, 0)
	if c != 2 || len(tbl.Checks) != 1 {
		t.Errorf("inline check not added: consumed=%d checks=%d", c, len(tbl.Checks))
	}
}

func TestApplySerialDefault_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "id", Nullable: true}
	applySerialDefault(tbl, col, true)
	if col.Default == "" || col.Nullable {
		t.Errorf("serial default not applied: %#v", col)
	}
}

func TestApplyTypeParams_ZeroCov(t *testing.T) {
	ct := &CanonicalType{Base: "NUMERIC"}
	applyTypeParams(ct, "10,2")
	if ct.Precision != 10 || ct.Scale != 2 {
		t.Errorf("numeric params wrong: %#v", ct)
	}
	ct2 := &CanonicalType{Base: "VARCHAR"}
	applyTypeParams(ct2, "255")
	if ct2.Length != 255 {
		t.Errorf("varchar length wrong")
	}
}

func TestBuildAddColumnOp_ZeroCov(t *testing.T) {
	col := &Column{Name: "age", Type: CanonicalType{Base: "INTEGER"}}
	h := newEmptyHints()
	h.Backfills[colKey{Table: "t", Column: "age"}] = "0"
	op := buildAddColumnOp("t", "age", col, h)
	if op.Backfill != "0" || op.Table != "t" {
		t.Errorf("add col op wrong: %#v", op)
	}
}

func TestCollectDiags_ZeroCov(t *testing.T) {
	prev := NewSchema()
	curr := NewSchema()
	diags := collectDiags([]diagnostic.Diagnostic{{Message: "X"}}, prev, curr, newEmptyHints(), nil, nil)
	if len(diags) == 0 {
		t.Errorf("expected at least the seeded diag")
	}
}

func TestColumnAddOps_ZeroCov(t *testing.T) {
	currMap := map[string]*Column{"age": {Name: "age"}}
	ops := columnAddOps("t", []string{"age"}, map[string]*Column{}, map[string]bool{}, currMap, newEmptyHints())
	if len(ops) != 1 {
		t.Errorf("expected one add op, got %d", len(ops))
	}
}

func TestColumnAlterForPair_ZeroCov(t *testing.T) {
	pc := &Column{Name: "c", Type: CanonicalType{Base: "INTEGER"}, Nullable: true}
	cc := &Column{Name: "c", Type: CanonicalType{Base: "BIGINT"}, Nullable: false, Default: "0"}
	ops := columnAlterForPair("t", "c", pc, cc, newEmptyHints())
	if len(ops) != 3 {
		t.Errorf("expected type+nullable+default ops, got %d", len(ops))
	}
}

func TestColumnAlterOps_ZeroCov(t *testing.T) {
	prevMap := map[string]*Column{"c": {Name: "c", Type: CanonicalType{Base: "INTEGER"}}}
	currMap := map[string]*Column{"c": {Name: "c", Type: CanonicalType{Base: "BIGINT"}}}
	ops := columnAlterOps("t", []string{"c"}, prevMap, currMap, map[string]bool{}, newEmptyHints())
	if len(ops) == 0 {
		t.Errorf("expected alter ops")
	}
}

func TestColumnDropOps_ZeroCov(t *testing.T) {
	ops := columnDropOps("t", []string{"old"}, map[string]*Column{}, map[string]bool{}, newEmptyHints())
	if len(ops) != 1 {
		t.Errorf("expected one drop op, got %d", len(ops))
	}
}

func TestDiffAddOrRenameTarget_ZeroCov(t *testing.T) {
	prev := NewSchema()
	c := &Table{Name: "newt", Columns: []*Column{{Name: "id"}}}
	ops := diffAddOrRenameTarget("newt", prev, c, newEmptyHints(), map[string]string{})
	if len(ops) == 0 {
		t.Errorf("expected create table ops")
	}
}

func TestDiffChecks_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ops := diffChecks(prev, curr, "users")
	if len(ops) == 0 {
		t.Errorf("expected check diff ops")
	}
}

func TestDiffColumns_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ops := diffColumns(prev, curr, newEmptyHints(), "users")
	if len(ops) == 0 {
		t.Errorf("expected column diff ops")
	}
}

func TestDiffForeignKeys_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ops := diffForeignKeys(prev, curr, "users")
	if len(ops) == 0 {
		t.Errorf("expected fk diff ops")
	}
}

func TestDiffIndexes_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ops := diffIndexes(prev, curr, "users")
	if len(ops) == 0 {
		t.Errorf("expected index diff ops")
	}
}

func TestDiffOneTable_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ps := NewSchema()
	ps.Tables["users"] = prev
	cs := NewSchema()
	cs.Tables["users"] = curr
	ops := diffOneTable("users", ps, ps, cs, newEmptyHints(), map[string]string{}, map[string]string{})
	if len(ops) == 0 {
		t.Errorf("expected body diff ops")
	}
}

func TestDiffTableBody_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	if ops := diffTableBody(prev, curr, newEmptyHints(), "users"); len(ops) == 0 {
		t.Errorf("expected table body ops")
	}
	if ops := diffTableBody(nil, curr, nil, "users"); ops != nil {
		t.Errorf("nil prev should give nil")
	}
}

func TestDiffTables_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ps := NewSchema()
	ps.Tables["users"] = prev
	cs := NewSchema()
	cs.Tables["users"] = curr
	if ops := diffTables(ps, ps, cs, newEmptyHints()); len(ops) == 0 {
		t.Errorf("expected diff ops")
	}
}

func TestDispatchColumnAttr_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "c", Nullable: true}
	if step := dispatchColumnAttr(tbl, col, []string{"PRIMARY", "KEY"}, 0); step != 2 {
		t.Errorf("PRIMARY KEY step=%d want 2", step)
	}
	if len(tbl.PrimaryKey) != 1 {
		t.Errorf("PK not set")
	}
}

func TestDispatchStatement_ZeroCov(t *testing.T) {
	s := NewSchema()
	if err := dispatchStatement(s, "CREATE TABLE t (id BIGINT PRIMARY KEY)"); err != nil {
		t.Fatalf("dispatch: %v", err)
	}
	if _, ok := s.Tables["t"]; !ok {
		t.Errorf("table not created")
	}
	if err := dispatchStatement(s, ""); err != nil {
		t.Errorf("empty stmt should be nil")
	}
}

func TestEmitMigration_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	curr := NewSchema()
	curr.Tables["t"] = &Table{Name: "t", Columns: []*Column{{Name: "id", Type: CanonicalType{Base: "BIGINT"}}}}
	ops := []Operation{CreateTable{Table: curr.Tables["t"]}}
	res, _, err := emitMigration(dir, dir, curr, ops, newEmptyHints(), ModeInitial, "v0.0.0", time.Unix(0, 0).UTC(), nil)
	if err != nil {
		t.Fatalf("emitMigration: %v", err)
	}
	if res.MigrationFile == "" {
		t.Errorf("no migration file recorded")
	}
}

func TestFkAlterOrAddOps_ZeroCov(t *testing.T) {
	currMap := map[string]*ForeignKey{"fk": {Name: "fk", Columns: []string{"a"}, RefTable: "o", RefColumns: []string{"id"}}}
	ops := fkAlterOrAddOps("t", []string{"fk"}, map[string]*ForeignKey{}, currMap)
	if len(ops) != 1 {
		t.Errorf("expected add fk op, got %d", len(ops))
	}
}

func TestFkDiffForOne_ZeroCov(t *testing.T) {
	prevMap := map[string]*ForeignKey{"fk": {Name: "fk", Columns: []string{"a"}, RefTable: "o", RefColumns: []string{"id"}}}
	currMap := map[string]*ForeignKey{"fk": {Name: "fk", Columns: []string{"a"}, RefTable: "o", RefColumns: []string{"id"}, OnDelete: "CASCADE"}}
	ops := fkDiffForOne("t", "fk", prevMap, currMap)
	if len(ops) != 2 {
		t.Errorf("changed fk should drop+add, got %d", len(ops))
	}
	// equal → nil
	if got := fkDiffForOne("t", "fk", prevMap, prevMap); got != nil {
		t.Errorf("equal fk should be nil")
	}
}

func TestFkDropOps_ZeroCov(t *testing.T) {
	ops := fkDropOps("t", []string{"fk"}, map[string]*ForeignKey{})
	if len(ops) != 1 {
		t.Errorf("expected drop fk op, got %d", len(ops))
	}
}

func TestMig001From_ZeroCov(t *testing.T) {
	prev := NewSchema()
	curr := NewSchema()
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "missing_from", To: "missing_to"}}
	diags := mig001From(prev, curr, h)
	_ = diags
	if got := mig001From(prev, curr, nil); got != nil {
		t.Errorf("nil hints should give nil")
	}
}

func TestParseColumn_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseColumn(tbl, "id BIGINT NOT NULL")
	if len(tbl.Columns) != 1 || tbl.Columns[0].Name != "id" {
		t.Errorf("column not parsed: %#v", tbl.Columns)
	}
}

func TestParseInlineRef_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	fk, consumed := parseInlineRef(tbl, "org_id", []string{"orgs(id)"})
	if fk == nil || fk.RefTable != "orgs" || consumed == 0 {
		t.Errorf("inline ref wrong: %#v consumed=%d", fk, consumed)
	}
}

func TestParseInlineRefTarget_ZeroCov(t *testing.T) {
	rt, rc, c := parseInlineRefTarget([]string{"orgs(id)"})
	if rt != "orgs" || rc != "id" || c != 1 {
		t.Errorf("target wrong: %s %s %d", rt, rc, c)
	}
}

func TestParseNamedConstraint_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseNamedConstraint(tbl, "CONSTRAINT t_pkey PRIMARY KEY (id)")
	if len(tbl.PrimaryKey) == 0 {
		t.Errorf("named PK not parsed")
	}
}

func TestParseTableCheck_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	c := parseTableCheck(tbl, "", "CHECK (age > 0)")
	if c == nil || c.Name == "" || c.Expression == "" {
		t.Errorf("check not parsed: %#v", c)
	}
}

func TestParseTableFK_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	fk := parseTableFK(tbl, "FOREIGN KEY (org_id) REFERENCES orgs (id)")
	if fk == nil || fk.RefTable != "orgs" {
		t.Errorf("table fk not parsed: %#v", fk)
	}
}

func TestParseTableItem_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseTableItem(tbl, "PRIMARY KEY (id)")
	if len(tbl.PrimaryKey) == 0 {
		t.Errorf("PK item not handled")
	}
	parseTableItem(tbl, "name TEXT")
	if len(tbl.Columns) == 0 {
		t.Errorf("column item not handled")
	}
}

func TestParseTableItems_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseTableItems(tbl, "id BIGINT, name TEXT, PRIMARY KEY (id)")
	if len(tbl.Columns) != 2 || len(tbl.PrimaryKey) == 0 {
		t.Errorf("items not parsed: cols=%d pk=%v", len(tbl.Columns), tbl.PrimaryKey)
	}
}

func TestStepTopLevel_ZeroCov(t *testing.T) {
	st := newSplitState()
	s := "a,b"
	i := 0
	for i < len(s) {
		i = stepTopLevel(st, s, i, ',')
		i++
	}
	out := st.finish()
	if len(out) != 2 {
		t.Errorf("expected 2 parts, got %v", out)
	}
}
