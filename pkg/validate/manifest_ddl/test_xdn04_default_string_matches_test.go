//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN04_DefaultStringClaimMatchesVarchar_NoDiag — Go 타입 생략 시 string 으로 간주

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN04_DefaultStringClaimMatchesVarchar_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						// GoType empty → defaults to "string".
						"Email": {Key: "email"},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	if d := xdn04ClaimColumnType(fs); len(d) != 0 {
		t.Fatalf("default string claim should match VARCHAR-derived string: %+v", d)
	}
}
