//ff:func feature=report type=util control=sequence topic=sarif
//ff:what buildResult — 단일 diagnostic → SARIF Result 변환 (locations, ruleIndex 포함)

package sarif

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

// buildResult converts a single diagnostic into a fully populated SARIF
// Result and returns it alongside the extracted rule ID (for fired-rule
// tracking).
func buildResult(d diagnostic.Diagnostic, specsDir, absSpecs string, cat *rulecatalog.Catalog) (Result, string) {
	ruleID, msgText := extractRuleID(d.Message)
	res := Result{
		RuleID:    ruleID,
		Level:     mapLevel(d.Level),
		Message:   Message{Text: msgText},
		Locations: buildResultLocations(d, specsDir, absSpecs),
	}
	attachRuleIndex(&res, cat, ruleID)
	return res, ruleID
}
