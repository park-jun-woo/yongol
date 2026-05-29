//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_RenameColumnHint — hints 없으면 drop+add / 있으면 RenameColumn
package migration

import "testing"

func TestDiff_RenameColumnHint(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email VARCHAR(255) NOT NULL);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email_address VARCHAR(255) NOT NULL);`)

	assertDropAndAddWithoutHints(t, prev, curr)

	hints := &Hints{
		RenameColumns: []RenameColumnHint{{Table: "users", From: "email", To: "email_address"}},
	}
	assertRenameColumnWithHints(t, prev, curr, hints)
}
