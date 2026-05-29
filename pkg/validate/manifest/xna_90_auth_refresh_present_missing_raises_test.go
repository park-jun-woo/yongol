//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXna90_AuthRefresh_Present_Missing_Raises — backend.auth 구성 시 refresh_tokens 누락 진단

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXna90_AuthRefresh_Present_Missing_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{SecretEnv: "JWT_SECRET"},
			},
		},
	}
	diags := xna90RefreshRequiresSQLC(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "refresh_tokens") {
		t.Errorf("diagnostic missing refresh_tokens: %s", diags[0].Message)
	}
}
