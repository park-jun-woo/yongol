//ff:func feature=agent type=test control=sequence
//ff:what TestTableFromDDLFile — DDL 파일명/경로에서 테이블명(.sql 제거) 추출 검증

package agent

import "testing"

func TestTableFromDDLFile(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"db/users.sql", "users"},
		{"users.sql", "users"},
		{"db/order_items.sql", "order_items"},
		{"plain", "plain"},
	}
	for _, c := range cases {
		if got := tableFromDDLFile(c.in); got != c.want {
			t.Errorf("tableFromDDLFile(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
