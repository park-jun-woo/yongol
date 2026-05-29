//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN05_InvalidType_RaisesError — 허용되지 않는 타입 → XDN-05 ERROR

package manifest_ddl

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN05_InvalidType_RaisesError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID": {Key: "org_id", GoType: "float64", Typed: true},
					},
				},
			},
		},
	}
	d := xdn05ClaimTypeRequired(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDN-05]") {
		t.Errorf("expected XDN-05, got %s", d[0].Message)
	}
	if !strings.Contains(d[0].Message, "float64") {
		t.Errorf("expected message to mention float64: %s", d[0].Message)
	}
}
