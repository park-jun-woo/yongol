//ff:func feature=orchestrator type=util control=iteration
//ff:what stripModelPrefix 테이블 테스트 — prefix 일치/불일치/빈 model
package sqlc

import "testing"

func TestStripModelPrefix_Table(t *testing.T) {
	cases := []struct {
		name  string
		query string
		model string
		want  string
	}{
		{"match", "UserFindByID", "User", "FindByID"},
		{"no-match", "ListUsers", "Workflow", "ListUsers"},
		{"empty-model", "UserCreate", "", "UserCreate"},
		{"exact-match-returns-original", "User", "User", "User"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stripModelPrefix(tc.query, tc.model)
			if got != tc.want {
				t.Errorf("stripModelPrefix(%q, %q) = %q, want %q", tc.query, tc.model, got, tc.want)
			}
		})
	}
}
