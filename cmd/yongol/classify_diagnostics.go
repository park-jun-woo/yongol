//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what classifyDiagnostics — 진단 슬라이스를 ERROR/WARNING 으로 분류
package main

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func classifyDiagnostics(diags []diagnostic.Diagnostic) (errors, warnings []diagnostic.Diagnostic) {
	for _, d := range diags {
		switch d.Level {
		case diagnostic.LevelError:
			errors = append(errors, d)
		case diagnostic.LevelWarning:
			warnings = append(warnings, d)
		}
	}
	return errors, warnings
}
