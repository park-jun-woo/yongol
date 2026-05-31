//ff:func feature=gen-ir type=test control=sequence
//ff:what ir 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ir

import (
	"testing"

	pddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestEnrichFieldArgDDL_ZeroCov(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{Model: "User", Args: []FieldArg{
			{Key: "ID", Field: ".OrgID"},
		}}},
	}
	fs := &yongol.Fullstack{
		DDLTables: []pddl.Table{{
			Name:       "users",
			PrimaryKey: []string{"id"},
			Columns:    map[string]pddl.Column{"id": {Name: "id"}},
		}},
	}
	enrichFieldArgDDL(ops, fs)
	if ops[0].Get.Args[0].SourceColumn != "org_id" {
		t.Errorf("SourceColumn=%q want org_id", ops[0].Get.Args[0].SourceColumn)
	}
	// nil fs -> only Pass1 runs, no panic
	enrichFieldArgDDL(ops, nil)
}
