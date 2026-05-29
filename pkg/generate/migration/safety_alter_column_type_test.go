//ff:func feature=migration type=test control=sequence
//ff:what TestSafetyAlterColumnType — risky cast + @cast 없으면 MIG-005 WARNING
package migration

import "testing"

func TestSafetyAlterColumnType(t *testing.T) {
	risky := AlterColumnType{Table: "u", Column: "c",
		From: CanonicalType{Base: "VARCHAR", Length: 255},
		To:   CanonicalType{Base: "VARCHAR", Length: 50}}

	// risky without @cast → MIG-005
	issues := safetyAlterColumnType(risky)
	if len(issues) != 1 || issues[0].RuleID != "MIG-005" || issues[0].Level != SafetyWarning {
		t.Errorf("got %+v, want one MIG-005 warning", issues)
	}

	// risky but @cast supplied → suppressed
	withCast := risky
	withCast.Using = "c::varchar(50)"
	if got := safetyAlterColumnType(withCast); got != nil {
		t.Errorf("with @cast want nil, got %v", got)
	}

	// non-risky widen → nil
	safe := AlterColumnType{Table: "u", Column: "c",
		From: CanonicalType{Base: "INTEGER"}, To: CanonicalType{Base: "BIGINT"}}
	if got := safetyAlterColumnType(safe); got != nil {
		t.Errorf("widen want nil, got %v", got)
	}
}
