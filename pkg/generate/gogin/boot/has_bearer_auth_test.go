//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what hasBearerAuth — manifest.backend.middleware 에 bearerAuth 포함 여부

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasBearerAuth(t *testing.T) {
	if hasBearerAuth(&yongol.Fullstack{}) {
		t.Errorf("nil manifest should be false")
	}
	none := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Middleware: []string{"cors", "logger"}},
	}}
	if hasBearerAuth(none) {
		t.Errorf("middleware without bearerAuth should be false")
	}
	present := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Middleware: []string{"cors", "bearerAuth"}},
	}}
	if !hasBearerAuth(present) {
		t.Errorf("middleware with bearerAuth should be true")
	}
}
