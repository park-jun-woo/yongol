//ff:func feature=migration type=util control=iteration dimension=1
//ff:what hasErrorDiag — 진단 중 ERROR 레벨 존재 여부
package migration

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// hasErrorDiag reports whether diags contains at least one ERROR-level
// entry.
func hasErrorDiag(diags []diagnostic.Diagnostic) bool {
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			return true
		}
	}
	return false
}
