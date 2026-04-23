//ff:func feature=report type=emitter control=sequence topic=sarif
//ff:what Emit — validate.Report + catalog → SARIF 2.1.0 full-catalog JSON ([]byte)
package sarif

import (
	"encoding/json"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// informationURI for tool.driver.informationUri (hardcoded per PhaseF01).
const informationURI = "https://github.com/park-jun-woo/yongol"

// schemaURI for the top-level $schema field.
const schemaURI = "https://json.schemastore.org/sarif-2.1.0.json"

// Emit converts a validate.Report into a SARIF 2.1.0 document and returns
// the pretty-printed JSON bytes.
//
//   - yongolVersion: goes into tool.driver.version
//   - specsDir: used to rebase physicalLocation.artifactLocation.uri to a
//     specsDir-relative path when possible.
//   - cat: the full rule catalog (parsed from rulebook.md). When non-nil, all
//     catalogued rules are emitted under tool.driver.rules[] — independent
//     of whether they fired — with shortDescription / helpUri / defaultConfiguration
//     populated from the catalog. results[].ruleIndex is filled from the
//     catalog's stable slice position.
//
// When cat is nil the emitter degrades gracefully to the PhaseF01 "fired only"
// behaviour (used by callers that don't have a catalog, e.g. unit tests).
func Emit(report *validate.Report, yongolVersion, specsDir string, cat *rulecatalog.Catalog) ([]byte, error) {
	results, fired := collectResults(report, specsDir, cat)
	doc := Document{
		Schema:  schemaURI,
		Version: "2.1.0",
		Runs: []Run{{
			Tool: Tool{
				Driver: Driver{
					Name:           "yongol",
					Version:        yongolVersion,
					InformationURI: informationURI,
					Rules:          buildDriverRules(cat, fired),
				},
			},
			Results: results,
		}},
	}
	return json.MarshalIndent(doc, "", "  ")
}
