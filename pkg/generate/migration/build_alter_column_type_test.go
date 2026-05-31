//ff:func feature=migration type=test control=sequence
//ff:what build_and_mig001_unit_test — buildAlterColumnNullable/Type + mig001CheckRename(Columns/Tables) + removeLegacyBaseline 단위 테스트
package migration

import (
	"testing"
)

func TestBuildAlterColumnType(t *testing.T) {
	from := CanonicalType{Base: "INTEGER"}
	to := CanonicalType{Base: "BIGINT"}
	op := buildAlterColumnType("t", "id", from, to, nil)
	if op.Using != "" {
		t.Errorf("nil hints should leave Using empty: %+v", op)
	}
	hints := &Hints{Casts: map[colKey]string{{Table: "t", Column: "id"}: "id::bigint"}}
	op = buildAlterColumnType("t", "id", from, to, hints)
	if op.Using != "id::bigint" {
		t.Errorf("cast USING not applied: %+v", op)
	}
}
