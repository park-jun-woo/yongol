//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN03_TwoMissingColumns_TwoDiagsStableOrder — 2건 부재 시 정렬된 안정 순서로 진단

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN03_TwoMissingColumns_TwoDiagsStableOrder(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID":   {Key: "org_idx", GoType: "int64"},
						"EmailGh": {Key: "ghost_email", GoType: "string"},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	d := xdn03ClaimColumnExists(fs)
	if len(d) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "EmailGh") {
		t.Errorf("first diag should be EmailGh (sorted), got %s", d[0].Message)
	}
	if !strings.Contains(d[1].Message, "OrgID") {
		t.Errorf("second diag should be OrgID, got %s", d[1].Message)
	}
}
