//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what findDDLTableByModelName 단위 테스트 (sqlc 모델명 → DDL 테이블)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestFindDDLTableByModelName(t *testing.T) {
	tables := []ddl.Table{
		{Name: "workflows"},
		{Name: "users"},
		{Name: "execution_logs"},
	}
	cases := []struct {
		name      string
		modelName string
		wantTable string // "" => nil
	}{
		{"Workflow → workflows", "Workflow", "workflows"},
		{"User → users", "User", "users"},
		{"ExecutionLog → execution_logs", "ExecutionLog", "execution_logs"},
		{"unknown → nil", "Nonexistent", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findDDLTableByModelName(tables, tc.modelName)
			if tc.wantTable == "" {
				if got != nil {
					t.Errorf("expected nil, got %q", got.Name)
				}
				return
			}
			if got == nil || got.Name != tc.wantTable {
				t.Errorf("findDDLTableByModelName(%q) = %v, want table %q", tc.modelName, got, tc.wantTable)
			}
		})
	}
}
