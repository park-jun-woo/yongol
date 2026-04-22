//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateAuthz — default CheckRequest 필드 집합 vs custom authz package

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestPopulateAuthz_Defaults verifies the default field set when no custom
// authz package is configured.
func TestPopulateAuthz_Defaults(t *testing.T) {
	fs := newMinimalFullstack()
	g := newGround()

	populateAuthz(g, fs)

	set := g.Lookup["Authz.checkRequest"]
	for _, f := range []string{"Action", "Resource", "UserID", "Role", "ResourceID"} {
		if !set[f] {
			t.Errorf("Authz.checkRequest missing %q: %v", f, set)
		}
	}
}

// TestPopulateAuthz_CustomPackage empties the default set so XAS-60 skips.
func TestPopulateAuthz_CustomPackage(t *testing.T) {
	fs := newMinimalFullstack(func(fs *yongol.Fullstack) {
		fs.Manifest = &manifest.ProjectConfig{
			Authz: &manifest.AuthzConfig{Package: "customauthz"},
		}
	})
	g := newGround()

	populateAuthz(g, fs)

	set := g.Lookup["Authz.checkRequest"]
	if len(set) != 0 {
		t.Errorf("Authz.checkRequest should be empty when custom authz package, got %v", set)
	}
}
