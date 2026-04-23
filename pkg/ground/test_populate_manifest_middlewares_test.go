//ff:func feature=rule type=test control=sequence
//ff:what populateManifest — middleware/claims/roles/queue 등록 및 Config flags

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPopulateManifest_FullSettings covers middleware + auth claims + roles +
// queue backend flag.
func TestPopulateManifest_FullSettings(t *testing.T) {
	pc := &manifest.ProjectConfig{
		Backend: manifest.Backend{
			Middleware: []string{"bearerAuth", "cors"},
			Auth: &manifest.Auth{
				Type: "jwt",
				Claims: map[string]manifest.ClaimDef{
					"UserID": {Key: "user_id", GoType: "int64"},
					"OrgID":  {Key: "org_id", GoType: "int64"},
				},
				Roles: []string{"admin", "member"},
			},
		},
		Queue: &manifest.QueueBackend{Backend: "redis"},
	}
	fs := newMinimalFullstack(func(fs *yongol.Fullstack) { fs.Manifest = pc })
	g := newGround()

	populateManifest(g, fs)

	if !g.Lookup["Manifest.middleware"]["bearerAuth"] {
		t.Errorf("Manifest.middleware missing bearerAuth: %v", g.Lookup["Manifest.middleware"])
	}
	if !g.Lookup["Manifest.claims"]["UserID"] {
		t.Errorf("Manifest.claims missing UserID: %v", g.Lookup["Manifest.claims"])
	}
	if !g.Lookup["Manifest.claims.keys"]["user_id"] {
		t.Errorf("Manifest.claims.keys missing user_id: %v", g.Lookup["Manifest.claims.keys"])
	}
	if !g.Lookup["Manifest.roles"]["admin"] {
		t.Errorf("Manifest.roles missing admin: %v", g.Lookup["Manifest.roles"])
	}
	if !g.Config["queue.backend"] {
		t.Errorf("Config[queue.backend] = false, want true")
	}
	if !g.Config["backend.middleware"] {
		t.Errorf("Config[backend.middleware] = false, want true")
	}
	if !g.Config["backend.auth.claims"] {
		t.Errorf("Config[backend.auth.claims] = false, want true")
	}
}
