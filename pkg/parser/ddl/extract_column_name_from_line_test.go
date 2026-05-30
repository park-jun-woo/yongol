//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what extractColumnNameFromLine — 첫 bareword 추출 / 인용 식별자 / 빈 입력

package ddl

import "testing"

func TestExtractColumnNameFromLine(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"id BIGINT NOT NULL", "id"},
		{`"user" TEXT`, "user"},
		{"email VARCHAR(255),", "email"},
		{"count(*)", "count"},
		{"   ", ""},
		{"", ""},
		{"single", "single"},
		{"trailing,", "trailing"},
	}
	for _, c := range cases {
		if got := extractColumnNameFromLine(c.in); got != c.want {
			t.Errorf("extractColumnNameFromLine(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
