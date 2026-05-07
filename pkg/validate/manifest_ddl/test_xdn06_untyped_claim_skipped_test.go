//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN06_UntypedClaim_Skipped — 타입 미선언 claim 은 XDN-06 건너뜀 (XDN-05 영역)

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN06_UntypedClaim_Skipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID": {Key: "org_id", GoType: "string", Typed: false},
					},
				},
			},
		},
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"org_id": {Name: "org_id", RawType: "UUID"},
			},
		}},
	}
	if d := xdn06ClaimDDLType(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics for untyped claim, got %+v", d)
	}
}
