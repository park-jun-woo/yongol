//ff:func feature=generate type=util control=sequence
//ff:what runMigrationStep — Generate 내부에서 migration 실행 + warning 로깅 래퍼
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runMigrationStep(fs *yongol.Fullstack, artifactsDir string, cfg *generateConfig) error {
	if fs.SpecsDir == "" {
		return nil
	}
	diags, err := runMigration(fs.SpecsDir, artifactsDir, cfg.migration)
	if err != nil {
		return fmt.Errorf("migration: %w", err)
	}
	logMigrationWarnings(cfg.migration, diags)
	return nil
}
