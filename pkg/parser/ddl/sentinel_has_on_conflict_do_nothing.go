//ff:func feature=manifest type=parser control=sequence
//ff:what sentinelHasOnConflictDoNothing — INSERT body 에 ON CONFLICT DO NOTHING 포함 여부
package ddl

// sentinelHasOnConflictDoNothing reports whether the collected INSERT
// text contains an `ON CONFLICT ... DO NOTHING` clause.
func sentinelHasOnConflictDoNothing(sql string) bool {
	return onConflictDoNothingRe.MatchString(sql)
}
