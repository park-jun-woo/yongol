//ff:func feature=report type=util control=iteration dimension=2 topic=sarif
//ff:what collectResults — validate.Report → SARIF Results + fired ruleID 집합

package sarif

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// collectResults walks every ERROR/WARNING diagnostic across all steps and
// returns the SARIF Results plus a set of rule IDs that actually fired in
// this run (used to build tool.driver.rules[] fallback list).
func collectResults(report *validate.Report, specsDir string, cat *rulecatalog.Catalog) ([]Result, map[string]struct{}) {
	results := []Result{}
	fired := map[string]struct{}{}
	if report == nil {
		return results, fired
	}
	absSpecs, _ := filepath.Abs(specsDir)
	for _, step := range report.Steps {
		for _, d := range step.Diagnostics {
			if d.Level != diagnostic.LevelError && d.Level != diagnostic.LevelWarning {
				continue
			}
			res, ruleID := buildResult(d, specsDir, absSpecs, cat)
			if ruleID != "" {
				fired[ruleID] = struct{}{}
			}
			results = append(results, res)
		}
	}
	return results, fired
}
