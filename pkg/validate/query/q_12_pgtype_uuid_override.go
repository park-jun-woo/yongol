//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-12 — DDL 에 UUID 컬럼이 있으면 sqlc.yaml 에 pgtype.UUID overrides (NULL/NOT NULL) 강제

package query

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

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
// Both entries missing collapse into a single diagnostic (one rule = one
// message; the advice block carries both YAML stanzas so the user pastes
// once). One missing entry produces one diagnostic that names which
// nullable side is absent. Skipped entirely when DDL has no UUID columns
// or when sqlc.yaml is unreadable (D-4 already reports the latter).
func q12PgtypeUuidOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.SpecsDir == "" {
		return nil
	}
	if !ddlHasUUIDColumn(fs.SpecsDir) {
		return nil
	}
	sqlcPath := filepath.Join(fs.SpecsDir, "db", "sqlc.yaml")
	data, err := os.ReadFile(sqlcPath)
	if err != nil {
		return nil
	}
	var cfg sqlcOverridesConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil
	}
	hasNotNull, hasNullable := scanUUIDOverrides(cfg)
	return diagnoseUUIDOverrideGaps(hasNotNull, hasNullable)
}
