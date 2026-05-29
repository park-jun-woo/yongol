//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN03_AllClaimsPresent_NoDiag — 모든 claim 컬럼이 user_table 에 있을 때 진단 0

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN03_AllClaimsPresent_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"ID":    {Key: "user_id", GoType: "int64"},
						"Email": {Key: "email", GoType: "string"},
						"Role":  {Key: "role", GoType: "string"},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	if d := xdn03ClaimColumnExists(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
