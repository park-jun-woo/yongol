//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteFeatureModules — feature별 service+router 일괄 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteFeatureModules(t *testing.T) {
	t.Run("WritesAllFeatures", func(t *testing.T) {
		appDir := t.TempDir()
		plansByFeature := map[string][]*ir.ServicePlan{
			"workflow": {
				{
					OperationID: "CreateWorkflow",
					HTTPMethod:  "POST",
					TriggerKind: ir.TriggerHTTP,
				},
			},
		}
		names, err := writeFeatureModules(plansByFeature, appDir, nil)
		if err != nil {
			t.Fatalf("writeFeatureModules error: %v", err)
		}
		if len(names) != 1 || names[0] != "workflow" {
			t.Errorf("feature names = %v, want [workflow]", names)
		}
		if _, err := os.Stat(filepath.Join(appDir, "services", "workflow.py")); err != nil {
			t.Errorf("expected services/workflow.py: %v", err)
		}
	})

	t.Run("PropagatesWriteOneFeatureError", func(t *testing.T) {
		// appDir/services collides with a regular file -> writeOneFeature fails.
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "services"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		plansByFeature := map[string][]*ir.ServicePlan{"workflow": nil}
		names, err := writeFeatureModules(plansByFeature, appDir, nil)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if names != nil {
			t.Errorf("expected nil names on error, got %v", names)
		}
	})
}
