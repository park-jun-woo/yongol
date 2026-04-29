//ff:func feature=validate type=util control=iteration dimension=2 topic=query-structural
//ff:what scanOverridesFor — sqlc.yaml overrides 에서 (db_type, pgtype) NOT NULL/NULLABLE entry 존재 여부 집계

package query

// scanOverridesFor walks cfg.SQL[*].Gen.Go.Overrides and reports whether
// the matching overrides for both nullable=false and nullable=true are
// present. dbType is matched case-insensitively against the override's
// `db_type` field; pgPackage / pgType are matched against the
// `go_type.package` / `go_type.type` fields. The import path is also
// constrained to the canonical pgx/v5 pgtype location.
func scanOverridesFor(cfg sqlcOverridesConfig, dbType, pgPackage, pgType string) (hasNotNull, hasNullable bool) {
	for _, sql := range cfg.SQL {
		for _, ov := range sql.Gen.Go.Overrides {
			if !isPgtypeOverrideFor(ov, dbType, pgPackage, pgType) {
				continue
			}
			if ov.Nullable {
				hasNullable = true
			} else {
				hasNotNull = true
			}
		}
	}
	return hasNotNull, hasNullable
}
