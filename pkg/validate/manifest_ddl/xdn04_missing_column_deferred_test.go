//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN04_MissingColumn_DeferredToXDN03 — 컬럼 부재는 XDN-03 의 영역, XDN-04 는 silent

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN04_MissingColumn_DeferredToXDN03(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"Ghost": {Key: "missing_column", GoType: "int64"},
					},
				},
			},
		},
		DDLTables: []ddl.Table{usersTableFixture()},
	}
	if d := xdn04ClaimColumnType(fs); len(d) != 0 {
		t.Fatalf("missing column is XDN-03's job, XDN-04 must stay silent: %+v", d)
	}
}
