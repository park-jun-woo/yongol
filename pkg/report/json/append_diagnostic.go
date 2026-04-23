//ff:func feature=report type=util control=selection topic=json
//ff:what appendDiagnostic — ERROR/WARNING diagnostic 1건을 Document 에 추가하고 summary 카운터 갱신

package json

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// appendDiagnostic inspects a single diagnostic and, when it is an ERROR or
// WARNING, appends the converted entry to doc.Diagnostics and increments the
// matching summary counter. Lower-severity diagnostics are ignored.
func appendDiagnostic(doc *Document, d diagnostic.Diagnostic, specsDir, absSpecs string) {
	switch d.Level {
	case diagnostic.LevelError:
		doc.Summary.Errors++
	case diagnostic.LevelWarning:
		doc.Summary.Warnings++
	default:
		return
	}
	ruleID, msgText := extractRuleID(d.Message)
	doc.Diagnostics = append(doc.Diagnostics, Diagnostic{
		RuleID:  ruleID,
		Level:   string(d.Level),
		File:    relativeFile(d.File, specsDir, absSpecs),
		Line:    d.Line,
		Col:     0,
		Message: msgText,
	})
}
