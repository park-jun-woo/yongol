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

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteSchemas — feature 별 Pydantic 스키마 파일 기록 검증
//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestSchemaFormatToPython — OpenAPI format → Python 타입 매핑 검증
func TestSchemaFormatToPython(t *testing.T) {
	cases := map[string]string{
		"email":     "str",
		"uuid":      "str",
		"uri":       "str",
		"url":       "str",
		"":          "str",
		"date-time": "str",
		"date":      "str",
		"int32":     "int",
		"int64":     "int",
		"float":     "float",
		"double":    "float",
		"boolean":   "bool",
		"unknown":   "str",
	}
	for format, want := range cases {
		if got := schemaFormatToPython(format); got != want {
			t.Errorf("schemaFormatToPython(%q) = %q, want %q", format, got, want)
		}
	}
}

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestSchemaPascalCase — camelCase/snake_case/PascalCase → PascalCase 변환 검증
func TestSchemaPascalCase(t *testing.T) {
	cases := map[string]string{
		"":               "",
		"createWorkflow": "CreateWorkflow",
		"CreateWorkflow": "CreateWorkflow",
		"x":              "X",
	}
	for in, want := range cases {
		if got := schemaPascalCase(in); got != want {
			t.Errorf("schemaPascalCase(%q) = %q, want %q", in, got, want)
		}
	}
}

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

func TestRenderFeatureSchemas(t *testing.T) {
	t.Run("POSTWithBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateWorkflow",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				BodyFields: []ir.BodyFieldMeta{
					{Name: "title", Required: true},
					{Name: "description", Required: false},
				},
			},
		}
		got := renderFeatureSchemas(plans)
		if !strings.Contains(got, "from pydantic import BaseModel") {
			t.Errorf("expected pydantic import, got: %s", got)
		}
		if !strings.Contains(got, "class CreateWorkflowRequest(BaseModel):") {
			t.Errorf("expected class definition, got: %s", got)
		}
		if !strings.Contains(got, "title: str") {
			t.Errorf("expected required field, got: %s", got)
		}
		if !strings.Contains(got, "description: Optional[str] = None") {
			t.Errorf("expected optional field, got: %s", got)
		}
	})

	t.Run("GETNoBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "GetWorkflow",
				HTTPMethod:  "GET",
				TriggerKind: ir.TriggerHTTP,
			},
		}
		got := renderFeatureSchemas(plans)
		if got != "" {
			t.Errorf("expected empty for GET plan, got: %s", got)
		}
	})

	t.Run("SubscribeNoBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "HandleEvent",
				TriggerKind: ir.TriggerSubscribe,
				Topic:       "event.created",
			},
		}
		got := renderFeatureSchemas(plans)
		if got != "" {
			t.Errorf("expected empty for subscribe plan, got: %s", got)
		}
	})
}
