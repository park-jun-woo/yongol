//ff:func feature=rule type=test control=sequence
//ff:what populateAuthz — custom authz package 설정 시 기본 필드 셋을 비움

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
