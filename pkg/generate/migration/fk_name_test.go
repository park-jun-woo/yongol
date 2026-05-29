//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestFKName — 기본 FK 제약 이름은 <table>_<col...>_fkey (소문자)
package migration

import "testing"

func TestFKName(t *testing.T) {
	cases := []struct {
		name    string
		table   string
		columns []string
		want    string
	}{
		{"single", "posts", []string{"user_id"}, "posts_user_id_fkey"},
		{"multi", "edges", []string{"a", "b"}, "edges_a_b_fkey"},
		{"uppercase", "Posts", []string{"User_ID"}, "posts_user_id_fkey"},
		{"no columns", "posts", nil, "posts_fkey"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := FKName(c.table, c.columns); got != c.want {
				t.Errorf("FKName(%q,%v) = %q, want %q", c.table, c.columns, got, c.want)
			}
		})
	}
}
