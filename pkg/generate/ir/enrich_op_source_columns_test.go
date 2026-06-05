//ff:func feature=gen-ir type=test control=sequence
//ff:what TestEnrichOpSourceColumns -- Op FieldArg.Field→SourceColumn(snake_case) 파생 검증
package ir

import "testing"

func TestEnrichOpSourceColumns(t *testing.T) {
	op := &Op{
		Kind: OpGet,
		Get: &GetOp{
			Model: "Course",
			Args: []FieldArg{
				{Key: "ID", Field: ".OrgID"},
				{Key: "Title"}, // no Field -> unchanged
			},
		},
	}
	enrichOpSourceColumns(op)
	if op.Get.Args[0].SourceColumn != "org_id" {
		t.Errorf("SourceColumn = %q, want org_id", op.Get.Args[0].SourceColumn)
	}
	if op.Get.Args[1].SourceColumn != "" {
		t.Errorf("arg without Field should stay empty, got %q", op.Get.Args[1].SourceColumn)
	}
}
