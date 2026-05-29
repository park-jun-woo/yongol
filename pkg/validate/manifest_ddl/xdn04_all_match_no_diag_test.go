//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN04_AllTypesMatch_NoDiag — claim Go 타입 ↔ 컬럼 Go 타입 모두 일치 → 진단 0

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN04_AllTypesMatch_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID":   {Key: "org_id", GoType: "int64"},
						"Email":   {Key: "email", GoType: "string"},
						"IsAdmin": {Key: "is_admin", GoType: "bool"},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	if d := xdn04ClaimColumnType(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
