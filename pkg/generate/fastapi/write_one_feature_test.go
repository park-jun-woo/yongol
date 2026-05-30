//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteOneFeature — 단일 feature service+router 파일 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteOneFeature(t *testing.T) {
	t.Run("WritesServiceAndRouter", func(t *testing.T) {
		appDir := t.TempDir()
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateWorkflow",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				BodyFields: []ir.BodyFieldMeta{
					{Name: "title", Required: true},
				},
			},
		}
		if err := writeOneFeature("workflow", plans, appDir, nil); err != nil {
			t.Fatalf("writeOneFeature error: %v", err)
		}

		svcPath := filepath.Join(appDir, "services", "workflow.py")
		svc, err := os.ReadFile(svcPath)
		if err != nil {
			t.Fatalf("read service file: %v", err)
		}
		if len(svc) == 0 {
			t.Errorf("service file is empty")
		}

		routerPath := filepath.Join(appDir, "routers", "workflow.py")
		router, err := os.ReadFile(routerPath)
		if err != nil {
			t.Fatalf("read router file: %v", err)
		}
		if !strings.Contains(string(router), "APIRouter") {
			t.Errorf("router file missing APIRouter, got: %s", router)
		}
	})

	t.Run("MkdirServicesFails", func(t *testing.T) {
		// appDir/services already exists as a regular file -> MkdirAll fails.
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "services"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeOneFeature("workflow", nil, appDir, nil)
		if err == nil || !strings.Contains(err.Error(), "mkdir services") {
			t.Errorf("expected mkdir services error, got: %v", err)
		}
	})

	t.Run("MkdirRoutersFails", func(t *testing.T) {
		// services dir creatable, but routers path collides with a file.
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "routers"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeOneFeature("workflow", nil, appDir, nil)
		if err == nil || !strings.Contains(err.Error(), "mkdir routers") {
			t.Errorf("expected mkdir routers error, got: %v", err)
		}
	})

	t.Run("RenderRouterFailsOnEmptyFeature", func(t *testing.T) {
		appDir := t.TempDir()
		err := writeOneFeature("", nil, appDir, nil)
		if err == nil || !strings.Contains(err.Error(), "render router") {
			t.Errorf("expected render router error, got: %v", err)
		}
	})
}
