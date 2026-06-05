//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteFeatureSchemaFile — writeFeatureSchemaFile 스키마 파일 생성·빈 내용 skip 검증
package fastapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteFeatureSchemaFile(t *testing.T) {
	t.Run("WritesWhenSchemaPresent", func(t *testing.T) {
		dir := t.TempDir()
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateWorkflow",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				BodyFields:  []ir.BodyFieldMeta{{Name: "title", Required: true}},
			},
		}
		if err := writeFeatureSchemaFile(dir, "workflow", plans); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		path := filepath.Join(dir, "workflow.py")
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected schema file written: %v", err)
		}
		if len(data) == 0 {
			t.Error("expected non-empty schema content")
		}
	})

	t.Run("SkipsWhenEmpty", func(t *testing.T) {
		dir := t.TempDir()
		plans := []*ir.ServicePlan{
			{OperationID: "GetWorkflow", HTTPMethod: "GET", TriggerKind: ir.TriggerHTTP},
		}
		if err := writeFeatureSchemaFile(dir, "workflow", plans); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "workflow.py")); !os.IsNotExist(err) {
			t.Error("expected no file written for empty schema")
		}
	})
}
