//ff:func feature=migration type=util control=iteration dimension=1
//ff:what diffDropTable — prev 에만 있는 테이블: rename 소스면 skip, 아니면 FK 정리 후 DropTable
package migration

// diffDropTable emits DropForeignKey for each FK on p then DropTable
// unless the table is a rename source (handled by RenameTable already).
func diffDropTable(n string, p *Table, renamed map[string]string) []Operation {
	if _, isRenameSource := renamed[n]; isRenameSource {
		return nil
	}
	var ops []Operation
	for _, fk := range p.ForeignKeys {
		ops = append(ops, DropForeignKey{Table: p.Name, Name: fk.Name})
	}
	ops = append(ops, DropTable{Name: n})
	return ops
}
