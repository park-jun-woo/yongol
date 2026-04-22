//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateMiddleware — bearerAuth+auth.claims 시 claim KEY가 Middleware.claims에 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
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

// TestPopulateMiddleware_NoBearerAuth — no key is written when bearerAuth
// middleware is absent.
func TestPopulateMiddleware_NoBearerAuth(t *testing.T) {
	fs := newMinimalFullstack(func(fs *yongol.Fullstack) {
		fs.Manifest = &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Middleware: []string{"cors"},
				Auth:       &manifest.Auth{Claims: map[string]manifest.ClaimDef{"UserID": {Key: "user_id"}}},
			},
		}
	})
	g := newGround()

	populateMiddleware(g, fs)

	if _, ok := g.Lookup["Middleware.claims"]; ok {
		t.Errorf("Middleware.claims must not exist without bearerAuth")
	}
}

// TestPopulateMiddleware_NilManifest short-circuits safely.
func TestPopulateMiddleware_NilManifest(t *testing.T) {
	g := newGround()
	populateMiddleware(g, newMinimalFullstack())
	if len(g.Lookup) != 0 {
		t.Errorf("expected no keys, got %d", len(g.Lookup))
	}
}

// TestHasBearerAuthMiddleware covers the small pure helper.
func TestHasBearerAuthMiddleware(t *testing.T) {
	if !hasBearerAuthMiddleware([]string{"cors", "bearerAuth", "gzip"}) {
		t.Errorf("expected true when bearerAuth present")
	}
	if hasBearerAuthMiddleware([]string{"cors"}) {
		t.Errorf("expected false when absent")
	}
	if hasBearerAuthMiddleware(nil) {
		t.Errorf("expected false on nil")
	}
}
