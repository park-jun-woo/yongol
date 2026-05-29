//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN02_AuthInactive_Skipped — auth.type=none 이면 XDN-02 skip

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN02_AuthInactive_Skipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "none", UserTable: "users"},
			},
		},
	}
	if d := xdn02UserTableExists(fs); len(d) != 0 {
		t.Fatalf("auth inactive must skip XDN-02: %+v", d)
	}
}
