//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeModels — DDL → SQLAlchemy 모델 파일 기록

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/fastapi/models"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// writeModels renders and writes the SQLAlchemy models from DDL tables.
func writeModels(fs *yongol.Fullstack, appDir string) error {
	modelsDir := filepath.Join(appDir, "models")
	if err := os.MkdirAll(modelsDir, 0o755); err != nil {
		return fmt.Errorf("mkdir models: %w", err)
	}

	// Write __init__.py
	if err := os.WriteFile(filepath.Join(modelsDir, "__init__.py"), []byte(""), 0o644); err != nil {
		return err
	}

	if len(fs.DDLTables) == 0 {
		return nil
	}

	content, err := models.RenderModels(fs.DDLTables)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(modelsDir, "models.py"), []byte(content), 0o644)
}
