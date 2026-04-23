//ff:func feature=migration type=util control=sequence
//ff:what columnAlterForPair — prev/curr 한 쌍의 차이에서 AlterColumnType/Nullable/Default 생성
package migration

// columnAlterForPair compares one prev/curr column pair and returns the
// relevant ALTER ops.
func columnAlterForPair(tableName, name string, pc, cc *Column, hints *Hints) []Operation {
	var ops []Operation
	if !pc.Type.Equal(cc.Type) {
		ops = append(ops, buildAlterColumnType(tableName, name, pc.Type, cc.Type, hints))
	}
	if pc.Nullable != cc.Nullable {
		ops = append(ops, buildAlterColumnNullable(tableName, name, pc.Nullable, cc.Nullable, hints))
	}
	if pc.Default != cc.Default {
		ops = append(ops, AlterColumnDefault{
			Table: tableName, Column: name,
			From: pc.Default, To: cc.Default,
		})
	}
	return ops
}
