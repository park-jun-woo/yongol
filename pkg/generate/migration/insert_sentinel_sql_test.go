//ff:func feature=migration type=test control=sequence
//ff:what TestInsertSentinel_SQL — Body 를 그대로 반환
package migration

import "testing"

func TestInsertSentinel_SQL(t *testing.T) {
	body := "INSERT INTO roles (id, name) VALUES (0, 'none') ON CONFLICT DO NOTHING;"
	op := InsertSentinel{Table: "roles", Body: body}
	if got := op.SQL(); got != body {
		t.Errorf("SQL() = %q, want %q", got, body)
	}
}
