//ff:func feature=generate type=util control=selection
//ff:what runMigration — migration.Generate 호출 후 mode 별 로그 라인 분기
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func runMigration(specsDir, artifactsDir string, h MigrationHook) ([]diagnostic.Diagnostic, error) {
	opt := migration.Options{
		YongolVersion: h.Version,
	}
	res, diags, err := migration.Generate(specsDir, artifactsDir, opt)
	if err != nil {
		return diags, err
	}
	if h.Logger == nil {
		return diags, nil
	}
	switch res.Mode {
	case migration.ModeInitial:
		fmt.Fprintf(h.Logger, "[migration] mode=initial file=%s ops=%d\n",
			res.MigrationFile, res.OpsCount)
	case migration.ModeIncremental:
		logIncrementalMigration(h, res)
	case migration.ModeNoop:
		fmt.Fprintln(h.Logger, "[migration] mode=noop (no schema changes)")
	}
	return diags, nil
}
