//ff:func feature=report type=emitter control=iteration dimension=2 topic=sarif
//ff:what Emit — validate.Report + catalog → SARIF 2.1.0 full-catalog JSON ([]byte)
package sarif

import (
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// informationURI for tool.driver.informationUri (hardcoded per PhaseF01).
const informationURI = "https://github.com/park-jun-woo/yongol"

// helpURIBase is the GitHub blob URL prefix used to link each rule back to
// its section in the canonical rulebook.md.
const helpURIBase = "https://github.com/park-jun-woo/yongol/blob/master/rulebook.md"

// schemaURI for the top-level $schema field.
const schemaURI = "https://json.schemastore.org/sarif-2.1.0.json"

// ruleIDPattern matches a leading "[RULE-ID]" prefix in diagnostic messages,
// e.g. "[S-27] Variable foo not declared" or "[XOS-15] ...".
// Rule IDs: uppercase letters + dash + digits (S-27, XOS-15, Q-10, M-2, …).
var ruleIDPattern = regexp.MustCompile(`^\[([A-Z]+(?:-[A-Z]+)*-\d+)\]\s*`)

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

// buildDriverRules materialises the tool.driver.rules[] array.
//
// When a catalog is provided every catalogued rule is included (full
// catalog mode). Without a catalog only rules that fired are included
// (legacy fallback for tests / callers without a catalog).
func buildDriverRules(cat *rulecatalog.Catalog, fired map[string]struct{}) []Rule {
	if cat != nil && cat.Len() > 0 {
		out := make([]Rule, 0, cat.Len())
		for _, m := range cat.Rules {
			out = append(out, ruleFromMeta(m))
		}
		return out
	}
	if len(fired) == 0 {
		return nil
	}
	out := make([]Rule, 0, len(fired))
	// Stable order for deterministic output.
	ids := make([]string, 0, len(fired))
	for id := range fired {
		ids = append(ids, id)
	}
	// Caller test already asserts contents, not ordering — sort for safety.
	sortStrings(ids)
	for _, id := range ids {
		out = append(out, Rule{ID: id})
	}
	return out
}

// ruleFromMeta renders a catalog RuleMeta as a SARIF Rule entry.
func ruleFromMeta(m rulecatalog.RuleMeta) Rule {
	r := Rule{
		ID:   m.ID,
		Name: m.ID, // canonical rulebook uses ID as the display name today
	}
	if m.Description != "" {
		r.ShortDescription = &Message{Text: m.Description}
	}
	if m.SectionAnchor != "" {
		r.HelpURI = helpURIBase + "#" + m.SectionAnchor
	}
	switch m.Level {
	case "ERROR":
		r.DefaultConfiguration = &DefaultConfiguration{Level: "error"}
	case "WARNING":
		r.DefaultConfiguration = &DefaultConfiguration{Level: "warning"}
	}
	if m.Source != "" {
		r.Properties = map[string]interface{}{"source": m.Source}
	}
	return r
}

// sortStrings sorts s in place. Tiny helper to avoid a package-level
// `import "sort"` when buildDriverRules is exercised only in fallback mode.
func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}

// extractRuleID pulls a leading "[RULE-ID]" token out of a diagnostic message.
// Returns (ruleID, remainingMessage). When no prefix is present the original
// message is returned with an empty ruleID.
func extractRuleID(msg string) (string, string) {
	m := ruleIDPattern.FindStringSubmatchIndex(msg)
	if m == nil {
		return "", msg
	}
	id := msg[m[2]:m[3]]
	rest := strings.TrimSpace(msg[m[1]:])
	return id, rest
}

// mapLevel converts our internal severity into SARIF's level enum.
func mapLevel(l diagnostic.Level) string {
	switch l {
	case diagnostic.LevelError:
		return "error"
	case diagnostic.LevelWarning:
		return "warning"
	}
	return "note"
}

// relativeArtifactURI returns file relative to specsDir when possible.
// Falls back to file as-is when either path is empty or not comparable.
func relativeArtifactURI(file, specsDir, absSpecs string) string {
	if file == "" {
		return ""
	}
	if specsDir == "" {
		return filepath.ToSlash(file)
	}
	// Try: file is already relative or under specsDir literally.
	if rel, err := filepath.Rel(specsDir, file); err == nil && !strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(rel)
	}
	// Try: absolute-versus-absolute comparison.
	if absSpecs != "" {
		if absFile, err := filepath.Abs(file); err == nil {
			if rel, err := filepath.Rel(absSpecs, absFile); err == nil && !strings.HasPrefix(rel, "..") {
				return filepath.ToSlash(rel)
			}
		}
	}
	return filepath.ToSlash(file)
}

// regionOrNil returns a *Region for a positive line number, nil otherwise.
func regionOrNil(line int) *Region {
	if line <= 0 {
		return nil
	}
	return &Region{StartLine: line}
}
