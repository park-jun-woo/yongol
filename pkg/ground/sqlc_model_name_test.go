//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what sqlcModelName — 복수형 테이블명 → PascalCase 단수 모델명 변환 테스트

package ground

import "testing"

func TestSqlcModelName(t *testing.T) {
	tests := []struct {
		table string
		want  string
	}{
		{"workflows", "Workflow"},
		{"organizations", "Organization"},
		{"users", "User"},
		{"actions", "Action"},
		{"execution_logs", "ExecutionLog"},
		{"categories", "Category"}, // ies → y
		{"addresses", "Address"},   // sses → ss
		{"indexes", "Index"},       // xes → x
		{"boss", "Boss"},           // "ss" suffix — no strip
	}
	for _, tt := range tests {
		got := sqlcModelName(tt.table)
		if got != tt.want {
			t.Errorf("sqlcModelName(%q) = %q, want %q", tt.table, got, tt.want)
		}
	}
}
