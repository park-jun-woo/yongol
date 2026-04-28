//ff:func feature=validate type=util control=iteration dimension=2 topic=query-structural
//ff:what scanUUIDOverrides — sqlc.yaml overrides 에서 pgtype.UUID NULL/NOT NULL entry 존재 여부 집계

package query

// scanUUIDOverrides walks cfg.SQL[*].Gen.Go.Overrides and reports whether
// the two required pgtype.UUID entries (nullable=false and nullable=true)
// are each present. Q-12 uses the booleans to decide which side(s) to flag
// in the diagnostic message.
func scanUUIDOverrides(cfg sqlcOverridesConfig) (hasNotNull, hasNullable bool) {
	for _, sql := range cfg.SQL {
		for _, ov := range sql.Gen.Go.Overrides {
			if !isPgtypeUUIDOverride(ov) {
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
