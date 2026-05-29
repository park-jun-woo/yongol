//ff:func feature=rule type=test control=sequence
//ff:what populateMiddleware — bearerAuth 부재 시 Middleware.claims 키 미등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
