//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafety_CastHint_Silences — @cast 힌트 적용 시 MIG-005 미발생
package migration

import "testing"

func TestCheckSafety_CastHint_Silences(t *testing.T) {
	ops := []Operation{
		AlterColumnType{
			Table: "t", Column: "c",
			From: CanonicalType{Base: "INTEGER"},
			To:   CanonicalType{Base: "TEXT"},
		},
	}
	hints := &Hints{
		Casts: map[colKey]string{
			{Table: "t", Column: "c"}: "c::text",
		},
	}
	ops = ApplyHintsToOps(ops, hints)
	issues := CheckSafety(ops)
	if len(issues) != 0 {
		t.Errorf("expected no issues with @cast, got %+v", issues)
	}
}
