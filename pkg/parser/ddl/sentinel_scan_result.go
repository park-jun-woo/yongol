//ff:type feature=manifest type=model
//ff:what sentinelScanResult — parseSentinelInserts 내부 결과 구조체
package ddl

// sentinelScanResult carries one INSERT occurrence plus a flag telling
// the caller whether `-- @sentinel` preceded it.
type sentinelScanResult struct {
	Table     string
	SQL       string
	StartLine int
	Annotated bool
}
