//ff:func feature=migration type=test control=sequence
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

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
