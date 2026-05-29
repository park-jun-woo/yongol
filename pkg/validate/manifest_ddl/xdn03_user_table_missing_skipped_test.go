//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN03_UserTableMissing_Skipped — XDN-02 가 잡는 케이스에서는 XDN-03 silent

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN03_UserTableMissing_Skipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"Email": {Key: "email", GoType: "string"},
					},
				},
			},
		},
		DDLTables: nil,
	}
	if d := xdn03ClaimColumnExists(fs); len(d) != 0 {
		t.Fatalf("missing table → XDN-03 must skip: %+v", d)
	}
}
