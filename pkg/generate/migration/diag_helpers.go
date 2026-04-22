//ff:func feature=migration type=util control=iteration dimension=1
//ff:what diag helpers — generate.go 에서 validate 패키지 의존 없이 MIG-001/MIG-002~005 Diagnostic 생성
package migration

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// mig001From is a generate-side mirror of validate/migration.Mig001RenameMismatch.
// Avoids the validate -> generate import cycle by reimplementing the
// small check here. Intentional duplication: the validate package is
// the normative owner; this version runs inside the generate pipeline.
func mig001From(prev, curr *Schema, hints *Hints) []diagnostic.Diagnostic {
	if hints == nil {
		return nil
	}
	var out []diagnostic.Diagnostic
	for _, r := range hints.RenameTables {
		if _, ok := prev.Tables[r.From]; !ok {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename from=%s (table) not in previous snapshot", r.From),
				Advice:  "fix the 'from' value or remove the hint",
			})
		}
		if _, ok := curr.Tables[r.To]; !ok {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename to=%s (table) not in current DDL", r.To),
				Advice:  "rename the CREATE TABLE to match 'to'",
			})
		}
	}
	for _, r := range hints.RenameColumns {
		if pt, ok := prev.Tables[r.Table]; ok && !tableHasColumn(pt, r.From) {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename from=%s not in previous snapshot %s", r.From, r.Table),
				Advice:  "fix the 'from' value to match the old column name",
			})
		}
		if ct, ok := curr.Tables[r.Table]; ok && !tableHasColumn(ct, r.To) {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename to=%s not in current DDL %s", r.To, r.Table),
				Advice:  "rename the column in DDL to match 'to'",
			})
		}
	}
	return out
}

func tableHasColumn(t *Table, name string) bool {
	for _, c := range t.Columns {
		if c.Name == name {
			return true
		}
	}
	return false
}

// issuesToDiags converts SafetyIssue records (produced by CheckSafety)
// and missing sidecar file paths into diagnostics. ERROR / WARNING
// level derives from SafetyIssue.Level.
func issuesToDiags(issues []SafetyIssue, missing []string) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for _, is := range issues {
		lvl := diagnostic.LevelWarning
		if is.Level == SafetyError {
			lvl = diagnostic.LevelError
		}
		out = append(out, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   lvl,
			Message: "[" + is.RuleID + "] " + is.Message,
			Advice:  is.Advice,
		})
	}
	for _, p := range missing {
		out = append(out, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[MIG-003] @data_migration file not found: " + p,
			Advice:  "create the sidecar file or fix the path in the hint",
		})
	}
	return out
}
