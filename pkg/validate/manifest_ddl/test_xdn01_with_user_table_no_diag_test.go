//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN01_AuthJWT_WithUserTable_NoDiag — auth jwt 활성 + user_table 지정 → 진단 0

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN01_AuthJWT_WithUserTable_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users"},
			},
		},
	}
	if d := xdn01UserTableRequired(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
