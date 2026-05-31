//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestSqlcFieldName — snake_case→PascalCase + id→ID 이니셜리즘 + 빈 파트 스킵 검증
package sqlcpost

import (
	"testing"
)

func TestSqlcFieldName(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"org_id", "OrgID"},
		{"password_hash", "PasswordHash"},
		{"url", "Url"},
		{"id", "ID"},
		{"user_id_id", "UserIDID"},
		{"created_at", "CreatedAt"},
		{"__name__", "Name"}, // consecutive/edge underscores produce empty parts
		{"", ""},
	}
	for _, c := range cases {
		if got := sqlcFieldName(c.in); got != c.want {
			t.Errorf("sqlcFieldName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
