//ff:func feature=report type=emitter control=sequence topic=sarif
//ff:what Emit — validate.Report + catalog → SARIF 2.1.0 full-catalog JSON ([]byte)
package sarif

import (
	"encoding/json"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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
	doc := Document{
		Schema:  schemaURI,
		Version: "2.1.0",
		Runs: []Run{{
			Tool: Tool{
				Driver: Driver{
					Name:           "yongol",
					Version:        yongolVersion,
					InformationURI: informationURI,
				},
			},
			Results: []Result{},
		}},
	}

	absSpecs, _ := filepath.Abs(specsDir)

	// Build the driver rules array.
	//
	// Full catalog mode: emit every catalogued rule so that consumers can
	// browse the entire ruleset (e.g. GitHub Security tab rule list).
	// Fallback mode (no catalog): emit only rules that fired in this run,
	// preserving the earlier PhaseF01 shape.
	firedRules := map[string]struct{}{}

	if report != nil {
		for _, step := range report.Steps {
			for _, d := range step.Diagnostics {
				if d.Level != diagnostic.LevelError && d.Level != diagnostic.LevelWarning {
					continue
				}
				ruleID, msgText := extractRuleID(d.Message)
				if ruleID != "" {
					firedRules[ruleID] = struct{}{}
				}
				res := Result{
					RuleID:  ruleID,
					Level:   mapLevel(d.Level),
					Message: Message{Text: msgText},
				}
				if d.File != "" {
					res.Locations = []Location{{
						PhysicalLocation: PhysicalLocation{
							ArtifactLocation: ArtifactLocation{URI: relativeArtifactURI(d.File, specsDir, absSpecs)},
							Region:           regionOrNil(d.Line),
						},
					}}
				}
				if cat != nil && ruleID != "" {
					if idx := cat.Index(ruleID); idx >= 0 {
						i := idx
						res.RuleIndex = &i
					}
				}
				doc.Runs[0].Results = append(doc.Runs[0].Results, res)
			}
		}
	}

	doc.Runs[0].Tool.Driver.Rules = buildDriverRules(cat, firedRules)

	return json.MarshalIndent(doc, "", "  ")
}
