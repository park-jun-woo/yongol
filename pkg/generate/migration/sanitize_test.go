//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestSanitize — 소문자화 + [a-z0-9_] 외 문자를 _ 로 치환
package migration

import "testing"

func TestSanitize(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"Users", "users"},
		{"my-table", "my_table"},
		{"a.b c", "a_b_c"},
		{"ok_123", "ok_123"},
	}
	for _, c := range cases {
		if got := sanitize(c.in); got != c.want {
			t.Errorf("sanitize(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
