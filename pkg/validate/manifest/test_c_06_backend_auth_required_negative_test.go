//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-6 테스트 — backend 블록은 있으나 auth 부재 시 ERROR 1건

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestC06BackendAuthRequired_AuthAbsent_Error — backend 블록은 있지만 auth 가
// 없으면 [C-6] ERROR 1건. 메시지 prefix / Level / advice 키워드 검사.
func TestC06BackendAuthRequired_AuthAbsent_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "github.com/park-jun-woo/no-auth",
			},
		},
	}
	got := c06BackendAuthRequired(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[C-6]") {
		t.Fatalf("message missing [C-6] prefix: %q", got[0].Message)
	}
	if got[0].Level != diagnostic.LevelError {
		t.Fatalf("expected LevelError, got %q", got[0].Level)
	}
	if !strings.Contains(got[0].Advice, "auth-free") {
		t.Fatalf("advice missing 'auth-free' guidance: %q", got[0].Advice)
	}
}
