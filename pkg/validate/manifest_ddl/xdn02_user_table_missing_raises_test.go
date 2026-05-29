//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN02_UserTableMissing_RaisesError — user_table 이름의 테이블이 DDL 에 없을 때 ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN02_UserTableMissing_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users", UserTableLine: 7},
			},
		},
		DDLTables: []ddl.Table{{Name: "organizations", Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT"}}}},
	}
	d := xdn02UserTableExists(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-02]") {
		t.Errorf("missing XDN-02 prefix: %s", d[0].Message)
	}
	if d[0].Line != 7 {
		t.Errorf("expected Line=7 (UserTableLine), got %d", d[0].Line)
	}
}
