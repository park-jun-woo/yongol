//ff:func feature=migration type=util control=iteration dimension=1
//ff:what CheckSafety — Operation 리스트 전수 점검 → warnings / errors 분리
package migration

import (
	"fmt"
	"sort"
)

// SafetyIssue carries one diagnostic produced by CheckSafety. The Rule
// ID lets the CLI / validate pipeline map it to MIG-00N entries.
type SafetyIssue struct {
	Level   SafetyLevel
	RuleID  string // "MIG-002" / "MIG-004" / "MIG-005"
	Message string
	Advice  string // suggested fix (hint to add, etc)
}

// ApplyHintsToOps mutates each operation (in place) so its
// AllowDestructive / Backfill / Using fields reflect the supplied
// hints. Call this after Diff so ops' SafetyLevel() respects hints.
func ApplyHintsToOps(ops []Operation, hints *Hints) []Operation {
	if hints == nil {
		return ops
	}
	out := make([]Operation, len(ops))
	for i, op := range ops {
		out[i] = applyHint(op, hints)
	}
	return out
}

func applyHint(op Operation, hints *Hints) Operation {
	switch v := op.(type) {
	case DropTable:
		if hints.AllowDestructive[v.Name] {
			v.AllowDestructive = true
		}
		return v
	case DropColumn:
		if hints.AllowDestructive[v.Table] {
			v.AllowDestructive = true
		}
		return v
	case AddColumn:
		if b, ok := hints.Backfills[colKey{Table: v.Table, Column: v.Column.Name}]; ok {
			v.Backfill = b
		}
		return v
	case AlterColumnNullable:
		if b, ok := hints.Backfills[colKey{Table: v.Table, Column: v.Column}]; ok {
			v.Backfill = b
		}
		return v
	case AlterColumnType:
		if using, ok := hints.Casts[colKey{Table: v.Table, Column: v.Column}]; ok {
			v.Using = using
		}
		return v
	}
	return op
}

// CheckSafety returns the list of issues found across the op list.
// Callers decide whether to abort (any SafetyError) or just warn
// (SafetyWarning only).
func CheckSafety(ops []Operation) []SafetyIssue {
	var issues []SafetyIssue
	for _, op := range ops {
		switch v := op.(type) {
		case AlterColumnNullable:
			if !v.To && v.Backfill == "" {
				issues = append(issues, SafetyIssue{
					Level:   SafetyError,
					RuleID:  "MIG-002",
					Message: fmt.Sprintf("NOT NULL added to %s.%s without @backfill hint", v.Table, v.Column),
					Advice:  fmt.Sprintf("add `-- @backfill default=<value>` on the %s column line in specs/db/*.sql", v.Column),
				})
			}
		case AddColumn:
			if !v.Column.Nullable && v.Column.Default == "" && v.Backfill == "" {
				issues = append(issues, SafetyIssue{
					Level:   SafetyError,
					RuleID:  "MIG-002",
					Message: fmt.Sprintf("NOT NULL column %s.%s added without DEFAULT or @backfill", v.Table, v.Column.Name),
					Advice:  fmt.Sprintf("set DEFAULT in DDL or add `-- @backfill default=<value>` on %s line", v.Column.Name),
				})
			}
		case DropTable:
			if !v.AllowDestructive {
				issues = append(issues, SafetyIssue{
					Level:   SafetyWarning,
					RuleID:  "MIG-004",
					Message: fmt.Sprintf("DROP TABLE %s without @allow_destructive", v.Name),
					Advice:  "add `-- @allow_destructive` above the (removed) CREATE TABLE or keep the table",
				})
			}
		case DropColumn:
			if !v.AllowDestructive {
				issues = append(issues, SafetyIssue{
					Level:   SafetyWarning,
					RuleID:  "MIG-004",
					Message: fmt.Sprintf("DROP COLUMN %s.%s without @allow_destructive", v.Table, v.Column),
					Advice:  fmt.Sprintf("add `-- @allow_destructive` above CREATE TABLE %s", v.Table),
				})
			}
		case AlterColumnType:
			if v.Using == "" && riskyCast(v.From, v.To) {
				issues = append(issues, SafetyIssue{
					Level:   SafetyWarning,
					RuleID:  "MIG-005",
					Message: fmt.Sprintf("type change %s.%s %s → %s without @cast USING", v.Table, v.Column, v.From.SQL(), v.To.SQL()),
					Advice:  fmt.Sprintf("add `-- @cast using=<expr>` on %s line if default USING cast is unsafe", v.Column),
				})
			}
		}
	}
	// Deterministic ordering.
	sort.SliceStable(issues, func(i, j int) bool {
		if issues[i].RuleID != issues[j].RuleID {
			return issues[i].RuleID < issues[j].RuleID
		}
		return issues[i].Message < issues[j].Message
	})
	return issues
}

func riskyCast(from, to CanonicalType) bool {
	if from.Base == to.Base {
		// VARCHAR(N) shrink is risky; NUMERIC(p,s) reduce too.
		if from.Base == "VARCHAR" && from.Length > 0 && to.Length > 0 && to.Length < from.Length {
			return true
		}
		if from.Base == "NUMERIC" && from.Precision > 0 && to.Precision > 0 && to.Precision < from.Precision {
			return true
		}
		return false
	}
	// Numeric ↔ text transitions are always risky.
	num := map[string]bool{"INTEGER": true, "BIGINT": true, "SMALLINT": true, "NUMERIC": true, "REAL": true, "DOUBLE PRECISION": true}
	text := map[string]bool{"TEXT": true, "VARCHAR": true, "CHAR": true}
	if (num[from.Base] && text[to.Base]) || (text[from.Base] && num[to.Base]) {
		return true
	}
	return false
}
