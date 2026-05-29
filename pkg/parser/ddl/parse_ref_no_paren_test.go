//ff:func feature=manifest type=test control=sequence
//ff:what parseRef — users → (users, "")

package ddl

import "testing"

func TestParseRef_NoParen(t *testing.T) {
	table, col := parseRef("users")
	if table != "users" || col != "" {
		t.Errorf("got (%q,%q), want (users,\"\")", table, col)
	}
}
