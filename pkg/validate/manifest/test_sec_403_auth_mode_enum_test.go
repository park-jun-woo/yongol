//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-403 테스트 — auth.mode enum 값 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec403_ValidModes_NoDiag(t *testing.T) {
	// The empty string resolves to "cookie" via ResolvedMode() and must
	// not be flagged — requiring authors to spell out the default would
	// force churn on every project that adopts Phase020 defaults.
	for _, mode := range []string{"", "cookie", "bearer", "hybrid"} {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					Auth: &pmanifest.Auth{Mode: mode},
				},
			},
		}
		if got := sec403AuthModeEnum(fs); len(got) != 0 {
			t.Fatalf("mode=%q should not emit SEC-403, got %+v", mode, got)
		}
	}
}

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

func TestSec403_NoAuth_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if got := sec403AuthModeEnum(fs); len(got) != 0 {
		t.Fatalf("no auth block should not emit SEC-403, got %+v", got)
	}
}
