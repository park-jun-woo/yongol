//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestTableFromSQLcFile — sqlc 파일명/경로에서 테이블명(.sql 제거) 추출 검증
package agent

import (
	"testing"
)

func TestTableFromSQLcFile(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"db/queries/users.sql", "users"},
		{"users.sql", "users"},
		{"db/queries/order_items.sql", "order_items"},
		{"plain", "plain"},
	}
	for _, c := range cases {
		if got := tableFromSQLcFile(c.in); got != c.want {
			t.Errorf("tableFromSQLcFile(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
