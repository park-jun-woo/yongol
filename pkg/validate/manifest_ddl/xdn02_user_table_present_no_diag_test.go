//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN02_UserTablePresent_NoDiag — user_table 가 DDL 에 있으면 진단 0

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN02_UserTablePresent_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users"},
			},
		},
		DDLTables: []ddl.Table{{Name: "users", Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT"}}}},
	}
	if d := xdn02UserTableExists(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
