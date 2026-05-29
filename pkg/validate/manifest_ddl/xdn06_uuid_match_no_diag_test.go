//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN06_UUIDMatch_NoDiag — claim uuid ↔ DDL UUID → PASS

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN06_UUIDMatch_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"ID":    {Key: "id", GoType: "uuid", Typed: true},
						"OrgID": {Key: "org_id", GoType: "uuid", Typed: true},
					},
				},
			},
		},
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"id":     {Name: "id", RawType: "UUID"},
				"org_id": {Name: "org_id", RawType: "UUID"},
			},
		}},
	}
	if d := xdn06ClaimDDLType(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
