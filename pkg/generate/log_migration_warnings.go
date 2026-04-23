//ff:func feature=generate type=util control=iteration dimension=1
//ff:what logMigrationWarnings — migration 진단의 WARNING 항목을 로거에 출력
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func logMigrationWarnings(h MigrationHook, diags []diagnostic.Diagnostic) {
	if h.Logger == nil {
		return
	}
	for _, d := range diags {
		if d.Level == diagnostic.LevelWarning {
			fmt.Fprintf(h.Logger, "[migration] WARNING %s\n", d.Message)
		}
	}
}
