//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what Run — manifest 검증 전체 실행 통합 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	t.Run("nil_manifest_no_panic", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags for nil manifest, got %d", len(diags))
		}
	})

	t.Run("valid_manifest_runs_all_rules", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pm.ProjectConfig{
				APIVersion: "yongol/v1",
				Kind:       "Project",
				Metadata:   pm.Metadata{Name: "myproject"},
				Backend: pm.Backend{
					Module: "github.com/org/myproject",
					Auth:   &pm.Auth{Mode: "bearer"},
				},
			},
		}
		// Should not panic; aggregates all sub-rule results
		_ = Run(fs)
	})

	t.Run("invalid_manifest_produces_errors", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pm.ProjectConfig{
				APIVersion: "wrong",
				Kind:       "wrong",
			},
		}
		diags := Run(fs)
		if len(diags) == 0 {
			t.Error("expected errors for invalid manifest")
		}
	})
}
