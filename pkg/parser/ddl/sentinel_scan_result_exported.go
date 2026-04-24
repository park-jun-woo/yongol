//ff:type feature=manifest type=model
//ff:what SentinelScanResult — 외부 공개용 sentinelScanResult 래퍼
package ddl

// SentinelScanResult is the exported mirror of sentinelScanResult for
// downstream packages (pkg/generate/migration) that need to reuse the
// scan logic without duplicating the quote-aware terminator state.
type SentinelScanResult struct {
	Table     string // target table name (raw, un-canonicalised)
	SQL       string // raw SQL including INSERT through final ";"
	StartLine int    // 1-based line number of the INSERT keyword
	Annotated bool   // true when `-- @sentinel` preceded the INSERT
}
