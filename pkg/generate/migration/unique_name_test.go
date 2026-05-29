//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestUniqueName — 기본 UNIQUE 제약 이름은 <table>_<col...>_key (소문자)
package migration

import "testing"

func TestUniqueName(t *testing.T) {
	cases := []struct {
		name    string
		table   string
		columns []string
		want    string
	}{
		{"single", "users", []string{"email"}, "users_email_key"},
		{"multi", "users", []string{"org", "email"}, "users_org_email_key"},
		{"uppercase", "Users", []string{"Email"}, "users_email_key"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := UniqueName(c.table, c.columns); got != c.want {
				t.Errorf("UniqueName(%q,%v) = %q, want %q", c.table, c.columns, got, c.want)
			}
		})
	}
}
