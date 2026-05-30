//ff:func feature=migration type=test control=iteration dimension=1
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestCheckDiffForOne(t *testing.T) {
	cur := map[string]*CheckConstraint{"c": {Name: "c", Expression: "x > 0"}}

	// added (not in prev)
	ops := checkDiffForOne("t", "c", map[string]*CheckConstraint{}, cur)
	if len(ops) != 1 {
		t.Fatalf("added: got %d ops, want 1", len(ops))
	}
	if _, ok := ops[0].(AddCheck); !ok {
		t.Errorf("added: op0 should be AddCheck, got %T", ops[0])
	}

	// unchanged -> no ops
	prevSame := map[string]*CheckConstraint{"c": {Name: "c", Expression: "x > 0"}}
	if ops := checkDiffForOne("t", "c", prevSame, cur); ops != nil {
		t.Errorf("unchanged: expected nil ops, got %#v", ops)
	}

	// changed expression -> Drop + Add
	prevDiff := map[string]*CheckConstraint{"c": {Name: "c", Expression: "x > 5"}}
	ops = checkDiffForOne("t", "c", prevDiff, cur)
	if len(ops) != 2 {
		t.Fatalf("changed: got %d ops, want 2", len(ops))
	}
	if _, ok := ops[0].(DropCheck); !ok {
		t.Errorf("changed: op0 should be DropCheck, got %T", ops[0])
	}
	if _, ok := ops[1].(AddCheck); !ok {
		t.Errorf("changed: op1 should be AddCheck, got %T", ops[1])
	}
}

func TestFindPrevViaRenameHint(t *testing.T) {
	old := &Column{Name: "old_name"}
	prevMap := map[string]*Column{"old_name": old}
	rules := []RenameColumnHint{{Table: "users", From: "old_name", To: "new_name"}}

	if got := findPrevViaRenameHint("new_name", prevMap, rules, "users"); got != old {
		t.Errorf("matching rule should return prev column, got %v", got)
	}
	if got := findPrevViaRenameHint("new_name", prevMap, rules, "other_table"); got != nil {
		t.Errorf("table mismatch should return nil")
	}
	if got := findPrevViaRenameHint("unknown", prevMap, rules, "users"); got != nil {
		t.Errorf("no matching To should return nil")
	}
	// rule matches but From column missing in prevMap
	if got := findPrevViaRenameHint("new_name", map[string]*Column{}, rules, "users"); got != nil {
		t.Errorf("missing From in prevMap should return nil")
	}
}

func TestCheckDropOps(t *testing.T) {
	prevNames := []string{"keep", "drop_me"}
	curr := map[string]*CheckConstraint{"keep": {Name: "keep"}}
	ops := checkDropOps("t", prevNames, curr)
	if len(ops) != 1 {
		t.Fatalf("got %d ops, want 1", len(ops))
	}
	dc, ok := ops[0].(DropCheck)
	if !ok || dc.Name != "drop_me" || dc.Table != "t" {
		t.Errorf("expected DropCheck drop_me on t, got %#v", ops[0])
	}
}

func TestCheckAlterOrAddOps(t *testing.T) {
	prev := map[string]*CheckConstraint{"same": {Name: "same", Expression: "x > 0"}}
	curr := map[string]*CheckConstraint{
		"same": {Name: "same", Expression: "x > 0"},
		"new":  {Name: "new", Expression: "y > 0"},
	}
	ops := checkAlterOrAddOps("t", []string{"new", "same"}, prev, curr)
	// "new" -> 1 Add; "same" -> 0
	if len(ops) != 1 {
		t.Fatalf("got %d ops, want 1: %#v", len(ops), ops)
	}
	if _, ok := ops[0].(AddCheck); !ok {
		t.Errorf("expected AddCheck, got %T", ops[0])
	}
}

func TestAllCreateTableOps(t *testing.T) {
	allCreate := []Operation{
		CreateTable{Table: &Table{Name: "a"}},
		CreateTable{Table: &Table{Name: "b"}},
	}
	if !allCreateTableOps(allCreate) {
		t.Errorf("all-CreateTable slice should return true")
	}
	mixed := []Operation{
		CreateTable{Table: &Table{Name: "a"}},
		DropTable{Name: "b"},
	}
	if allCreateTableOps(mixed) {
		t.Errorf("mixed slice should return false")
	}
	if !allCreateTableOps(nil) {
		t.Errorf("empty slice should vacuously return true")
	}
}

func TestCreateTableWithDeps(t *testing.T) {
	c := &Table{
		Name:        "orders",
		Sentinels:   []SentinelInsert{{SQL: "INSERT INTO orders ..."}},
		Indexes:     []*Index{{Name: "idx_a"}},
		ForeignKeys: []*ForeignKey{{Name: "fk_u"}},
		Checks:      []*CheckConstraint{{Name: "chk_q"}},
	}
	ops := createTableWithDeps(c)
	// CreateTable, InsertSentinel, CreateIndex, AddForeignKey, AddCheck = 5
	if len(ops) != 5 {
		t.Fatalf("got %d ops, want 5: %#v", len(ops), ops)
	}
	if _, ok := ops[0].(CreateTable); !ok {
		t.Errorf("op0 should be CreateTable, got %T", ops[0])
	}
	if _, ok := ops[1].(InsertSentinel); !ok {
		t.Errorf("op1 should be InsertSentinel, got %T", ops[1])
	}
	if _, ok := ops[2].(CreateIndex); !ok {
		t.Errorf("op2 should be CreateIndex, got %T", ops[2])
	}
	if _, ok := ops[3].(AddForeignKey); !ok {
		t.Errorf("op3 should be AddForeignKey, got %T", ops[3])
	}
	if _, ok := ops[4].(AddCheck); !ok {
		t.Errorf("op4 should be AddCheck, got %T", ops[4])
	}
}

func TestParseTableFKRef(t *testing.T) {
	cases := []struct {
		name         string
		toks         []string
		wantTable    string
		wantCols     []string
		wantConsumed int
	}{
		{"target with parens", []string{"(a)", "REFERENCES", "users(id)"}, "users", []string{"id"}, 3},
		{"target then separate parens", []string{"(a)", "REFERENCES", "users", "(id, tenant)"}, "users", []string{"id", "tenant"}, 4},
		{"target without cols", []string{"(a)", "REFERENCES", "users"}, "users", nil, 3},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			gotT, gotCols, gotC := parseTableFKRef(c.toks)
			if gotT != c.wantTable {
				t.Errorf("table = %q, want %q", gotT, c.wantTable)
			}
			if !reflect.DeepEqual(gotCols, c.wantCols) {
				t.Errorf("cols = %#v, want %#v", gotCols, c.wantCols)
			}
			if gotC != c.wantConsumed {
				t.Errorf("consumed = %d, want %d", gotC, c.wantConsumed)
			}
		})
	}
}
