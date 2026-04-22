//ff:func feature=cli type=util control=iteration dimension=1
//ff:what countLevels — diagnostic 슬라이스에서 ERROR/WARN 개수 집계
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
