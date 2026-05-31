//ff:func feature=gen-react type=test control=sequence
//ff:what theme_picks_unit_test — pick_* 색상 접근자(via writeTailwindConfig) + orDefault + resolveTheme 단위 테스트
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveTheme(t *testing.T) {
	if got := resolveTheme(nil); got != nil {
		t.Errorf("nil fullstack -> %v, want nil", got)
	}
	if got := resolveTheme(&yongol.Fullstack{}); got != nil {
		t.Errorf("nil manifest -> %v, want nil", got)
	}
	theme := &manifest.FrontendTheme{Primary: "#zzz"}
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Theme: theme},
		},
	}
	if got := resolveTheme(fs); got != theme {
		t.Errorf("resolveTheme should return the embedded theme pointer, got %v", got)
	}
}
