//ff:func feature=cli type=util control=iteration dimension=1
//ff:what countLevels — tallies ERROR and WARNING counts from a diagnostic slice
package main

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func countLevels(diags []diagnostic.Diagnostic) (errs, warns int) {
	for _, d := range diags {
		switch d.Level {
		case diagnostic.LevelError:
			errs++
		case diagnostic.LevelWarning:
			warns++
		}
	}
	return
}
