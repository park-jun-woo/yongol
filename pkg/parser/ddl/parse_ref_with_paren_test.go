//ff:func feature=manifest type=test control=sequence
//ff:what parseRef — users(id) → (users, id)

package ddl

import "testing"

func TestParseRef_WithParen(t *testing.T) {
	table, col := parseRef("users(id)")
	if table != "users" || col != "id" {
		t.Errorf("got (%q,%q), want (users,id)", table, col)
	}
}
