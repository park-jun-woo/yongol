//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeBootFiles — main.py + config.py + database.py 파일 기록

package fastapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/fastapi/boot"
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeBootFiles writes main.py, config.py, and database.py.
func writeBootFiles(appDir string, bootPlan *ir.BootPlan, featureNames []string) error {
	mainContent, err := boot.RenderMain(bootPlan, featureNames)
	if err != nil {
		return fmt.Errorf("render main.py: %w", err)
	}
	if err := os.WriteFile(filepath.Join(appDir, "main.py"), []byte(mainContent), 0o644); err != nil {
		return fmt.Errorf("write main.py: %w", err)
	}

	configContent, err := boot.RenderConfig()
	if err != nil {
		return fmt.Errorf("render config.py: %w", err)
	}
	if err := os.WriteFile(filepath.Join(appDir, "config.py"), []byte(configContent), 0o644); err != nil {
		return fmt.Errorf("write config.py: %w", err)
	}

	dbContent, err := boot.RenderDatabase()
	if err != nil {
		return fmt.Errorf("render database.py: %w", err)
	}
	return os.WriteFile(filepath.Join(appDir, "database.py"), []byte(dbContent), 0o644)
}
