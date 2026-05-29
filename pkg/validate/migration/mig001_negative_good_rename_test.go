//ff:func feature=validate type=test control=sequence topic=migration-hints
//ff:what MIG-001 negative — 정상 @rename 은 진단을 생성하지 않는다

package migration

import (
	"testing"

	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG001_Negative_GoodRename(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email VARCHAR(255));`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email_address VARCHAR(255));`)
	hints := &migration.Hints{
		RenameColumns: []migration.RenameColumnHint{
			{Table: "users", From: "email", To: "email_address"},
		},
	}
	diags := Mig001RenameMismatch(prev, curr, hints)
	if len(diags) != 0 {
		t.Errorf("expected no MIG-001 diags, got %+v", diags)
	}
}
