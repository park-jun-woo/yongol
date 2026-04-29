//ff:type feature=validate type=model topic=query-structural
//ff:what pgtypeOverrideRule — Q-NN 룰별 sqlc.yaml override 검사 파라미터 (RuleID/DBType/PgPackage/PgType/Filter/Advice)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// pgtypeOverrideRule wraps the per-rule parameters checkPgtypeOverride
// needs. Each per-type Q-NN rule supplies a value of this struct; the
// helper performs the load → scan → diagnose pipeline uniformly. The
// struct is shared by Q-12 (UUID), Q-13 (NUMERIC), Q-14~16 (timestamp
// family), Q-17 (INET), Q-18 (INTERVAL).
type pgtypeOverrideRule struct {
	// RuleID is the diagnostic prefix (e.g. "Q-12", "Q-13").
	RuleID string

	// DBType is the lowercased PG type that sqlc.yaml `db_type` keys must
	// match (e.g. "uuid", "numeric", "inet").
	DBType string

	// PgPackage / PgType identify the expected pgtype Go type
	// (e.g. "pgtype" / "UUID", "pgtype" / "Numeric").
	PgPackage string
	PgType    string

	// Filter selects DDL columns this rule applies to. The helper invokes
	// it per column; if any column matches, the rule fires when the
	// override is missing.
	Filter func(col ddl.Column) bool

	// Advice is the per-rule advice block appended to the diagnostic
	// message. Use a YAML stanza the user can paste into sqlc.yaml.
	Advice string
}
