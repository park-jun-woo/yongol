//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what countErrors — ERROR 레벨 진단 수 반환

package agent

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// countErrors counts ERROR-level diagnostics.
func countErrors(diags []diagnostic.Diagnostic) int {
	n := 0
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			n++
		}
	}
	return n
}
