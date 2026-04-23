//ff:func feature=migration type=util control=iteration dimension=1
//ff:what countErrors — 진단 중 ERROR 레벨 개수
package migration

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func countErrors(diags []diagnostic.Diagnostic) int {
	n := 0
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			n++
		}
	}
	return n
}
