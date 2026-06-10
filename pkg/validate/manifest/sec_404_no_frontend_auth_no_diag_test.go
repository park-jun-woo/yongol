//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-404 테스트 — frontend.auth 블록 없으면 미발화

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec404_NoFrontendAuth_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if got := sec404FrontendAuthStoreEnum(fs); len(got) != 0 {
		t.Fatalf("no frontend.auth block should not emit SEC-404, got %+v", got)
	}
	if got := sec404FrontendAuthStoreEnum(&yongol.Fullstack{}); len(got) != 0 {
		t.Fatalf("nil manifest should not emit SEC-404, got %+v", got)
	}
}
