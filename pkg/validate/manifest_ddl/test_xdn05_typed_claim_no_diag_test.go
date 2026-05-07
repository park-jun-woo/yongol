//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN05_TypedClaim_NoDiag — 타입 명시 claim → XDN-05 PASS

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN05_TypedClaim_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"OrgID": {Key: "org_id", GoType: "uuid", Typed: true},
						"Email": {Key: "email", GoType: "string", Typed: true},
					},
				},
			},
		},
	}
	if d := xdn05ClaimTypeRequired(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
