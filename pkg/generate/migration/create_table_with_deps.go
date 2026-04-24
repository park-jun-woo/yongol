//ff:func feature=migration type=util control=iteration dimension=1
//ff:what createTableWithDeps — 신규 테이블을 CreateTable + CreateIndex/AddForeignKey/AddCheck 로 변환
package migration

// createTableWithDeps returns the ops needed to materialise a brand-new
// table: CreateTable, its sentinel INSERTs, then its indexes, FKs and
// check constraints. Phase sorting (sortByDependency) guarantees that
// every CreateTable executes before any InsertSentinel, and every
// InsertSentinel executes before any AddForeignKey — so FK references
// always find their sentinel target rows pre-seeded.
func createTableWithDeps(c *Table) []Operation {
	ops := []Operation{CreateTable{Table: c}}
	for _, s := range c.Sentinels {
		ops = append(ops, InsertSentinel{Table: c.Name, Body: s.SQL})
	}
	for _, idx := range c.Indexes {
		ops = append(ops, CreateIndex{Table: c.Name, Index: idx})
	}
	for _, fk := range c.ForeignKeys {
		ops = append(ops, AddForeignKey{Table: c.Name, FK: fk})
	}
	for _, chk := range c.Checks {
		ops = append(ops, AddCheck{Table: c.Name, Check: chk})
	}
	return ops
}
