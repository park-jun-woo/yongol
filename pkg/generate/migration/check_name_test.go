//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCheckName — 기본 CHECK 제약 이름은 소문자 <table>_<column>_check
package migration

import "testing"

func TestCheckName(t *testing.T) {
	cases := []struct {
		name, table, column, want string
	}{
		{"lower", "users", "age", "users_age_check"},
		{"uppercase normalised", "USERS", "AGE", "users_age_check"},
		{"mixed", "Posts", "Score", "posts_score_check"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := CheckName(c.table, c.column); got != c.want {
				t.Errorf("CheckName(%q,%q) = %q, want %q", c.table, c.column, got, c.want)
			}
		})
	}
}
