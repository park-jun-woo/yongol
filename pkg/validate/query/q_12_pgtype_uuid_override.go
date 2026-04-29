//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-12 — DDL 에 UUID 컬럼이 있으면 sqlc.yaml 에 pgtype.UUID overrides (NULL/NOT NULL) 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q12PgtypeUuidOverride validates Q-12: when any DDL file declares a UUID
// column, db/sqlc.yaml must register two pgtype.UUID overrides — one for
// nullable=false, one for nullable=true. sqlc's pgx/v5 mode has no default
// mapping for UUID, so without the explicit override sqlc may emit
// `types.UUID` (or pgtype.UUID inconsistently across nullability), breaking
// the yongol-generated handler that imports `github.com/jackc/pgx/v5/pgtype`.
//
// Implementation: thin wrapper over checkPgtypeOverride. The shared
// helper is also used by per-type Q-NN rules (NUMERIC / INET / INTERVAL /
// timestamp family) so all override-required pgtypes share one
// diagnostic pipeline and one drift-free policy.
func q12PgtypeUuidOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-12",
		DBType:    "uuid",
		PgPackage: "pgtype",
		PgType:    "UUID",
		Filter:    isUUIDColumn,
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
	})
}
