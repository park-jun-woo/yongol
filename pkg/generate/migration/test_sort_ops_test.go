//ff:func feature=migration type=test control=iteration dimension=1
//ff:what sortByDependency — 11단계 정렬 보장
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
	}
	got := sortByDependency(raw)
	// Expected order by phase: DropFK(2), DropTable(6), CreateTable(7), AddColumn(8), AddFK(12).
	wantPhase := []int{2, 6, 7, 8, 12}
	if len(got) != len(wantPhase) {
		t.Fatalf("length mismatch: %d vs %d", len(got), len(wantPhase))
	}
	for i, op := range got {
		if phaseOf(op) != wantPhase[i] {
			t.Errorf("op %d: phase %d, want %d (%T)", i, phaseOf(op), wantPhase[i], op)
		}
	}
}
