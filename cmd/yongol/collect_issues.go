//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what collectIssues — Report에서 모든 ERROR/WARNING 수집 (ERROR 우선 정렬)
package main

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func collectIssues(r *validate.Report) []diagnostic.Diagnostic {
	var errors, warnings []diagnostic.Diagnostic
	for _, s := range r.Steps {
		e, w := classifyDiagnostics(s.Diagnostics)
		errors = append(errors, e...)
		warnings = append(warnings, w...)
	}
	return append(errors, warnings...)
}
