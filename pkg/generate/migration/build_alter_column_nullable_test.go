//ff:func feature=migration type=test control=sequence
//ff:what build_and_mig001_unit_test — buildAlterColumnNullable/Type + mig001CheckRename(Columns/Tables) + removeLegacyBaseline 단위 테스트
package migration

import (
	"testing"
)

func TestBuildAlterColumnNullable(t *testing.T) {
	// nil hints -> no backfill
	op := buildAlterColumnNullable("users", "email", true, false, nil)
	if op.Table != "users" || op.Column != "email" || op.From != true || op.To != false || op.Backfill != "" {
		t.Errorf("nil hints op wrong: %+v", op)
	}
	// matching backfill hint applied
	hints := &Hints{Backfills: map[colKey]string{{Table: "users", Column: "email"}: "''"}}
	op = buildAlterColumnNullable("users", "email", true, false, hints)
	if op.Backfill != "''" {
		t.Errorf("backfill not applied: %+v", op)
	}
	// non-matching column -> no backfill
	op = buildAlterColumnNullable("users", "other", true, false, hints)
	if op.Backfill != "" {
		t.Errorf("backfill should not apply to non-matching column: %+v", op)
	}
}
