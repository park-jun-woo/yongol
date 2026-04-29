//ff:func feature=validate type=util control=selection topic=query-structural
//ff:what diagnoseOverrideGaps — pgtype override 누락을 1건 진단으로 묶어 emit

package query

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// diagnoseOverrideGaps emits at most one diagnostic per rule that names
// the missing nullability side(s). When both NULL and NOT NULL entries
// are present, returns nil. When both are absent, both are listed in a
// single message — one rule, one diagnostic.
func diagnoseOverrideGaps(rule pgtypeOverrideRule, hasNotNull, hasNullable bool) []diagnostic.Diagnostic {
	var missing string
	switch {
	case hasNotNull && hasNullable:
		return nil
	case !hasNotNull && !hasNullable:
		missing = "nullable=false and nullable=true"
	case !hasNotNull:
		missing = "nullable=false"
	default:
		missing = "nullable=true"
	}
	return []diagnostic.Diagnostic{{
		File:    "db/sqlc.yaml",
		Line:    0,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[" + rule.RuleID + "] DDL has " + rule.DBType + " column(s) but sqlc.yaml is missing " + rule.PgPackage + "." + rule.PgType + " overrides (" + missing + ")",
		Advice:  rule.Advice,
	}}
}
