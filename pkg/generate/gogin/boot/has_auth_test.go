//ff:func feature=gen-gogin type=test control=sequence
//ff:what hasAuth — manifest.backend.auth 존재 여부

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasAuth(t *testing.T) {
	if hasAuth(&yongol.Fullstack{}) {
		t.Errorf("nil manifest should be false")
	}
	noAuth := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if hasAuth(noAuth) {
		t.Errorf("nil auth should be false")
	}
	emptyClaims := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Auth: &pmanifest.Auth{}},
	}}
	if hasAuth(emptyClaims) {
		t.Errorf("auth without claims should be false")
	}
	withClaims := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Auth: &pmanifest.Auth{
			Claims: map[string]pmanifest.ClaimDef{"sub": {}},
		}},
	}}
	if !hasAuth(withClaims) {
		t.Errorf("auth with claims should be true")
	}
}
