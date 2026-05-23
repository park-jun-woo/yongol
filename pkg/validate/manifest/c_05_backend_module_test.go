//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c05BackendModule — manifest backend.module 비어있는지 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC05BackendModule(t *testing.T) {
	cases := []TestC05BackendModuleCase{
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "valid_module", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Module: "github.com/org/project"}}}, wantCount: 0},
		{name: "empty_module", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 1},
		{name: "whitespace_module", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Module: "  "}}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC05BackendModule(t, c)
		})
	}
}
