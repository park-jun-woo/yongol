//ff:func feature=report type=util control=iteration dimension=2 topic=json
//ff:what collectReportDiagnostics — validate.Report 의 모든 step × diagnostic 을 순회해 Document 에 누적

package json

import (
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// collectReportDiagnostics walks every step's diagnostics, delegating
// per-diagnostic conversion to appendDiagnostic.
func collectReportDiagnostics(doc *Document, report *validate.Report, specsDir, absSpecs string) {
	if report == nil {
		return
	}
	for _, step := range report.Steps {
		for _, d := range step.Diagnostics {
			appendDiagnostic(doc, d, specsDir, absSpecs)
		}
	}
}
