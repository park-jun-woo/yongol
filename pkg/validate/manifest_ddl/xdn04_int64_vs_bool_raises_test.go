//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN04_Int64ClaimMismatchedColumnType_RaisesError — int64 claim ↔ bool 컬럼 → ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN04_Int64ClaimMismatchedColumnType_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					// profile_id is bool in DDL fixture — claim says int64.
					Claims: map[string]pmanifest.ClaimDef{
						"ProfileID": {Key: "profile_id", GoType: "int64", SourceLine: 11},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	d := xdn04ClaimColumnType(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-04]") || !strings.Contains(d[0].Message, "ProfileID") {
		t.Errorf("unexpected message: %s", d[0].Message)
	}
	if d[0].Line != 11 {
		t.Errorf("expected Line=11, got %d", d[0].Line)
	}
}
