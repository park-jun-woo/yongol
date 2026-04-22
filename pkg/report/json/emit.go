//ff:func feature=report type=emitter control=iteration dimension=2 topic=json
//ff:what Emit — validate.Report → yongol bespoke flat JSON ([]byte) 변환
package json

import (
	stdjson "encoding/json"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// ruleIDPattern matches a leading "[RULE-ID]" prefix in diagnostic messages
// (same shape as the SARIF emitter's pattern: uppercase letters + dash +
// digits; supports compound prefixes like XOS-15).
var ruleIDPattern = regexp.MustCompile(`^\[([A-Z]+(?:-[A-Z]+)*-\d+)\]\s*`)

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
		SpecsDir:       specsDir,
		Summary:        Summary{Checks: checks},
		Diagnostics:    []Diagnostic{},
	}

	absSpecs, _ := filepath.Abs(specsDir)

	if report != nil {
		for _, step := range report.Steps {
			for _, d := range step.Diagnostics {
				if d.Level != diagnostic.LevelError && d.Level != diagnostic.LevelWarning {
					continue
				}
				ruleID, msgText := extractRuleID(d.Message)
				entry := Diagnostic{
					RuleID:  ruleID,
					Level:   string(d.Level), // "ERROR" or "WARNING"
					File:    relativeFile(d.File, specsDir, absSpecs),
					Line:    d.Line,
					Col:     0, // diagnostic.Diagnostic has no column today
					Message: msgText,
				}
				doc.Diagnostics = append(doc.Diagnostics, entry)
				switch d.Level {
				case diagnostic.LevelError:
					doc.Summary.Errors++
				case diagnostic.LevelWarning:
					doc.Summary.Warnings++
				}
			}
		}
	}

	return stdjson.MarshalIndent(doc, "", "  ")
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

// relativeFile rebases a file path against specsDir when possible. Mirrors
// the SARIF emitter behaviour for consistency across formats.
func relativeFile(file, specsDir, absSpecs string) string {
	if file == "" {
		return ""
	}
	if specsDir == "" {
		return filepath.ToSlash(file)
	}
	if rel, err := filepath.Rel(specsDir, file); err == nil && !strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(rel)
	}
	if absSpecs != "" {
		if absFile, err := filepath.Abs(file); err == nil {
			if rel, err := filepath.Rel(absSpecs, absFile); err == nil && !strings.HasPrefix(rel, "..") {
				return filepath.ToSlash(rel)
			}
		}
	}
	return filepath.ToSlash(file)
}
