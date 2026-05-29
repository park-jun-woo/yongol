//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what kebabToPascal 변환 테이블 검증

package react

import "testing"

func TestKebabToPascal(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"login", "Login"},
		{"workflow-detail", "WorkflowDetail"},
		{"audit-logs", "AuditLogs"},
		{"my-long-page-name", "MyLongPageName"},
		{"dashboard", "Dashboard"},
	}
	for _, tt := range tests {
		got := kebabToPascal(tt.in)
		if got != tt.want {
			t.Errorf("kebabToPascal(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
