//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafety_RiskyCast — INTEGER→TEXT 같은 risky 타입 변경에 대한 MIG-005 발생
package migration

import "testing"

func TestCheckSafety_RiskyCast(t *testing.T) {
	ops := []Operation{
		AlterColumnType{
			Table: "t", Column: "c",
			From: CanonicalType{Base: "INTEGER"},
			To:   CanonicalType{Base: "TEXT"},
		},
	}
	issues := CheckSafety(ops)
	if len(issues) != 1 || issues[0].RuleID != "MIG-005" {
		t.Errorf("expected MIG-005, got %+v", issues)
	}
}
