//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-403 테스트 — 오타 mode 값은 ERROR

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec403_UnknownMode_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Mode: "cookies"}, // typo
			},
		},
	}
	got := sec403AuthModeEnum(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic for typo mode, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[SEC-403]") {
		t.Fatalf("message missing [SEC-403] prefix: %q", got[0].Message)
	}
}
