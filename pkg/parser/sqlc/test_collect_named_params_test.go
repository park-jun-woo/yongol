//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what collectNamedParams 테이블 테스트 — @name / sqlc.arg(name) / 혼합
package sqlc

import (
	"reflect"
	"testing"
)

func TestCollectNamedParams_Table(t *testing.T) {
	cases := []struct {
		name string
		line string
		want []string // PascalCase sorted
	}{
		{
			name: "at-name",
			line: "SELECT * FROM users WHERE id = @user_id;",
			want: []string{"UserID"},
		},
		{
			name: "sqlc-arg",
			line: "SELECT * FROM users WHERE id = sqlc.arg(user_id);",
			want: []string{"UserID"},
		},
		{
			name: "mixed-at-and-sqlc-arg",
			line: "SELECT * FROM t WHERE a = @email AND b = sqlc.arg(org_id);",
			want: []string{"Email", "OrgID"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			paramSet := map[string]bool{}
			collectNamedParams(tc.line, paramSet)
			got := sortedKeys(paramSet)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("collectNamedParams(%q) = %v, want %v", tc.line, got, tc.want)
			}
		})
	}
}
