//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=dos-guard
//ff:what resolveHTTPLimits — manifest.backend.http 에서 global + per-op limit 추출
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveHTTPLimits_Defaults(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{nil, {}, {Manifest: &pmanifest.ProjectConfig{}}} {
		body, multipart, bo, mo := resolveHTTPLimits(fs)
		if body != defaultBodyLimit || multipart != defaultMultipartLimit {
			t.Errorf("expected defaults, got body=%d multipart=%d", body, multipart)
		}
		if len(bo) != 0 || len(mo) != 0 {
			t.Errorf("expected empty overrides, got bo=%v mo=%v", bo, mo)
		}
	}
}
