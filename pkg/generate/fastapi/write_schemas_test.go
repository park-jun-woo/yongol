//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderFeatureSchemas — Pydantic BaseModel 스키마 생성 검증
package fastapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteSchemas(t *testing.T) {
	t.Run("WritesSchemaForBodyFeatureSkipsOthers", func(t *testing.T) {
		appDir := t.TempDir()
		plansByFeature := map[string][]*ir.ServicePlan{
			"workflow": {
				{
					OperationID: "CreateWorkflow",
					HTTPMethod:  "POST",
					TriggerKind: ir.TriggerHTTP,
					BodyFields:  []ir.BodyFieldMeta{{Name: "title", Required: true}},
				},
			},
			"reader": {
				{
					OperationID: "GetWorkflow",
					HTTPMethod:  "GET",
					TriggerKind: ir.TriggerHTTP,
				},
			},
		}
		if err := writeSchemas(plansByFeature, appDir); err != nil {
			t.Fatalf("writeSchemas error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(appDir, "schemas", "workflow.py")); err != nil {
			t.Errorf("expected workflow.py schema: %v", err)
		}
		// reader has no body -> no file written.
		if _, err := os.Stat(filepath.Join(appDir, "schemas", "reader.py")); !os.IsNotExist(err) {
			t.Errorf("expected reader.py to be absent, stat err: %v", err)
		}
	})

	t.Run("MkdirFails", func(t *testing.T) {
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "schemas"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeSchemas(map[string][]*ir.ServicePlan{}, appDir)
		if err == nil || !strings.Contains(err.Error(), "mkdir schemas") {
			t.Errorf("expected mkdir schemas error, got: %v", err)
		}
	})

	t.Run("WriteFileFails", func(t *testing.T) {
		appDir := t.TempDir()
		// Pre-create schemas dir and a directory colliding with the target file.
		schemasDir := filepath.Join(appDir, "schemas")
		if err := os.MkdirAll(filepath.Join(schemasDir, "workflow.py"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		plansByFeature := map[string][]*ir.ServicePlan{
			"workflow": {
				{
					OperationID: "CreateWorkflow",
					HTTPMethod:  "POST",
					TriggerKind: ir.TriggerHTTP,
					BodyFields:  []ir.BodyFieldMeta{{Name: "title", Required: true}},
				},
			},
		}
		err := writeSchemas(plansByFeature, appDir)
		if err == nil || !strings.Contains(err.Error(), "write schema") {
			t.Errorf("expected write schema error, got: %v", err)
		}
	})
}
