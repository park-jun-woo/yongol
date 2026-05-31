//ff:func feature=validate type=test control=sequence
//ff:what TestCheckStateFieldColumn_ZeroCov — XDM-27 상태 필드↔DDL 컬럼 매칭 분기 직접 호출

package ddl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestCheckStateFieldColumn_ZeroCov(t *testing.T) {
	// no underscore → skipped (nil).
	g := &rule.Ground{Lookup: map[string]rule.StringSet{}}
	if d := checkStateFieldColumn("course", g); d != nil {
		t.Errorf("no-underscore diagram should be skipped, got %v", d)
	}
	// trailing underscore → skipped.
	if d := checkStateFieldColumn("course_", g); d != nil {
		t.Errorf("trailing-underscore should be skipped")
	}

	// underscore, column exists (orders.status) → nil.
	g2 := &rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.column.orders": {"status": true},
	}}
	if d := checkStateFieldColumn("order_status", g2); d != nil {
		t.Errorf("existing column should yield nil, got %v", d)
	}

	// underscore, column missing → XDM-27 diagnostic.
	g3 := &rule.Ground{Lookup: map[string]rule.StringSet{}}
	d := checkStateFieldColumn("order_status", g3)
	if d == nil {
		t.Fatalf("missing column should yield diagnostic")
	}
	if d.Message == "" {
		t.Errorf("diagnostic missing message")
	}
}
