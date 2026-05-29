//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN01_AuthJWT_WithoutUserTable_RaisesError — jwt 활성 + user_table 누락 → ERROR 1건

package manifest_ddl

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN01_AuthJWT_WithoutUserTable_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "jwt"},
			},
		},
	}
	d := xdn01UserTableRequired(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-01]") {
		t.Errorf("missing XDN-01 prefix: %s", d[0].Message)
	}
}
