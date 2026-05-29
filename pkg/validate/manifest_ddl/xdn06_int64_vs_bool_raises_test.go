//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN06_Int64VsBool_RaisesError — claim int64 ↔ DDL BOOLEAN → ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN06_Int64VsBool_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"IsAdmin": {Key: "is_admin", GoType: "int64", Typed: true},
					},
				},
			},
		},
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"is_admin": {Name: "is_admin", RawType: "BOOLEAN"},
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
}
