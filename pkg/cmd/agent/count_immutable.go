//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what countImmutable — immutable 파일을 가리키는 ERROR 진단 수 반환

package agent

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// countImmutable returns how many diagnostics target immutable files.
func countImmutable(diags []diagnostic.Diagnostic) int {
	n := 0
	for _, d := range diags {
		if d.Level == diagnostic.LevelError && isImmutable(d.File) {
			n++
		}
	}
	return n
}
