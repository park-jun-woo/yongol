//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteModels — DDL → SQLAlchemy 모델 파일 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestWriteModels(t *testing.T) {
	t.Run("EmptyDDLTables", func(t *testing.T) {
		appDir := t.TempDir()
		fs := &yongol.Fullstack{}
		if err := writeModels(fs, appDir); err != nil {
			t.Fatalf("writeModels error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(appDir, "models", "__init__.py")); err != nil {
			t.Errorf("expected models/__init__.py: %v", err)
		}
		if _, err := os.Stat(filepath.Join(appDir, "models", "models.py")); !os.IsNotExist(err) {
			t.Errorf("models.py should not exist for empty DDLTables")
		}
	})

	t.Run("WithDDLTables", func(t *testing.T) {
		appDir := t.TempDir()
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{
					Name: "courses",
					Columns: map[string]ddl.Column{
						"id":    {Name: "id", RawType: "BIGINT", NotNull: true},
						"title": {Name: "title", RawType: "TEXT", NotNull: true},
					},
					PrimaryKey: []string{"id"},
				},
			},
		}
		if err := writeModels(fs, appDir); err != nil {
			t.Fatalf("writeModels error: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(appDir, "models", "models.py"))
		if err != nil {
			t.Fatalf("expected models.py: %v", err)
		}
		if !strings.Contains(string(data), "courses") {
			t.Errorf("models.py missing table content: %s", data)
		}
	})

	t.Run("MkdirModelsFails", func(t *testing.T) {
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "models"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeModels(&yongol.Fullstack{}, appDir)
		if err == nil || !strings.Contains(err.Error(), "mkdir models") {
			t.Errorf("expected mkdir models error, got: %v", err)
		}
	})
}
