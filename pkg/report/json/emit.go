//ff:func feature=report type=emitter control=sequence topic=json
//ff:what Emit — validate.Report → yongol bespoke flat JSON ([]byte) 변환
package json

import (
	stdjson "encoding/json"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

// Emit converts a validate.Report into the yongol flat JSON document and
// returns pretty-printed bytes.
//
// - yongolVersion: set on the top-level `yongol_version` field
// - specsDir: echoed on `specs_dir`, also used to rebase diagnostic file
//   paths to specsDir-relative slash form
// - checks: number of rules in the active catalog (echoed on summary.checks)
//
// Every ERROR or WARNING diagnostic becomes one Diagnostic entry. Rule IDs
// are extracted from the "[RULE-ID] …" message prefix. When no prefix is
// present, rule_id is "" — consumers decide how to handle.
func Emit(report *validate.Report, yongolVersion, specsDir string, checks int) ([]byte, error) {
	doc := Document{
		YongolVersion: yongolVersion,
		SpecsDir:      specsDir,
		Summary:       Summary{Checks: checks},
		Diagnostics:   []Diagnostic{},
	}
	absSpecs, _ := filepath.Abs(specsDir)
	collectReportDiagnostics(&doc, report, specsDir, absSpecs)
	return stdjson.MarshalIndent(doc, "", "  ")
}
