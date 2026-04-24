//ff:func feature=migration type=test control=iteration dimension=1
//ff:what sortByDependency — 13단계 정렬 보장 (InsertSentinel 포함)
package migration

import "testing"

func TestSortByDependency_PhaseOrder(t *testing.T) {
	// Intentionally shuffled order.
	raw := []Operation{
		AddForeignKey{Table: "a", FK: &ForeignKey{Name: "fk_a"}},
		CreateTable{Table: &Table{Name: "a"}},
		DropForeignKey{Table: "b", Name: "fk_old"},
		DropTable{Name: "x"},
		AddColumn{Table: "a", Column: &Column{Name: "c"}},
		InsertSentinel{Table: "a", Body: "INSERT INTO a (id) VALUES (0) ON CONFLICT DO NOTHING;"},
	}
	got := sortByDependency(raw)
	// Expected phases:
	//   DropFK(2), DropTable(6), CreateTable(7), InsertSentinel(8), AddColumn(9), AddFK(13).
	wantPhase := []int{2, 6, 7, 8, 9, 13}
	if len(got) != len(wantPhase) {
		t.Fatalf("length mismatch: %d vs %d", len(got), len(wantPhase))
	}
	for i, op := range got {
		if phaseOf(op) != wantPhase[i] {
			t.Errorf("op %d: phase %d, want %d (%T)", i, phaseOf(op), wantPhase[i], op)
		}
	}
}
