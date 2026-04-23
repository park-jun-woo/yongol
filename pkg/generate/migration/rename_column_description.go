//ff:func feature=migration type=accessor control=sequence
//ff:what RenameColumn.Description — 파일명/헤더용 설명 문자열
package migration

func (op RenameColumn) Description() string {
	return "rename column " + op.Table + "." + op.From + " → " + op.To
}
