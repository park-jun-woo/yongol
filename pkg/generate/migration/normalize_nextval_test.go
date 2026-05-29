//ff:func feature=migration type=test control=sequence
//ff:what TestNormalizeNextval — nextval(...) 내부 ::regclass 캐스트 제거
package migration

import "testing"

func TestNormalizeNextval(t *testing.T) {
	cases := []struct{ in, want string }{
		{"nextval('users_id_seq'::regclass)", "nextval('users_id_seq')"},
		{"nextval('users_id_seq')", "nextval('users_id_seq')"},
		{"no_parens", "no_parens"},
	}
	for _, c := range cases {
		if got := normalizeNextval(c.in); got != c.want {
			t.Errorf("normalizeNextval(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
