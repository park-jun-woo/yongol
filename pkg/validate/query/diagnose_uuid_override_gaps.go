//ff:func feature=validate type=rule control=selection topic=query-structural
//ff:what diagnoseUUIDOverrideGaps — Q-12 진단 1건 생성 (양쪽 누락은 합산 1건)

package query

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// diagnoseUUIDOverrideGaps emits at most one Q-12 diagnostic that captures
// the missing nullable side(s). When both NULL and NOT NULL entries are
// present, returns nil. When both are absent, the message marks both as
// missing in a single diagnostic — one rule, one message.
func diagnoseUUIDOverrideGaps(hasNotNull, hasNullable bool) []diagnostic.Diagnostic {
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
		Message: "[Q-12] DDL has UUID column(s) but sqlc.yaml is missing pgtype.UUID overrides (" + missing + ")",
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"uuid\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"UUID\"\n" +
			"  - db_type: \"uuid\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"UUID\"\n" +
			"PostgreSQL types without a Go native equivalent (UUID, NUMERIC, JSONB, INET, INTERVAL) require explicit pgtype overrides in sqlc.yaml.",
	}}
}
