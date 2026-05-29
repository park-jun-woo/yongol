//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN01_AuthTypeNone_Skipped — auth.type=none 이면 XDN-01 skip

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN01_AuthTypeNone_Skipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "none"},
			},
		},
	}
	if d := xdn01UserTableRequired(fs); len(d) != 0 {
		t.Fatalf("type=none must skip XDN-01: %+v", d)
	}
}
