//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-402 테스트 — access_token_ttl 상한 경고

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func fsWithAccessTTL(ttl string) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{AccessTokenTTL: ttl},
			},
		},
	}
}

func TestSEC402_ExceedsUpperBound_Warns(t *testing.T) {
	diags := sec402AccessTTLUpperBound(fsWithAccessTTL("2h"))
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if diags[0].Level != diagnostic.LevelWarning {
		t.Fatalf("expected WARNING, got %v", diags[0].Level)
	}
}

func TestSEC402_AtUpperBound_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("30m")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics at boundary 30m, got %d", len(diags))
	}
}

func TestSEC402_BelowUpperBound_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("15m")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics at 15m, got %d", len(diags))
	}
}

func TestSEC402_Empty_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when empty, got %d", len(diags))
	}
}

func TestSEC402_Unparseable_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("nonsense")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when unparseable, got %d", len(diags))
	}
}
