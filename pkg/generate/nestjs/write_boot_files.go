//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeBootFiles — main.ts + app.module.ts 파일 기록

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/boot"
)

// writeBootFiles writes main.ts and app.module.ts.
func writeBootFiles(srcDir string, bootPlan *ir.BootPlan, featureNames []string) error {
	mainContent, err := boot.RenderMain(bootPlan)
	if err != nil {
		return fmt.Errorf("render main.ts: %w", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "main.ts"), []byte(mainContent), 0o644); err != nil {
		return fmt.Errorf("write main.ts: %w", err)
	}
	appModContent, err := boot.RenderAppModule(featureNames)
	if err != nil {
		return fmt.Errorf("render app.module.ts: %w", err)
	}
	return os.WriteFile(filepath.Join(srcDir, "app.module.ts"), []byte(appModContent), 0o644)
}
