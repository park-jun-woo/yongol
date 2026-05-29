//ff:func feature=rule type=test control=iteration dimension=1
//ff:what populateManifest — claim 필드 GoType을 g.Types["Manifest.claim.<Field>"]에 등록 (uuid → pgtype.UUID 변환 포함)

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestPopulateManifest_ClaimTypes(t *testing.T) {
	pc := &manifest.ProjectConfig{
		Backend: manifest.Backend{
			Auth: &manifest.Auth{
				Type: "jwt",
				Claims: map[string]manifest.ClaimDef{
					"UserID": {Key: "user_id", GoType: "int64"},
					"OrgID":  {Key: "org_id", GoType: "uuid"},
					"Role":   {Key: "role", GoType: "string"},
				},
			},
		},
	}
	fs := newMinimalFullstack(func(fs *yongol.Fullstack) { fs.Manifest = pc })
	g := newGround()

	populateManifest(g, fs)

	tests := []struct {
		field string
		want  string
	}{
		{"UserID", "int64"},
		{"OrgID", "pgtype.UUID"},
		{"Role", "string"},
	}
	for _, tt := range tests {
		key := "Manifest.claim." + tt.field
		got := g.Types[key]
		if got != tt.want {
			t.Errorf("Types[%q] = %q, want %q", key, got, tt.want)
		}
	}
}
