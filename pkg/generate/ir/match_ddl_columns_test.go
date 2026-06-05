//ff:func feature=gen-ir type=test control=sequence
//ff:what TestMatchDDLColumns -- FieldArg.Key→snake 컬럼 매칭 시 ColumnName/IsPK 세팅, 미매칭은 무변경 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestMatchDDLColumns(t *testing.T) {
	tbl := &ddl.Table{
		Columns: map[string]ddl.Column{
			"id":        {},
			"owner_id":  {},
			"course_id": {},
		},
	}
	pkSet := map[string]bool{"id": true}

	args := []FieldArg{
		{Key: "Id"},      // -> id, is PK
		{Key: "OwnerId"}, // -> owner_id, not PK
		{Key: "Missing"}, // no column -> untouched
	}

	matchDDLColumns(args, tbl, pkSet)

	if args[0].ColumnName != "id" || !args[0].IsPK {
		t.Errorf("args[0] = %+v, want ColumnName=id IsPK=true", args[0])
	}
	if args[1].ColumnName != "owner_id" || args[1].IsPK {
		t.Errorf("args[1] = %+v, want ColumnName=owner_id IsPK=false", args[1])
	}
	if args[2].ColumnName != "" || args[2].IsPK {
		t.Errorf("args[2] = %+v, want untouched (empty ColumnName, IsPK=false)", args[2])
	}
}
