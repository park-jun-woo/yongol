//ff:func feature=manifest type=parser control=sequence
//ff:what SentinelHasOnConflictDoNothing — ON CONFLICT DO NOTHING 탐지기 공개 API
package ddl

// SentinelHasOnConflictDoNothing exposes the ON CONFLICT DO NOTHING
// detector so migration / tests can reuse it.
func SentinelHasOnConflictDoNothing(sql string) bool {
	return sentinelHasOnConflictDoNothing(sql)
}
