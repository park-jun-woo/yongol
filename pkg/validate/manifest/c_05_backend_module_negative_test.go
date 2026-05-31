//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-5 테스트 — backend.module negative

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC05BackendModule_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: ""},
		},
	}
	got := c05BackendModule(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[C-5]") {
		t.Fatalf("message missing [C-5] prefix: %q", got[0].Message)
	}
}
