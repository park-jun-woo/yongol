//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXna90_AuthAbsent_NoDiagnostic — backend.auth 미구성 시 XNA-90 트리거 없음

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXna90_AuthAbsent_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{},
	}
	if diags := xna90RefreshRequiresSQLC(fs); len(diags) != 0 {
		t.Errorf("no backend.auth must not trigger XNA-90: %+v", diags)
	}
}
