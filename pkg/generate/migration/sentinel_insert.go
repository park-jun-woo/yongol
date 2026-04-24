//ff:type feature=migration type=model
//ff:what SentinelInsert — @sentinel 어노테이션 INSERT 블록 verbatim 래퍼
package migration

// SentinelInsert captures one `-- @sentinel` INSERT statement verbatim,
// so the migration emitter can embed it unchanged after all CREATE TABLE
// and before CREATE INDEX / ALTER TABLE ADD FOREIGN KEY statements.
type SentinelInsert struct {
	SQL string // raw SQL through the final `;`
}
