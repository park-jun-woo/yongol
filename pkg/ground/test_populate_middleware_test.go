//ff:func feature=rule type=test control=sequence
//ff:what populateMiddleware — bearerAuth+auth.claims 시 claim KEY가 Middleware.claims에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestPopulateMiddleware_BearerAuthWithClaims verifies claim KEYs (not field
// names) populate Middleware.claims.
func TestPopulateMiddleware_BearerAuthWithClaims(t *testing.T) {
	fs := newMinimalFullstack(func(fs *yongol.Fullstack) {
		fs.Manifest = &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Middleware: []string{"bearerAuth"},
				Auth: &manifest.Auth{
					Claims: map[string]manifest.ClaimDef{
						"UserID": {Key: "user_id"},
						"OrgID":  {Key: "org_id"},
					},
				},
			},
		}
	})
	g := newGround()

	populateMiddleware(g, fs)

	set := g.Lookup["Middleware.claims"]
	if !set["user_id"] {
		t.Errorf("Middleware.claims missing 'user_id' (claim KEY not field name): %v", set)
	}
	if !set["org_id"] {
		t.Errorf("Middleware.claims missing 'org_id': %v", set)
	}
}
