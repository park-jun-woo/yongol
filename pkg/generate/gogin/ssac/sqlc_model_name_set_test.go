//ff:func feature=gen-gogin type=test control=sequence
//ff:what sqlcModelNameSet 단위 테스트 (DDL 테이블명 → sqlc row struct 이름 set)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSQLcModelNameSet(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{Name: "workflows"},
			{Name: "execution_logs"},
			{Name: ""}, // skipped
		},
	}
	got := sqlcModelNameSet(fs)
	if !got["Workflow"] {
		t.Errorf("expected Workflow in set, got %v", got)
	}
	if !got["ExecutionLog"] {
		t.Errorf("expected ExecutionLog in set, got %v", got)
	}
	if got[""] {
		t.Errorf("empty-name table should be skipped")
	}
	if len(got) != 2 {
		t.Errorf("expected 2 entries, got %d (%v)", len(got), got)
	}
}
