//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN02_UserTableEmpty_Skipped — user_table 빈 문자열은 XDN-01 의 영역이므로 XDN-02 skip

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN02_UserTableEmpty_Skipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Type: "jwt"},
			},
		},
	}
	if d := xdn02UserTableExists(fs); len(d) != 0 {
		t.Fatalf("empty user_table is XDN-01's job: %+v", d)
	}
}
