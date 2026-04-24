//ff:type feature=manifest type=model
//ff:what SentinelInsert — verbatim 으로 보존되는 `-- @sentinel` INSERT 블록
package ddl

// SentinelInsert captures a single `-- @sentinel` INSERT block verbatim.
// The emitter writes SQL untouched; no further parsing of the INSERT body
// occurs.
type SentinelInsert struct {
	SQL  string // raw SQL including "INSERT" through the final ";"
	Line int    // source line of the INSERT keyword (1-based)
	File string // source .sql file path
}
