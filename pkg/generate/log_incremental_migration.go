//ff:func feature=generate type=util control=iteration dimension=1
//ff:what logIncrementalMigration — incremental migration 결과(파일/ops 라인)를 로거에 쓴다
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func logIncrementalMigration(h MigrationHook, res *migration.Result) {
	fmt.Fprintf(h.Logger, "[migration] mode=incremental file=%s ops=%d\n",
		res.MigrationFile, res.OpsCount)
	for _, op := range res.Operations {
		fmt.Fprintf(h.Logger, "[migration]   * %s\n", op.Description())
	}
}
