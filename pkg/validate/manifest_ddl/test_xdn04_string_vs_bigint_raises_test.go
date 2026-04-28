//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN04_StringClaimAgainstBigint_RaisesError — string claim ↔ int64 (BIGINT) 컬럼 → ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN04_StringClaimAgainstBigint_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					// org_id is int64 in DDL — claim defaults to string → mismatch.
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID": {Key: "org_id"},
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
	if !strings.Contains(d[0].Message, "string") || !strings.Contains(d[0].Message, "int64") {
		t.Errorf("expected message to mention both types: %s", d[0].Message)
	}
}
