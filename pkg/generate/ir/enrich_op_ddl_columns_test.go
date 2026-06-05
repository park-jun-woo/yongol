//ff:func feature=gen-ir type=test control=sequence
//ff:what TestEnrichOpDDLColumns — TestEnrichOpDDLColumns -- CRUD Op DDL 매칭으로 ColumnName/IsPK 세팅·미매칭/빈모델 분기 검증

package ir

import (
	"testing"
)

func TestEnrichOpDDLColumns(t *testing.T) {
	t.Run("MatchSetsColumns", func(t *testing.T) {
		op := &Op{
			Kind: OpGet,
			Get: &GetOp{
				Model: "Course",
				Args: []FieldArg{
					{Key: "ID"},
					{Key: "Title"},
				},
			},
		}
		enrichOpDDLColumns(op, enrichTestFS())
		id := findArgByKey(op.Get.Args, "ID")
		if id == nil || id.ColumnName != "id" || !id.IsPK {
			t.Errorf("ID arg = %+v, want column id pk", id)
		}
		title := findArgByKey(op.Get.Args, "Title")
		if title == nil || title.ColumnName != "title" || title.IsPK {
			t.Errorf("Title arg = %+v, want column title non-pk", title)
		}
	})

	t.Run("EmptyModelNoop", func(t *testing.T) {
		op := &Op{Kind: OpGet, Get: &GetOp{Model: "", Args: []FieldArg{{Key: "ID"}}}}
		enrichOpDDLColumns(op, enrichTestFS())
		if op.Get.Args[0].ColumnName != "" {
			t.Errorf("empty model should not set ColumnName, got %q", op.Get.Args[0].ColumnName)
		}
	})

	t.Run("NoTableMatchNoop", func(t *testing.T) {
		op := &Op{Kind: OpGet, Get: &GetOp{Model: "Unknown", Args: []FieldArg{{Key: "ID"}}}}
		enrichOpDDLColumns(op, enrichTestFS())
		if op.Get.Args[0].ColumnName != "" {
			t.Errorf("unmatched table should not set ColumnName, got %q", op.Get.Args[0].ColumnName)
		}
	})
}
