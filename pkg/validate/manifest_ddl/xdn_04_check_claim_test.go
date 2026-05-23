//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn04CheckClaim — column 누락/타입 일치/타입 불일치 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXdn04CheckClaim(t *testing.T) {
	columns := map[string]ddl.Column{
		"id":    {RawType: "BIGINT"},
		"email": {RawType: "TEXT"},
		"role":  {RawType: "VARCHAR(50)"},
	}

	tests := []struct {
		name     string
		field    string
		def      pmanifest.ClaimDef
		wantFire bool
	}{
		{
			name:     "column not found defers to XDN-03",
			field:    "OrgID",
			def:      pmanifest.ClaimDef{Key: "org_id", GoType: "int64"},
			wantFire: false,
		},
		{
			name:     "int64 claim matches BIGINT column",
			field:    "UserID",
			def:      pmanifest.ClaimDef{Key: "id", GoType: "int64"},
			wantFire: false,
		},
		{
			name:     "empty GoType defaults to string, matches TEXT",
			field:    "Email",
			def:      pmanifest.ClaimDef{Key: "email", GoType: ""},
			wantFire: false,
		},
		{
			name:     "string claim matches VARCHAR column",
			field:    "Role",
			def:      pmanifest.ClaimDef{Key: "role", GoType: "string"},
			wantFire: false,
		},
		{
			name:     "int64 claim vs TEXT column fires mismatch",
			field:    "Email",
			def:      pmanifest.ClaimDef{Key: "email", GoType: "int64"},
			wantFire: true,
		},
		{
			name:     "string claim vs BIGINT column fires mismatch",
			field:    "UserID",
			def:      pmanifest.ClaimDef{Key: "id", GoType: "string"},
			wantFire: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diag, fired := xdn04CheckClaim(tt.field, "users", tt.def, columns)
			if fired != tt.wantFire {
				t.Errorf("fired = %v, want %v; diag = %+v", fired, tt.wantFire, diag)
			}
		})
	}
}
