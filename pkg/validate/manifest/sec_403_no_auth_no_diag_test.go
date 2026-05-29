//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-403 테스트 — auth 블록 없으면 미발화

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec403_NoAuth_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if got := sec403AuthModeEnum(fs); len(got) != 0 {
		t.Fatalf("no auth block should not emit SEC-403, got %+v", got)
	}
}
