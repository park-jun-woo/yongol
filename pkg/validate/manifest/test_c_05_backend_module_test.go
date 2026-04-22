//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-5 테스트 — backend.module golden

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestC05BackendModule_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "github.com/park-jun-woo/zenflow"},
		},
	}
	if got := c05BackendModule(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
