//ff:func feature=gen-gogin type=test control=iteration topic=sqlc-post
//ff:what TestStructNameFor — 테이블명 단수화+PascalCase 매핑 + 빈 파트 스킵 검증

package sqlcpost

import "testing"

func TestStructNameFor(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"users", "User"},
		{"audit_logs", "AuditLog"},
		{"execution_logs", "ExecutionLog"},
		{"categories", "Category"}, // ies -> y via singularize
		{"data", "Data"},           // no plural change
		{"_orphan_", "Orphan"},     // empty parts skipped
		{"", ""},
	}
	for _, c := range cases {
		if got := structNameFor(c.in); got != c.want {
			t.Errorf("structNameFor(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
