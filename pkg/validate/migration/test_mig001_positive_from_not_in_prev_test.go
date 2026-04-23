//ff:func feature=validate type=test control=sequence topic=migration-hints
//ff:what MIG-001 positive — @rename from 이 이전 스냅샷에 없으면 ERROR

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG001_Positive_FromNotInPrev(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email VARCHAR(255));`)
	hints := &migration.Hints{
		RenameColumns: []migration.RenameColumnHint{
			{Table: "users", From: "nonexistent", To: "email"},
		},
	}
	diags := Mig001RenameMismatch(prev, curr, hints)
	if len(diags) == 0 {
		t.Errorf("expected MIG-001 diagnostic, got none")
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR level, got %v", diags[0].Level)
	}
}
