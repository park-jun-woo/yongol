//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what c05BackendModule — manifest backend.module 비어있는지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC05BackendModule_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{Module: "github.com/park-jun-woo/zenflow"},
		},
	}
	if got := c05BackendModule(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
