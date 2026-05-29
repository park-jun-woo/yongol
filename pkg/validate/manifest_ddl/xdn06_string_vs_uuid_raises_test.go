//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN06_StringVsUUID_RaisesError — claim string ↔ DDL UUID → ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN06_StringVsUUID_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID": {Key: "org_id", GoType: "string", Typed: true},
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
	d := xdn06ClaimDDLType(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-06]") {
		t.Errorf("expected XDN-06, got %s", d[0].Message)
	}
	if !strings.Contains(d[0].Message, "string") || !strings.Contains(d[0].Message, "UUID") {
		t.Errorf("expected both types in message: %s", d[0].Message)
	}
}
