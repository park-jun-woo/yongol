//ff:func feature=validate type=test control=sequence
//ff:what TestCollectReferencedTable_ZeroCov — 시퀀스에서 참조 테이블 수집 분기 직접 호출

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectReferencedTable_ZeroCov(t *testing.T) {
	// package-qualified seq → skipped entirely.
	tables := map[string]bool{}
	collectReferencedTable(ssac.Sequence{Package: "auth", Model: "Foo.Bar"}, tables)
	if len(tables) != 0 {
		t.Errorf("package seq should be skipped, got %v", tables)
	}

	// Model with table prefix + Result type.
	tables = map[string]bool{}
	seq := ssac.Sequence{
		Model:  "Course.FindByID",
		Result: &ssac.Result{Type: "Reservation"},
	}
	collectReferencedTable(seq, tables)
	if len(tables) == 0 {
		t.Errorf("expected referenced tables collected, got %v", tables)
	}
	if !tables[canonicalTableKey("Course")] {
		t.Errorf("Course not collected: %v", tables)
	}
	if !tables[canonicalTableKey("Reservation")] {
		t.Errorf("Reservation not collected: %v", tables)
	}

	// primitive result type → not added.
	tables = map[string]bool{}
	collectReferencedTable(ssac.Sequence{Result: &ssac.Result{Type: "int64"}}, tables)
	if len(tables) != 0 {
		t.Errorf("primitive type should not be collected, got %v", tables)
	}
}
