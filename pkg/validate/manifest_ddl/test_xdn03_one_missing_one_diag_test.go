//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN03_OneMissingColumn_OneDiag — claim 1개의 컬럼이 부재할 때 진단 1건 + 라인 매핑 확인

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN03_OneMissingColumn_OneDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgIDX": {Key: "org_idx", GoType: "int64", SourceLine: 9},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	d := xdn03ClaimColumnExists(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-03]") || !strings.Contains(d[0].Message, "org_idx") {
		t.Errorf("unexpected message: %s", d[0].Message)
	}
	if d[0].Line != 9 {
		t.Errorf("expected Line=9, got %d", d[0].Line)
	}
}
