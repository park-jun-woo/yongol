//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateDDLArchived — archived/sensitive 어노테이션이 Flags에 투영

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLArchived_TableAndColumns covers the table-level archived
// flag, per-column archived, and sensitive column flags.
func TestPopulateDDLArchived_TableAndColumns(t *testing.T) {
	tab := ddl.Table{
		Name:             "users",
		Archived:         true,
		ArchivedColumns:  map[string]bool{"legacy_col": true},
		SensitiveColumns: map[string]bool{"password_hash": true, "email": true},
	}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDLArchived(g, fs)

	if !g.Flags["archived.users"] {
		t.Errorf("archived.users flag missing")
	}
	if !g.Flags["archived.users.legacy_col"] {
		t.Errorf("archived.users.legacy_col flag missing")
	}
	if !g.Flags["sensitive.users.password_hash"] {
		t.Errorf("sensitive.users.password_hash flag missing")
	}
	if !g.Flags["sensitive.users.email"] {
		t.Errorf("sensitive.users.email flag missing")
	}
}

// TestPopulateDDLArchived_NotArchived ensures no table-level flag is set when
// Archived=false.
func TestPopulateDDLArchived_NotArchived(t *testing.T) {
	tab := ddl.Table{Name: "plain", Archived: false}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDLArchived(g, fs)

	if g.Flags["archived.plain"] {
		t.Errorf("archived.plain should not be set when Archived=false")
	}
}
